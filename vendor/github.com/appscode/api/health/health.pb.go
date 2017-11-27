// Code generated by protoc-gen-go. DO NOT EDIT.
// source: health.proto

/*
Package health is a generated protocol buffer package.

It is generated from these files:
	health.proto

It has these top-level messages:
	StatusResponse
	URLBase
	NetConfig
	Metadata
*/
package health

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"
import _ "github.com/grpc-ecosystem/grpc-gateway/third_party/appscodeapis/appscode/api"
import appscode_dtypes "github.com/appscode/api/dtypes"
import appscode_version "github.com/appscode/api/version"

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

type StatusResponse struct {
	Version  *appscode_version.Version `protobuf:"bytes,1,opt,name=version" json:"version,omitempty"`
	Metadata *Metadata                 `protobuf:"bytes,2,opt,name=metadata" json:"metadata,omitempty"`
}

func (m *StatusResponse) Reset()                    { *m = StatusResponse{} }
func (m *StatusResponse) String() string            { return proto.CompactTextString(m) }
func (*StatusResponse) ProtoMessage()               {}
func (*StatusResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *StatusResponse) GetVersion() *appscode_version.Version {
	if m != nil {
		return m.Version
	}
	return nil
}

func (m *StatusResponse) GetMetadata() *Metadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

type URLBase struct {
	Scheme   string `protobuf:"bytes,1,opt,name=scheme" json:"scheme,omitempty"`
	BaseAddr string `protobuf:"bytes,2,opt,name=base_addr,json=baseAddr" json:"base_addr,omitempty"`
}

func (m *URLBase) Reset()                    { *m = URLBase{} }
func (m *URLBase) String() string            { return proto.CompactTextString(m) }
func (*URLBase) ProtoMessage()               {}
func (*URLBase) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *URLBase) GetScheme() string {
	if m != nil {
		return m.Scheme
	}
	return ""
}

func (m *URLBase) GetBaseAddr() string {
	if m != nil {
		return m.BaseAddr
	}
	return ""
}

type NetConfig struct {
	TeamAddr         string   `protobuf:"bytes,1,opt,name=team_addr,json=teamAddr" json:"team_addr,omitempty"`
	PublicUrls       *URLBase `protobuf:"bytes,2,opt,name=public_urls,json=publicUrls" json:"public_urls,omitempty"`
	TeamUrls         *URLBase `protobuf:"bytes,3,opt,name=team_urls,json=teamUrls" json:"team_urls,omitempty"`
	ClusterUrls      *URLBase `protobuf:"bytes,4,opt,name=cluster_urls,json=clusterUrls" json:"cluster_urls,omitempty"`
	InClusterUrls    *URLBase `protobuf:"bytes,5,opt,name=in_cluster_urls,json=inClusterUrls" json:"in_cluster_urls,omitempty"`
	URLShortenerUrls *URLBase `protobuf:"bytes,6,opt,name=URL_shortener_urls,json=URLShortenerUrls" json:"URL_shortener_urls,omitempty"`
	FileUrls         *URLBase `protobuf:"bytes,7,opt,name=file_urls,json=fileUrls" json:"file_urls,omitempty"`
}

func (m *NetConfig) Reset()                    { *m = NetConfig{} }
func (m *NetConfig) String() string            { return proto.CompactTextString(m) }
func (*NetConfig) ProtoMessage()               {}
func (*NetConfig) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *NetConfig) GetTeamAddr() string {
	if m != nil {
		return m.TeamAddr
	}
	return ""
}

func (m *NetConfig) GetPublicUrls() *URLBase {
	if m != nil {
		return m.PublicUrls
	}
	return nil
}

func (m *NetConfig) GetTeamUrls() *URLBase {
	if m != nil {
		return m.TeamUrls
	}
	return nil
}

func (m *NetConfig) GetClusterUrls() *URLBase {
	if m != nil {
		return m.ClusterUrls
	}
	return nil
}

func (m *NetConfig) GetInClusterUrls() *URLBase {
	if m != nil {
		return m.InClusterUrls
	}
	return nil
}

func (m *NetConfig) GetURLShortenerUrls() *URLBase {
	if m != nil {
		return m.URLShortenerUrls
	}
	return nil
}

func (m *NetConfig) GetFileUrls() *URLBase {
	if m != nil {
		return m.FileUrls
	}
	return nil
}

type Metadata struct {
	Env       string     `protobuf:"bytes,1,opt,name=env" json:"env,omitempty"`
	TeamId    string     `protobuf:"bytes,2,opt,name=team_id,json=teamId" json:"team_id,omitempty"`
	NetConfig *NetConfig `protobuf:"bytes,3,opt,name=net_config,json=netConfig" json:"net_config,omitempty"`
}

func (m *Metadata) Reset()                    { *m = Metadata{} }
func (m *Metadata) String() string            { return proto.CompactTextString(m) }
func (*Metadata) ProtoMessage()               {}
func (*Metadata) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Metadata) GetEnv() string {
	if m != nil {
		return m.Env
	}
	return ""
}

func (m *Metadata) GetTeamId() string {
	if m != nil {
		return m.TeamId
	}
	return ""
}

func (m *Metadata) GetNetConfig() *NetConfig {
	if m != nil {
		return m.NetConfig
	}
	return nil
}

func init() {
	proto.RegisterType((*StatusResponse)(nil), "appscode.health.StatusResponse")
	proto.RegisterType((*URLBase)(nil), "appscode.health.URLBase")
	proto.RegisterType((*NetConfig)(nil), "appscode.health.NetConfig")
	proto.RegisterType((*Metadata)(nil), "appscode.health.Metadata")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Health service

type HealthClient interface {
	Status(ctx context.Context, in *appscode_dtypes.VoidRequest, opts ...grpc.CallOption) (*StatusResponse, error)
}

type healthClient struct {
	cc *grpc.ClientConn
}

func NewHealthClient(cc *grpc.ClientConn) HealthClient {
	return &healthClient{cc}
}

func (c *healthClient) Status(ctx context.Context, in *appscode_dtypes.VoidRequest, opts ...grpc.CallOption) (*StatusResponse, error) {
	out := new(StatusResponse)
	err := grpc.Invoke(ctx, "/appscode.health.Health/Status", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Health service

type HealthServer interface {
	Status(context.Context, *appscode_dtypes.VoidRequest) (*StatusResponse, error)
}

func RegisterHealthServer(s *grpc.Server, srv HealthServer) {
	s.RegisterService(&_Health_serviceDesc, srv)
}

func _Health_Status_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(appscode_dtypes.VoidRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HealthServer).Status(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appscode.health.Health/Status",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HealthServer).Status(ctx, req.(*appscode_dtypes.VoidRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Health_serviceDesc = grpc.ServiceDesc{
	ServiceName: "appscode.health.Health",
	HandlerType: (*HealthServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Status",
			Handler:    _Health_Status_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "health.proto",
}

func init() { proto.RegisterFile("health.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 504 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x93, 0x41, 0x6f, 0xd3, 0x30,
	0x14, 0xc7, 0xd5, 0x16, 0xd2, 0xf6, 0x75, 0xb0, 0xc9, 0x07, 0x56, 0x4a, 0xc5, 0x50, 0x4e, 0x3b,
	0x35, 0x68, 0xd3, 0x0e, 0x13, 0x12, 0x82, 0x4e, 0x42, 0x20, 0x15, 0x34, 0x79, 0xea, 0x0e, 0x5c,
	0x22, 0x37, 0x7e, 0x6b, 0x8d, 0x52, 0x3b, 0x8b, 0x9d, 0x49, 0x08, 0x4e, 0xfb, 0x0a, 0x5c, 0xf8,
	0x22, 0x7c, 0x12, 0xbe, 0x02, 0x1f, 0x04, 0xc5, 0x76, 0x3a, 0x42, 0x35, 0xe5, 0x12, 0xdb, 0xcf,
	0xff, 0xdf, 0xdf, 0xcf, 0x7e, 0x79, 0xb0, 0xb3, 0x42, 0x96, 0x9a, 0xd5, 0x24, 0xcb, 0x95, 0x51,
	0x64, 0x97, 0x65, 0x99, 0x4e, 0x14, 0xc7, 0x89, 0x0b, 0x8f, 0xc6, 0x4b, 0xa5, 0x96, 0x29, 0x46,
	0x2c, 0x13, 0x11, 0x93, 0x52, 0x19, 0x66, 0x84, 0x92, 0xda, 0xc9, 0x47, 0xcf, 0x2b, 0xf9, 0x3d,
	0xfb, 0x07, 0xb5, 0x7d, 0x6e, 0xbe, 0x66, 0xa8, 0x23, 0xfb, 0xf5, 0x82, 0xb0, 0x26, 0xb8, 0xc1,
	0x5c, 0x0b, 0x25, 0xab, 0xd1, 0x69, 0xc2, 0xef, 0xf0, 0xf8, 0xc2, 0x30, 0x53, 0x68, 0x8a, 0x3a,
	0x53, 0x52, 0x23, 0x39, 0x86, 0xae, 0x97, 0x0c, 0x5b, 0x2f, 0x5a, 0x87, 0x83, 0xa3, 0xa7, 0x93,
	0x4d, 0xde, 0x15, 0x7b, 0xe9, 0x46, 0x5a, 0x29, 0xc9, 0x09, 0xf4, 0xd6, 0x68, 0x18, 0x67, 0x86,
	0x0d, 0xdb, 0xff, 0x53, 0xfe, 0x11, 0x3e, 0x7a, 0x01, 0xdd, 0x48, 0xc3, 0xd7, 0xd0, 0x9d, 0xd3,
	0xd9, 0x94, 0x69, 0x24, 0x4f, 0x20, 0xd0, 0xc9, 0x0a, 0xd7, 0x68, 0x4f, 0xed, 0x53, 0xbf, 0x22,
	0xcf, 0xa0, 0xbf, 0x60, 0x1a, 0x63, 0xc6, 0x79, 0x6e, 0xad, 0xfb, 0xb4, 0x57, 0x06, 0xde, 0x72,
	0x9e, 0x87, 0x3f, 0x3b, 0xd0, 0xff, 0x84, 0xe6, 0x4c, 0xc9, 0x2b, 0xb1, 0x2c, 0xa5, 0x06, 0xd9,
	0xda, 0x49, 0x9d, 0x4b, 0xaf, 0x0c, 0x94, 0x52, 0x72, 0x0a, 0x83, 0xac, 0x58, 0xa4, 0x22, 0x89,
	0x8b, 0x3c, 0xd5, 0x3e, 0xc9, 0xe1, 0x56, 0x92, 0x3e, 0x1d, 0x0a, 0x4e, 0x3c, 0xcf, 0x53, 0x4d,
	0x4e, 0xbc, 0xaf, 0x05, 0x3b, 0x0d, 0xa0, 0x3d, 0xd1, 0x62, 0xaf, 0x60, 0x27, 0x49, 0x0b, 0x6d,
	0x30, 0x77, 0xe4, 0x83, 0x06, 0x72, 0xe0, 0xd5, 0x16, 0x7e, 0x03, 0xbb, 0x42, 0xc6, 0x35, 0xfe,
	0x61, 0x03, 0xff, 0x48, 0xc8, 0xb3, 0x7f, 0x1c, 0xde, 0x01, 0x99, 0xd3, 0x59, 0xac, 0x57, 0x2a,
	0x37, 0x28, 0x2b, 0x93, 0xa0, 0xc1, 0x64, 0x6f, 0x4e, 0x67, 0x17, 0x15, 0x52, 0xdd, 0xfe, 0x4a,
	0xa4, 0xe8, 0xf0, 0x6e, 0xd3, 0xed, 0x4b, 0x69, 0x89, 0x85, 0x19, 0xf4, 0xaa, 0x82, 0x93, 0x3d,
	0xe8, 0xa0, 0xbc, 0xf1, 0x25, 0x29, 0xa7, 0x64, 0x1f, 0xba, 0xf6, 0x49, 0x05, 0xf7, 0x35, 0x0d,
	0xca, 0xe5, 0x07, 0x4e, 0x4e, 0x01, 0x24, 0x9a, 0x38, 0xb1, 0x15, 0xf5, 0x8f, 0x3d, 0xda, 0x3a,
	0x6e, 0x53, 0x73, 0xda, 0x97, 0xd5, 0xf4, 0xe8, 0x1b, 0x04, 0xef, 0xed, 0x36, 0xb9, 0x86, 0xc0,
	0xfd, 0xd4, 0x64, 0x7c, 0x87, 0xba, 0x06, 0x99, 0x5c, 0x2a, 0xc1, 0x29, 0x5e, 0x17, 0xa8, 0xcd,
	0xe8, 0x60, 0xcb, 0xb8, 0xde, 0x0b, 0xe1, 0xe1, 0xed, 0xaf, 0x61, 0xbb, 0xd7, 0xba, 0xfd, 0xfd,
	0xe7, 0x47, 0x7b, 0x4c, 0x46, 0x51, 0x5c, 0xeb, 0x28, 0xc7, 0x44, 0x5f, 0xb4, 0x92, 0xd3, 0x97,
	0xb0, 0x9f, 0xa8, 0xf5, 0x9d, 0x1f, 0xcb, 0x84, 0xf7, 0x9c, 0x0e, 0x5c, 0x56, 0xe7, 0x65, 0xbf,
	0x9d, 0xb7, 0x3e, 0x07, 0x2e, 0xbc, 0x08, 0x6c, 0x03, 0x1e, 0xff, 0x0d, 0x00, 0x00, 0xff, 0xff,
	0xf9, 0x8f, 0x47, 0x6d, 0x24, 0x04, 0x00, 0x00,
}