// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.0
// source: proto/command.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CommandList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Commands []*Command `protobuf:"bytes,1,rep,name=commands,proto3" json:"commands,omitempty"`
}

func (x *CommandList) Reset() {
	*x = CommandList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_command_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommandList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommandList) ProtoMessage() {}

func (x *CommandList) ProtoReflect() protoreflect.Message {
	mi := &file_proto_command_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommandList.ProtoReflect.Descriptor instead.
func (*CommandList) Descriptor() ([]byte, []int) {
	return file_proto_command_proto_rawDescGZIP(), []int{0}
}

func (x *CommandList) GetCommands() []*Command {
	if x != nil {
		return x.Commands
	}
	return nil
}

type Command struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id              string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Text            string                 `protobuf:"bytes,2,opt,name=text,proto3" json:"text,omitempty"`
	ExecutionStatus int32                  `protobuf:"varint,3,opt,name=execution_status,json=executionStatus,proto3" json:"execution_status,omitempty"`
	ExecutionTime   *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=execution_time,json=executionTime,proto3" json:"execution_time,omitempty"`
	Pid             int32                  `protobuf:"varint,5,opt,name=pid,proto3" json:"pid,omitempty"`
	Ppid            int32                  `protobuf:"varint,6,opt,name=ppid,proto3" json:"ppid,omitempty"`
	Env             *Environment           `protobuf:"bytes,7,opt,name=env,proto3" json:"env,omitempty"`
}

func (x *Command) Reset() {
	*x = Command{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_command_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Command) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Command) ProtoMessage() {}

func (x *Command) ProtoReflect() protoreflect.Message {
	mi := &file_proto_command_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Command.ProtoReflect.Descriptor instead.
func (*Command) Descriptor() ([]byte, []int) {
	return file_proto_command_proto_rawDescGZIP(), []int{1}
}

func (x *Command) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Command) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *Command) GetExecutionStatus() int32 {
	if x != nil {
		return x.ExecutionStatus
	}
	return 0
}

func (x *Command) GetExecutionTime() *timestamppb.Timestamp {
	if x != nil {
		return x.ExecutionTime
	}
	return nil
}

func (x *Command) GetPid() int32 {
	if x != nil {
		return x.Pid
	}
	return 0
}

func (x *Command) GetPpid() int32 {
	if x != nil {
		return x.Ppid
	}
	return 0
}

func (x *Command) GetEnv() *Environment {
	if x != nil {
		return x.Env
	}
	return nil
}

type Environment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User             *User  `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	WorkingDirectory string `protobuf:"bytes,2,opt,name=working_directory,json=workingDirectory,proto3" json:"working_directory,omitempty"`
	Debug            bool   `protobuf:"varint,3,opt,name=debug,proto3" json:"debug,omitempty"`
}

func (x *Environment) Reset() {
	*x = Environment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_command_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Environment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Environment) ProtoMessage() {}

func (x *Environment) ProtoReflect() protoreflect.Message {
	mi := &file_proto_command_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Environment.ProtoReflect.Descriptor instead.
func (*Environment) Descriptor() ([]byte, []int) {
	return file_proto_command_proto_rawDescGZIP(), []int{2}
}

func (x *Environment) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *Environment) GetWorkingDirectory() string {
	if x != nil {
		return x.WorkingDirectory
	}
	return ""
}

func (x *Environment) GetDebug() bool {
	if x != nil {
		return x.Debug
	}
	return false
}

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Name     string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_command_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_proto_command_proto_msgTypes[3]
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
	return file_proto_command_proto_rawDescGZIP(), []int{3}
}

func (x *User) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *User) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_proto_command_proto protoreflect.FileDescriptor

var file_proto_command_proto_rawDesc = []byte{
	0x0a, 0x13, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1a, 0x73, 0x68, 0x65, 0x6c, 0x6c, 0x5f, 0x68, 0x69, 0x73,
	0x74, 0x6f, 0x72, 0x79, 0x5f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x4e, 0x0a, 0x0b, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x4c, 0x69, 0x73,
	0x74, 0x12, 0x3f, 0x0a, 0x08, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x73, 0x68, 0x65, 0x6c, 0x6c, 0x5f, 0x68, 0x69, 0x73, 0x74,
	0x6f, 0x72, 0x79, 0x5f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x52, 0x08, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e,
	0x64, 0x73, 0x22, 0xfc, 0x01, 0x0a, 0x07, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65,
	0x78, 0x74, 0x12, 0x29, 0x0a, 0x10, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0f, 0x65, 0x78,
	0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x41, 0x0a,
	0x0e, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x0d, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65,
	0x12, 0x10, 0x0a, 0x03, 0x70, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x70,
	0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x70, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x04, 0x70, 0x70, 0x69, 0x64, 0x12, 0x39, 0x0a, 0x03, 0x65, 0x6e, 0x76, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x73, 0x68, 0x65, 0x6c, 0x6c, 0x5f, 0x68, 0x69, 0x73, 0x74,
	0x6f, 0x72, 0x79, 0x5f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x45, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x03, 0x65, 0x6e,
	0x76, 0x22, 0x86, 0x01, 0x0a, 0x0b, 0x45, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e,
	0x74, 0x12, 0x34, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x20, 0x2e, 0x73, 0x68, 0x65, 0x6c, 0x6c, 0x5f, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x5f,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x73, 0x65,
	0x72, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x2b, 0x0a, 0x11, 0x77, 0x6f, 0x72, 0x6b, 0x69,
	0x6e, 0x67, 0x5f, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x10, 0x77, 0x6f, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x44, 0x69, 0x72, 0x65, 0x63,
	0x74, 0x6f, 0x72, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x64, 0x65, 0x62, 0x75, 0x67, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x05, 0x64, 0x65, 0x62, 0x75, 0x67, 0x22, 0x36, 0x0a, 0x04, 0x55, 0x73,
	0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x42, 0x1c, 0x5a, 0x1a, 0x73, 0x68, 0x65, 0x6c, 0x6c, 0x5f, 0x68, 0x69, 0x73, 0x74,
	0x6f, 0x72, 0x79, 0x5f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_command_proto_rawDescOnce sync.Once
	file_proto_command_proto_rawDescData = file_proto_command_proto_rawDesc
)

func file_proto_command_proto_rawDescGZIP() []byte {
	file_proto_command_proto_rawDescOnce.Do(func() {
		file_proto_command_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_command_proto_rawDescData)
	})
	return file_proto_command_proto_rawDescData
}

var file_proto_command_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_proto_command_proto_goTypes = []any{
	(*CommandList)(nil),           // 0: shell_history_client.proto.CommandList
	(*Command)(nil),               // 1: shell_history_client.proto.Command
	(*Environment)(nil),           // 2: shell_history_client.proto.Environment
	(*User)(nil),                  // 3: shell_history_client.proto.User
	(*timestamppb.Timestamp)(nil), // 4: google.protobuf.Timestamp
}
var file_proto_command_proto_depIdxs = []int32{
	1, // 0: shell_history_client.proto.CommandList.commands:type_name -> shell_history_client.proto.Command
	4, // 1: shell_history_client.proto.Command.execution_time:type_name -> google.protobuf.Timestamp
	2, // 2: shell_history_client.proto.Command.env:type_name -> shell_history_client.proto.Environment
	3, // 3: shell_history_client.proto.Environment.user:type_name -> shell_history_client.proto.User
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_proto_command_proto_init() }
func file_proto_command_proto_init() {
	if File_proto_command_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_command_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*CommandList); i {
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
		file_proto_command_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*Command); i {
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
		file_proto_command_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*Environment); i {
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
		file_proto_command_proto_msgTypes[3].Exporter = func(v any, i int) any {
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
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_command_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_command_proto_goTypes,
		DependencyIndexes: file_proto_command_proto_depIdxs,
		MessageInfos:      file_proto_command_proto_msgTypes,
	}.Build()
	File_proto_command_proto = out.File
	file_proto_command_proto_rawDesc = nil
	file_proto_command_proto_goTypes = nil
	file_proto_command_proto_depIdxs = nil
}