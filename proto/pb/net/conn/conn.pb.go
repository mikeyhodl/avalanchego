// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: net/conn/conn.proto

package conn

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// ErrorCode provides information for special sentinel error types
type ErrorCode int32

const (
	ErrorCode_ERROR_CODE_UNSPECIFIED              ErrorCode = 0
	ErrorCode_ERROR_CODE_EOF                      ErrorCode = 1
	ErrorCode_ERROR_CODE_OS_ERR_DEADLINE_EXCEEDED ErrorCode = 2
)

// Enum value maps for ErrorCode.
var (
	ErrorCode_name = map[int32]string{
		0: "ERROR_CODE_UNSPECIFIED",
		1: "ERROR_CODE_EOF",
		2: "ERROR_CODE_OS_ERR_DEADLINE_EXCEEDED",
	}
	ErrorCode_value = map[string]int32{
		"ERROR_CODE_UNSPECIFIED":              0,
		"ERROR_CODE_EOF":                      1,
		"ERROR_CODE_OS_ERR_DEADLINE_EXCEEDED": 2,
	}
)

func (x ErrorCode) Enum() *ErrorCode {
	p := new(ErrorCode)
	*p = x
	return p
}

func (x ErrorCode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ErrorCode) Descriptor() protoreflect.EnumDescriptor {
	return file_net_conn_conn_proto_enumTypes[0].Descriptor()
}

func (ErrorCode) Type() protoreflect.EnumType {
	return &file_net_conn_conn_proto_enumTypes[0]
}

func (x ErrorCode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ErrorCode.Descriptor instead.
func (ErrorCode) EnumDescriptor() ([]byte, []int) {
	return file_net_conn_conn_proto_rawDescGZIP(), []int{0}
}

type ReadRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// length of the request in bytes
	Length        int32 `protobuf:"varint,1,opt,name=length,proto3" json:"length,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ReadRequest) Reset() {
	*x = ReadRequest{}
	mi := &file_net_conn_conn_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ReadRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadRequest) ProtoMessage() {}

func (x *ReadRequest) ProtoReflect() protoreflect.Message {
	mi := &file_net_conn_conn_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadRequest.ProtoReflect.Descriptor instead.
func (*ReadRequest) Descriptor() ([]byte, []int) {
	return file_net_conn_conn_proto_rawDescGZIP(), []int{0}
}

func (x *ReadRequest) GetLength() int32 {
	if x != nil {
		return x.Length
	}
	return 0
}

type ReadResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// read is the payload in bytes
	Read []byte `protobuf:"bytes,1,opt,name=read,proto3" json:"read,omitempty"`
	// error is an error message
	Error         *Error `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ReadResponse) Reset() {
	*x = ReadResponse{}
	mi := &file_net_conn_conn_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ReadResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadResponse) ProtoMessage() {}

func (x *ReadResponse) ProtoReflect() protoreflect.Message {
	mi := &file_net_conn_conn_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadResponse.ProtoReflect.Descriptor instead.
func (*ReadResponse) Descriptor() ([]byte, []int) {
	return file_net_conn_conn_proto_rawDescGZIP(), []int{1}
}

func (x *ReadResponse) GetRead() []byte {
	if x != nil {
		return x.Read
	}
	return nil
}

func (x *ReadResponse) GetError() *Error {
	if x != nil {
		return x.Error
	}
	return nil
}

type WriteRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// payload is the write request in bytes
	Payload       []byte `protobuf:"bytes,1,opt,name=payload,proto3" json:"payload,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *WriteRequest) Reset() {
	*x = WriteRequest{}
	mi := &file_net_conn_conn_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *WriteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WriteRequest) ProtoMessage() {}

func (x *WriteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_net_conn_conn_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WriteRequest.ProtoReflect.Descriptor instead.
func (*WriteRequest) Descriptor() ([]byte, []int) {
	return file_net_conn_conn_proto_rawDescGZIP(), []int{2}
}

func (x *WriteRequest) GetPayload() []byte {
	if x != nil {
		return x.Payload
	}
	return nil
}

type WriteResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// length of the response in bytes
	Length int32 `protobuf:"varint,1,opt,name=length,proto3" json:"length,omitempty"`
	// error is an error message
	Error         *string `protobuf:"bytes,2,opt,name=error,proto3,oneof" json:"error,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *WriteResponse) Reset() {
	*x = WriteResponse{}
	mi := &file_net_conn_conn_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *WriteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WriteResponse) ProtoMessage() {}

func (x *WriteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_net_conn_conn_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WriteResponse.ProtoReflect.Descriptor instead.
func (*WriteResponse) Descriptor() ([]byte, []int) {
	return file_net_conn_conn_proto_rawDescGZIP(), []int{3}
}

func (x *WriteResponse) GetLength() int32 {
	if x != nil {
		return x.Length
	}
	return 0
}

func (x *WriteResponse) GetError() string {
	if x != nil && x.Error != nil {
		return *x.Error
	}
	return ""
}

type SetDeadlineRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// time represents an instant in time in bytes
	Time          []byte `protobuf:"bytes,1,opt,name=time,proto3" json:"time,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SetDeadlineRequest) Reset() {
	*x = SetDeadlineRequest{}
	mi := &file_net_conn_conn_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SetDeadlineRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetDeadlineRequest) ProtoMessage() {}

func (x *SetDeadlineRequest) ProtoReflect() protoreflect.Message {
	mi := &file_net_conn_conn_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetDeadlineRequest.ProtoReflect.Descriptor instead.
func (*SetDeadlineRequest) Descriptor() ([]byte, []int) {
	return file_net_conn_conn_proto_rawDescGZIP(), []int{4}
}

func (x *SetDeadlineRequest) GetTime() []byte {
	if x != nil {
		return x.Time
	}
	return nil
}

type Error struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ErrorCode     ErrorCode              `protobuf:"varint,1,opt,name=error_code,json=errorCode,proto3,enum=net.conn.ErrorCode" json:"error_code,omitempty"`
	Message       string                 `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Error) Reset() {
	*x = Error{}
	mi := &file_net_conn_conn_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Error) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Error) ProtoMessage() {}

func (x *Error) ProtoReflect() protoreflect.Message {
	mi := &file_net_conn_conn_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Error.ProtoReflect.Descriptor instead.
func (*Error) Descriptor() ([]byte, []int) {
	return file_net_conn_conn_proto_rawDescGZIP(), []int{5}
}

func (x *Error) GetErrorCode() ErrorCode {
	if x != nil {
		return x.ErrorCode
	}
	return ErrorCode_ERROR_CODE_UNSPECIFIED
}

func (x *Error) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_net_conn_conn_proto protoreflect.FileDescriptor

const file_net_conn_conn_proto_rawDesc = "" +
	"\n" +
	"\x13net/conn/conn.proto\x12\bnet.conn\x1a\x1bgoogle/protobuf/empty.proto\"%\n" +
	"\vReadRequest\x12\x16\n" +
	"\x06length\x18\x01 \x01(\x05R\x06length\"I\n" +
	"\fReadResponse\x12\x12\n" +
	"\x04read\x18\x01 \x01(\fR\x04read\x12%\n" +
	"\x05error\x18\x02 \x01(\v2\x0f.net.conn.ErrorR\x05error\"(\n" +
	"\fWriteRequest\x12\x18\n" +
	"\apayload\x18\x01 \x01(\fR\apayload\"L\n" +
	"\rWriteResponse\x12\x16\n" +
	"\x06length\x18\x01 \x01(\x05R\x06length\x12\x19\n" +
	"\x05error\x18\x02 \x01(\tH\x00R\x05error\x88\x01\x01B\b\n" +
	"\x06_error\"(\n" +
	"\x12SetDeadlineRequest\x12\x12\n" +
	"\x04time\x18\x01 \x01(\fR\x04time\"U\n" +
	"\x05Error\x122\n" +
	"\n" +
	"error_code\x18\x01 \x01(\x0e2\x13.net.conn.ErrorCodeR\terrorCode\x12\x18\n" +
	"\amessage\x18\x02 \x01(\tR\amessage*d\n" +
	"\tErrorCode\x12\x1a\n" +
	"\x16ERROR_CODE_UNSPECIFIED\x10\x00\x12\x12\n" +
	"\x0eERROR_CODE_EOF\x10\x01\x12'\n" +
	"#ERROR_CODE_OS_ERR_DEADLINE_EXCEEDED\x10\x022\x88\x03\n" +
	"\x04Conn\x125\n" +
	"\x04Read\x12\x15.net.conn.ReadRequest\x1a\x16.net.conn.ReadResponse\x128\n" +
	"\x05Write\x12\x16.net.conn.WriteRequest\x1a\x17.net.conn.WriteResponse\x127\n" +
	"\x05Close\x12\x16.google.protobuf.Empty\x1a\x16.google.protobuf.Empty\x12C\n" +
	"\vSetDeadline\x12\x1c.net.conn.SetDeadlineRequest\x1a\x16.google.protobuf.Empty\x12G\n" +
	"\x0fSetReadDeadline\x12\x1c.net.conn.SetDeadlineRequest\x1a\x16.google.protobuf.Empty\x12H\n" +
	"\x10SetWriteDeadline\x12\x1c.net.conn.SetDeadlineRequest\x1a\x16.google.protobuf.EmptyB3Z1github.com/ava-labs/avalanchego/proto/pb/net/connb\x06proto3"

var (
	file_net_conn_conn_proto_rawDescOnce sync.Once
	file_net_conn_conn_proto_rawDescData []byte
)

func file_net_conn_conn_proto_rawDescGZIP() []byte {
	file_net_conn_conn_proto_rawDescOnce.Do(func() {
		file_net_conn_conn_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_net_conn_conn_proto_rawDesc), len(file_net_conn_conn_proto_rawDesc)))
	})
	return file_net_conn_conn_proto_rawDescData
}

var file_net_conn_conn_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_net_conn_conn_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_net_conn_conn_proto_goTypes = []any{
	(ErrorCode)(0),             // 0: net.conn.ErrorCode
	(*ReadRequest)(nil),        // 1: net.conn.ReadRequest
	(*ReadResponse)(nil),       // 2: net.conn.ReadResponse
	(*WriteRequest)(nil),       // 3: net.conn.WriteRequest
	(*WriteResponse)(nil),      // 4: net.conn.WriteResponse
	(*SetDeadlineRequest)(nil), // 5: net.conn.SetDeadlineRequest
	(*Error)(nil),              // 6: net.conn.Error
	(*emptypb.Empty)(nil),      // 7: google.protobuf.Empty
}
var file_net_conn_conn_proto_depIdxs = []int32{
	6, // 0: net.conn.ReadResponse.error:type_name -> net.conn.Error
	0, // 1: net.conn.Error.error_code:type_name -> net.conn.ErrorCode
	1, // 2: net.conn.Conn.Read:input_type -> net.conn.ReadRequest
	3, // 3: net.conn.Conn.Write:input_type -> net.conn.WriteRequest
	7, // 4: net.conn.Conn.Close:input_type -> google.protobuf.Empty
	5, // 5: net.conn.Conn.SetDeadline:input_type -> net.conn.SetDeadlineRequest
	5, // 6: net.conn.Conn.SetReadDeadline:input_type -> net.conn.SetDeadlineRequest
	5, // 7: net.conn.Conn.SetWriteDeadline:input_type -> net.conn.SetDeadlineRequest
	2, // 8: net.conn.Conn.Read:output_type -> net.conn.ReadResponse
	4, // 9: net.conn.Conn.Write:output_type -> net.conn.WriteResponse
	7, // 10: net.conn.Conn.Close:output_type -> google.protobuf.Empty
	7, // 11: net.conn.Conn.SetDeadline:output_type -> google.protobuf.Empty
	7, // 12: net.conn.Conn.SetReadDeadline:output_type -> google.protobuf.Empty
	7, // 13: net.conn.Conn.SetWriteDeadline:output_type -> google.protobuf.Empty
	8, // [8:14] is the sub-list for method output_type
	2, // [2:8] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_net_conn_conn_proto_init() }
func file_net_conn_conn_proto_init() {
	if File_net_conn_conn_proto != nil {
		return
	}
	file_net_conn_conn_proto_msgTypes[3].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_net_conn_conn_proto_rawDesc), len(file_net_conn_conn_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_net_conn_conn_proto_goTypes,
		DependencyIndexes: file_net_conn_conn_proto_depIdxs,
		EnumInfos:         file_net_conn_conn_proto_enumTypes,
		MessageInfos:      file_net_conn_conn_proto_msgTypes,
	}.Build()
	File_net_conn_conn_proto = out.File
	file_net_conn_conn_proto_goTypes = nil
	file_net_conn_conn_proto_depIdxs = nil
}
