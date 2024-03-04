// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.3
// source: register.proto

package objects

import (
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

type Command int32

const (
	Command_Unkwn   Command = 0
	Command_Ping    Command = 1
	Command_Pong    Command = 2
	Command_Reg     Command = 3
	Command_Log     Command = 4
	Command_PrivMsg Command = 5
	Command_GrpMsg  Command = 6
)

// Enum value maps for Command.
var (
	Command_name = map[int32]string{
		0: "Unkwn",
		1: "Ping",
		2: "Pong",
		3: "Reg",
		4: "Log",
		5: "PrivMsg",
		6: "GrpMsg",
	}
	Command_value = map[string]int32{
		"Unkwn":   0,
		"Ping":    1,
		"Pong":    2,
		"Reg":     3,
		"Log":     4,
		"PrivMsg": 5,
		"GrpMsg":  6,
	}
)

func (x Command) Enum() *Command {
	p := new(Command)
	*p = x
	return p
}

func (x Command) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Command) Descriptor() protoreflect.EnumDescriptor {
	return file_register_proto_enumTypes[0].Descriptor()
}

func (Command) Type() protoreflect.EnumType {
	return &file_register_proto_enumTypes[0]
}

func (x Command) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Command.Descriptor instead.
func (Command) EnumDescriptor() ([]byte, []int) {
	return file_register_proto_rawDescGZIP(), []int{0}
}

type Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Command     Command      `protobuf:"varint,2,opt,name=command,proto3,enum=Command" json:"command,omitempty"`
	User        *User        `protobuf:"bytes,3,opt,name=user,proto3" json:"user,omitempty"`
	HostAddress *HostAddress `protobuf:"bytes,4,opt,name=hostAddress,proto3" json:"hostAddress,omitempty"`
}

func (x *Message) Reset() {
	*x = Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_register_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
	mi := &file_register_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message.ProtoReflect.Descriptor instead.
func (*Message) Descriptor() ([]byte, []int) {
	return file_register_proto_rawDescGZIP(), []int{0}
}

func (x *Message) GetCommand() Command {
	if x != nil {
		return x.Command
	}
	return Command_Unkwn
}

func (x *Message) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *Message) GetHostAddress() *HostAddress {
	if x != nil {
		return x.HostAddress
	}
	return nil
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status string `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Msg    string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_register_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_register_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_register_proto_rawDescGZIP(), []int{1}
}

func (x *Response) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *Response) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_register_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_register_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_register_proto_rawDescGZIP(), []int{2}
}

func (x *User) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *User) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type HostAddress struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HostIp   string `protobuf:"bytes,1,opt,name=hostIp,proto3" json:"hostIp,omitempty"`
	HostPort string `protobuf:"bytes,2,opt,name=hostPort,proto3" json:"hostPort,omitempty"`
}

func (x *HostAddress) Reset() {
	*x = HostAddress{}
	if protoimpl.UnsafeEnabled {
		mi := &file_register_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HostAddress) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HostAddress) ProtoMessage() {}

func (x *HostAddress) ProtoReflect() protoreflect.Message {
	mi := &file_register_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HostAddress.ProtoReflect.Descriptor instead.
func (*HostAddress) Descriptor() ([]byte, []int) {
	return file_register_proto_rawDescGZIP(), []int{3}
}

func (x *HostAddress) GetHostIp() string {
	if x != nil {
		return x.HostIp
	}
	return ""
}

func (x *HostAddress) GetHostPort() string {
	if x != nil {
		return x.HostPort
	}
	return ""
}

var File_register_proto protoreflect.FileDescriptor

var file_register_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x78, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x22, 0x0a, 0x07, 0x63,
	0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x08, 0x2e, 0x43,
	0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12,
	0x19, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x05, 0x2e,
	0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x2e, 0x0a, 0x0b, 0x68, 0x6f,
	0x73, 0x74, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0c, 0x2e, 0x48, 0x6f, 0x73, 0x74, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x0b, 0x68,
	0x6f, 0x73, 0x74, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0x34, 0x0a, 0x08, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x10,
	0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67,
	0x22, 0x3e, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x22, 0x41, 0x0a, 0x0b, 0x48, 0x6f, 0x73, 0x74, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12,
	0x16, 0x0a, 0x06, 0x68, 0x6f, 0x73, 0x74, 0x49, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x68, 0x6f, 0x73, 0x74, 0x49, 0x70, 0x12, 0x1a, 0x0a, 0x08, 0x68, 0x6f, 0x73, 0x74, 0x50,
	0x6f, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x68, 0x6f, 0x73, 0x74, 0x50,
	0x6f, 0x72, 0x74, 0x2a, 0x53, 0x0a, 0x07, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12, 0x09,
	0x0a, 0x05, 0x55, 0x6e, 0x6b, 0x77, 0x6e, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x50, 0x69, 0x6e,
	0x67, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x50, 0x6f, 0x6e, 0x67, 0x10, 0x02, 0x12, 0x07, 0x0a,
	0x03, 0x52, 0x65, 0x67, 0x10, 0x03, 0x12, 0x07, 0x0a, 0x03, 0x4c, 0x6f, 0x67, 0x10, 0x04, 0x12,
	0x0b, 0x0a, 0x07, 0x50, 0x72, 0x69, 0x76, 0x4d, 0x73, 0x67, 0x10, 0x05, 0x12, 0x0a, 0x0a, 0x06,
	0x47, 0x72, 0x70, 0x4d, 0x73, 0x67, 0x10, 0x06, 0x42, 0x14, 0x5a, 0x12, 0x69, 0x72, 0x63, 0x5f,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_register_proto_rawDescOnce sync.Once
	file_register_proto_rawDescData = file_register_proto_rawDesc
)

func file_register_proto_rawDescGZIP() []byte {
	file_register_proto_rawDescOnce.Do(func() {
		file_register_proto_rawDescData = protoimpl.X.CompressGZIP(file_register_proto_rawDescData)
	})
	return file_register_proto_rawDescData
}

var file_register_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_register_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_register_proto_goTypes = []interface{}{
	(Command)(0),        // 0: Command
	(*Message)(nil),     // 1: Message
	(*Response)(nil),    // 2: Response
	(*User)(nil),        // 3: User
	(*HostAddress)(nil), // 4: HostAddress
}
var file_register_proto_depIdxs = []int32{
	0, // 0: Message.command:type_name -> Command
	3, // 1: Message.user:type_name -> User
	4, // 2: Message.hostAddress:type_name -> HostAddress
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_register_proto_init() }
func file_register_proto_init() {
	if File_register_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_register_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Message); i {
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
		file_register_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
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
		file_register_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
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
		file_register_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HostAddress); i {
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
			RawDescriptor: file_register_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_register_proto_goTypes,
		DependencyIndexes: file_register_proto_depIdxs,
		EnumInfos:         file_register_proto_enumTypes,
		MessageInfos:      file_register_proto_msgTypes,
	}.Build()
	File_register_proto = out.File
	file_register_proto_rawDesc = nil
	file_register_proto_goTypes = nil
	file_register_proto_depIdxs = nil
}