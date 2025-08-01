// Copyright (C) 2019-2025, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package nat

import (
	"net/netip"
	"sync"
	"time"

	"go.uber.org/zap"

	"github.com/ava-labs/avalanchego/utils"
	"github.com/ava-labs/avalanchego/utils/logging"
)

const (
	mapTimeout        = 30 * time.Minute
	maxRefreshRetries = 3
)

// Router describes the functionality that a network device must support to be
// able to open ports to an external IP.
type Router interface {
	// True iff this router supports NAT
	SupportsNAT() bool
	// Map external port [extPort] to internal port [intPort] for [duration]
	MapPort(intPort, extPort uint16, desc string, duration time.Duration) error
	// Undo a port mapping
	UnmapPort(intPort, extPort uint16) error
	// Return our external IP
	ExternalIP() (netip.Addr, error)
}

// GetRouter returns a router on the current network.
func GetRouter() Router {
	if r := getUPnPRouter(); r != nil {
		return r
	}
	if r := getPMPRouter(); r != nil {
		return r
	}

	return NewNoRouter()
}

// Mapper attempts to open a set of ports on a router
type Mapper struct {
	log    logging.Logger
	r      Router
	closer chan struct{}
	wg     sync.WaitGroup
}

// NewPortMapper returns an initialized mapper
func NewPortMapper(log logging.Logger, r Router) *Mapper {
	return &Mapper{
		log:    log,
		r:      r,
		closer: make(chan struct{}),
	}
}

// Map external port [extPort] (exposed to the internet) to internal port [intPort] (where our process is listening)
// and set [ip]. Does this every [updateTime]. [ip] may be nil.
func (m *Mapper) Map(
	intPort uint16,
	extPort uint16,
	desc string,
	ip *utils.Atomic[netip.AddrPort],
	updateTime time.Duration,
) {
	if !m.r.SupportsNAT() {
		return
	}

	// we attempt a port map, and log an Error if it fails.
	err := m.retryMapPort(intPort, extPort, desc, mapTimeout)
	if err != nil {
		m.log.Error("NAT traversal failed",
			zap.Uint16("externalPort", extPort),
			zap.Uint16("internalPort", intPort),
			zap.Error(err),
		)
	} else {
		m.log.Info("NAT traversal successful",
			zap.Uint16("externalPort", extPort),
			zap.Uint16("internalPort", intPort),
		)
	}

	m.wg.Add(1)
	go m.keepPortMapping(intPort, extPort, desc, ip, updateTime)
}

// Retry port map up to maxRefreshRetries with a 1 second delay
func (m *Mapper) retryMapPort(intPort, extPort uint16, desc string, timeout time.Duration) error {
	var err error
	for retryCnt := 0; retryCnt < maxRefreshRetries; retryCnt++ {
		err = m.r.MapPort(intPort, extPort, desc, timeout)
		if err == nil {
			return nil
		}

		// log a message, sleep a second and retry.
		m.log.Warn("renewing port mapping failed",
			zap.Int("attempt", retryCnt+1),
			zap.Uint16("externalPort", extPort),
			zap.Uint16("internalPort", intPort),
			zap.Error(err),
		)
		time.Sleep(1 * time.Second)
	}
	return err
}

// keepPortMapping runs in the background to keep a port mapped. It renews the mapping from [extPort]
// to [intPort]] every [updateTime]. Updates [ip] every [updateTime].
func (m *Mapper) keepPortMapping(
	intPort uint16,
	extPort uint16,
	desc string,
	ip *utils.Atomic[netip.AddrPort],
	updateTime time.Duration,
) {
	updateTimer := time.NewTimer(updateTime)

	defer func(extPort uint16) {
		updateTimer.Stop()

		m.log.Debug("unmapping port",
			zap.Uint16("externalPort", extPort),
		)

		if err := m.r.UnmapPort(intPort, extPort); err != nil {
			m.log.Debug("error unmapping port",
				zap.Uint16("externalPort", extPort),
				zap.Uint16("internalPort", intPort),
				zap.Error(err),
			)
		}

		m.wg.Done()
	}(extPort)

	for {
		select {
		case <-updateTimer.C:
			err := m.retryMapPort(intPort, extPort, desc, mapTimeout)
			if err != nil {
				m.log.Warn("renew NAT traversal failed",
					zap.Uint16("externalPort", extPort),
					zap.Uint16("internalPort", intPort),
					zap.Error(err),
				)
			}
			m.updateIP(ip)
			updateTimer.Reset(updateTime)
		case <-m.closer:
			return
		}
	}
}

func (m *Mapper) updateIP(ip *utils.Atomic[netip.AddrPort]) {
	if ip == nil {
		return
	}
	newAddr, err := m.r.ExternalIP()
	if err != nil {
		m.log.Error("failed to get external IP",
			zap.Error(err),
		)
		return
	}
	oldAddrPort := ip.Get()
	oldAddr := oldAddrPort.Addr()
	if newAddr != oldAddr {
		port := oldAddrPort.Port()
		ip.Set(netip.AddrPortFrom(newAddr, port))
		m.log.Info("external IP updated",
			zap.Stringer("oldIP", oldAddr),
			zap.Stringer("newIP", newAddr),
		)
	}
}

// UnmapAllPorts stops mapping all ports from this mapper and attempts to unmap
// them.
func (m *Mapper) UnmapAllPorts() {
	close(m.closer)
	m.wg.Wait()
	m.log.Info("Unmapped all ports")
}
