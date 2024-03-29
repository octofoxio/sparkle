// Code generated by protoc-gen-go. DO NOT EDIT.
// source: entities/v1/user.proto

package entitiesv1

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	v1 "github.com/octofoxio/sparkle/pkg/common/v1"
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

type UserStatus int32

const (
	UserStatus_None                        UserStatus = 0
	UserStatus_WaitingForEmailVerification UserStatus = 1
	UserStatus_Active                      UserStatus = 2
)

var UserStatus_name = map[int32]string{
	0: "None",
	1: "WaitingForEmailVerification",
	2: "Active",
}

var UserStatus_value = map[string]int32{
	"None":                        0,
	"WaitingForEmailVerification": 1,
	"Active":                      2,
}

func (x UserStatus) String() string {
	return proto.EnumName(UserStatus_name, int32(x))
}

func (UserStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_0d9dc2379926ce20, []int{0}
}

type User struct {
	// @inject_tag: bson:"_id,omitempty"
	ID *v1.String `protobuf:"bytes,100,opt,name=ID,proto3" json:"ID,omitempty" bson:"_id,omitempty"`
	// @inject_tag: bson:",omitempty"
	Email *v1.String `protobuf:"bytes,101,opt,name=Email,proto3" json:"Email,omitempty" bson:",omitempty"`
	// @inject_tag: bson:",omitempty"
	FullName *v1.String `protobuf:"bytes,102,opt,name=FullName,proto3" json:"FullName,omitempty" bson:",omitempty"`
	// @inject_tag: bson:",omitempty"
	PhoneNumber *v1.String `protobuf:"bytes,103,opt,name=PhoneNumber,proto3" json:"PhoneNumber,omitempty" bson:",omitempty"`
	// @inject_tag: bson:",omitempty"
	BirthDay *v1.Timestamp `protobuf:"bytes,104,opt,name=BirthDay,proto3" json:"BirthDay,omitempty" bson:",omitempty"`
	// @inject_tag: bson:",omitempty"
	Gender v1.Gender `protobuf:"varint,105,opt,name=Gender,proto3,enum=sparkle.common.v1.Gender" json:"Gender,omitempty" bson:",omitempty"`
	// @inject_tag: bson:",omitempty"
	Status UserStatus `protobuf:"varint,10,opt,name=Status,proto3,enum=sparkle.entities.v1.UserStatus" json:"Status,omitempty" bson:",omitempty"`
	// @inject_tag: bson:",omitempty"
	CreatedAt            *v1.Timestamp `protobuf:"bytes,11,opt,name=CreatedAt,proto3" json:"CreatedAt,omitempty" bson:",omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-" gorm:"-" bson:"-"`
	XXX_unrecognized     []byte        `json:"-" gorm:"-" bson:"-"`
	XXX_sizecache        int32         `json:"-" gorm:"-" bson:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_0d9dc2379926ce20, []int{0}
}

func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (m *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(m, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetID() *v1.String {
	if m != nil {
		return m.ID
	}
	return nil
}

func (m *User) GetEmail() *v1.String {
	if m != nil {
		return m.Email
	}
	return nil
}

func (m *User) GetFullName() *v1.String {
	if m != nil {
		return m.FullName
	}
	return nil
}

func (m *User) GetPhoneNumber() *v1.String {
	if m != nil {
		return m.PhoneNumber
	}
	return nil
}

func (m *User) GetBirthDay() *v1.Timestamp {
	if m != nil {
		return m.BirthDay
	}
	return nil
}

func (m *User) GetGender() v1.Gender {
	if m != nil {
		return m.Gender
	}
	return v1.Gender_Male
}

func (m *User) GetStatus() UserStatus {
	if m != nil {
		return m.Status
	}
	return UserStatus_None
}

func (m *User) GetCreatedAt() *v1.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func init() {
	proto.RegisterEnum("sparkle.entities.v1.UserStatus", UserStatus_name, UserStatus_value)
	proto.RegisterType((*User)(nil), "sparkle.entities.v1.User")
}

func init() { proto.RegisterFile("entities/v1/user.proto", fileDescriptor_0d9dc2379926ce20) }

var fileDescriptor_0d9dc2379926ce20 = []byte{
	// 409 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xcf, 0x6e, 0xd3, 0x40,
	0x10, 0x87, 0xb1, 0x5b, 0xa2, 0x30, 0x91, 0x50, 0xb4, 0x95, 0x60, 0x0b, 0x48, 0xad, 0x38, 0x15,
	0x0e, 0x6b, 0x19, 0x84, 0x8a, 0xe8, 0xa9, 0x69, 0x52, 0xd4, 0x4b, 0x14, 0x25, 0x10, 0x10, 0xea,
	0x65, 0xe3, 0x4c, 0x9c, 0x51, 0xb3, 0xbb, 0xd6, 0x7a, 0x6d, 0xd1, 0xe7, 0xe0, 0x0d, 0x38, 0xf2,
	0x04, 0x3c, 0x03, 0x4f, 0x85, 0xfc, 0x2f, 0xed, 0xc1, 0x52, 0x7a, 0xb3, 0xc7, 0xdf, 0x37, 0x33,
	0xf2, 0xfc, 0xe0, 0x19, 0x6a, 0x47, 0x8e, 0x30, 0x0d, 0xf2, 0x30, 0xc8, 0x52, 0xb4, 0x22, 0xb1,
	0xc6, 0x19, 0x76, 0x90, 0x26, 0xd2, 0xde, 0x6c, 0x50, 0x34, 0xdf, 0x45, 0x1e, 0xbe, 0x38, 0x8c,
	0x8c, 0x52, 0x46, 0x17, 0x68, 0x62, 0x49, 0x91, 0xa3, 0x1c, 0x2b, 0xfe, 0xf5, 0xdf, 0x3d, 0xd8,
	0xff, 0x9a, 0xa2, 0x65, 0x6f, 0xc0, 0xbf, 0x1a, 0xf2, 0xe5, 0xb1, 0x77, 0xd2, 0x7b, 0x77, 0x28,
	0x9a, 0x2e, 0x95, 0x28, 0xf2, 0x50, 0xcc, 0x9c, 0x25, 0x1d, 0x4f, 0xfd, 0xab, 0x21, 0x0b, 0xe0,
	0xf1, 0x48, 0x49, 0xda, 0x70, 0xdc, 0x45, 0x57, 0x1c, 0xfb, 0x00, 0xdd, 0xcb, 0x6c, 0xb3, 0x19,
	0x4b, 0x85, 0x7c, 0xb5, 0xcb, 0xd9, 0xa2, 0xec, 0x0c, 0x7a, 0x93, 0xb5, 0xd1, 0x38, 0xce, 0xd4,
	0x02, 0x2d, 0x8f, 0x77, 0x99, 0xf7, 0x69, 0xf6, 0x11, 0xba, 0x03, 0xb2, 0x6e, 0x3d, 0x94, 0xb7,
	0x7c, 0x5d, 0x9a, 0xaf, 0x5a, 0xcc, 0x2f, 0xa4, 0x30, 0x75, 0x52, 0x25, 0xd3, 0x2d, 0xcd, 0x42,
	0xe8, 0x7c, 0x46, 0xbd, 0x44, 0xcb, 0xe9, 0xd8, 0x3b, 0x79, 0xda, 0x3a, 0xb1, 0x02, 0xa6, 0x35,
	0xc8, 0x4e, 0xa1, 0x33, 0x73, 0xd2, 0x65, 0x29, 0x87, 0x52, 0x39, 0x12, 0x2d, 0x67, 0x10, 0xc5,
	0x7f, 0xae, 0xb0, 0x69, 0x8d, 0xb3, 0x4f, 0xf0, 0xe4, 0xc2, 0xa2, 0x74, 0xb8, 0x3c, 0x77, 0xbc,
	0xf7, 0x80, 0x35, 0xef, 0xf0, 0xb7, 0x17, 0x00, 0x77, 0x1d, 0x59, 0x17, 0xf6, 0xc7, 0x46, 0x63,
	0xff, 0x11, 0x3b, 0x82, 0x97, 0xdf, 0x24, 0x39, 0xd2, 0xf1, 0xa5, 0xb1, 0xe5, 0x01, 0xe6, 0x68,
	0x69, 0x45, 0x91, 0x74, 0x64, 0x74, 0xdf, 0x63, 0x00, 0x9d, 0xf3, 0xa8, 0xc8, 0x40, 0xdf, 0x1f,
	0xfc, 0xf2, 0xe0, 0x79, 0x64, 0x54, 0xdb, 0xbe, 0x83, 0xde, 0xa8, 0x78, 0xb9, 0x9d, 0x14, 0x41,
	0x99, 0x78, 0x3f, 0x4e, 0x63, 0x72, 0xeb, 0x6c, 0x51, 0xac, 0x15, 0x98, 0xc8, 0x99, 0x95, 0xf9,
	0x49, 0x26, 0xa8, 0xc5, 0x20, 0xb9, 0x89, 0x83, 0x7b, 0x99, 0x3c, 0x6b, 0x9e, 0xf3, 0xf0, 0xb7,
	0xbf, 0x37, 0x1b, 0x7d, 0xff, 0xe3, 0x1f, 0xcc, 0xea, 0x19, 0xa3, 0x66, 0xc6, 0x3c, 0xfc, 0xb7,
	0xad, 0x5e, 0x37, 0xd5, 0xeb, 0x79, 0xb8, 0xe8, 0x94, 0xe1, 0x7c, 0xff, 0x3f, 0x00, 0x00, 0xff,
	0xff, 0xa8, 0xe4, 0x90, 0xfe, 0xe6, 0x02, 0x00, 0x00,
}
