// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.21.0
// 	protoc        v3.2.0
// source: xc.proto

package xcpb

import (
	proto "github.com/golang/protobuf/proto"
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

type PublicKeyAlgorithm int32

const (
	PublicKeyAlgorithm_UNKNOWN PublicKeyAlgorithm = 0
	PublicKeyAlgorithm_NACL    PublicKeyAlgorithm = 1
)

// Enum value maps for PublicKeyAlgorithm.
var (
	PublicKeyAlgorithm_name = map[int32]string{
		0: "UNKNOWN",
		1: "NACL",
	}
	PublicKeyAlgorithm_value = map[string]int32{
		"UNKNOWN": 0,
		"NACL":    1,
	}
)

func (x PublicKeyAlgorithm) Enum() *PublicKeyAlgorithm {
	p := new(PublicKeyAlgorithm)
	*p = x
	return p
}

func (x PublicKeyAlgorithm) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PublicKeyAlgorithm) Descriptor() protoreflect.EnumDescriptor {
	return file_xc_proto_enumTypes[0].Descriptor()
}

func (PublicKeyAlgorithm) Type() protoreflect.EnumType {
	return &file_xc_proto_enumTypes[0]
}

func (x PublicKeyAlgorithm) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PublicKeyAlgorithm.Descriptor instead.
func (PublicKeyAlgorithm) EnumDescriptor() ([]byte, []int) {
	return file_xc_proto_rawDescGZIP(), []int{0}
}

type Header struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sender     string            `protobuf:"bytes,1,opt,name=Sender,proto3" json:"Sender,omitempty"`
	Recipients map[string][]byte `protobuf:"bytes,2,rep,name=Recipients,proto3" json:"Recipients,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Header) Reset() {
	*x = Header{}
	if protoimpl.UnsafeEnabled {
		mi := &file_xc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Header) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Header) ProtoMessage() {}

func (x *Header) ProtoReflect() protoreflect.Message {
	mi := &file_xc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Header.ProtoReflect.Descriptor instead.
func (*Header) Descriptor() ([]byte, []int) {
	return file_xc_proto_rawDescGZIP(), []int{0}
}

func (x *Header) GetSender() string {
	if x != nil {
		return x.Sender
	}
	return ""
}

func (x *Header) GetRecipients() map[string][]byte {
	if x != nil {
		return x.Recipients
	}
	return nil
}

type Chunk struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Body []byte `protobuf:"bytes,1,opt,name=Body,proto3" json:"Body,omitempty"`
}

func (x *Chunk) Reset() {
	*x = Chunk{}
	if protoimpl.UnsafeEnabled {
		mi := &file_xc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Chunk) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Chunk) ProtoMessage() {}

func (x *Chunk) ProtoReflect() protoreflect.Message {
	mi := &file_xc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Chunk.ProtoReflect.Descriptor instead.
func (*Chunk) Descriptor() ([]byte, []int) {
	return file_xc_proto_rawDescGZIP(), []int{1}
}

func (x *Chunk) GetBody() []byte {
	if x != nil {
		return x.Body
	}
	return nil
}

type Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Version    uint32   `protobuf:"varint,1,opt,name=Version,proto3" json:"Version,omitempty"`
	Header     *Header  `protobuf:"bytes,2,opt,name=Header,proto3" json:"Header,omitempty"`
	Chunks     []*Chunk `protobuf:"bytes,3,rep,name=Chunks,proto3" json:"Chunks,omitempty"`
	Compressed bool     `protobuf:"varint,4,opt,name=Compressed,proto3" json:"Compressed,omitempty"`
}

func (x *Message) Reset() {
	*x = Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_xc_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
	mi := &file_xc_proto_msgTypes[2]
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
	return file_xc_proto_rawDescGZIP(), []int{2}
}

func (x *Message) GetVersion() uint32 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *Message) GetHeader() *Header {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *Message) GetChunks() []*Chunk {
	if x != nil {
		return x.Chunks
	}
	return nil
}

func (x *Message) GetCompressed() bool {
	if x != nil {
		return x.Compressed
	}
	return false
}

type Identity struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name    string `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	Comment string `protobuf:"bytes,2,opt,name=Comment,proto3" json:"Comment,omitempty"`
	Email   string `protobuf:"bytes,3,opt,name=Email,proto3" json:"Email,omitempty"`
}

func (x *Identity) Reset() {
	*x = Identity{}
	if protoimpl.UnsafeEnabled {
		mi := &file_xc_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Identity) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Identity) ProtoMessage() {}

func (x *Identity) ProtoReflect() protoreflect.Message {
	mi := &file_xc_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Identity.ProtoReflect.Descriptor instead.
func (*Identity) Descriptor() ([]byte, []int) {
	return file_xc_proto_rawDescGZIP(), []int{3}
}

func (x *Identity) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Identity) GetComment() string {
	if x != nil {
		return x.Comment
	}
	return ""
}

func (x *Identity) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type PublicKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PubKeyAlgo   PublicKeyAlgorithm `protobuf:"varint,1,opt,name=PubKeyAlgo,proto3,enum=xcpb.PublicKeyAlgorithm" json:"PubKeyAlgo,omitempty"`
	CreationTime uint64             `protobuf:"varint,2,opt,name=CreationTime,proto3" json:"CreationTime,omitempty"`
	PublicKey    []byte             `protobuf:"bytes,3,opt,name=PublicKey,proto3" json:"PublicKey,omitempty"`
	Identity     *Identity          `protobuf:"bytes,4,opt,name=Identity,proto3" json:"Identity,omitempty"`
	Fingerprint  string             `protobuf:"bytes,5,opt,name=Fingerprint,proto3" json:"Fingerprint,omitempty"`
}

func (x *PublicKey) Reset() {
	*x = PublicKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_xc_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PublicKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublicKey) ProtoMessage() {}

func (x *PublicKey) ProtoReflect() protoreflect.Message {
	mi := &file_xc_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PublicKey.ProtoReflect.Descriptor instead.
func (*PublicKey) Descriptor() ([]byte, []int) {
	return file_xc_proto_rawDescGZIP(), []int{4}
}

func (x *PublicKey) GetPubKeyAlgo() PublicKeyAlgorithm {
	if x != nil {
		return x.PubKeyAlgo
	}
	return PublicKeyAlgorithm_UNKNOWN
}

func (x *PublicKey) GetCreationTime() uint64 {
	if x != nil {
		return x.CreationTime
	}
	return 0
}

func (x *PublicKey) GetPublicKey() []byte {
	if x != nil {
		return x.PublicKey
	}
	return nil
}

func (x *PublicKey) GetIdentity() *Identity {
	if x != nil {
		return x.Identity
	}
	return nil
}

func (x *PublicKey) GetFingerprint() string {
	if x != nil {
		return x.Fingerprint
	}
	return ""
}

type PrivateKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PublicKey  *PublicKey `protobuf:"bytes,1,opt,name=PublicKey,proto3" json:"PublicKey,omitempty"`
	Ciphertext []byte     `protobuf:"bytes,2,opt,name=Ciphertext,proto3" json:"Ciphertext,omitempty"`
	Nonce      []byte     `protobuf:"bytes,3,opt,name=Nonce,proto3" json:"Nonce,omitempty"`
	Salt       []byte     `protobuf:"bytes,4,opt,name=Salt,proto3" json:"Salt,omitempty"`
}

func (x *PrivateKey) Reset() {
	*x = PrivateKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_xc_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PrivateKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PrivateKey) ProtoMessage() {}

func (x *PrivateKey) ProtoReflect() protoreflect.Message {
	mi := &file_xc_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PrivateKey.ProtoReflect.Descriptor instead.
func (*PrivateKey) Descriptor() ([]byte, []int) {
	return file_xc_proto_rawDescGZIP(), []int{5}
}

func (x *PrivateKey) GetPublicKey() *PublicKey {
	if x != nil {
		return x.PublicKey
	}
	return nil
}

func (x *PrivateKey) GetCiphertext() []byte {
	if x != nil {
		return x.Ciphertext
	}
	return nil
}

func (x *PrivateKey) GetNonce() []byte {
	if x != nil {
		return x.Nonce
	}
	return nil
}

func (x *PrivateKey) GetSalt() []byte {
	if x != nil {
		return x.Salt
	}
	return nil
}

type Secring struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Version     uint32        `protobuf:"varint,1,opt,name=Version,proto3" json:"Version,omitempty"`
	PrivateKeys []*PrivateKey `protobuf:"bytes,2,rep,name=PrivateKeys,proto3" json:"PrivateKeys,omitempty"`
}

func (x *Secring) Reset() {
	*x = Secring{}
	if protoimpl.UnsafeEnabled {
		mi := &file_xc_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Secring) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Secring) ProtoMessage() {}

func (x *Secring) ProtoReflect() protoreflect.Message {
	mi := &file_xc_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Secring.ProtoReflect.Descriptor instead.
func (*Secring) Descriptor() ([]byte, []int) {
	return file_xc_proto_rawDescGZIP(), []int{6}
}

func (x *Secring) GetVersion() uint32 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *Secring) GetPrivateKeys() []*PrivateKey {
	if x != nil {
		return x.PrivateKeys
	}
	return nil
}

type Pubring struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Version    uint32       `protobuf:"varint,1,opt,name=Version,proto3" json:"Version,omitempty"`
	PublicKeys []*PublicKey `protobuf:"bytes,2,rep,name=PublicKeys,proto3" json:"PublicKeys,omitempty"`
}

func (x *Pubring) Reset() {
	*x = Pubring{}
	if protoimpl.UnsafeEnabled {
		mi := &file_xc_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Pubring) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Pubring) ProtoMessage() {}

func (x *Pubring) ProtoReflect() protoreflect.Message {
	mi := &file_xc_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Pubring.ProtoReflect.Descriptor instead.
func (*Pubring) Descriptor() ([]byte, []int) {
	return file_xc_proto_rawDescGZIP(), []int{7}
}

func (x *Pubring) GetVersion() uint32 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *Pubring) GetPublicKeys() []*PublicKey {
	if x != nil {
		return x.PublicKeys
	}
	return nil
}

var File_xc_proto protoreflect.FileDescriptor

var file_xc_proto_rawDesc = []byte{
	0x0a, 0x08, 0x78, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x78, 0x63, 0x70, 0x62,
	0x22, 0x9d, 0x01, 0x0a, 0x06, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x53,
	0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x53, 0x65, 0x6e,
	0x64, 0x65, 0x72, 0x12, 0x3c, 0x0a, 0x0a, 0x52, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74,
	0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x78, 0x63, 0x70, 0x62, 0x2e, 0x48,
	0x65, 0x61, 0x64, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x73,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0a, 0x52, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74,
	0x73, 0x1a, 0x3d, 0x0a, 0x0f, 0x52, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01,
	0x22, 0x1b, 0x0a, 0x05, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x12, 0x12, 0x0a, 0x04, 0x42, 0x6f, 0x64,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x42, 0x6f, 0x64, 0x79, 0x22, 0x8e, 0x01,
	0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x56, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x56, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x12, 0x24, 0x0a, 0x06, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x78, 0x63, 0x70, 0x62, 0x2e, 0x48, 0x65, 0x61, 0x64, 0x65,
	0x72, 0x52, 0x06, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x23, 0x0a, 0x06, 0x43, 0x68, 0x75,
	0x6e, 0x6b, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x78, 0x63, 0x70, 0x62,
	0x2e, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x52, 0x06, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x73, 0x12, 0x1e,
	0x0a, 0x0a, 0x43, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x0a, 0x43, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x65, 0x64, 0x22, 0x4e,
	0x0a, 0x08, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x45, 0x6d, 0x61, 0x69,
	0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x22, 0xd5,
	0x01, 0x0a, 0x09, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x12, 0x38, 0x0a, 0x0a,
	0x50, 0x75, 0x62, 0x4b, 0x65, 0x79, 0x41, 0x6c, 0x67, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x18, 0x2e, 0x78, 0x63, 0x70, 0x62, 0x2e, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65,
	0x79, 0x41, 0x6c, 0x67, 0x6f, 0x72, 0x69, 0x74, 0x68, 0x6d, 0x52, 0x0a, 0x50, 0x75, 0x62, 0x4b,
	0x65, 0x79, 0x41, 0x6c, 0x67, 0x6f, 0x12, 0x22, 0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0c, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x50, 0x75,
	0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x50,
	0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x12, 0x2a, 0x0a, 0x08, 0x49, 0x64, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x78, 0x63, 0x70,
	0x62, 0x2e, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x08, 0x49, 0x64, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x12, 0x20, 0x0a, 0x0b, 0x46, 0x69, 0x6e, 0x67, 0x65, 0x72, 0x70, 0x72,
	0x69, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x46, 0x69, 0x6e, 0x67, 0x65,
	0x72, 0x70, 0x72, 0x69, 0x6e, 0x74, 0x22, 0x85, 0x01, 0x0a, 0x0a, 0x50, 0x72, 0x69, 0x76, 0x61,
	0x74, 0x65, 0x4b, 0x65, 0x79, 0x12, 0x2d, 0x0a, 0x09, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x78, 0x63, 0x70, 0x62, 0x2e,
	0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x52, 0x09, 0x50, 0x75, 0x62, 0x6c, 0x69,
	0x63, 0x4b, 0x65, 0x79, 0x12, 0x1e, 0x0a, 0x0a, 0x43, 0x69, 0x70, 0x68, 0x65, 0x72, 0x74, 0x65,
	0x78, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0a, 0x43, 0x69, 0x70, 0x68, 0x65, 0x72,
	0x74, 0x65, 0x78, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x4e, 0x6f, 0x6e, 0x63, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x05, 0x4e, 0x6f, 0x6e, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x53, 0x61,
	0x6c, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x53, 0x61, 0x6c, 0x74, 0x22, 0x57,
	0x0a, 0x07, 0x53, 0x65, 0x63, 0x72, 0x69, 0x6e, 0x67, 0x12, 0x18, 0x0a, 0x07, 0x56, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x56, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x12, 0x32, 0x0a, 0x0b, 0x50, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x4b, 0x65,
	0x79, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x78, 0x63, 0x70, 0x62, 0x2e,
	0x50, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x4b, 0x65, 0x79, 0x52, 0x0b, 0x50, 0x72, 0x69, 0x76,
	0x61, 0x74, 0x65, 0x4b, 0x65, 0x79, 0x73, 0x22, 0x54, 0x0a, 0x07, 0x50, 0x75, 0x62, 0x72, 0x69,
	0x6e, 0x67, 0x12, 0x18, 0x0a, 0x07, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x07, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x2f, 0x0a, 0x0a,
	0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x0f, 0x2e, 0x78, 0x63, 0x70, 0x62, 0x2e, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65,
	0x79, 0x52, 0x0a, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x73, 0x2a, 0x2b, 0x0a,
	0x12, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x41, 0x6c, 0x67, 0x6f, 0x72, 0x69,
	0x74, 0x68, 0x6d, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00,
	0x12, 0x08, 0x0a, 0x04, 0x4e, 0x41, 0x43, 0x4c, 0x10, 0x01, 0x42, 0x08, 0x5a, 0x06, 0x2e, 0x3b,
	0x78, 0x63, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_xc_proto_rawDescOnce sync.Once
	file_xc_proto_rawDescData = file_xc_proto_rawDesc
)

func file_xc_proto_rawDescGZIP() []byte {
	file_xc_proto_rawDescOnce.Do(func() {
		file_xc_proto_rawDescData = protoimpl.X.CompressGZIP(file_xc_proto_rawDescData)
	})
	return file_xc_proto_rawDescData
}

var file_xc_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_xc_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_xc_proto_goTypes = []interface{}{
	(PublicKeyAlgorithm)(0), // 0: xcpb.PublicKeyAlgorithm
	(*Header)(nil),          // 1: xcpb.Header
	(*Chunk)(nil),           // 2: xcpb.Chunk
	(*Message)(nil),         // 3: xcpb.Message
	(*Identity)(nil),        // 4: xcpb.Identity
	(*PublicKey)(nil),       // 5: xcpb.PublicKey
	(*PrivateKey)(nil),      // 6: xcpb.PrivateKey
	(*Secring)(nil),         // 7: xcpb.Secring
	(*Pubring)(nil),         // 8: xcpb.Pubring
	nil,                     // 9: xcpb.Header.RecipientsEntry
}
var file_xc_proto_depIdxs = []int32{
	9, // 0: xcpb.Header.Recipients:type_name -> xcpb.Header.RecipientsEntry
	1, // 1: xcpb.Message.Header:type_name -> xcpb.Header
	2, // 2: xcpb.Message.Chunks:type_name -> xcpb.Chunk
	0, // 3: xcpb.PublicKey.PubKeyAlgo:type_name -> xcpb.PublicKeyAlgorithm
	4, // 4: xcpb.PublicKey.Identity:type_name -> xcpb.Identity
	5, // 5: xcpb.PrivateKey.PublicKey:type_name -> xcpb.PublicKey
	6, // 6: xcpb.Secring.PrivateKeys:type_name -> xcpb.PrivateKey
	5, // 7: xcpb.Pubring.PublicKeys:type_name -> xcpb.PublicKey
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_xc_proto_init() }
func file_xc_proto_init() {
	if File_xc_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_xc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Header); i {
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
		file_xc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Chunk); i {
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
		file_xc_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
		file_xc_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Identity); i {
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
		file_xc_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PublicKey); i {
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
		file_xc_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PrivateKey); i {
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
		file_xc_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Secring); i {
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
		file_xc_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Pubring); i {
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
			RawDescriptor: file_xc_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_xc_proto_goTypes,
		DependencyIndexes: file_xc_proto_depIdxs,
		EnumInfos:         file_xc_proto_enumTypes,
		MessageInfos:      file_xc_proto_msgTypes,
	}.Build()
	File_xc_proto = out.File
	file_xc_proto_rawDesc = nil
	file_xc_proto_goTypes = nil
	file_xc_proto_depIdxs = nil
}