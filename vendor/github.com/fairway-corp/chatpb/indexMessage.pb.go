// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: indexMessage.proto

package chatpb

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type StatusRequest struct {
}

func (m *StatusRequest) Reset()                    { *m = StatusRequest{} }
func (m *StatusRequest) String() string            { return proto.CompactTextString(m) }
func (*StatusRequest) ProtoMessage()               {}
func (*StatusRequest) Descriptor() ([]byte, []int) { return fileDescriptorIndexMessage, []int{0} }

type StatusResponse struct {
	AppName string `protobuf:"bytes,1,opt,name=appName,proto3" json:"appName,omitempty"`
	Status  string `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
}

func (m *StatusResponse) Reset()                    { *m = StatusResponse{} }
func (m *StatusResponse) String() string            { return proto.CompactTextString(m) }
func (*StatusResponse) ProtoMessage()               {}
func (*StatusResponse) Descriptor() ([]byte, []int) { return fileDescriptorIndexMessage, []int{1} }

func (m *StatusResponse) GetAppName() string {
	if m != nil {
		return m.AppName
	}
	return ""
}

func (m *StatusResponse) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func init() {
	proto.RegisterType((*StatusRequest)(nil), "fairway.chatpb.StatusRequest")
	proto.RegisterType((*StatusResponse)(nil), "fairway.chatpb.StatusResponse")
}

func init() { proto.RegisterFile("indexMessage.proto", fileDescriptorIndexMessage) }

var fileDescriptorIndexMessage = []byte{
	// 152 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xca, 0xcc, 0x4b, 0x49,
	0xad, 0xf0, 0x4d, 0x2d, 0x2e, 0x4e, 0x4c, 0x4f, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2,
	0x4b, 0x4b, 0xcc, 0x2c, 0x2a, 0x4f, 0xac, 0xd4, 0x4b, 0xce, 0x48, 0x2c, 0x29, 0x48, 0x52, 0xe2,
	0xe7, 0xe2, 0x0d, 0x2e, 0x49, 0x2c, 0x29, 0x2d, 0x0e, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x51,
	0x72, 0xe2, 0xe2, 0x83, 0x09, 0x14, 0x17, 0xe4, 0xe7, 0x15, 0xa7, 0x0a, 0x49, 0x70, 0xb1, 0x27,
	0x16, 0x14, 0xf8, 0x25, 0xe6, 0xa6, 0x4a, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0xc1, 0xb8, 0x42,
	0x62, 0x5c, 0x6c, 0xc5, 0x60, 0xb5, 0x12, 0x4c, 0x60, 0x09, 0x28, 0xcf, 0x49, 0x21, 0x4a, 0x2e,
	0x3d, 0xb3, 0x24, 0xa3, 0x34, 0x49, 0x2f, 0x39, 0x3f, 0x57, 0x1f, 0x6a, 0xa3, 0x6e, 0x72, 0x7e,
	0x51, 0x81, 0x3e, 0xc4, 0xda, 0x24, 0x36, 0xb0, 0x6b, 0x8c, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff,
	0x45, 0x69, 0x1a, 0x94, 0xa3, 0x00, 0x00, 0x00,
}