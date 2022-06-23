// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/museum/exhibit.proto

package omo_msp_museum

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type ExhibitInfo struct {
	Id                   uint64         `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Uid                  string         `protobuf:"bytes,2,opt,name=uid,proto3" json:"uid,omitempty"`
	Name                 string         `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Created              uint64         `protobuf:"varint,4,opt,name=created,proto3" json:"created,omitempty"`
	Updated              uint64         `protobuf:"varint,5,opt,name=updated,proto3" json:"updated,omitempty"`
	Creator              string         `protobuf:"bytes,6,opt,name=creator,proto3" json:"creator,omitempty"`
	Operator             string         `protobuf:"bytes,7,opt,name=operator,proto3" json:"operator,omitempty"`
	Sn                   string         `protobuf:"bytes,8,opt,name=sn,proto3" json:"sn,omitempty"`
	Owner                string         `protobuf:"bytes,9,opt,name=owner,proto3" json:"owner,omitempty"`
	Entity               string         `protobuf:"bytes,10,opt,name=entity,proto3" json:"entity,omitempty"`
	Size                 *Vector3       `protobuf:"bytes,11,opt,name=size,proto3" json:"size,omitempty"`
	Locals               []*LocalInfo   `protobuf:"bytes,12,rep,name=locals,proto3" json:"locals,omitempty"`
	Specials             []*SpecialInfo `protobuf:"bytes,13,rep,name=specials,proto3" json:"specials,omitempty"`
	Tags                 []string       `protobuf:"bytes,14,rep,name=tags,proto3" json:"tags,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *ExhibitInfo) Reset()         { *m = ExhibitInfo{} }
func (m *ExhibitInfo) String() string { return proto.CompactTextString(m) }
func (*ExhibitInfo) ProtoMessage()    {}
func (*ExhibitInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_f26dec03d71c05fd, []int{0}
}

func (m *ExhibitInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExhibitInfo.Unmarshal(m, b)
}
func (m *ExhibitInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExhibitInfo.Marshal(b, m, deterministic)
}
func (m *ExhibitInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExhibitInfo.Merge(m, src)
}
func (m *ExhibitInfo) XXX_Size() int {
	return xxx_messageInfo_ExhibitInfo.Size(m)
}
func (m *ExhibitInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ExhibitInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ExhibitInfo proto.InternalMessageInfo

func (m *ExhibitInfo) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *ExhibitInfo) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *ExhibitInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ExhibitInfo) GetCreated() uint64 {
	if m != nil {
		return m.Created
	}
	return 0
}

func (m *ExhibitInfo) GetUpdated() uint64 {
	if m != nil {
		return m.Updated
	}
	return 0
}

func (m *ExhibitInfo) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *ExhibitInfo) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

func (m *ExhibitInfo) GetSn() string {
	if m != nil {
		return m.Sn
	}
	return ""
}

func (m *ExhibitInfo) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *ExhibitInfo) GetEntity() string {
	if m != nil {
		return m.Entity
	}
	return ""
}

func (m *ExhibitInfo) GetSize() *Vector3 {
	if m != nil {
		return m.Size
	}
	return nil
}

func (m *ExhibitInfo) GetLocals() []*LocalInfo {
	if m != nil {
		return m.Locals
	}
	return nil
}

func (m *ExhibitInfo) GetSpecials() []*SpecialInfo {
	if m != nil {
		return m.Specials
	}
	return nil
}

func (m *ExhibitInfo) GetTags() []string {
	if m != nil {
		return m.Tags
	}
	return nil
}

type SpecialInfo struct {
	Id                   uint32   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Key                  string   `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	Value                string   `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SpecialInfo) Reset()         { *m = SpecialInfo{} }
func (m *SpecialInfo) String() string { return proto.CompactTextString(m) }
func (*SpecialInfo) ProtoMessage()    {}
func (*SpecialInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_f26dec03d71c05fd, []int{1}
}

func (m *SpecialInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SpecialInfo.Unmarshal(m, b)
}
func (m *SpecialInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SpecialInfo.Marshal(b, m, deterministic)
}
func (m *SpecialInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SpecialInfo.Merge(m, src)
}
func (m *SpecialInfo) XXX_Size() int {
	return xxx_messageInfo_SpecialInfo.Size(m)
}
func (m *SpecialInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_SpecialInfo.DiscardUnknown(m)
}

var xxx_messageInfo_SpecialInfo proto.InternalMessageInfo

func (m *SpecialInfo) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *SpecialInfo) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *SpecialInfo) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type ReqExhibitAdd struct {
	Owner                string   `protobuf:"bytes,1,opt,name=owner,proto3" json:"owner,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Remark               string   `protobuf:"bytes,3,opt,name=remark,proto3" json:"remark,omitempty"`
	Operator             string   `protobuf:"bytes,4,opt,name=operator,proto3" json:"operator,omitempty"`
	Entity               string   `protobuf:"bytes,5,opt,name=entity,proto3" json:"entity,omitempty"`
	Sn                   string   `protobuf:"bytes,6,opt,name=sn,proto3" json:"sn,omitempty"`
	Tags                 []string `protobuf:"bytes,7,rep,name=tags,proto3" json:"tags,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReqExhibitAdd) Reset()         { *m = ReqExhibitAdd{} }
func (m *ReqExhibitAdd) String() string { return proto.CompactTextString(m) }
func (*ReqExhibitAdd) ProtoMessage()    {}
func (*ReqExhibitAdd) Descriptor() ([]byte, []int) {
	return fileDescriptor_f26dec03d71c05fd, []int{2}
}

func (m *ReqExhibitAdd) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqExhibitAdd.Unmarshal(m, b)
}
func (m *ReqExhibitAdd) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqExhibitAdd.Marshal(b, m, deterministic)
}
func (m *ReqExhibitAdd) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqExhibitAdd.Merge(m, src)
}
func (m *ReqExhibitAdd) XXX_Size() int {
	return xxx_messageInfo_ReqExhibitAdd.Size(m)
}
func (m *ReqExhibitAdd) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqExhibitAdd.DiscardUnknown(m)
}

var xxx_messageInfo_ReqExhibitAdd proto.InternalMessageInfo

func (m *ReqExhibitAdd) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *ReqExhibitAdd) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ReqExhibitAdd) GetRemark() string {
	if m != nil {
		return m.Remark
	}
	return ""
}

func (m *ReqExhibitAdd) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

func (m *ReqExhibitAdd) GetEntity() string {
	if m != nil {
		return m.Entity
	}
	return ""
}

func (m *ReqExhibitAdd) GetSn() string {
	if m != nil {
		return m.Sn
	}
	return ""
}

func (m *ReqExhibitAdd) GetTags() []string {
	if m != nil {
		return m.Tags
	}
	return nil
}

type ReqExhibitLocals struct {
	Uid                  string       `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Operator             string       `protobuf:"bytes,2,opt,name=operator,proto3" json:"operator,omitempty"`
	Locals               []*LocalInfo `protobuf:"bytes,3,rep,name=locals,proto3" json:"locals,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ReqExhibitLocals) Reset()         { *m = ReqExhibitLocals{} }
func (m *ReqExhibitLocals) String() string { return proto.CompactTextString(m) }
func (*ReqExhibitLocals) ProtoMessage()    {}
func (*ReqExhibitLocals) Descriptor() ([]byte, []int) {
	return fileDescriptor_f26dec03d71c05fd, []int{3}
}

func (m *ReqExhibitLocals) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqExhibitLocals.Unmarshal(m, b)
}
func (m *ReqExhibitLocals) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqExhibitLocals.Marshal(b, m, deterministic)
}
func (m *ReqExhibitLocals) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqExhibitLocals.Merge(m, src)
}
func (m *ReqExhibitLocals) XXX_Size() int {
	return xxx_messageInfo_ReqExhibitLocals.Size(m)
}
func (m *ReqExhibitLocals) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqExhibitLocals.DiscardUnknown(m)
}

var xxx_messageInfo_ReqExhibitLocals proto.InternalMessageInfo

func (m *ReqExhibitLocals) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *ReqExhibitLocals) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

func (m *ReqExhibitLocals) GetLocals() []*LocalInfo {
	if m != nil {
		return m.Locals
	}
	return nil
}

type ReqExhibitSpecials struct {
	Uid                  string         `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Operator             string         `protobuf:"bytes,2,opt,name=operator,proto3" json:"operator,omitempty"`
	Specials             []*SpecialInfo `protobuf:"bytes,3,rep,name=specials,proto3" json:"specials,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *ReqExhibitSpecials) Reset()         { *m = ReqExhibitSpecials{} }
func (m *ReqExhibitSpecials) String() string { return proto.CompactTextString(m) }
func (*ReqExhibitSpecials) ProtoMessage()    {}
func (*ReqExhibitSpecials) Descriptor() ([]byte, []int) {
	return fileDescriptor_f26dec03d71c05fd, []int{4}
}

func (m *ReqExhibitSpecials) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqExhibitSpecials.Unmarshal(m, b)
}
func (m *ReqExhibitSpecials) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqExhibitSpecials.Marshal(b, m, deterministic)
}
func (m *ReqExhibitSpecials) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqExhibitSpecials.Merge(m, src)
}
func (m *ReqExhibitSpecials) XXX_Size() int {
	return xxx_messageInfo_ReqExhibitSpecials.Size(m)
}
func (m *ReqExhibitSpecials) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqExhibitSpecials.DiscardUnknown(m)
}

var xxx_messageInfo_ReqExhibitSpecials proto.InternalMessageInfo

func (m *ReqExhibitSpecials) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *ReqExhibitSpecials) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

func (m *ReqExhibitSpecials) GetSpecials() []*SpecialInfo {
	if m != nil {
		return m.Specials
	}
	return nil
}

type ReplyExhibitInfo struct {
	Info                 *ExhibitInfo `protobuf:"bytes,1,opt,name=info,proto3" json:"info,omitempty"`
	Status               *ReplyStatus `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ReplyExhibitInfo) Reset()         { *m = ReplyExhibitInfo{} }
func (m *ReplyExhibitInfo) String() string { return proto.CompactTextString(m) }
func (*ReplyExhibitInfo) ProtoMessage()    {}
func (*ReplyExhibitInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_f26dec03d71c05fd, []int{5}
}

func (m *ReplyExhibitInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReplyExhibitInfo.Unmarshal(m, b)
}
func (m *ReplyExhibitInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReplyExhibitInfo.Marshal(b, m, deterministic)
}
func (m *ReplyExhibitInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReplyExhibitInfo.Merge(m, src)
}
func (m *ReplyExhibitInfo) XXX_Size() int {
	return xxx_messageInfo_ReplyExhibitInfo.Size(m)
}
func (m *ReplyExhibitInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ReplyExhibitInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ReplyExhibitInfo proto.InternalMessageInfo

func (m *ReplyExhibitInfo) GetInfo() *ExhibitInfo {
	if m != nil {
		return m.Info
	}
	return nil
}

func (m *ReplyExhibitInfo) GetStatus() *ReplyStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

type ReplyExhibitList struct {
	Total                uint32         `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`
	Pages                uint32         `protobuf:"varint,2,opt,name=pages,proto3" json:"pages,omitempty"`
	Page                 uint32         `protobuf:"varint,3,opt,name=page,proto3" json:"page,omitempty"`
	Number               uint32         `protobuf:"varint,4,opt,name=number,proto3" json:"number,omitempty"`
	Status               *ReplyStatus   `protobuf:"bytes,5,opt,name=status,proto3" json:"status,omitempty"`
	List                 []*ExhibitInfo `protobuf:"bytes,6,rep,name=list,proto3" json:"list,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *ReplyExhibitList) Reset()         { *m = ReplyExhibitList{} }
func (m *ReplyExhibitList) String() string { return proto.CompactTextString(m) }
func (*ReplyExhibitList) ProtoMessage()    {}
func (*ReplyExhibitList) Descriptor() ([]byte, []int) {
	return fileDescriptor_f26dec03d71c05fd, []int{6}
}

func (m *ReplyExhibitList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReplyExhibitList.Unmarshal(m, b)
}
func (m *ReplyExhibitList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReplyExhibitList.Marshal(b, m, deterministic)
}
func (m *ReplyExhibitList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReplyExhibitList.Merge(m, src)
}
func (m *ReplyExhibitList) XXX_Size() int {
	return xxx_messageInfo_ReplyExhibitList.Size(m)
}
func (m *ReplyExhibitList) XXX_DiscardUnknown() {
	xxx_messageInfo_ReplyExhibitList.DiscardUnknown(m)
}

var xxx_messageInfo_ReplyExhibitList proto.InternalMessageInfo

func (m *ReplyExhibitList) GetTotal() uint32 {
	if m != nil {
		return m.Total
	}
	return 0
}

func (m *ReplyExhibitList) GetPages() uint32 {
	if m != nil {
		return m.Pages
	}
	return 0
}

func (m *ReplyExhibitList) GetPage() uint32 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *ReplyExhibitList) GetNumber() uint32 {
	if m != nil {
		return m.Number
	}
	return 0
}

func (m *ReplyExhibitList) GetStatus() *ReplyStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *ReplyExhibitList) GetList() []*ExhibitInfo {
	if m != nil {
		return m.List
	}
	return nil
}

func init() {
	proto.RegisterType((*ExhibitInfo)(nil), "omo.msp.museum.ExhibitInfo")
	proto.RegisterType((*SpecialInfo)(nil), "omo.msp.museum.SpecialInfo")
	proto.RegisterType((*ReqExhibitAdd)(nil), "omo.msp.museum.ReqExhibitAdd")
	proto.RegisterType((*ReqExhibitLocals)(nil), "omo.msp.museum.ReqExhibitLocals")
	proto.RegisterType((*ReqExhibitSpecials)(nil), "omo.msp.museum.ReqExhibitSpecials")
	proto.RegisterType((*ReplyExhibitInfo)(nil), "omo.msp.museum.ReplyExhibitInfo")
	proto.RegisterType((*ReplyExhibitList)(nil), "omo.msp.museum.ReplyExhibitList")
}

func init() {
	proto.RegisterFile("proto/museum/exhibit.proto", fileDescriptor_f26dec03d71c05fd)
}

var fileDescriptor_f26dec03d71c05fd = []byte{
	// 718 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x55, 0xd1, 0x4e, 0xdb, 0x4a,
	0x10, 0xc5, 0x71, 0x62, 0x60, 0x42, 0x72, 0xd1, 0xea, 0x8a, 0x6b, 0x82, 0xb8, 0xb2, 0xfc, 0x14,
	0xa9, 0x52, 0x50, 0xe1, 0xa1, 0x0f, 0x7d, 0x02, 0x89, 0x22, 0xd4, 0x54, 0xa8, 0x4e, 0xdb, 0x87,
	0xbe, 0x20, 0x63, 0x0f, 0xb0, 0xc2, 0xf6, 0x1a, 0xef, 0x9a, 0x92, 0xf6, 0x47, 0xfa, 0x0b, 0xfd,
	0x9b, 0xf6, 0x8f, 0xaa, 0xdd, 0xb5, 0x63, 0x27, 0x35, 0x10, 0xd4, 0xb7, 0x9d, 0x9d, 0x33, 0x67,
	0xc6, 0xe7, 0xcc, 0x26, 0x30, 0x48, 0x33, 0x26, 0xd8, 0x5e, 0x9c, 0x73, 0xcc, 0xe3, 0x3d, 0xbc,
	0xbf, 0xa6, 0x17, 0x54, 0x8c, 0xd4, 0x25, 0xe9, 0xb3, 0x98, 0x8d, 0x62, 0x9e, 0x8e, 0x74, 0x76,
	0xb0, 0x3d, 0x87, 0x0d, 0x58, 0x1c, 0xb3, 0x44, 0x43, 0xdd, 0xef, 0x26, 0x74, 0x8f, 0x75, 0xf1,
	0x69, 0x72, 0xc9, 0x48, 0x1f, 0x5a, 0x34, 0xb4, 0x0d, 0xc7, 0x18, 0xb6, 0xbd, 0x16, 0x0d, 0xc9,
	0x26, 0x98, 0x39, 0x0d, 0xed, 0x96, 0x63, 0x0c, 0xd7, 0x3d, 0x79, 0x24, 0x04, 0xda, 0x89, 0x1f,
	0xa3, 0x6d, 0xaa, 0x2b, 0x75, 0x26, 0x36, 0xac, 0x06, 0x19, 0xfa, 0x02, 0x43, 0xbb, 0xad, 0x4a,
	0xcb, 0x50, 0x66, 0xf2, 0x34, 0x54, 0x99, 0x8e, 0xce, 0x14, 0xe1, 0xac, 0x86, 0x65, 0xb6, 0xa5,
	0xa8, 0xca, 0x90, 0x0c, 0x60, 0x8d, 0xa5, 0x98, 0xa9, 0xd4, 0xaa, 0x4a, 0xcd, 0x62, 0x39, 0x1f,
	0x4f, 0xec, 0x35, 0x75, 0xdb, 0xe2, 0x09, 0xf9, 0x17, 0x3a, 0xec, 0x4b, 0x82, 0x99, 0xbd, 0xae,
	0xae, 0x74, 0x40, 0xb6, 0xc0, 0xc2, 0x44, 0x50, 0x31, 0xb5, 0x41, 0x5d, 0x17, 0x11, 0x79, 0x01,
	0x6d, 0x4e, 0xbf, 0xa2, 0xdd, 0x75, 0x8c, 0x61, 0x77, 0xff, 0xbf, 0xd1, 0xbc, 0x4e, 0xa3, 0x4f,
	0x18, 0x08, 0x96, 0x1d, 0x78, 0x0a, 0x44, 0x5e, 0x82, 0x15, 0xb1, 0xc0, 0x8f, 0xb8, 0xbd, 0xe1,
	0x98, 0xc3, 0xee, 0xfe, 0xf6, 0x22, 0x7c, 0x2c, 0xb3, 0x52, 0x35, 0xaf, 0x00, 0x92, 0x57, 0xb0,
	0xc6, 0x53, 0x0c, 0xa8, 0x2c, 0xea, 0xa9, 0xa2, 0x9d, 0xc5, 0xa2, 0x89, 0xce, 0xab, 0xb2, 0x19,
	0x58, 0x8a, 0x2a, 0xfc, 0x2b, 0x6e, 0xf7, 0x1d, 0x53, 0x8a, 0x2a, 0xcf, 0xee, 0x31, 0x74, 0x6b,
	0xe0, 0x9a, 0x33, 0xbd, 0xd2, 0x99, 0x1b, 0x9c, 0x96, 0xce, 0xdc, 0xe0, 0x54, 0x6a, 0x71, 0xe7,
	0x47, 0x79, 0x69, 0x8d, 0x0e, 0xdc, 0x1f, 0x06, 0xf4, 0x3c, 0xbc, 0x2d, 0x4c, 0x3e, 0x0c, 0xc3,
	0x4a, 0x33, 0xa3, 0xae, 0x59, 0xe9, 0x6b, 0xab, 0xe6, 0xeb, 0x16, 0x58, 0x19, 0xc6, 0x7e, 0x76,
	0x53, 0x50, 0x16, 0xd1, 0x9c, 0x43, 0xed, 0x05, 0x87, 0x2a, 0xed, 0x3b, 0x73, 0xda, 0x6b, 0xe7,
	0xac, 0x99, 0x73, 0xe5, 0x27, 0xaf, 0xd6, 0x3e, 0x99, 0xc3, 0x66, 0x35, 0xea, 0x58, 0x6b, 0x5a,
	0x6c, 0xa0, 0x51, 0x6d, 0x60, 0xbd, 0x7b, 0x6b, 0xa1, 0x7b, 0x65, 0x9a, 0xb9, 0xa4, 0x69, 0xee,
	0x37, 0x20, 0x55, 0xd3, 0x49, 0xe9, 0xc8, 0xf3, 0xda, 0xd6, 0x8d, 0x37, 0x9f, 0x61, 0xbc, 0x7b,
	0x2f, 0xbf, 0x38, 0x8d, 0xa6, 0xf5, 0x37, 0xb8, 0x07, 0x6d, 0x9a, 0x5c, 0x32, 0xd5, 0xbb, 0x81,
	0xa8, 0x06, 0xf5, 0x14, 0x90, 0x1c, 0x80, 0xc5, 0x85, 0x2f, 0x72, 0xae, 0xe6, 0x6a, 0x28, 0x51,
	0x2d, 0x26, 0x0a, 0xe2, 0x15, 0x50, 0xf7, 0xa7, 0x31, 0xdf, 0x7a, 0x4c, 0xb9, 0x90, 0xab, 0x21,
	0x98, 0xf0, 0xa3, 0x62, 0xcf, 0x74, 0x20, 0x6f, 0x53, 0xff, 0x0a, 0x35, 0x7d, 0xcf, 0xd3, 0x81,
	0x34, 0x50, 0x1e, 0xd4, 0x6a, 0xf4, 0x3c, 0x75, 0x96, 0xe6, 0x27, 0x79, 0x7c, 0x81, 0x7a, 0x2d,
	0x7a, 0x5e, 0x11, 0xd5, 0x26, 0xec, 0x2c, 0x3d, 0xa1, 0xd4, 0x21, 0xa2, 0x5c, 0xd8, 0x56, 0xb3,
	0xa0, 0x73, 0x3a, 0x48, 0xe0, 0xfe, 0xaf, 0x0e, 0xf4, 0x4b, 0x1f, 0x31, 0xbb, 0xa3, 0x01, 0x92,
	0xb7, 0x60, 0x1d, 0x86, 0xe1, 0x59, 0x82, 0x64, 0xf7, 0xcf, 0x96, 0xb5, 0x47, 0x31, 0x70, 0x1a,
	0x27, 0xaa, 0xf5, 0x70, 0x57, 0xc8, 0x29, 0x58, 0x27, 0x28, 0x24, 0x59, 0xc3, 0xfc, 0xb7, 0x39,
	0x72, 0x05, 0x5c, 0x8a, 0xea, 0x18, 0xd6, 0x3d, 0x8c, 0xd9, 0x1d, 0x3e, 0xc9, 0xb6, 0xdd, 0xc8,
	0x56, 0x4d, 0x34, 0x41, 0x3f, 0x0b, 0xae, 0xff, 0x62, 0x22, 0x69, 0xbc, 0xbb, 0x42, 0x3e, 0xc0,
	0x3f, 0x27, 0xa8, 0x82, 0xa3, 0xe9, 0x1b, 0x1a, 0x09, 0xcc, 0x1a, 0x25, 0x93, 0x9c, 0x3a, 0xbd,
	0x14, 0xeb, 0x19, 0x6c, 0x9c, 0xa0, 0x90, 0xc6, 0x52, 0x2e, 0x68, 0xf0, 0x14, 0xe5, 0xff, 0x0f,
	0xee, 0x85, 0x2a, 0x77, 0x57, 0xc8, 0x3b, 0xd8, 0xf8, 0xa8, 0xfe, 0x41, 0x8a, 0x9f, 0x07, 0xe7,
	0x61, 0x5b, 0x35, 0xe2, 0x71, 0x01, 0xdf, 0x43, 0x5f, 0xd3, 0xcd, 0x1e, 0xbe, 0xfb, 0x30, 0x61,
	0x89, 0x79, 0x9c, 0x72, 0x5c, 0x52, 0x3e, 0xa9, 0xa3, 0x86, 0x3d, 0xca, 0x76, 0xb4, 0xfb, 0x79,
	0xa7, 0xfe, 0xef, 0xfd, 0x9a, 0xc5, 0xec, 0x3c, 0xe6, 0xe9, 0xb9, 0x0e, 0x2f, 0x2c, 0x95, 0x3c,
	0xf8, 0x1d, 0x00, 0x00, 0xff, 0xff, 0x08, 0xe4, 0x9a, 0xb7, 0x0f, 0x08, 0x00, 0x00,
}
