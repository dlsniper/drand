// Code generated by protoc-gen-go. DO NOT EDIT.
// source: crypto/vss/vss.proto

package vss

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import crypto "github.com/drand/drand/protobuf/crypto"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// EncryptedDeal holds a share encrypted towards the share holder's longterm
// public key..
type EncryptedDeal struct {
	// ephemereal diffie hellman key
	Dhkey []byte `protobuf:"bytes,1,opt,name=dhkey,proto3" json:"dhkey,omitempty"`
	// schnorr signature over the dhkey by the longterm key of the dealer
	Signature []byte `protobuf:"bytes,2,opt,name=signature,proto3" json:"signature,omitempty"`
	// nonce used in the generation of the ephemereal key
	Nonce []byte `protobuf:"bytes,3,opt,name=nonce,proto3" json:"nonce,omitempty"`
	// cipher of the deal marshalled by protobuf2. TODO: marshal in a
	// more explicit and easier way.
	Cipher               []byte   `protobuf:"bytes,4,opt,name=cipher,proto3" json:"cipher,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EncryptedDeal) Reset()         { *m = EncryptedDeal{} }
func (m *EncryptedDeal) String() string { return proto.CompactTextString(m) }
func (*EncryptedDeal) ProtoMessage()    {}
func (*EncryptedDeal) Descriptor() ([]byte, []int) {
	return fileDescriptor_vss_0694a90ccd2794fa, []int{0}
}
func (m *EncryptedDeal) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EncryptedDeal.Unmarshal(m, b)
}
func (m *EncryptedDeal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EncryptedDeal.Marshal(b, m, deterministic)
}
func (dst *EncryptedDeal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EncryptedDeal.Merge(dst, src)
}
func (m *EncryptedDeal) XXX_Size() int {
	return xxx_messageInfo_EncryptedDeal.Size(m)
}
func (m *EncryptedDeal) XXX_DiscardUnknown() {
	xxx_messageInfo_EncryptedDeal.DiscardUnknown(m)
}

var xxx_messageInfo_EncryptedDeal proto.InternalMessageInfo

func (m *EncryptedDeal) GetDhkey() []byte {
	if m != nil {
		return m.Dhkey
	}
	return nil
}

func (m *EncryptedDeal) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

func (m *EncryptedDeal) GetNonce() []byte {
	if m != nil {
		return m.Nonce
	}
	return nil
}

func (m *EncryptedDeal) GetCipher() []byte {
	if m != nil {
		return m.Cipher
	}
	return nil
}

//
// Deal holds the share created by a dealer for a round of a vss or dkg protocol
// It is always meant to be encrypted when on transit because it contains
// private information (the share).
type Deal struct {
	// session id of the current protocol run
	SessionId []byte `protobuf:"bytes,1,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`
	// private share of the deal
	Share *crypto.PrivateShare `protobuf:"bytes,2,opt,name=share,proto3" json:"share,omitempty"`
	// threshold of the secret sharing protocol run
	Threshold uint32 `protobuf:"varint,3,opt,name=threshold,proto3" json:"threshold,omitempty"`
	// commitments of the polynomial used to derive the share
	Commitments          [][]byte `protobuf:"bytes,4,rep,name=commitments,proto3" json:"commitments,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Deal) Reset()         { *m = Deal{} }
func (m *Deal) String() string { return proto.CompactTextString(m) }
func (*Deal) ProtoMessage()    {}
func (*Deal) Descriptor() ([]byte, []int) {
	return fileDescriptor_vss_0694a90ccd2794fa, []int{1}
}
func (m *Deal) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Deal.Unmarshal(m, b)
}
func (m *Deal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Deal.Marshal(b, m, deterministic)
}
func (dst *Deal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Deal.Merge(dst, src)
}
func (m *Deal) XXX_Size() int {
	return xxx_messageInfo_Deal.Size(m)
}
func (m *Deal) XXX_DiscardUnknown() {
	xxx_messageInfo_Deal.DiscardUnknown(m)
}

var xxx_messageInfo_Deal proto.InternalMessageInfo

func (m *Deal) GetSessionId() []byte {
	if m != nil {
		return m.SessionId
	}
	return nil
}

func (m *Deal) GetShare() *crypto.PrivateShare {
	if m != nil {
		return m.Share
	}
	return nil
}

func (m *Deal) GetThreshold() uint32 {
	if m != nil {
		return m.Threshold
	}
	return 0
}

func (m *Deal) GetCommitments() [][]byte {
	if m != nil {
		return m.Commitments
	}
	return nil
}

//
// Response is the response of a participant after having received an encrypted
// deal. It is meant to be broadcasted to every participants.
type Response struct {
	// session id of the
	SessionId []byte `protobuf:"bytes,1,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`
	// index of the verifier issuing the response
	Index uint32 `protobuf:"varint,2,opt,name=index,proto3" json:"index,omitempty"`
	// status of the response. false = complaint, true = approval
	Status bool `protobuf:"varint,3,opt,name=status,proto3" json:"status,omitempty"`
	// signature over the packet using the longterm's key of the participant at
	// the given index
	Signature            []byte   `protobuf:"bytes,4,opt,name=signature,proto3" json:"signature,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_vss_0694a90ccd2794fa, []int{2}
}
func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (dst *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(dst, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetSessionId() []byte {
	if m != nil {
		return m.SessionId
	}
	return nil
}

func (m *Response) GetIndex() uint32 {
	if m != nil {
		return m.Index
	}
	return 0
}

func (m *Response) GetStatus() bool {
	if m != nil {
		return m.Status
	}
	return false
}

func (m *Response) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

//
// Justification enables a dealer to justify that it did not cheat in case some
// nodes complain about their received deal. It is NOT YET in production use
// though.
type Justification struct {
	// session id of the current run
	SessionId []byte `protobuf:"bytes,1,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`
	// index of the issuer of this justification
	Index uint32 `protobuf:"varint,2,opt,name=index,proto3" json:"index,omitempty"`
	// plaintext deal that the complaint response points to
	Deal *Deal `protobuf:"bytes,3,opt,name=deal,proto3" json:"deal,omitempty"`
	// signature over the whole packet
	Signature            []byte   `protobuf:"bytes,4,opt,name=signature,proto3" json:"signature,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Justification) Reset()         { *m = Justification{} }
func (m *Justification) String() string { return proto.CompactTextString(m) }
func (*Justification) ProtoMessage()    {}
func (*Justification) Descriptor() ([]byte, []int) {
	return fileDescriptor_vss_0694a90ccd2794fa, []int{3}
}
func (m *Justification) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Justification.Unmarshal(m, b)
}
func (m *Justification) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Justification.Marshal(b, m, deterministic)
}
func (dst *Justification) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Justification.Merge(dst, src)
}
func (m *Justification) XXX_Size() int {
	return xxx_messageInfo_Justification.Size(m)
}
func (m *Justification) XXX_DiscardUnknown() {
	xxx_messageInfo_Justification.DiscardUnknown(m)
}

var xxx_messageInfo_Justification proto.InternalMessageInfo

func (m *Justification) GetSessionId() []byte {
	if m != nil {
		return m.SessionId
	}
	return nil
}

func (m *Justification) GetIndex() uint32 {
	if m != nil {
		return m.Index
	}
	return 0
}

func (m *Justification) GetDeal() *Deal {
	if m != nil {
		return m.Deal
	}
	return nil
}

func (m *Justification) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

func init() {
	proto.RegisterType((*EncryptedDeal)(nil), "vss.EncryptedDeal")
	proto.RegisterType((*Deal)(nil), "vss.Deal")
	proto.RegisterType((*Response)(nil), "vss.Response")
	proto.RegisterType((*Justification)(nil), "vss.Justification")
}

func init() { proto.RegisterFile("github.com/drand/drand/protobuf/crypto/vss/vss.proto", fileDescriptor_vss_0694a90ccd2794fa) }

var fileDescriptor_vss_0694a90ccd2794fa = []byte{
	// 317 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0x31, 0x4f, 0xeb, 0x30,
	0x10, 0xc7, 0x95, 0xd7, 0xa4, 0x6a, 0xaf, 0xcd, 0xe2, 0x57, 0x3d, 0x45, 0x4f, 0x54, 0xaa, 0x32,
	0x95, 0xa5, 0x48, 0xe5, 0x1b, 0x20, 0x18, 0x60, 0x42, 0x66, 0x63, 0x41, 0x26, 0x3e, 0x88, 0x45,
	0x6b, 0x07, 0x9f, 0x53, 0xd1, 0x91, 0x6f, 0xc0, 0x47, 0x46, 0xbe, 0x44, 0x2a, 0x30, 0x80, 0xc4,
	0xe0, 0xe1, 0xff, 0x3f, 0x9f, 0xfe, 0x3f, 0x9f, 0x0f, 0x66, 0x95, 0xdf, 0x37, 0xc1, 0x9d, 0xec,
	0x88, 0xe2, 0x59, 0x35, 0xde, 0x05, 0x27, 0x06, 0x3b, 0xa2, 0xff, 0xa2, 0x2f, 0x51, 0xad, 0x3c,
	0x76, 0x85, 0xf2, 0x19, 0xf2, 0x0b, 0xcb, 0x3e, 0xea, 0x73, 0x54, 0x1b, 0x31, 0x83, 0x4c, 0xd7,
	0x4f, 0xb8, 0x2f, 0x92, 0x45, 0xb2, 0x9c, 0xca, 0x4e, 0x88, 0x23, 0x18, 0x93, 0x79, 0xb4, 0x2a,
	0xb4, 0x1e, 0x8b, 0x3f, 0x5c, 0x39, 0x18, 0xb1, 0xc7, 0x3a, 0x5b, 0x61, 0x31, 0xe8, 0x7a, 0x58,
	0x88, 0x7f, 0x30, 0xac, 0x4c, 0x53, 0xa3, 0x2f, 0x52, 0xb6, 0x7b, 0x55, 0xbe, 0x25, 0x90, 0x72,
	0xd4, 0x1c, 0x80, 0x90, 0xc8, 0x38, 0x7b, 0x67, 0x74, 0x9f, 0x37, 0xee, 0x9d, 0x4b, 0x2d, 0x8e,
	0x21, 0x63, 0x52, 0xce, 0x9b, 0xac, 0xff, 0xae, 0x3a, 0xee, 0x6b, 0x6f, 0x76, 0x2a, 0xe0, 0x4d,
	0x14, 0xb2, 0xbb, 0x11, 0xf1, 0x42, 0xed, 0x91, 0x6a, 0xb7, 0xd1, 0x0c, 0x91, 0xcb, 0x83, 0x21,
	0x16, 0x30, 0xa9, 0xdc, 0x76, 0x6b, 0xc2, 0x16, 0x6d, 0xa0, 0x22, 0x5d, 0x0c, 0x96, 0x53, 0xf9,
	0xd1, 0x2a, 0x5b, 0x18, 0x49, 0xa4, 0xc6, 0x59, 0xc2, 0x9f, 0xa8, 0x66, 0x90, 0x19, 0xab, 0xf1,
	0x85, 0xa9, 0x72, 0xd9, 0x89, 0xf8, 0x56, 0x0a, 0x2a, 0xb4, 0xc4, 0xe9, 0x23, 0xd9, 0xab, 0xcf,
	0x73, 0x4b, 0xbf, 0xcc, 0xad, 0x7c, 0x4d, 0x20, 0xbf, 0x6a, 0x29, 0x98, 0x07, 0x53, 0xa9, 0x60,
	0x9c, 0xfd, 0x5d, 0xf8, 0x1c, 0x52, 0x8d, 0x6a, 0xc3, 0xd1, 0x93, 0xf5, 0x78, 0x15, 0xbf, 0x3d,
	0x0e, 0x58, 0xb2, 0xfd, 0x3d, 0xc3, 0x59, 0x76, 0x1b, 0x77, 0xe3, 0x7e, 0xc8, 0xeb, 0x70, 0xfa,
	0x1e, 0x00, 0x00, 0xff, 0xff, 0x44, 0x75, 0x44, 0x0f, 0x3f, 0x02, 0x00, 0x00,
}