// Copyright (C) 2019-2025, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package main

import (
	"context"
	"log"
	"net/netip"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"google.golang.org/protobuf/proto"

	"github.com/ava-labs/avalanchego/api/info"
	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/network/p2p"
	"github.com/ava-labs/avalanchego/network/peer"
	"github.com/ava-labs/avalanchego/proto/pb/sdk"
	"github.com/ava-labs/avalanchego/snow/networking/router"
	"github.com/ava-labs/avalanchego/utils/compression"
	"github.com/ava-labs/avalanchego/utils/constants"
	"github.com/ava-labs/avalanchego/vms/platformvm/warp"
	"github.com/ava-labs/avalanchego/vms/platformvm/warp/payload"
	"github.com/ava-labs/avalanchego/wallet/subnet/primary"

	p2pmessage "github.com/ava-labs/avalanchego/message"
	warpmessage "github.com/ava-labs/avalanchego/vms/platformvm/warp/message"
)

func main() {
	uri := primary.LocalAPIURI
	subnetID := ids.FromStringOrPanic("2DeHa7Qb6sufPkmQcFWG2uCd4pBPv9WB6dkzroiMQhd1NSRtof")
	conversionID := ids.FromStringOrPanic("28tfqwucuoH7oWxmVYDVQ2C1ehdYecF5mzwNmX2t1dTu1S5vHE")
	infoClient := info.NewClient(uri)
	networkID, err := infoClient.GetNetworkID(context.Background())
	if err != nil {
		log.Fatalf("failed to fetch network ID: %s\n", err)
	}

	subnetToL1Conversion, err := warpmessage.NewSubnetToL1Conversion(conversionID)
	if err != nil {
		log.Fatalf("failed to create SubnetToL1Conversion message: %s\n", err)
	}

	addressedCall, err := payload.NewAddressedCall(
		nil,
		subnetToL1Conversion.Bytes(),
	)
	if err != nil {
		log.Fatalf("failed to create AddressedCall message: %s\n", err)
	}

	unsignedWarp, err := warp.NewUnsignedMessage(
		networkID,
		constants.PlatformChainID,
		addressedCall.Bytes(),
	)
	if err != nil {
		log.Fatalf("failed to create unsigned Warp message: %s\n", err)
	}

	p, err := peer.StartTestPeer(
		context.Background(),
		netip.AddrPortFrom(
			netip.AddrFrom4([4]byte{127, 0, 0, 1}),
			9651,
		),
		networkID,
		router.InboundHandlerFunc(func(_ context.Context, msg p2pmessage.InboundMessage) {
			log.Printf("received %s: %s", msg.Op(), msg.Message())
		}),
	)
	if err != nil {
		log.Fatalf("failed to start peer: %s\n", err)
	}

	messageBuilder, err := p2pmessage.NewCreator(
		prometheus.NewRegistry(),
		compression.TypeZstd,
		time.Hour,
	)
	if err != nil {
		log.Fatalf("failed to create message builder: %s\n", err)
	}

	appRequestPayload, err := proto.Marshal(&sdk.SignatureRequest{
		Message:       unsignedWarp.Bytes(),
		Justification: subnetID[:],
	})
	if err != nil {
		log.Fatalf("failed to marshal SignatureRequest: %s\n", err)
	}

	appRequest, err := messageBuilder.AppRequest(
		constants.PlatformChainID,
		0,
		time.Hour,
		p2p.PrefixMessage(
			p2p.ProtocolPrefix(p2p.SignatureRequestHandlerID),
			appRequestPayload,
		),
	)
	if err != nil {
		log.Fatalf("failed to create AppRequest: %s\n", err)
	}

	p.Send(context.Background(), appRequest)

	time.Sleep(5 * time.Second)

	p.StartClose()
	err = p.AwaitClosed(context.Background())
	if err != nil {
		log.Fatalf("failed to close peer: %s\n", err)
	}
}
