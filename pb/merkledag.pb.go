// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.0
// source: pb/merkledag.proto

package merkledag_pb

import (
	_ "github.com/gogo/protobuf/gogoproto"
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

// An IPFS MerkleDAG Link
type PBLink struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// multihash of the target object
	Hash []byte `protobuf:"bytes,1,opt,name=Hash" json:"Hash,omitempty"`
	// utf string name. should be unique per object
	Name *string `protobuf:"bytes,2,opt,name=Name" json:"Name,omitempty"`
	// cumulative size of target object
	Tsize *uint64 `protobuf:"varint,3,opt,name=Tsize" json:"Tsize,omitempty"`
}

func (x *PBLink) Reset() {
	*x = PBLink{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_merkledag_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PBLink) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PBLink) ProtoMessage() {}

func (x *PBLink) ProtoReflect() protoreflect.Message {
	mi := &file_pb_merkledag_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PBLink.ProtoReflect.Descriptor instead.
func (*PBLink) Descriptor() ([]byte, []int) {
	return file_pb_merkledag_proto_rawDescGZIP(), []int{0}
}

func (x *PBLink) GetHash() []byte {
	if x != nil {
		return x.Hash
	}
	return nil
}

func (x *PBLink) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *PBLink) GetTsize() uint64 {
	if x != nil && x.Tsize != nil {
		return *x.Tsize
	}
	return 0
}

// An IPFS MerkleDAG Node
type PBNode struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// refs to other objects
	Links []*PBLink `protobuf:"bytes,2,rep,name=Links" json:"Links,omitempty"`
	// opaque user data
	Data []byte `protobuf:"bytes,1,opt,name=Data" json:"Data,omitempty"`
}

func (x *PBNode) Reset() {
	*x = PBNode{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_merkledag_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PBNode) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PBNode) ProtoMessage() {}

func (x *PBNode) ProtoReflect() protoreflect.Message {
	mi := &file_pb_merkledag_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PBNode.ProtoReflect.Descriptor instead.
func (*PBNode) Descriptor() ([]byte, []int) {
	return file_pb_merkledag_proto_rawDescGZIP(), []int{1}
}

func (x *PBNode) GetLinks() []*PBLink {
	if x != nil {
		return x.Links
	}
	return nil
}

func (x *PBNode) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_pb_merkledag_proto protoreflect.FileDescriptor

var file_pb_merkledag_proto_rawDesc = []byte{
	0x0a, 0x12, 0x70, 0x62, 0x2f, 0x6d, 0x65, 0x72, 0x6b, 0x6c, 0x65, 0x64, 0x61, 0x67, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x6d, 0x65, 0x72, 0x6b, 0x6c, 0x65, 0x64, 0x61, 0x67, 0x2e,
	0x70, 0x62, 0x1a, 0x2d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67,
	0x6f, 0x67, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x67, 0x6f, 0x67,
	0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x67, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x46, 0x0a, 0x06, 0x50, 0x42, 0x4c, 0x69, 0x6e, 0x6b, 0x12, 0x12, 0x0a, 0x04, 0x48,
	0x61, 0x73, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x48, 0x61, 0x73, 0x68, 0x12,
	0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x05, 0x54, 0x73, 0x69, 0x7a, 0x65, 0x22, 0x48, 0x0a, 0x06, 0x50, 0x42, 0x4e,
	0x6f, 0x64, 0x65, 0x12, 0x2a, 0x0a, 0x05, 0x4c, 0x69, 0x6e, 0x6b, 0x73, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x14, 0x2e, 0x6d, 0x65, 0x72, 0x6b, 0x6c, 0x65, 0x64, 0x61, 0x67, 0x2e, 0x70,
	0x62, 0x2e, 0x50, 0x42, 0x4c, 0x69, 0x6e, 0x6b, 0x52, 0x05, 0x4c, 0x69, 0x6e, 0x6b, 0x73, 0x12,
	0x12, 0x0a, 0x04, 0x44, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x44,
	0x61, 0x74, 0x61, 0x42, 0x3f, 0x5a, 0x11, 0x2e, 0x2f, 0x70, 0x62, 0x3b, 0x6d, 0x65, 0x72, 0x6b,
	0x6c, 0x65, 0x64, 0x61, 0x67, 0x5f, 0x70, 0x62, 0xf0, 0xe1, 0x1e, 0x01, 0xa8, 0xe2, 0x1e, 0x01,
	0xe0, 0xe1, 0x1e, 0x01, 0xd8, 0xe1, 0x1e, 0x00, 0x80, 0xe2, 0x1e, 0x01, 0xf8, 0xe1, 0x1e, 0x01,
	0xb8, 0xe2, 0x1e, 0x01, 0xc0, 0xe2, 0x1e, 0x01, 0xc8, 0xe2, 0x1e, 0x01, 0xe0, 0xe2, 0x1e, 0x01,
	0xd0, 0xe2, 0x1e, 0x01,
}

var (
	file_pb_merkledag_proto_rawDescOnce sync.Once
	file_pb_merkledag_proto_rawDescData = file_pb_merkledag_proto_rawDesc
)

func file_pb_merkledag_proto_rawDescGZIP() []byte {
	file_pb_merkledag_proto_rawDescOnce.Do(func() {
		file_pb_merkledag_proto_rawDescData = protoimpl.X.CompressGZIP(file_pb_merkledag_proto_rawDescData)
	})
	return file_pb_merkledag_proto_rawDescData
}

var file_pb_merkledag_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_pb_merkledag_proto_goTypes = []interface{}{
	(*PBLink)(nil), // 0: merkledag.pb.PBLink
	(*PBNode)(nil), // 1: merkledag.pb.PBNode
}
var file_pb_merkledag_proto_depIdxs = []int32{
	0, // 0: merkledag.pb.PBNode.Links:type_name -> merkledag.pb.PBLink
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_pb_merkledag_proto_init() }
func file_pb_merkledag_proto_init() {
	if File_pb_merkledag_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pb_merkledag_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PBLink); i {
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
		file_pb_merkledag_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PBNode); i {
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
			RawDescriptor: file_pb_merkledag_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pb_merkledag_proto_goTypes,
		DependencyIndexes: file_pb_merkledag_proto_depIdxs,
		MessageInfos:      file_pb_merkledag_proto_msgTypes,
	}.Build()
	File_pb_merkledag_proto = out.File
	file_pb_merkledag_proto_rawDesc = nil
	file_pb_merkledag_proto_goTypes = nil
	file_pb_merkledag_proto_depIdxs = nil
}
