// Code generated by protoc-gen-go. DO NOT EDIT.
// source: email.proto

package email

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
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

type EmailRequest struct {
	From                 string            `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	To                   []string          `protobuf:"bytes,2,rep,name=to,proto3" json:"to,omitempty"`
	Subject              string            `protobuf:"bytes,3,opt,name=subject,proto3" json:"subject,omitempty"`
	PlainText            string            `protobuf:"bytes,4,opt,name=plain_text,json=plainText,proto3" json:"plain_text,omitempty"`
	HtmlAlternate        string            `protobuf:"bytes,5,opt,name=html_alternate,json=htmlAlternate,proto3" json:"html_alternate,omitempty"`
	Attachments          []*Attachment     `protobuf:"bytes,6,rep,name=attachments,proto3" json:"attachments,omitempty"`
	Headers              map[string]string `protobuf:"bytes,7,rep,name=headers,proto3" json:"headers,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *EmailRequest) Reset()         { *m = EmailRequest{} }
func (m *EmailRequest) String() string { return proto.CompactTextString(m) }
func (*EmailRequest) ProtoMessage()    {}
func (*EmailRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6175298cb4ed6faa, []int{0}
}

func (m *EmailRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EmailRequest.Unmarshal(m, b)
}
func (m *EmailRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EmailRequest.Marshal(b, m, deterministic)
}
func (m *EmailRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EmailRequest.Merge(m, src)
}
func (m *EmailRequest) XXX_Size() int {
	return xxx_messageInfo_EmailRequest.Size(m)
}
func (m *EmailRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_EmailRequest.DiscardUnknown(m)
}

var xxx_messageInfo_EmailRequest proto.InternalMessageInfo

func (m *EmailRequest) GetFrom() string {
	if m != nil {
		return m.From
	}
	return ""
}

func (m *EmailRequest) GetTo() []string {
	if m != nil {
		return m.To
	}
	return nil
}

func (m *EmailRequest) GetSubject() string {
	if m != nil {
		return m.Subject
	}
	return ""
}

func (m *EmailRequest) GetPlainText() string {
	if m != nil {
		return m.PlainText
	}
	return ""
}

func (m *EmailRequest) GetHtmlAlternate() string {
	if m != nil {
		return m.HtmlAlternate
	}
	return ""
}

func (m *EmailRequest) GetAttachments() []*Attachment {
	if m != nil {
		return m.Attachments
	}
	return nil
}

func (m *EmailRequest) GetHeaders() map[string]string {
	if m != nil {
		return m.Headers
	}
	return nil
}

type Attachment struct {
	Filename             string   `protobuf:"bytes,1,opt,name=filename,proto3" json:"filename,omitempty"`
	Body                 []byte   `protobuf:"bytes,2,opt,name=body,proto3" json:"body,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Attachment) Reset()         { *m = Attachment{} }
func (m *Attachment) String() string { return proto.CompactTextString(m) }
func (*Attachment) ProtoMessage()    {}
func (*Attachment) Descriptor() ([]byte, []int) {
	return fileDescriptor_6175298cb4ed6faa, []int{1}
}

func (m *Attachment) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Attachment.Unmarshal(m, b)
}
func (m *Attachment) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Attachment.Marshal(b, m, deterministic)
}
func (m *Attachment) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Attachment.Merge(m, src)
}
func (m *Attachment) XXX_Size() int {
	return xxx_messageInfo_Attachment.Size(m)
}
func (m *Attachment) XXX_DiscardUnknown() {
	xxx_messageInfo_Attachment.DiscardUnknown(m)
}

var xxx_messageInfo_Attachment proto.InternalMessageInfo

func (m *Attachment) GetFilename() string {
	if m != nil {
		return m.Filename
	}
	return ""
}

func (m *Attachment) GetBody() []byte {
	if m != nil {
		return m.Body
	}
	return nil
}

type EmailResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EmailResponse) Reset()         { *m = EmailResponse{} }
func (m *EmailResponse) String() string { return proto.CompactTextString(m) }
func (*EmailResponse) ProtoMessage()    {}
func (*EmailResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6175298cb4ed6faa, []int{2}
}

func (m *EmailResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EmailResponse.Unmarshal(m, b)
}
func (m *EmailResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EmailResponse.Marshal(b, m, deterministic)
}
func (m *EmailResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EmailResponse.Merge(m, src)
}
func (m *EmailResponse) XXX_Size() int {
	return xxx_messageInfo_EmailResponse.Size(m)
}
func (m *EmailResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_EmailResponse.DiscardUnknown(m)
}

var xxx_messageInfo_EmailResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*EmailRequest)(nil), "email.EmailRequest")
	proto.RegisterMapType((map[string]string)(nil), "email.EmailRequest.HeadersEntry")
	proto.RegisterType((*Attachment)(nil), "email.Attachment")
	proto.RegisterType((*EmailResponse)(nil), "email.EmailResponse")
}

func init() { proto.RegisterFile("email.proto", fileDescriptor_6175298cb4ed6faa) }

var fileDescriptor_6175298cb4ed6faa = []byte{
	// 339 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x52, 0x4d, 0x6b, 0xc2, 0x40,
	0x10, 0xad, 0x89, 0xd1, 0x3a, 0x7e, 0xb4, 0xdd, 0x7a, 0x58, 0x84, 0x42, 0x08, 0x14, 0x3c, 0x45,
	0xd0, 0x4b, 0x11, 0x2f, 0x16, 0x84, 0x9e, 0xd3, 0x9e, 0x7a, 0x91, 0x8d, 0x8e, 0x4d, 0xda, 0x4d,
	0x36, 0xcd, 0x4e, 0x8a, 0xfe, 0xba, 0xfe, 0xb5, 0x92, 0x4d, 0xb4, 0x16, 0x7a, 0x7b, 0xef, 0xcd,
	0x7b, 0xcb, 0x9b, 0x61, 0xa1, 0x8b, 0x89, 0x88, 0xa5, 0x9f, 0xe5, 0x8a, 0x14, 0x73, 0x0c, 0xf1,
	0xbe, 0x2d, 0xe8, 0xad, 0x4a, 0x14, 0xe0, 0x67, 0x81, 0x9a, 0x18, 0x83, 0xe6, 0x2e, 0x57, 0x09,
	0x6f, 0xb8, 0x8d, 0x71, 0x27, 0x30, 0x98, 0x0d, 0xc0, 0x22, 0xc5, 0x2d, 0xd7, 0x1e, 0x77, 0x02,
	0x8b, 0x14, 0xe3, 0xd0, 0xd6, 0x45, 0xf8, 0x8e, 0x1b, 0xe2, 0xb6, 0xb1, 0x1d, 0x29, 0xbb, 0x03,
	0xc8, 0xa4, 0x88, 0xd3, 0x35, 0xe1, 0x9e, 0x78, 0xd3, 0x0c, 0x3b, 0x46, 0x79, 0xc1, 0x3d, 0xb1,
	0x7b, 0x18, 0x44, 0x94, 0xc8, 0xb5, 0x90, 0x84, 0x79, 0x2a, 0x08, 0xb9, 0x63, 0x2c, 0xfd, 0x52,
	0x5d, 0x1e, 0x45, 0x36, 0x83, 0xae, 0x20, 0x12, 0x9b, 0x28, 0xc1, 0x94, 0x34, 0x6f, 0xb9, 0xf6,
	0xb8, 0x3b, 0xbd, 0xf1, 0xab, 0xfa, 0xcb, 0xd3, 0x24, 0x38, 0x77, 0xb1, 0x39, 0xb4, 0x23, 0x14,
	0x5b, 0xcc, 0x35, 0x6f, 0x9b, 0x80, 0x5b, 0x07, 0xce, 0xd7, 0xf3, 0x9f, 0x2a, 0xcb, 0x2a, 0xa5,
	0xfc, 0x10, 0x1c, 0x03, 0xa3, 0x39, 0xf4, 0xce, 0x07, 0xec, 0x1a, 0xec, 0x0f, 0x3c, 0xd4, 0x37,
	0x28, 0x21, 0x1b, 0x82, 0xf3, 0x25, 0x64, 0x81, 0xdc, 0x32, 0x5a, 0x45, 0xe6, 0xd6, 0x43, 0xc3,
	0x5b, 0x00, 0xfc, 0x56, 0x62, 0x23, 0xb8, 0xdc, 0xc5, 0x12, 0x53, 0x91, 0x60, 0x1d, 0x3f, 0xf1,
	0xf2, 0xb4, 0xa1, 0xda, 0x1e, 0xcc, 0x13, 0xbd, 0xc0, 0x60, 0xef, 0x0a, 0xfa, 0x75, 0x3f, 0x9d,
	0xa9, 0x54, 0xe3, 0x74, 0x01, 0x8e, 0x11, 0xd8, 0x0c, 0x9a, 0xcf, 0x98, 0x6e, 0xd9, 0xed, 0x3f,
	0x6b, 0x8c, 0x86, 0x7f, 0xc5, 0x2a, 0xeb, 0x5d, 0x3c, 0x7a, 0xaf, 0xee, 0x5b, 0x4c, 0x51, 0x11,
	0xfa, 0x1b, 0x95, 0x4c, 0x64, 0x2c, 0x31, 0x56, 0x13, 0xdc, 0x8b, 0x24, 0x93, 0xa8, 0x27, 0x26,
	0x13, 0xb6, 0xcc, 0x07, 0x98, 0xfd, 0x04, 0x00, 0x00, 0xff, 0xff, 0xb1, 0x0a, 0xf4, 0x75, 0x0f,
	0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// EmailClient is the client API for Email service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type EmailClient interface {
	Send(ctx context.Context, in *EmailRequest, opts ...grpc.CallOption) (*EmailResponse, error)
}

type emailClient struct {
	cc *grpc.ClientConn
}

func NewEmailClient(cc *grpc.ClientConn) EmailClient {
	return &emailClient{cc}
}

func (c *emailClient) Send(ctx context.Context, in *EmailRequest, opts ...grpc.CallOption) (*EmailResponse, error) {
	out := new(EmailResponse)
	err := c.cc.Invoke(ctx, "/email.Email/Send", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EmailServer is the server API for Email service.
type EmailServer interface {
	Send(context.Context, *EmailRequest) (*EmailResponse, error)
}

func RegisterEmailServer(s *grpc.Server, srv EmailServer) {
	s.RegisterService(&_Email_serviceDesc, srv)
}

func _Email_Send_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmailServer).Send(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/email.Email/Send",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmailServer).Send(ctx, req.(*EmailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Email_serviceDesc = grpc.ServiceDesc{
	ServiceName: "email.Email",
	HandlerType: (*EmailServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Send",
			Handler:    _Email_Send_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "email.proto",
}
