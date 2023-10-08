// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: duat.proto

package pb

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

type DocumentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DocumentTypeId uint64  `protobuf:"varint,1,opt,name=document_type_id,json=documentTypeId,proto3" json:"document_type_id,omitempty"`
	DepartmentId   uint64  `protobuf:"varint,2,opt,name=department_id,json=departmentId,proto3" json:"department_id,omitempty"`
	Code           string  `protobuf:"bytes,3,opt,name=code,proto3" json:"code,omitempty"`
	Identity       string  `protobuf:"bytes,4,opt,name=identity,proto3" json:"identity,omitempty"`
	Name           string  `protobuf:"bytes,5,opt,name=name,proto3" json:"name,omitempty"`
	Comment        *string `protobuf:"bytes,6,opt,name=comment,proto3,oneof" json:"comment,omitempty"`
	Storage        *string `protobuf:"bytes,7,opt,name=storage,proto3,oneof" json:"storage,omitempty"`
	DateDocument   string  `protobuf:"bytes,8,opt,name=date_document,json=dateDocument,proto3" json:"date_document,omitempty"`
}

func (x *DocumentRequest) Reset() {
	*x = DocumentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_duat_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DocumentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DocumentRequest) ProtoMessage() {}

func (x *DocumentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_duat_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DocumentRequest.ProtoReflect.Descriptor instead.
func (*DocumentRequest) Descriptor() ([]byte, []int) {
	return file_duat_proto_rawDescGZIP(), []int{0}
}

func (x *DocumentRequest) GetDocumentTypeId() uint64 {
	if x != nil {
		return x.DocumentTypeId
	}
	return 0
}

func (x *DocumentRequest) GetDepartmentId() uint64 {
	if x != nil {
		return x.DepartmentId
	}
	return 0
}

func (x *DocumentRequest) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *DocumentRequest) GetIdentity() string {
	if x != nil {
		return x.Identity
	}
	return ""
}

func (x *DocumentRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *DocumentRequest) GetComment() string {
	if x != nil && x.Comment != nil {
		return *x.Comment
	}
	return ""
}

func (x *DocumentRequest) GetStorage() string {
	if x != nil && x.Storage != nil {
		return *x.Storage
	}
	return ""
}

func (x *DocumentRequest) GetDateDocument() string {
	if x != nil {
		return x.DateDocument
	}
	return ""
}

type DocumentImageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DocumentId  string      `protobuf:"bytes,1,opt,name=document_id,json=documentId,proto3" json:"document_id,omitempty"`
	Page        uint32      `protobuf:"varint,2,opt,name=page,proto3" json:"page,omitempty"`
	ImageExt    string      `protobuf:"bytes,3,opt,name=image_ext,json=imageExt,proto3" json:"image_ext,omitempty"`
	StorageType StorageType `protobuf:"varint,4,opt,name=storage_type,json=storageType,proto3,enum=duat.StorageType" json:"storage_type,omitempty"`
	Data        []byte      `protobuf:"bytes,5,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *DocumentImageRequest) Reset() {
	*x = DocumentImageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_duat_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DocumentImageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DocumentImageRequest) ProtoMessage() {}

func (x *DocumentImageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_duat_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DocumentImageRequest.ProtoReflect.Descriptor instead.
func (*DocumentImageRequest) Descriptor() ([]byte, []int) {
	return file_duat_proto_rawDescGZIP(), []int{1}
}

func (x *DocumentImageRequest) GetDocumentId() string {
	if x != nil {
		return x.DocumentId
	}
	return ""
}

func (x *DocumentImageRequest) GetPage() uint32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *DocumentImageRequest) GetImageExt() string {
	if x != nil {
		return x.ImageExt
	}
	return ""
}

func (x *DocumentImageRequest) GetStorageType() StorageType {
	if x != nil {
		return x.StorageType
	}
	return StorageType_LOCAL
}

func (x *DocumentImageRequest) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_duat_proto protoreflect.FileDescriptor

var file_duat_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x64, 0x75, 0x61, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x64, 0x75,
	0x61, 0x74, 0x1a, 0x0e, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x14, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x6d, 0x61,
	0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9f, 0x02, 0x0a, 0x0f, 0x44, 0x6f, 0x63,
	0x75, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x28, 0x0a, 0x10,
	0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0e, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74,
	0x54, 0x79, 0x70, 0x65, 0x49, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x64, 0x65, 0x70, 0x61, 0x72, 0x74,
	0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0c, 0x64,
	0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x63,
	0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x1d, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x48, 0x00, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x88, 0x01, 0x01, 0x12, 0x1d,
	0x0a, 0x07, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x48,
	0x01, 0x52, 0x07, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x88, 0x01, 0x01, 0x12, 0x23, 0x0a,
	0x0d, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x64, 0x61, 0x74, 0x65, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65,
	0x6e, 0x74, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x42, 0x0a,
	0x0a, 0x08, 0x5f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x22, 0xb2, 0x01, 0x0a, 0x14, 0x44,
	0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65,
	0x6e, 0x74, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x6d, 0x61, 0x67,
	0x65, 0x5f, 0x65, 0x78, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x69, 0x6d, 0x61,
	0x67, 0x65, 0x45, 0x78, 0x74, 0x12, 0x34, 0x0a, 0x0c, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x11, 0x2e, 0x64, 0x75,
	0x61, 0x74, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0b,
	0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x32,
	0x46, 0x0a, 0x0f, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x33, 0x0a, 0x05, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x12, 0x15, 0x2e, 0x64, 0x75,
	0x61, 0x74, 0x2e, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x13, 0x2e, 0x64, 0x75, 0x61, 0x74, 0x2e, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65,
	0x6e, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x32, 0x57, 0x0a, 0x14, 0x44, 0x6f, 0x63, 0x75, 0x6d,
	0x65, 0x6e, 0x74, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x3f, 0x0a, 0x05, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x12, 0x1a, 0x2e, 0x64, 0x75, 0x61, 0x74, 0x2e,
	0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x64, 0x75, 0x61, 0x74, 0x2e, 0x44, 0x6f, 0x63, 0x75,
	0x6d, 0x65, 0x6e, 0x74, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x28, 0x01,
	0x42, 0x32, 0x0a, 0x1f, 0x62, 0x72, 0x2e, 0x6f, 0x70, 0x74, 0x69, 0x6d, 0x75, 0x73, 0x2e, 0x6f,
	0x73, 0x61, 0x73, 0x63, 0x6f, 0x2e, 0x64, 0x75, 0x61, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x50, 0x01, 0x5a, 0x07, 0x64, 0x75, 0x61, 0x74, 0x2f, 0x70, 0x62, 0xa2, 0x02,
	0x03, 0x48, 0x4c, 0x57, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_duat_proto_rawDescOnce sync.Once
	file_duat_proto_rawDescData = file_duat_proto_rawDesc
)

func file_duat_proto_rawDescGZIP() []byte {
	file_duat_proto_rawDescOnce.Do(func() {
		file_duat_proto_rawDescData = protoimpl.X.CompressGZIP(file_duat_proto_rawDescData)
	})
	return file_duat_proto_rawDescData
}

var file_duat_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_duat_proto_goTypes = []interface{}{
	(*DocumentRequest)(nil),      // 0: duat.DocumentRequest
	(*DocumentImageRequest)(nil), // 1: duat.DocumentImageRequest
	(StorageType)(0),             // 2: duat.StorageType
	(*DocumentReply)(nil),        // 3: duat.DocumentReply
	(*DocumentImageReply)(nil),   // 4: duat.DocumentImageReply
}
var file_duat_proto_depIdxs = []int32{
	2, // 0: duat.DocumentImageRequest.storage_type:type_name -> duat.StorageType
	0, // 1: duat.DocumentService.Store:input_type -> duat.DocumentRequest
	1, // 2: duat.DocumentImageService.Store:input_type -> duat.DocumentImageRequest
	3, // 3: duat.DocumentService.Store:output_type -> duat.DocumentReply
	4, // 4: duat.DocumentImageService.Store:output_type -> duat.DocumentImageReply
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_duat_proto_init() }
func file_duat_proto_init() {
	if File_duat_proto != nil {
		return
	}
	file_document_proto_init()
	file_document_image_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_duat_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DocumentRequest); i {
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
		file_duat_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DocumentImageRequest); i {
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
	file_duat_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_duat_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_duat_proto_goTypes,
		DependencyIndexes: file_duat_proto_depIdxs,
		MessageInfos:      file_duat_proto_msgTypes,
	}.Build()
	File_duat_proto = out.File
	file_duat_proto_rawDesc = nil
	file_duat_proto_goTypes = nil
	file_duat_proto_depIdxs = nil
}