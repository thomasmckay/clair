// Code generated by protoc-gen-go. DO NOT EDIT.
// source: clair.proto

/*
Package clairpb is a generated protocol buffer package.

It is generated from these files:
	clair.proto

It has these top-level messages:
	Vulnerability
	Feature
	Ancestry
	LayersIntroducingVulnerabilty
	OrderedLayerName
	Layer
	Notification
	Page
	PostAncestryRequest
	PostAncestryResponse
	GetAncestryRequest
	GetAncestryResponse
	GetNotificationRequest
	GetNotificationResponse
	DeleteNotificationRequest
*/
package clairpb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"
import google_protobuf1 "github.com/golang/protobuf/ptypes/empty"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Vulnerability struct {
	Name            string     `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	NamespaceName   string     `protobuf:"bytes,2,opt,name=namespace_name,json=namespaceName" json:"namespace_name,omitempty"`
	Description     string     `protobuf:"bytes,3,opt,name=description" json:"description,omitempty"`
	Link            string     `protobuf:"bytes,4,opt,name=link" json:"link,omitempty"`
	Severity        string     `protobuf:"bytes,5,opt,name=severity" json:"severity,omitempty"`
	Metadata        string     `protobuf:"bytes,6,opt,name=metadata" json:"metadata,omitempty"`
	FixedBy         string     `protobuf:"bytes,7,opt,name=fixed_by,json=fixedBy" json:"fixed_by,omitempty"`
	FixedInFeatures []*Feature `protobuf:"bytes,8,rep,name=fixed_in_features,json=fixedInFeatures" json:"fixed_in_features,omitempty"`
}

func (m *Vulnerability) Reset()                    { *m = Vulnerability{} }
func (m *Vulnerability) String() string            { return proto.CompactTextString(m) }
func (*Vulnerability) ProtoMessage()               {}
func (*Vulnerability) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Vulnerability) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Vulnerability) GetNamespaceName() string {
	if m != nil {
		return m.NamespaceName
	}
	return ""
}

func (m *Vulnerability) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Vulnerability) GetLink() string {
	if m != nil {
		return m.Link
	}
	return ""
}

func (m *Vulnerability) GetSeverity() string {
	if m != nil {
		return m.Severity
	}
	return ""
}

func (m *Vulnerability) GetMetadata() string {
	if m != nil {
		return m.Metadata
	}
	return ""
}

func (m *Vulnerability) GetFixedBy() string {
	if m != nil {
		return m.FixedBy
	}
	return ""
}

func (m *Vulnerability) GetFixedInFeatures() []*Feature {
	if m != nil {
		return m.FixedInFeatures
	}
	return nil
}

type Feature struct {
	Name            string           `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	NamespaceName   string           `protobuf:"bytes,2,opt,name=namespace_name,json=namespaceName" json:"namespace_name,omitempty"`
	Version         string           `protobuf:"bytes,3,opt,name=version" json:"version,omitempty"`
	VersionFormat   string           `protobuf:"bytes,4,opt,name=version_format,json=versionFormat" json:"version_format,omitempty"`
	AddedBy         string           `protobuf:"bytes,5,opt,name=added_by,json=addedBy" json:"added_by,omitempty"`
	Vulnerabilities []*Vulnerability `protobuf:"bytes,6,rep,name=vulnerabilities" json:"vulnerabilities,omitempty"`
}

func (m *Feature) Reset()                    { *m = Feature{} }
func (m *Feature) String() string            { return proto.CompactTextString(m) }
func (*Feature) ProtoMessage()               {}
func (*Feature) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Feature) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Feature) GetNamespaceName() string {
	if m != nil {
		return m.NamespaceName
	}
	return ""
}

func (m *Feature) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *Feature) GetVersionFormat() string {
	if m != nil {
		return m.VersionFormat
	}
	return ""
}

func (m *Feature) GetAddedBy() string {
	if m != nil {
		return m.AddedBy
	}
	return ""
}

func (m *Feature) GetVulnerabilities() []*Vulnerability {
	if m != nil {
		return m.Vulnerabilities
	}
	return nil
}

type Ancestry struct {
	Name          string   `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	EngineVersion int32    `protobuf:"varint,2,opt,name=engine_version,json=engineVersion" json:"engine_version,omitempty"`
	Layers        []*Layer `protobuf:"bytes,3,rep,name=layers" json:"layers,omitempty"`
}

func (m *Ancestry) Reset()                    { *m = Ancestry{} }
func (m *Ancestry) String() string            { return proto.CompactTextString(m) }
func (*Ancestry) ProtoMessage()               {}
func (*Ancestry) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Ancestry) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Ancestry) GetEngineVersion() int32 {
	if m != nil {
		return m.EngineVersion
	}
	return 0
}

func (m *Ancestry) GetLayers() []*Layer {
	if m != nil {
		return m.Layers
	}
	return nil
}

type LayersIntroducingVulnerabilty struct {
	Vulnerability *Vulnerability      `protobuf:"bytes,1,opt,name=vulnerability" json:"vulnerability,omitempty"`
	Layers        []*OrderedLayerName `protobuf:"bytes,2,rep,name=layers" json:"layers,omitempty"`
}

func (m *LayersIntroducingVulnerabilty) Reset()                    { *m = LayersIntroducingVulnerabilty{} }
func (m *LayersIntroducingVulnerabilty) String() string            { return proto.CompactTextString(m) }
func (*LayersIntroducingVulnerabilty) ProtoMessage()               {}
func (*LayersIntroducingVulnerabilty) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *LayersIntroducingVulnerabilty) GetVulnerability() *Vulnerability {
	if m != nil {
		return m.Vulnerability
	}
	return nil
}

func (m *LayersIntroducingVulnerabilty) GetLayers() []*OrderedLayerName {
	if m != nil {
		return m.Layers
	}
	return nil
}

type OrderedLayerName struct {
	Index     int32  `protobuf:"varint,1,opt,name=index" json:"index,omitempty"`
	LayerName string `protobuf:"bytes,2,opt,name=layer_name,json=layerName" json:"layer_name,omitempty"`
}

func (m *OrderedLayerName) Reset()                    { *m = OrderedLayerName{} }
func (m *OrderedLayerName) String() string            { return proto.CompactTextString(m) }
func (*OrderedLayerName) ProtoMessage()               {}
func (*OrderedLayerName) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *OrderedLayerName) GetIndex() int32 {
	if m != nil {
		return m.Index
	}
	return 0
}

func (m *OrderedLayerName) GetLayerName() string {
	if m != nil {
		return m.LayerName
	}
	return ""
}

type Layer struct {
	Name           string   `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	NamespaceNames []string `protobuf:"bytes,2,rep,name=namespace_names,json=namespaceNames" json:"namespace_names,omitempty"`
}

func (m *Layer) Reset()                    { *m = Layer{} }
func (m *Layer) String() string            { return proto.CompactTextString(m) }
func (*Layer) ProtoMessage()               {}
func (*Layer) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *Layer) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Layer) GetNamespaceNames() []string {
	if m != nil {
		return m.NamespaceNames
	}
	return nil
}

type Notification struct {
	Name     string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Created  string `protobuf:"bytes,2,opt,name=created" json:"created,omitempty"`
	Notified string `protobuf:"bytes,3,opt,name=notified" json:"notified,omitempty"`
	Deleted  string `protobuf:"bytes,4,opt,name=deleted" json:"deleted,omitempty"`
	Limit    int32  `protobuf:"varint,5,opt,name=limit" json:"limit,omitempty"`
	Page     *Page  `protobuf:"bytes,6,opt,name=page" json:"page,omitempty"`
}

func (m *Notification) Reset()                    { *m = Notification{} }
func (m *Notification) String() string            { return proto.CompactTextString(m) }
func (*Notification) ProtoMessage()               {}
func (*Notification) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *Notification) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Notification) GetCreated() string {
	if m != nil {
		return m.Created
	}
	return ""
}

func (m *Notification) GetNotified() string {
	if m != nil {
		return m.Notified
	}
	return ""
}

func (m *Notification) GetDeleted() string {
	if m != nil {
		return m.Deleted
	}
	return ""
}

func (m *Notification) GetLimit() int32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *Notification) GetPage() *Page {
	if m != nil {
		return m.Page
	}
	return nil
}

type Page struct {
	ThisToken string                         `protobuf:"bytes,1,opt,name=this_token,json=thisToken" json:"this_token,omitempty"`
	NextToken string                         `protobuf:"bytes,2,opt,name=next_token,json=nextToken" json:"next_token,omitempty"`
	Old       *LayersIntroducingVulnerabilty `protobuf:"bytes,3,opt,name=old" json:"old,omitempty"`
	New       *LayersIntroducingVulnerabilty `protobuf:"bytes,4,opt,name=new" json:"new,omitempty"`
}

func (m *Page) Reset()                    { *m = Page{} }
func (m *Page) String() string            { return proto.CompactTextString(m) }
func (*Page) ProtoMessage()               {}
func (*Page) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *Page) GetThisToken() string {
	if m != nil {
		return m.ThisToken
	}
	return ""
}

func (m *Page) GetNextToken() string {
	if m != nil {
		return m.NextToken
	}
	return ""
}

func (m *Page) GetOld() *LayersIntroducingVulnerabilty {
	if m != nil {
		return m.Old
	}
	return nil
}

func (m *Page) GetNew() *LayersIntroducingVulnerabilty {
	if m != nil {
		return m.New
	}
	return nil
}

type PostAncestryRequest struct {
	AncestryName string                           `protobuf:"bytes,1,opt,name=ancestry_name,json=ancestryName" json:"ancestry_name,omitempty"`
	Format       string                           `protobuf:"bytes,2,opt,name=format" json:"format,omitempty"`
	Layers       []*PostAncestryRequest_PostLayer `protobuf:"bytes,3,rep,name=layers" json:"layers,omitempty"`
}

func (m *PostAncestryRequest) Reset()                    { *m = PostAncestryRequest{} }
func (m *PostAncestryRequest) String() string            { return proto.CompactTextString(m) }
func (*PostAncestryRequest) ProtoMessage()               {}
func (*PostAncestryRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *PostAncestryRequest) GetAncestryName() string {
	if m != nil {
		return m.AncestryName
	}
	return ""
}

func (m *PostAncestryRequest) GetFormat() string {
	if m != nil {
		return m.Format
	}
	return ""
}

func (m *PostAncestryRequest) GetLayers() []*PostAncestryRequest_PostLayer {
	if m != nil {
		return m.Layers
	}
	return nil
}

type PostAncestryRequest_PostLayer struct {
	Name    string            `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Path    string            `protobuf:"bytes,2,opt,name=path" json:"path,omitempty"`
	Headers map[string]string `protobuf:"bytes,3,rep,name=headers" json:"headers,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *PostAncestryRequest_PostLayer) Reset()         { *m = PostAncestryRequest_PostLayer{} }
func (m *PostAncestryRequest_PostLayer) String() string { return proto.CompactTextString(m) }
func (*PostAncestryRequest_PostLayer) ProtoMessage()    {}
func (*PostAncestryRequest_PostLayer) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{8, 0}
}

func (m *PostAncestryRequest_PostLayer) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *PostAncestryRequest_PostLayer) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *PostAncestryRequest_PostLayer) GetHeaders() map[string]string {
	if m != nil {
		return m.Headers
	}
	return nil
}

type PostAncestryResponse struct {
	EngineVersion int32 `protobuf:"varint,1,opt,name=engine_version,json=engineVersion" json:"engine_version,omitempty"`
}

func (m *PostAncestryResponse) Reset()                    { *m = PostAncestryResponse{} }
func (m *PostAncestryResponse) String() string            { return proto.CompactTextString(m) }
func (*PostAncestryResponse) ProtoMessage()               {}
func (*PostAncestryResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *PostAncestryResponse) GetEngineVersion() int32 {
	if m != nil {
		return m.EngineVersion
	}
	return 0
}

type GetAncestryRequest struct {
	AncestryName        string `protobuf:"bytes,1,opt,name=ancestry_name,json=ancestryName" json:"ancestry_name,omitempty"`
	WithVulnerabilities bool   `protobuf:"varint,2,opt,name=with_vulnerabilities,json=withVulnerabilities" json:"with_vulnerabilities,omitempty"`
	WithFeatures        bool   `protobuf:"varint,3,opt,name=with_features,json=withFeatures" json:"with_features,omitempty"`
}

func (m *GetAncestryRequest) Reset()                    { *m = GetAncestryRequest{} }
func (m *GetAncestryRequest) String() string            { return proto.CompactTextString(m) }
func (*GetAncestryRequest) ProtoMessage()               {}
func (*GetAncestryRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *GetAncestryRequest) GetAncestryName() string {
	if m != nil {
		return m.AncestryName
	}
	return ""
}

func (m *GetAncestryRequest) GetWithVulnerabilities() bool {
	if m != nil {
		return m.WithVulnerabilities
	}
	return false
}

func (m *GetAncestryRequest) GetWithFeatures() bool {
	if m != nil {
		return m.WithFeatures
	}
	return false
}

type GetAncestryResponse struct {
	Ancestry *Ancestry  `protobuf:"bytes,1,opt,name=ancestry" json:"ancestry,omitempty"`
	Features []*Feature `protobuf:"bytes,2,rep,name=features" json:"features,omitempty"`
}

func (m *GetAncestryResponse) Reset()                    { *m = GetAncestryResponse{} }
func (m *GetAncestryResponse) String() string            { return proto.CompactTextString(m) }
func (*GetAncestryResponse) ProtoMessage()               {}
func (*GetAncestryResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *GetAncestryResponse) GetAncestry() *Ancestry {
	if m != nil {
		return m.Ancestry
	}
	return nil
}

func (m *GetAncestryResponse) GetFeatures() []*Feature {
	if m != nil {
		return m.Features
	}
	return nil
}

type GetNotificationRequest struct {
	Page  string `protobuf:"bytes,1,opt,name=page" json:"page,omitempty"`
	Limit int32  `protobuf:"varint,2,opt,name=limit" json:"limit,omitempty"`
	Name  string `protobuf:"bytes,3,opt,name=name" json:"name,omitempty"`
}

func (m *GetNotificationRequest) Reset()                    { *m = GetNotificationRequest{} }
func (m *GetNotificationRequest) String() string            { return proto.CompactTextString(m) }
func (*GetNotificationRequest) ProtoMessage()               {}
func (*GetNotificationRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

func (m *GetNotificationRequest) GetPage() string {
	if m != nil {
		return m.Page
	}
	return ""
}

func (m *GetNotificationRequest) GetLimit() int32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *GetNotificationRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type GetNotificationResponse struct {
	Notification *Notification `protobuf:"bytes,1,opt,name=notification" json:"notification,omitempty"`
}

func (m *GetNotificationResponse) Reset()                    { *m = GetNotificationResponse{} }
func (m *GetNotificationResponse) String() string            { return proto.CompactTextString(m) }
func (*GetNotificationResponse) ProtoMessage()               {}
func (*GetNotificationResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{13} }

func (m *GetNotificationResponse) GetNotification() *Notification {
	if m != nil {
		return m.Notification
	}
	return nil
}

type DeleteNotificationRequest struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *DeleteNotificationRequest) Reset()                    { *m = DeleteNotificationRequest{} }
func (m *DeleteNotificationRequest) String() string            { return proto.CompactTextString(m) }
func (*DeleteNotificationRequest) ProtoMessage()               {}
func (*DeleteNotificationRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{14} }

func (m *DeleteNotificationRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func init() {
	proto.RegisterType((*Vulnerability)(nil), "clairpb.Vulnerability")
	proto.RegisterType((*Feature)(nil), "clairpb.Feature")
	proto.RegisterType((*Ancestry)(nil), "clairpb.Ancestry")
	proto.RegisterType((*LayersIntroducingVulnerabilty)(nil), "clairpb.LayersIntroducingVulnerabilty")
	proto.RegisterType((*OrderedLayerName)(nil), "clairpb.OrderedLayerName")
	proto.RegisterType((*Layer)(nil), "clairpb.Layer")
	proto.RegisterType((*Notification)(nil), "clairpb.Notification")
	proto.RegisterType((*Page)(nil), "clairpb.Page")
	proto.RegisterType((*PostAncestryRequest)(nil), "clairpb.PostAncestryRequest")
	proto.RegisterType((*PostAncestryRequest_PostLayer)(nil), "clairpb.PostAncestryRequest.PostLayer")
	proto.RegisterType((*PostAncestryResponse)(nil), "clairpb.PostAncestryResponse")
	proto.RegisterType((*GetAncestryRequest)(nil), "clairpb.GetAncestryRequest")
	proto.RegisterType((*GetAncestryResponse)(nil), "clairpb.GetAncestryResponse")
	proto.RegisterType((*GetNotificationRequest)(nil), "clairpb.GetNotificationRequest")
	proto.RegisterType((*GetNotificationResponse)(nil), "clairpb.GetNotificationResponse")
	proto.RegisterType((*DeleteNotificationRequest)(nil), "clairpb.DeleteNotificationRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for AncestryService service

type AncestryServiceClient interface {
	PostAncestry(ctx context.Context, in *PostAncestryRequest, opts ...grpc.CallOption) (*PostAncestryResponse, error)
	GetAncestry(ctx context.Context, in *GetAncestryRequest, opts ...grpc.CallOption) (*GetAncestryResponse, error)
}

type ancestryServiceClient struct {
	cc *grpc.ClientConn
}

func NewAncestryServiceClient(cc *grpc.ClientConn) AncestryServiceClient {
	return &ancestryServiceClient{cc}
}

func (c *ancestryServiceClient) PostAncestry(ctx context.Context, in *PostAncestryRequest, opts ...grpc.CallOption) (*PostAncestryResponse, error) {
	out := new(PostAncestryResponse)
	err := grpc.Invoke(ctx, "/clairpb.AncestryService/PostAncestry", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ancestryServiceClient) GetAncestry(ctx context.Context, in *GetAncestryRequest, opts ...grpc.CallOption) (*GetAncestryResponse, error) {
	out := new(GetAncestryResponse)
	err := grpc.Invoke(ctx, "/clairpb.AncestryService/GetAncestry", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for AncestryService service

type AncestryServiceServer interface {
	PostAncestry(context.Context, *PostAncestryRequest) (*PostAncestryResponse, error)
	GetAncestry(context.Context, *GetAncestryRequest) (*GetAncestryResponse, error)
}

func RegisterAncestryServiceServer(s *grpc.Server, srv AncestryServiceServer) {
	s.RegisterService(&_AncestryService_serviceDesc, srv)
}

func _AncestryService_PostAncestry_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PostAncestryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AncestryServiceServer).PostAncestry(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/clairpb.AncestryService/PostAncestry",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AncestryServiceServer).PostAncestry(ctx, req.(*PostAncestryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AncestryService_GetAncestry_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAncestryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AncestryServiceServer).GetAncestry(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/clairpb.AncestryService/GetAncestry",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AncestryServiceServer).GetAncestry(ctx, req.(*GetAncestryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AncestryService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "clairpb.AncestryService",
	HandlerType: (*AncestryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PostAncestry",
			Handler:    _AncestryService_PostAncestry_Handler,
		},
		{
			MethodName: "GetAncestry",
			Handler:    _AncestryService_GetAncestry_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "clair.proto",
}

// Client API for NotificationService service

type NotificationServiceClient interface {
	GetNotification(ctx context.Context, in *GetNotificationRequest, opts ...grpc.CallOption) (*GetNotificationResponse, error)
	DeleteNotification(ctx context.Context, in *DeleteNotificationRequest, opts ...grpc.CallOption) (*google_protobuf1.Empty, error)
}

type notificationServiceClient struct {
	cc *grpc.ClientConn
}

func NewNotificationServiceClient(cc *grpc.ClientConn) NotificationServiceClient {
	return &notificationServiceClient{cc}
}

func (c *notificationServiceClient) GetNotification(ctx context.Context, in *GetNotificationRequest, opts ...grpc.CallOption) (*GetNotificationResponse, error) {
	out := new(GetNotificationResponse)
	err := grpc.Invoke(ctx, "/clairpb.NotificationService/GetNotification", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notificationServiceClient) DeleteNotification(ctx context.Context, in *DeleteNotificationRequest, opts ...grpc.CallOption) (*google_protobuf1.Empty, error) {
	out := new(google_protobuf1.Empty)
	err := grpc.Invoke(ctx, "/clairpb.NotificationService/DeleteNotification", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for NotificationService service

type NotificationServiceServer interface {
	GetNotification(context.Context, *GetNotificationRequest) (*GetNotificationResponse, error)
	DeleteNotification(context.Context, *DeleteNotificationRequest) (*google_protobuf1.Empty, error)
}

func RegisterNotificationServiceServer(s *grpc.Server, srv NotificationServiceServer) {
	s.RegisterService(&_NotificationService_serviceDesc, srv)
}

func _NotificationService_GetNotification_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetNotificationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotificationServiceServer).GetNotification(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/clairpb.NotificationService/GetNotification",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotificationServiceServer).GetNotification(ctx, req.(*GetNotificationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NotificationService_DeleteNotification_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteNotificationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotificationServiceServer).DeleteNotification(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/clairpb.NotificationService/DeleteNotification",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotificationServiceServer).DeleteNotification(ctx, req.(*DeleteNotificationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _NotificationService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "clairpb.NotificationService",
	HandlerType: (*NotificationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetNotification",
			Handler:    _NotificationService_GetNotification_Handler,
		},
		{
			MethodName: "DeleteNotification",
			Handler:    _NotificationService_DeleteNotification_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "clair.proto",
}

func init() { proto.RegisterFile("clair.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 1042 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x56, 0xdd, 0x6e, 0x1b, 0x45,
	0x14, 0xd6, 0xda, 0x71, 0x6c, 0x1f, 0xdb, 0x49, 0x3a, 0x49, 0xd3, 0x8d, 0x93, 0x88, 0x74, 0x11,
	0xa5, 0xaa, 0xc0, 0x56, 0xd3, 0x9b, 0x12, 0x01, 0x82, 0xa8, 0x6d, 0xa8, 0x04, 0xa5, 0x5a, 0xaa,
	0x5c, 0x70, 0x63, 0x4d, 0xbc, 0x27, 0xce, 0x28, 0xeb, 0x59, 0xb3, 0x3b, 0x76, 0x62, 0x55, 0xdc,
	0xf0, 0x04, 0x54, 0x3c, 0x06, 0x2f, 0xc0, 0x15, 0x2f, 0xd1, 0x27, 0x00, 0xf1, 0x16, 0xdc, 0xa0,
	0xf9, 0xf5, 0xae, 0x63, 0x23, 0x7e, 0xae, 0x3c, 0xe7, 0x7c, 0xe7, 0xe7, 0x3b, 0x3f, 0x33, 0x5e,
	0x68, 0xf4, 0x63, 0xca, 0xd2, 0xce, 0x28, 0x4d, 0x44, 0x42, 0xaa, 0x4a, 0x18, 0x9d, 0xb5, 0xf7,
	0x06, 0x49, 0x32, 0x88, 0xb1, 0x4b, 0x47, 0xac, 0x4b, 0x39, 0x4f, 0x04, 0x15, 0x2c, 0xe1, 0x99,
	0x36, 0x6b, 0xef, 0x1a, 0x54, 0x49, 0x67, 0xe3, 0xf3, 0x2e, 0x0e, 0x47, 0x62, 0xaa, 0xc1, 0xe0,
	0x4d, 0x09, 0x5a, 0xa7, 0xe3, 0x98, 0x63, 0x4a, 0xcf, 0x58, 0xcc, 0xc4, 0x94, 0x10, 0x58, 0xe1,
	0x74, 0x88, 0xbe, 0x77, 0xe0, 0xdd, 0xaf, 0x87, 0xea, 0x4c, 0xde, 0x83, 0x35, 0xf9, 0x9b, 0x8d,
	0x68, 0x1f, 0x7b, 0x0a, 0x2d, 0x29, 0xb4, 0xe5, 0xb4, 0x2f, 0xa4, 0xd9, 0x01, 0x34, 0x22, 0xcc,
	0xfa, 0x29, 0x1b, 0xc9, 0xfc, 0x7e, 0x59, 0xd9, 0xe4, 0x55, 0x32, 0x78, 0xcc, 0xf8, 0xa5, 0xbf,
	0xa2, 0x83, 0xcb, 0x33, 0x69, 0x43, 0x2d, 0xc3, 0x09, 0xa6, 0x4c, 0x4c, 0xfd, 0x8a, 0xd2, 0x3b,
	0x59, 0x62, 0x43, 0x14, 0x34, 0xa2, 0x82, 0xfa, 0xab, 0x1a, 0xb3, 0x32, 0xd9, 0x81, 0xda, 0x39,
	0xbb, 0xc6, 0xa8, 0x77, 0x36, 0xf5, 0xab, 0x0a, 0xab, 0x2a, 0xf9, 0x78, 0x4a, 0x3e, 0x86, 0x5b,
	0x1a, 0x62, 0xbc, 0x77, 0x8e, 0x54, 0x8c, 0x53, 0xcc, 0xfc, 0xda, 0x41, 0xf9, 0x7e, 0xe3, 0x70,
	0xa3, 0x63, 0xba, 0xd6, 0x79, 0xa6, 0x81, 0x70, 0x5d, 0x99, 0x3e, 0xe7, 0x46, 0xce, 0x82, 0xdf,
	0x3d, 0xa8, 0x1a, 0xe1, 0xff, 0x74, 0xc3, 0x87, 0xea, 0x04, 0xd3, 0x6c, 0xd6, 0x09, 0x2b, 0xca,
	0x00, 0xe6, 0xd8, 0x3b, 0x4f, 0xd2, 0x21, 0x15, 0xa6, 0x1f, 0x2d, 0xa3, 0x7d, 0xa6, 0x94, 0xb2,
	0x40, 0x1a, 0x45, 0xba, 0x40, 0xdd, 0x98, 0xaa, 0x92, 0x8f, 0xa7, 0xe4, 0x33, 0x58, 0x9f, 0xe4,
	0xa6, 0xc6, 0x30, 0xf3, 0x57, 0x55, 0x79, 0xdb, 0xae, 0xbc, 0xc2, 0x54, 0xc3, 0x79, 0xf3, 0x60,
	0x08, 0xb5, 0xcf, 0x79, 0x1f, 0x33, 0x91, 0x2e, 0x1d, 0x39, 0xf2, 0x01, 0xe3, 0xd8, 0xb3, 0x45,
	0xc8, 0x22, 0x2b, 0x61, 0x4b, 0x6b, 0x4f, 0x4d, 0x29, 0xf7, 0x60, 0x35, 0xa6, 0x53, 0x4c, 0x33,
	0xbf, 0xac, 0xf2, 0xaf, 0xb9, 0xfc, 0x5f, 0x4a, 0x75, 0x68, 0xd0, 0xe0, 0x47, 0x0f, 0xf6, 0x95,
	0x26, 0x7b, 0xce, 0x45, 0x9a, 0x44, 0xe3, 0x3e, 0xe3, 0x83, 0x19, 0x45, 0x21, 0x67, 0xd6, 0xca,
	0x73, 0x9c, 0x2a, 0x36, 0xcb, 0x0b, 0x2a, 0x1a, 0x93, 0x87, 0x8e, 0x47, 0x49, 0xf1, 0xd8, 0x71,
	0x6e, 0x5f, 0xa7, 0x11, 0xa6, 0x18, 0xa9, 0xe4, 0x72, 0x2e, 0x8e, 0xd2, 0x09, 0x6c, 0xcc, 0x63,
	0x64, 0x0b, 0x2a, 0x8c, 0x47, 0x78, 0xad, 0x92, 0x57, 0x42, 0x2d, 0x90, 0x7d, 0x00, 0xe5, 0x93,
	0x1f, 0x76, 0x3d, 0xb6, 0x4e, 0xc1, 0x13, 0xa8, 0xa8, 0x08, 0x0b, 0xfb, 0xf8, 0x3e, 0xac, 0x17,
	0x97, 0x45, 0x33, 0xac, 0x87, 0x6b, 0x85, 0x6d, 0xc9, 0x82, 0x9f, 0x3d, 0x68, 0xbe, 0x48, 0x04,
	0x3b, 0x67, 0x7d, 0x6a, 0xef, 0xca, 0x8d, 0x68, 0x3e, 0x54, 0xfb, 0x29, 0x52, 0x81, 0x91, 0xa1,
	0x61, 0x45, 0x79, 0x53, 0xb8, 0xf2, 0xc6, 0xc8, 0xac, 0x9b, 0x93, 0xa5, 0x57, 0x84, 0x31, 0x4a,
	0x2f, 0xbd, 0x68, 0x56, 0x94, 0xf5, 0xc6, 0x6c, 0xc8, 0x84, 0xda, 0xaf, 0x4a, 0xa8, 0x05, 0x72,
	0x17, 0x56, 0x46, 0x74, 0x80, 0xea, 0xc6, 0x35, 0x0e, 0x5b, 0xae, 0x95, 0x2f, 0xe9, 0x00, 0x43,
	0x05, 0x05, 0xbf, 0x78, 0xb0, 0x22, 0x45, 0xd9, 0x1b, 0x71, 0xc1, 0xb2, 0x9e, 0x48, 0x2e, 0x91,
	0x1b, 0xae, 0x75, 0xa9, 0x79, 0x25, 0x15, 0x12, 0xe6, 0x78, 0x2d, 0x0c, 0x6c, 0x5a, 0x27, 0x35,
	0x1a, 0x7e, 0x0c, 0xe5, 0x24, 0xd6, 0x84, 0x1b, 0x87, 0xf7, 0x8a, 0xbb, 0xb3, 0x6c, 0x53, 0x42,
	0xe9, 0x22, 0x3d, 0x39, 0x5e, 0xa9, 0x7a, 0xfe, 0x85, 0x27, 0xc7, 0xab, 0xe0, 0x6d, 0x09, 0x36,
	0x5f, 0x26, 0x99, 0xb0, 0xeb, 0x1f, 0xe2, 0x77, 0x63, 0xcc, 0x04, 0x79, 0x17, 0x5a, 0xd4, 0xa8,
	0x7a, 0xb9, 0xc6, 0x37, 0xad, 0x52, 0x2d, 0xc8, 0x36, 0xac, 0x9a, 0x2b, 0xab, 0x6b, 0x31, 0x12,
	0xf9, 0x74, 0xee, 0x1e, 0xcc, 0x18, 0x2d, 0x48, 0xa5, 0x74, 0x85, 0xfb, 0xd1, 0xfe, 0xd5, 0x83,
	0xba, 0xd3, 0x2e, 0x1c, 0x3d, 0x91, 0x43, 0x11, 0x17, 0x26, 0xaf, 0x3a, 0x93, 0xaf, 0xa0, 0x7a,
	0x81, 0x34, 0x9a, 0xa5, 0x7d, 0xf4, 0xcf, 0xd2, 0x76, 0xbe, 0xd0, 0x5e, 0x4f, 0xb9, 0x44, 0x6d,
	0x8c, 0xf6, 0x11, 0x34, 0xf3, 0x00, 0xd9, 0x80, 0xf2, 0x25, 0x4e, 0x0d, 0x0b, 0x79, 0x94, 0xfb,
	0x32, 0xa1, 0xf1, 0xd8, 0x5e, 0x02, 0x2d, 0x1c, 0x95, 0x1e, 0x7b, 0xc1, 0x27, 0xb0, 0x55, 0x4c,
	0x99, 0x8d, 0x12, 0x9e, 0x2d, 0x7a, 0x47, 0xbc, 0x05, 0xef, 0x48, 0xf0, 0xc6, 0x03, 0x72, 0x82,
	0xff, 0x6d, 0x26, 0x0f, 0x61, 0xeb, 0x8a, 0x89, 0x8b, 0xde, 0xfc, 0x8b, 0x28, 0x39, 0xd6, 0xc2,
	0x4d, 0x89, 0x9d, 0x16, 0x21, 0x19, 0x57, 0xb9, 0xb8, 0x3f, 0x87, 0xb2, 0xb2, 0x6d, 0x4a, 0xa5,
	0xfb, 0x1f, 0x48, 0x61, 0xb3, 0x40, 0xc9, 0x54, 0xf4, 0x21, 0xd4, 0x6c, 0x7a, 0xf3, 0x46, 0xdd,
	0x72, 0x5d, 0x77, 0xc6, 0xce, 0x84, 0x7c, 0x00, 0x35, 0x97, 0xa5, 0xb4, 0xe4, 0x2f, 0xc8, 0x59,
	0x04, 0xa7, 0xb0, 0x7d, 0x82, 0x22, 0xff, 0x0e, 0xd8, 0x56, 0x10, 0x73, 0x29, 0x3d, 0x3b, 0xff,
	0x01, 0xce, 0xae, 0x6f, 0x29, 0x7f, 0x7d, 0xed, 0xf6, 0x94, 0x67, 0xdb, 0x13, 0xbc, 0x82, 0x3b,
	0x37, 0xe2, 0x9a, 0x7a, 0x3e, 0x82, 0x26, 0xcf, 0xe9, 0x4d, 0x4d, 0xb7, 0x1d, 0xc9, 0x82, 0x53,
	0xc1, 0x34, 0xe8, 0xc2, 0xce, 0x13, 0xf5, 0x92, 0x2c, 0x21, 0x3c, 0xbf, 0xc4, 0x87, 0xbf, 0x79,
	0xb0, 0x6e, 0x7b, 0xf4, 0x0d, 0xa6, 0x13, 0xd6, 0x47, 0x42, 0xa1, 0x99, 0xdf, 0x1c, 0xb2, 0xf7,
	0x77, 0x3b, 0xdc, 0xde, 0x5f, 0x82, 0xea, 0x62, 0x82, 0xad, 0x1f, 0xde, 0xfe, 0xf1, 0x53, 0x69,
	0x2d, 0xa8, 0x77, 0xed, 0x00, 0x8e, 0xbc, 0x07, 0xe4, 0x12, 0x1a, 0xb9, 0x49, 0x92, 0x5d, 0x17,
	0xe3, 0xe6, 0xca, 0xb5, 0xf7, 0x16, 0x83, 0x26, 0xfe, 0x5d, 0x15, 0x7f, 0x97, 0xec, 0xb8, 0xf8,
	0xdd, 0xd7, 0x85, 0x0d, 0xfd, 0xfe, 0xf0, 0x4f, 0x0f, 0x36, 0xf3, 0xfd, 0xb0, 0x75, 0x66, 0xb0,
	0x3e, 0x37, 0x02, 0xf2, 0x4e, 0x3e, 0xd7, 0x82, 0x1e, 0xb6, 0x0f, 0x96, 0x1b, 0x18, 0x42, 0xfb,
	0x8a, 0xd0, 0x1d, 0x72, 0xbb, 0x9b, 0x9f, 0x4c, 0xd6, 0x7d, 0xad, 0xc8, 0x90, 0x04, 0xc8, 0xcd,
	0x09, 0x91, 0xc0, 0x85, 0x5d, 0x3a, 0xbe, 0xf6, 0x76, 0x47, 0x7f, 0x37, 0x76, 0xec, 0x77, 0x63,
	0xe7, 0xa9, 0xfc, 0x6e, 0xb4, 0x09, 0x1f, 0x2c, 0x4e, 0x78, 0x5c, 0xff, 0xd6, 0x7e, 0x96, 0x9e,
	0xad, 0x2a, 0xcf, 0x47, 0x7f, 0x05, 0x00, 0x00, 0xff, 0xff, 0x5b, 0x9c, 0x1d, 0xc4, 0xb5, 0x0a,
	0x00, 0x00,
}
