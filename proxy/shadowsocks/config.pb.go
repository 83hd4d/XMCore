// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.23.1
// source: proxy/shadowsocks/config.proto

package shadowsocks

import (
	net "github.com/xcode75/xcore/common/net"
	protocol "github.com/xcode75/xcore/common/protocol"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CipherType int32

const (
	CipherType_UNKNOWN            CipherType = 0
	CipherType_AES_128_GCM        CipherType = 5
	CipherType_AES_256_GCM        CipherType = 6
	CipherType_CHACHA20_POLY1305  CipherType = 7
	CipherType_XCHACHA20_POLY1305 CipherType = 8
	CipherType_NONE               CipherType = 9
)

// Enum value maps for CipherType.
var (
	CipherType_name = map[int32]string{
		0: "UNKNOWN",
		5: "AES_128_GCM",
		6: "AES_256_GCM",
		7: "CHACHA20_POLY1305",
		8: "XCHACHA20_POLY1305",
		9: "NONE",
	}
	CipherType_value = map[string]int32{
		"UNKNOWN":            0,
		"AES_128_GCM":        5,
		"AES_256_GCM":        6,
		"CHACHA20_POLY1305":  7,
		"XCHACHA20_POLY1305": 8,
		"NONE":               9,
	}
)

func (x CipherType) Enum() *CipherType {
	p := new(CipherType)
	*p = x
	return p
}

func (x CipherType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CipherType) Descriptor() protoreflect.EnumDescriptor {
	return file_proxy_shadowsocks_config_proto_enumTypes[0].Descriptor()
}

func (CipherType) Type() protoreflect.EnumType {
	return &file_proxy_shadowsocks_config_proto_enumTypes[0]
}

func (x CipherType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CipherType.Descriptor instead.
func (CipherType) EnumDescriptor() ([]byte, []int) {
	return file_proxy_shadowsocks_config_proto_rawDescGZIP(), []int{0}
}

type Account struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Password   string     `protobuf:"bytes,1,opt,name=password,proto3" json:"password,omitempty"`
	CipherType CipherType `protobuf:"varint,2,opt,name=cipher_type,json=cipherType,proto3,enum=xray.proxy.shadowsocks.CipherType" json:"cipher_type,omitempty"`
	IvCheck    bool       `protobuf:"varint,3,opt,name=iv_check,json=ivCheck,proto3" json:"iv_check,omitempty"`
}

func (x *Account) Reset() {
	*x = Account{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proxy_shadowsocks_config_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Account) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Account) ProtoMessage() {}

func (x *Account) ProtoReflect() protoreflect.Message {
	mi := &file_proxy_shadowsocks_config_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Account.ProtoReflect.Descriptor instead.
func (*Account) Descriptor() ([]byte, []int) {
	return file_proxy_shadowsocks_config_proto_rawDescGZIP(), []int{0}
}

func (x *Account) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *Account) GetCipherType() CipherType {
	if x != nil {
		return x.CipherType
	}
	return CipherType_UNKNOWN
}

func (x *Account) GetIvCheck() bool {
	if x != nil {
		return x.IvCheck
	}
	return false
}

type ServerConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Users   []*protocol.User `protobuf:"bytes,1,rep,name=users,proto3" json:"users,omitempty"`
	Network []net.Network    `protobuf:"varint,2,rep,packed,name=network,proto3,enum=xray.common.net.Network" json:"network,omitempty"`
}

func (x *ServerConfig) Reset() {
	*x = ServerConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proxy_shadowsocks_config_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServerConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServerConfig) ProtoMessage() {}

func (x *ServerConfig) ProtoReflect() protoreflect.Message {
	mi := &file_proxy_shadowsocks_config_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServerConfig.ProtoReflect.Descriptor instead.
func (*ServerConfig) Descriptor() ([]byte, []int) {
	return file_proxy_shadowsocks_config_proto_rawDescGZIP(), []int{1}
}

func (x *ServerConfig) GetUsers() []*protocol.User {
	if x != nil {
		return x.Users
	}
	return nil
}

func (x *ServerConfig) GetNetwork() []net.Network {
	if x != nil {
		return x.Network
	}
	return nil
}

type ClientConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Server []*protocol.ServerEndpoint `protobuf:"bytes,1,rep,name=server,proto3" json:"server,omitempty"`
}

func (x *ClientConfig) Reset() {
	*x = ClientConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proxy_shadowsocks_config_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientConfig) ProtoMessage() {}

func (x *ClientConfig) ProtoReflect() protoreflect.Message {
	mi := &file_proxy_shadowsocks_config_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientConfig.ProtoReflect.Descriptor instead.
func (*ClientConfig) Descriptor() ([]byte, []int) {
	return file_proxy_shadowsocks_config_proto_rawDescGZIP(), []int{2}
}

func (x *ClientConfig) GetServer() []*protocol.ServerEndpoint {
	if x != nil {
		return x.Server
	}
	return nil
}

var File_proxy_shadowsocks_config_proto protoreflect.FileDescriptor

var file_proxy_shadowsocks_config_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f, 0x73, 0x68, 0x61, 0x64, 0x6f, 0x77, 0x73, 0x6f,
	0x63, 0x6b, 0x73, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x16, 0x78, 0x72, 0x61, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x73, 0x68, 0x61,
	0x64, 0x6f, 0x77, 0x73, 0x6f, 0x63, 0x6b, 0x73, 0x1a, 0x18, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2f, 0x6e, 0x65, 0x74, 0x2f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1a, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x6f, 0x6c, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x21,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x73, 0x70, 0x65, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x85, 0x01, 0x0a, 0x07, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1a, 0x0a,
	0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x43, 0x0a, 0x0b, 0x63, 0x69, 0x70,
	0x68, 0x65, 0x72, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x22,
	0x2e, 0x78, 0x72, 0x61, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x73, 0x68, 0x61, 0x64,
	0x6f, 0x77, 0x73, 0x6f, 0x63, 0x6b, 0x73, 0x2e, 0x43, 0x69, 0x70, 0x68, 0x65, 0x72, 0x54, 0x79,
	0x70, 0x65, 0x52, 0x0a, 0x63, 0x69, 0x70, 0x68, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x12, 0x19,
	0x0a, 0x08, 0x69, 0x76, 0x5f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x07, 0x69, 0x76, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x22, 0x74, 0x0a, 0x0c, 0x53, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x30, 0x0a, 0x05, 0x75, 0x73, 0x65,
	0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x78, 0x72, 0x61, 0x79, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e,
	0x55, 0x73, 0x65, 0x72, 0x52, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x12, 0x32, 0x0a, 0x07, 0x6e,
	0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x18, 0x2e, 0x78,
	0x72, 0x61, 0x79, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x6e, 0x65, 0x74, 0x2e, 0x4e,
	0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x52, 0x07, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x22,
	0x4c, 0x0a, 0x0c, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12,
	0x3c, 0x0a, 0x06, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x24, 0x2e, 0x78, 0x72, 0x61, 0x79, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x45, 0x6e, 0x64,
	0x70, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x06, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2a, 0x74, 0x0a,
	0x0a, 0x43, 0x69, 0x70, 0x68, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x55,
	0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x0f, 0x0a, 0x0b, 0x41, 0x45, 0x53, 0x5f,
	0x31, 0x32, 0x38, 0x5f, 0x47, 0x43, 0x4d, 0x10, 0x05, 0x12, 0x0f, 0x0a, 0x0b, 0x41, 0x45, 0x53,
	0x5f, 0x32, 0x35, 0x36, 0x5f, 0x47, 0x43, 0x4d, 0x10, 0x06, 0x12, 0x15, 0x0a, 0x11, 0x43, 0x48,
	0x41, 0x43, 0x48, 0x41, 0x32, 0x30, 0x5f, 0x50, 0x4f, 0x4c, 0x59, 0x31, 0x33, 0x30, 0x35, 0x10,
	0x07, 0x12, 0x16, 0x0a, 0x12, 0x58, 0x43, 0x48, 0x41, 0x43, 0x48, 0x41, 0x32, 0x30, 0x5f, 0x50,
	0x4f, 0x4c, 0x59, 0x31, 0x33, 0x30, 0x35, 0x10, 0x08, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x4f, 0x4e,
	0x45, 0x10, 0x09, 0x42, 0x64, 0x0a, 0x1a, 0x63, 0x6f, 0x6d, 0x2e, 0x78, 0x72, 0x61, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x73, 0x68, 0x61, 0x64, 0x6f, 0x77, 0x73, 0x6f, 0x63, 0x6b,
	0x73, 0x50, 0x01, 0x5a, 0x2b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x78, 0x74, 0x6c, 0x73, 0x2f, 0x78, 0x72, 0x61, 0x79, 0x2d, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x78, 0x79, 0x2f, 0x73, 0x68, 0x61, 0x64, 0x6f, 0x77, 0x73, 0x6f, 0x63, 0x6b, 0x73,
	0xaa, 0x02, 0x16, 0x58, 0x72, 0x61, 0x79, 0x2e, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x53, 0x68,
	0x61, 0x64, 0x6f, 0x77, 0x73, 0x6f, 0x63, 0x6b, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_proxy_shadowsocks_config_proto_rawDescOnce sync.Once
	file_proxy_shadowsocks_config_proto_rawDescData = file_proxy_shadowsocks_config_proto_rawDesc
)

func file_proxy_shadowsocks_config_proto_rawDescGZIP() []byte {
	file_proxy_shadowsocks_config_proto_rawDescOnce.Do(func() {
		file_proxy_shadowsocks_config_proto_rawDescData = protoimpl.X.CompressGZIP(file_proxy_shadowsocks_config_proto_rawDescData)
	})
	return file_proxy_shadowsocks_config_proto_rawDescData
}

var file_proxy_shadowsocks_config_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_proxy_shadowsocks_config_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_proxy_shadowsocks_config_proto_goTypes = []interface{}{
	(CipherType)(0),                 // 0: xray.proxy.shadowsocks.CipherType
	(*Account)(nil),                 // 1: xray.proxy.shadowsocks.Account
	(*ServerConfig)(nil),            // 2: xray.proxy.shadowsocks.ServerConfig
	(*ClientConfig)(nil),            // 3: xray.proxy.shadowsocks.ClientConfig
	(*protocol.User)(nil),           // 4: xray.common.protocol.User
	(net.Network)(0),                // 5: xray.common.net.Network
	(*protocol.ServerEndpoint)(nil), // 6: xray.common.protocol.ServerEndpoint
}
var file_proxy_shadowsocks_config_proto_depIdxs = []int32{
	0, // 0: xray.proxy.shadowsocks.Account.cipher_type:type_name -> xray.proxy.shadowsocks.CipherType
	4, // 1: xray.proxy.shadowsocks.ServerConfig.users:type_name -> xray.common.protocol.User
	5, // 2: xray.proxy.shadowsocks.ServerConfig.network:type_name -> xray.common.net.Network
	6, // 3: xray.proxy.shadowsocks.ClientConfig.server:type_name -> xray.common.protocol.ServerEndpoint
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_proxy_shadowsocks_config_proto_init() }
func file_proxy_shadowsocks_config_proto_init() {
	if File_proxy_shadowsocks_config_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proxy_shadowsocks_config_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Account); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proxy_shadowsocks_config_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServerConfig); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proxy_shadowsocks_config_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientConfig); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proxy_shadowsocks_config_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proxy_shadowsocks_config_proto_goTypes,
		DependencyIndexes: file_proxy_shadowsocks_config_proto_depIdxs,
		EnumInfos:         file_proxy_shadowsocks_config_proto_enumTypes,
		MessageInfos:      file_proxy_shadowsocks_config_proto_msgTypes,
	}.Build()
	File_proxy_shadowsocks_config_proto = out.File
	file_proxy_shadowsocks_config_proto_rawDesc = nil
	file_proxy_shadowsocks_config_proto_goTypes = nil
	file_proxy_shadowsocks_config_proto_depIdxs = nil
}
