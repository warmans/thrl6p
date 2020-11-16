// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.5.1
// source: proto/patch.proto

package v1

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type CreatePatchRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Patch       string `protobuf:"bytes,1,opt,name=patch,proto3" json:"patch,omitempty"`
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *CreatePatchRequest) Reset() {
	*x = CreatePatchRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_patch_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePatchRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePatchRequest) ProtoMessage() {}

func (x *CreatePatchRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_patch_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePatchRequest.ProtoReflect.Descriptor instead.
func (*CreatePatchRequest) Descriptor() ([]byte, []int) {
	return file_proto_patch_proto_rawDescGZIP(), []int{0}
}

func (x *CreatePatchRequest) GetPatch() string {
	if x != nil {
		return x.Patch
	}
	return ""
}

func (x *CreatePatchRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type Patch struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Patch       string `protobuf:"bytes,2,opt,name=patch,proto3" json:"patch,omitempty"`
	Name        string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Description string `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	Permalink   string `protobuf:"bytes,5,opt,name=permalink,proto3" json:"permalink,omitempty"`
}

func (x *Patch) Reset() {
	*x = Patch{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_patch_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Patch) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Patch) ProtoMessage() {}

func (x *Patch) ProtoReflect() protoreflect.Message {
	mi := &file_proto_patch_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Patch.ProtoReflect.Descriptor instead.
func (*Patch) Descriptor() ([]byte, []int) {
	return file_proto_patch_proto_rawDescGZIP(), []int{1}
}

func (x *Patch) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Patch) GetPatch() string {
	if x != nil {
		return x.Patch
	}
	return ""
}

func (x *Patch) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Patch) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Patch) GetPermalink() string {
	if x != nil {
		return x.Permalink
	}
	return ""
}

type GetPatchRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetPatchRequest) Reset() {
	*x = GetPatchRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_patch_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPatchRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPatchRequest) ProtoMessage() {}

func (x *GetPatchRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_patch_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPatchRequest.ProtoReflect.Descriptor instead.
func (*GetPatchRequest) Descriptor() ([]byte, []int) {
	return file_proto_patch_proto_rawDescGZIP(), []int{2}
}

func (x *GetPatchRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type ValidateNameRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *ValidateNameRequest) Reset() {
	*x = ValidateNameRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_patch_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValidateNameRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidateNameRequest) ProtoMessage() {}

func (x *ValidateNameRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_patch_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValidateNameRequest.ProtoReflect.Descriptor instead.
func (*ValidateNameRequest) Descriptor() ([]byte, []int) {
	return file_proto_patch_proto_rawDescGZIP(), []int{3}
}

func (x *ValidateNameRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_proto_patch_proto protoreflect.FileDescriptor

var file_proto_patch_proto_rawDesc = []byte{
	0x0a, 0x11, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x61, 0x74, 0x63, 0x68, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x06, 0x74, 0x68, 0x72, 0x6c, 0x36, 0x70, 0x1a, 0x1c, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2c, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x73, 0x77, 0x61, 0x67, 0x67, 0x65, 0x72, 0x2f, 0x6f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4c, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x61,
	0x74, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x61,
	0x74, 0x63, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x61, 0x74, 0x63, 0x68,
	0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x22, 0x81, 0x01, 0x0a, 0x05, 0x50, 0x61, 0x74, 0x63, 0x68, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05,
	0x70, 0x61, 0x74, 0x63, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x61, 0x74,
	0x63, 0x68, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x65, 0x72, 0x6d,
	0x61, 0x6c, 0x69, 0x6e, 0x6b, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x65, 0x72,
	0x6d, 0x61, 0x6c, 0x69, 0x6e, 0x6b, 0x22, 0x21, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x50, 0x61, 0x74,
	0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x29, 0x0a, 0x13, 0x56, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x32, 0xf6, 0x02, 0x0a, 0x0c, 0x50, 0x61, 0x74, 0x63, 0x68, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x68, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50,
	0x61, 0x74, 0x63, 0x68, 0x12, 0x1a, 0x2e, 0x74, 0x68, 0x72, 0x6c, 0x36, 0x70, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x50, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x0d, 0x2e, 0x74, 0x68, 0x72, 0x6c, 0x36, 0x70, 0x2e, 0x50, 0x61, 0x74, 0x63, 0x68, 0x22,
	0x2e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0f, 0x22, 0x0a, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x61,
	0x74, 0x63, 0x68, 0x3a, 0x01, 0x2a, 0x92, 0x41, 0x16, 0x1a, 0x14, 0x55, 0x70, 0x6c, 0x6f, 0x61,
	0x64, 0x73, 0x20, 0x61, 0x20, 0x6e, 0x65, 0x77, 0x20, 0x70, 0x61, 0x74, 0x63, 0x68, 0x2e, 0x12,
	0x83, 0x01, 0x0a, 0x0c, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x1b, 0x2e, 0x74, 0x68, 0x72, 0x6c, 0x36, 0x70, 0x2e, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x3e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1d, 0x22, 0x18, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x70, 0x61, 0x74, 0x63, 0x68, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x65, 0x2f, 0x6e, 0x61, 0x6d, 0x65, 0x3a, 0x01, 0x2a, 0x92, 0x41, 0x18, 0x1a, 0x16, 0x56,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x20, 0x61, 0x20, 0x70, 0x61, 0x74, 0x63, 0x68, 0x20,
	0x6e, 0x61, 0x6d, 0x65, 0x2e, 0x12, 0x76, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x50, 0x61, 0x74, 0x63,
	0x68, 0x12, 0x17, 0x2e, 0x74, 0x68, 0x72, 0x6c, 0x36, 0x70, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x61,
	0x74, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0d, 0x2e, 0x74, 0x68, 0x72,
	0x6c, 0x36, 0x70, 0x2e, 0x50, 0x61, 0x74, 0x63, 0x68, 0x22, 0x42, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x11, 0x12, 0x0f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x61, 0x74, 0x63, 0x68, 0x2f, 0x7b, 0x69,
	0x64, 0x7d, 0x92, 0x41, 0x28, 0x1a, 0x26, 0x47, 0x65, 0x74, 0x73, 0x20, 0x61, 0x20, 0x70, 0x61,
	0x74, 0x63, 0x68, 0x20, 0x77, 0x69, 0x74, 0x68, 0x20, 0x61, 0x73, 0x73, 0x6f, 0x63, 0x69, 0x61,
	0x74, 0x65, 0x64, 0x20, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x42, 0x08, 0x5a,
	0x06, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_patch_proto_rawDescOnce sync.Once
	file_proto_patch_proto_rawDescData = file_proto_patch_proto_rawDesc
)

func file_proto_patch_proto_rawDescGZIP() []byte {
	file_proto_patch_proto_rawDescOnce.Do(func() {
		file_proto_patch_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_patch_proto_rawDescData)
	})
	return file_proto_patch_proto_rawDescData
}

var file_proto_patch_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_proto_patch_proto_goTypes = []interface{}{
	(*CreatePatchRequest)(nil),  // 0: thrl6p.CreatePatchRequest
	(*Patch)(nil),               // 1: thrl6p.Patch
	(*GetPatchRequest)(nil),     // 2: thrl6p.GetPatchRequest
	(*ValidateNameRequest)(nil), // 3: thrl6p.ValidateNameRequest
	(*empty.Empty)(nil),         // 4: google.protobuf.Empty
}
var file_proto_patch_proto_depIdxs = []int32{
	0, // 0: thrl6p.PatchService.CreatePatch:input_type -> thrl6p.CreatePatchRequest
	3, // 1: thrl6p.PatchService.ValidateName:input_type -> thrl6p.ValidateNameRequest
	2, // 2: thrl6p.PatchService.GetPatch:input_type -> thrl6p.GetPatchRequest
	1, // 3: thrl6p.PatchService.CreatePatch:output_type -> thrl6p.Patch
	4, // 4: thrl6p.PatchService.ValidateName:output_type -> google.protobuf.Empty
	1, // 5: thrl6p.PatchService.GetPatch:output_type -> thrl6p.Patch
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_patch_proto_init() }
func file_proto_patch_proto_init() {
	if File_proto_patch_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_patch_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatePatchRequest); i {
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
		file_proto_patch_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Patch); i {
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
		file_proto_patch_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPatchRequest); i {
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
		file_proto_patch_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ValidateNameRequest); i {
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
			RawDescriptor: file_proto_patch_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_patch_proto_goTypes,
		DependencyIndexes: file_proto_patch_proto_depIdxs,
		MessageInfos:      file_proto_patch_proto_msgTypes,
	}.Build()
	File_proto_patch_proto = out.File
	file_proto_patch_proto_rawDesc = nil
	file_proto_patch_proto_goTypes = nil
	file_proto_patch_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// PatchServiceClient is the client API for PatchService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PatchServiceClient interface {
	CreatePatch(ctx context.Context, in *CreatePatchRequest, opts ...grpc.CallOption) (*Patch, error)
	ValidateName(ctx context.Context, in *ValidateNameRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	GetPatch(ctx context.Context, in *GetPatchRequest, opts ...grpc.CallOption) (*Patch, error)
}

type patchServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPatchServiceClient(cc grpc.ClientConnInterface) PatchServiceClient {
	return &patchServiceClient{cc}
}

func (c *patchServiceClient) CreatePatch(ctx context.Context, in *CreatePatchRequest, opts ...grpc.CallOption) (*Patch, error) {
	out := new(Patch)
	err := c.cc.Invoke(ctx, "/thrl6p.PatchService/CreatePatch", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *patchServiceClient) ValidateName(ctx context.Context, in *ValidateNameRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/thrl6p.PatchService/ValidateName", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *patchServiceClient) GetPatch(ctx context.Context, in *GetPatchRequest, opts ...grpc.CallOption) (*Patch, error) {
	out := new(Patch)
	err := c.cc.Invoke(ctx, "/thrl6p.PatchService/GetPatch", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PatchServiceServer is the server API for PatchService service.
type PatchServiceServer interface {
	CreatePatch(context.Context, *CreatePatchRequest) (*Patch, error)
	ValidateName(context.Context, *ValidateNameRequest) (*empty.Empty, error)
	GetPatch(context.Context, *GetPatchRequest) (*Patch, error)
}

// UnimplementedPatchServiceServer can be embedded to have forward compatible implementations.
type UnimplementedPatchServiceServer struct {
}

func (*UnimplementedPatchServiceServer) CreatePatch(context.Context, *CreatePatchRequest) (*Patch, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePatch not implemented")
}
func (*UnimplementedPatchServiceServer) ValidateName(context.Context, *ValidateNameRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ValidateName not implemented")
}
func (*UnimplementedPatchServiceServer) GetPatch(context.Context, *GetPatchRequest) (*Patch, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPatch not implemented")
}

func RegisterPatchServiceServer(s *grpc.Server, srv PatchServiceServer) {
	s.RegisterService(&_PatchService_serviceDesc, srv)
}

func _PatchService_CreatePatch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePatchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PatchServiceServer).CreatePatch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/thrl6p.PatchService/CreatePatch",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PatchServiceServer).CreatePatch(ctx, req.(*CreatePatchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PatchService_ValidateName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ValidateNameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PatchServiceServer).ValidateName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/thrl6p.PatchService/ValidateName",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PatchServiceServer).ValidateName(ctx, req.(*ValidateNameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PatchService_GetPatch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPatchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PatchServiceServer).GetPatch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/thrl6p.PatchService/GetPatch",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PatchServiceServer).GetPatch(ctx, req.(*GetPatchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _PatchService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "thrl6p.PatchService",
	HandlerType: (*PatchServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreatePatch",
			Handler:    _PatchService_CreatePatch_Handler,
		},
		{
			MethodName: "ValidateName",
			Handler:    _PatchService_ValidateName_Handler,
		},
		{
			MethodName: "GetPatch",
			Handler:    _PatchService_GetPatch_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/patch.proto",
}
