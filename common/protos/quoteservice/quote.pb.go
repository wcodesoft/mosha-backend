// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.27.0
// source: quoteservice/quote.proto

package quoteservice

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// The Quote message
type Quote struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	AuthorId  string `protobuf:"bytes,2,opt,name=author_id,json=authorId,proto3" json:"author_id,omitempty"`
	Text      string `protobuf:"bytes,3,opt,name=text,proto3" json:"text,omitempty"`
	Timestamp int64  `protobuf:"varint,4,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *Quote) Reset() {
	*x = Quote{}
	if protoimpl.UnsafeEnabled {
		mi := &file_quoteservice_quote_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Quote) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Quote) ProtoMessage() {}

func (x *Quote) ProtoReflect() protoreflect.Message {
	mi := &file_quoteservice_quote_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Quote.ProtoReflect.Descriptor instead.
func (*Quote) Descriptor() ([]byte, []int) {
	return file_quoteservice_quote_proto_rawDescGZIP(), []int{0}
}

func (x *Quote) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Quote) GetAuthorId() string {
	if x != nil {
		return x.AuthorId
	}
	return ""
}

func (x *Quote) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *Quote) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

// The GetQuoteRequest message
type GetQuoteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetQuoteRequest) Reset() {
	*x = GetQuoteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_quoteservice_quote_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetQuoteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetQuoteRequest) ProtoMessage() {}

func (x *GetQuoteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_quoteservice_quote_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetQuoteRequest.ProtoReflect.Descriptor instead.
func (*GetQuoteRequest) Descriptor() ([]byte, []int) {
	return file_quoteservice_quote_proto_rawDescGZIP(), []int{1}
}

func (x *GetQuoteRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

// The CreateQuoteRequest message
type CreateQuoteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Quote *Quote `protobuf:"bytes,1,opt,name=quote,proto3" json:"quote,omitempty"`
}

func (x *CreateQuoteRequest) Reset() {
	*x = CreateQuoteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_quoteservice_quote_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateQuoteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateQuoteRequest) ProtoMessage() {}

func (x *CreateQuoteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_quoteservice_quote_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateQuoteRequest.ProtoReflect.Descriptor instead.
func (*CreateQuoteRequest) Descriptor() ([]byte, []int) {
	return file_quoteservice_quote_proto_rawDescGZIP(), []int{2}
}

func (x *CreateQuoteRequest) GetQuote() *Quote {
	if x != nil {
		return x.Quote
	}
	return nil
}

// The UpdateQuoteRequest message
type UpdateQuoteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Quote *Quote `protobuf:"bytes,1,opt,name=quote,proto3" json:"quote,omitempty"`
}

func (x *UpdateQuoteRequest) Reset() {
	*x = UpdateQuoteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_quoteservice_quote_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateQuoteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateQuoteRequest) ProtoMessage() {}

func (x *UpdateQuoteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_quoteservice_quote_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateQuoteRequest.ProtoReflect.Descriptor instead.
func (*UpdateQuoteRequest) Descriptor() ([]byte, []int) {
	return file_quoteservice_quote_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateQuoteRequest) GetQuote() *Quote {
	if x != nil {
		return x.Quote
	}
	return nil
}

// The DeleteQuoteRequest message
type DeleteQuoteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteQuoteRequest) Reset() {
	*x = DeleteQuoteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_quoteservice_quote_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteQuoteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteQuoteRequest) ProtoMessage() {}

func (x *DeleteQuoteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_quoteservice_quote_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteQuoteRequest.ProtoReflect.Descriptor instead.
func (*DeleteQuoteRequest) Descriptor() ([]byte, []int) {
	return file_quoteservice_quote_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteQuoteRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

// The GetQuoteByAuthorRequest message
type GetQuotesByAuthorRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AuthorId string `protobuf:"bytes,1,opt,name=author_id,json=authorId,proto3" json:"author_id,omitempty"`
}

func (x *GetQuotesByAuthorRequest) Reset() {
	*x = GetQuotesByAuthorRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_quoteservice_quote_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetQuotesByAuthorRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetQuotesByAuthorRequest) ProtoMessage() {}

func (x *GetQuotesByAuthorRequest) ProtoReflect() protoreflect.Message {
	mi := &file_quoteservice_quote_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetQuotesByAuthorRequest.ProtoReflect.Descriptor instead.
func (*GetQuotesByAuthorRequest) Descriptor() ([]byte, []int) {
	return file_quoteservice_quote_proto_rawDescGZIP(), []int{5}
}

func (x *GetQuotesByAuthorRequest) GetAuthorId() string {
	if x != nil {
		return x.AuthorId
	}
	return ""
}

type DeleteQuotesByAuthorRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AuthorId string `protobuf:"bytes,1,opt,name=author_id,json=authorId,proto3" json:"author_id,omitempty"`
}

func (x *DeleteQuotesByAuthorRequest) Reset() {
	*x = DeleteQuotesByAuthorRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_quoteservice_quote_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteQuotesByAuthorRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteQuotesByAuthorRequest) ProtoMessage() {}

func (x *DeleteQuotesByAuthorRequest) ProtoReflect() protoreflect.Message {
	mi := &file_quoteservice_quote_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteQuotesByAuthorRequest.ProtoReflect.Descriptor instead.
func (*DeleteQuotesByAuthorRequest) Descriptor() ([]byte, []int) {
	return file_quoteservice_quote_proto_rawDescGZIP(), []int{6}
}

func (x *DeleteQuotesByAuthorRequest) GetAuthorId() string {
	if x != nil {
		return x.AuthorId
	}
	return ""
}

// The ListQuotesResponse message
type ListQuotesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Quotes []*Quote `protobuf:"bytes,1,rep,name=quotes,proto3" json:"quotes,omitempty"`
}

func (x *ListQuotesResponse) Reset() {
	*x = ListQuotesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_quoteservice_quote_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListQuotesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListQuotesResponse) ProtoMessage() {}

func (x *ListQuotesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_quoteservice_quote_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListQuotesResponse.ProtoReflect.Descriptor instead.
func (*ListQuotesResponse) Descriptor() ([]byte, []int) {
	return file_quoteservice_quote_proto_rawDescGZIP(), []int{7}
}

func (x *ListQuotesResponse) GetQuotes() []*Quote {
	if x != nil {
		return x.Quotes
	}
	return nil
}

// The DeleteQuoteResponse message
type DeleteQuoteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *DeleteQuoteResponse) Reset() {
	*x = DeleteQuoteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_quoteservice_quote_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteQuoteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteQuoteResponse) ProtoMessage() {}

func (x *DeleteQuoteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_quoteservice_quote_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteQuoteResponse.ProtoReflect.Descriptor instead.
func (*DeleteQuoteResponse) Descriptor() ([]byte, []int) {
	return file_quoteservice_quote_proto_rawDescGZIP(), []int{8}
}

func (x *DeleteQuoteResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

var File_quoteservice_quote_proto protoreflect.FileDescriptor

var file_quoteservice_quote_proto_rawDesc = []byte{
	0x0a, 0x18, 0x71, 0x75, 0x6f, 0x74, 0x65, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x71,
	0x75, 0x6f, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x71, 0x75, 0x6f, 0x74,
	0x65, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x66, 0x0a, 0x05, 0x51, 0x75, 0x6f, 0x74, 0x65, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1b,
	0x0a, 0x09, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74,
	0x65, 0x78, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x12,
	0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x22, 0x21, 0x0a,
	0x0f, 0x47, 0x65, 0x74, 0x51, 0x75, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x22, 0x3f, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x51, 0x75, 0x6f, 0x74, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x29, 0x0a, 0x05, 0x71, 0x75, 0x6f, 0x74, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x71, 0x75, 0x6f, 0x74, 0x65, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x51, 0x75, 0x6f, 0x74, 0x65, 0x52, 0x05, 0x71, 0x75, 0x6f, 0x74,
	0x65, 0x22, 0x3f, 0x0a, 0x12, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x51, 0x75, 0x6f, 0x74, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x29, 0x0a, 0x05, 0x71, 0x75, 0x6f, 0x74, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x71, 0x75, 0x6f, 0x74, 0x65, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x51, 0x75, 0x6f, 0x74, 0x65, 0x52, 0x05, 0x71, 0x75, 0x6f,
	0x74, 0x65, 0x22, 0x24, 0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x51, 0x75, 0x6f, 0x74,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x37, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x51,
	0x75, 0x6f, 0x74, 0x65, 0x73, 0x42, 0x79, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x49,
	0x64, 0x22, 0x3a, 0x0a, 0x1b, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x51, 0x75, 0x6f, 0x74, 0x65,
	0x73, 0x42, 0x79, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1b, 0x0a, 0x09, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x49, 0x64, 0x22, 0x41, 0x0a,
	0x12, 0x4c, 0x69, 0x73, 0x74, 0x51, 0x75, 0x6f, 0x74, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x2b, 0x0a, 0x06, 0x71, 0x75, 0x6f, 0x74, 0x65, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x71, 0x75, 0x6f, 0x74, 0x65, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x51, 0x75, 0x6f, 0x74, 0x65, 0x52, 0x06, 0x71, 0x75, 0x6f, 0x74, 0x65, 0x73,
	0x22, 0x2f, 0x0a, 0x13, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x51, 0x75, 0x6f, 0x74, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x32, 0xcc, 0x04, 0x0a, 0x0c, 0x51, 0x75, 0x6f, 0x74, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x40, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x51, 0x75, 0x6f, 0x74, 0x65, 0x12, 0x1d,
	0x2e, 0x71, 0x75, 0x6f, 0x74, 0x65, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x47, 0x65,
	0x74, 0x51, 0x75, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e,
	0x71, 0x75, 0x6f, 0x74, 0x65, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x51, 0x75, 0x6f,
	0x74, 0x65, 0x22, 0x00, 0x12, 0x48, 0x0a, 0x0a, 0x4c, 0x69, 0x73, 0x74, 0x51, 0x75, 0x6f, 0x74,
	0x65, 0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x20, 0x2e, 0x71, 0x75, 0x6f,
	0x74, 0x65, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x51, 0x75,
	0x6f, 0x74, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x46,
	0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x51, 0x75, 0x6f, 0x74, 0x65, 0x12, 0x20, 0x2e,
	0x71, 0x75, 0x6f, 0x74, 0x65, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x51, 0x75, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x13, 0x2e, 0x71, 0x75, 0x6f, 0x74, 0x65, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x51,
	0x75, 0x6f, 0x74, 0x65, 0x22, 0x00, 0x12, 0x46, 0x0a, 0x0b, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x51, 0x75, 0x6f, 0x74, 0x65, 0x12, 0x20, 0x2e, 0x71, 0x75, 0x6f, 0x74, 0x65, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x51, 0x75, 0x6f, 0x74, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x71, 0x75, 0x6f, 0x74, 0x65, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x51, 0x75, 0x6f, 0x74, 0x65, 0x22, 0x00, 0x12, 0x54,
	0x0a, 0x0b, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x51, 0x75, 0x6f, 0x74, 0x65, 0x12, 0x20, 0x2e,
	0x71, 0x75, 0x6f, 0x74, 0x65, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x51, 0x75, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x21, 0x2e, 0x71, 0x75, 0x6f, 0x74, 0x65, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x51, 0x75, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x5f, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x51, 0x75, 0x6f, 0x74, 0x65,
	0x73, 0x42, 0x79, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x12, 0x26, 0x2e, 0x71, 0x75, 0x6f, 0x74,
	0x65, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x51, 0x75, 0x6f, 0x74,
	0x65, 0x73, 0x42, 0x79, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x20, 0x2e, 0x71, 0x75, 0x6f, 0x74, 0x65, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x4c, 0x69, 0x73, 0x74, 0x51, 0x75, 0x6f, 0x74, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x69, 0x0a, 0x17, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41,
	0x6c, 0x6c, 0x51, 0x75, 0x6f, 0x74, 0x65, 0x73, 0x42, 0x79, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72,
	0x12, 0x29, 0x2e, 0x71, 0x75, 0x6f, 0x74, 0x65, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x51, 0x75, 0x6f, 0x74, 0x65, 0x73, 0x42, 0x79, 0x41, 0x75,
	0x74, 0x68, 0x6f, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x71, 0x75,
	0x6f, 0x74, 0x65, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x51, 0x75, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x42, 0x3f, 0x5a, 0x3d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x77,
	0x63, 0x6f, 0x64, 0x65, 0x73, 0x6f, 0x66, 0x74, 0x2f, 0x6d, 0x6f, 0x73, 0x68, 0x61, 0x2d, 0x62,
	0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x71, 0x75, 0x6f, 0x74, 0x65, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_quoteservice_quote_proto_rawDescOnce sync.Once
	file_quoteservice_quote_proto_rawDescData = file_quoteservice_quote_proto_rawDesc
)

func file_quoteservice_quote_proto_rawDescGZIP() []byte {
	file_quoteservice_quote_proto_rawDescOnce.Do(func() {
		file_quoteservice_quote_proto_rawDescData = protoimpl.X.CompressGZIP(file_quoteservice_quote_proto_rawDescData)
	})
	return file_quoteservice_quote_proto_rawDescData
}

var file_quoteservice_quote_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_quoteservice_quote_proto_goTypes = []interface{}{
	(*Quote)(nil),                       // 0: quoteservice.Quote
	(*GetQuoteRequest)(nil),             // 1: quoteservice.GetQuoteRequest
	(*CreateQuoteRequest)(nil),          // 2: quoteservice.CreateQuoteRequest
	(*UpdateQuoteRequest)(nil),          // 3: quoteservice.UpdateQuoteRequest
	(*DeleteQuoteRequest)(nil),          // 4: quoteservice.DeleteQuoteRequest
	(*GetQuotesByAuthorRequest)(nil),    // 5: quoteservice.GetQuotesByAuthorRequest
	(*DeleteQuotesByAuthorRequest)(nil), // 6: quoteservice.DeleteQuotesByAuthorRequest
	(*ListQuotesResponse)(nil),          // 7: quoteservice.ListQuotesResponse
	(*DeleteQuoteResponse)(nil),         // 8: quoteservice.DeleteQuoteResponse
	(*emptypb.Empty)(nil),               // 9: google.protobuf.Empty
}
var file_quoteservice_quote_proto_depIdxs = []int32{
	0,  // 0: quoteservice.CreateQuoteRequest.quote:type_name -> quoteservice.Quote
	0,  // 1: quoteservice.UpdateQuoteRequest.quote:type_name -> quoteservice.Quote
	0,  // 2: quoteservice.ListQuotesResponse.quotes:type_name -> quoteservice.Quote
	1,  // 3: quoteservice.QuoteService.GetQuote:input_type -> quoteservice.GetQuoteRequest
	9,  // 4: quoteservice.QuoteService.ListQuotes:input_type -> google.protobuf.Empty
	2,  // 5: quoteservice.QuoteService.CreateQuote:input_type -> quoteservice.CreateQuoteRequest
	3,  // 6: quoteservice.QuoteService.UpdateQuote:input_type -> quoteservice.UpdateQuoteRequest
	4,  // 7: quoteservice.QuoteService.DeleteQuote:input_type -> quoteservice.DeleteQuoteRequest
	5,  // 8: quoteservice.QuoteService.GetQuotesByAuthor:input_type -> quoteservice.GetQuotesByAuthorRequest
	6,  // 9: quoteservice.QuoteService.DeleteAllQuotesByAuthor:input_type -> quoteservice.DeleteQuotesByAuthorRequest
	0,  // 10: quoteservice.QuoteService.GetQuote:output_type -> quoteservice.Quote
	7,  // 11: quoteservice.QuoteService.ListQuotes:output_type -> quoteservice.ListQuotesResponse
	0,  // 12: quoteservice.QuoteService.CreateQuote:output_type -> quoteservice.Quote
	0,  // 13: quoteservice.QuoteService.UpdateQuote:output_type -> quoteservice.Quote
	8,  // 14: quoteservice.QuoteService.DeleteQuote:output_type -> quoteservice.DeleteQuoteResponse
	7,  // 15: quoteservice.QuoteService.GetQuotesByAuthor:output_type -> quoteservice.ListQuotesResponse
	8,  // 16: quoteservice.QuoteService.DeleteAllQuotesByAuthor:output_type -> quoteservice.DeleteQuoteResponse
	10, // [10:17] is the sub-list for method output_type
	3,  // [3:10] is the sub-list for method input_type
	3,  // [3:3] is the sub-list for extension type_name
	3,  // [3:3] is the sub-list for extension extendee
	0,  // [0:3] is the sub-list for field type_name
}

func init() { file_quoteservice_quote_proto_init() }
func file_quoteservice_quote_proto_init() {
	if File_quoteservice_quote_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_quoteservice_quote_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Quote); i {
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
		file_quoteservice_quote_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetQuoteRequest); i {
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
		file_quoteservice_quote_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateQuoteRequest); i {
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
		file_quoteservice_quote_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateQuoteRequest); i {
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
		file_quoteservice_quote_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteQuoteRequest); i {
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
		file_quoteservice_quote_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetQuotesByAuthorRequest); i {
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
		file_quoteservice_quote_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteQuotesByAuthorRequest); i {
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
		file_quoteservice_quote_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListQuotesResponse); i {
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
		file_quoteservice_quote_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteQuoteResponse); i {
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
			RawDescriptor: file_quoteservice_quote_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_quoteservice_quote_proto_goTypes,
		DependencyIndexes: file_quoteservice_quote_proto_depIdxs,
		MessageInfos:      file_quoteservice_quote_proto_msgTypes,
	}.Build()
	File_quoteservice_quote_proto = out.File
	file_quoteservice_quote_proto_rawDesc = nil
	file_quoteservice_quote_proto_goTypes = nil
	file_quoteservice_quote_proto_depIdxs = nil
}