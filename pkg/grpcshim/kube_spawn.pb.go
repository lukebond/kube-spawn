// Code generated by protoc-gen-go. DO NOT EDIT.
// source: kube_spawn.proto

/*
Package grpcshim is a generated protocol buffer package.

It is generated from these files:
	kube_spawn.proto

It has these top-level messages:
	Result
	ClusterProperties
	Cluster
	Empty
*/
package grpcshim

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

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

type Result struct {
	Success    bool   `protobuf:"varint,1,opt,name=success" json:"success,omitempty"`
	Error      string `protobuf:"bytes,2,opt,name=error" json:"error,omitempty"`
	Stdout     string `protobuf:"bytes,3,opt,name=stdout" json:"stdout,omitempty"`
	Stderr     string `protobuf:"bytes,4,opt,name=stderr" json:"stderr,omitempty"`
	Kubeconfig string `protobuf:"bytes,5,opt,name=kubeconfig" json:"kubeconfig,omitempty"`
}

func (m *Result) Reset()                    { *m = Result{} }
func (m *Result) String() string            { return proto.CompactTextString(m) }
func (*Result) ProtoMessage()               {}
func (*Result) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Result) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *Result) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *Result) GetStdout() string {
	if m != nil {
		return m.Stdout
	}
	return ""
}

func (m *Result) GetStderr() string {
	if m != nil {
		return m.Stderr
	}
	return ""
}

func (m *Result) GetKubeconfig() string {
	if m != nil {
		return m.Kubeconfig
	}
	return ""
}

type ClusterProperties struct {
	Name        string `protobuf:"bytes,1,opt,name=Name,json=name" json:"Name,omitempty"`
	NumberNodes uint32 `protobuf:"varint,2,opt,name=NumberNodes,json=numberNodes" json:"NumberNodes,omitempty"`
}

func (m *ClusterProperties) Reset()                    { *m = ClusterProperties{} }
func (m *ClusterProperties) String() string            { return proto.CompactTextString(m) }
func (*ClusterProperties) ProtoMessage()               {}
func (*ClusterProperties) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ClusterProperties) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ClusterProperties) GetNumberNodes() uint32 {
	if m != nil {
		return m.NumberNodes
	}
	return 0
}

type Cluster struct {
	Name string `protobuf:"bytes,1,opt,name=Name,json=name" json:"Name,omitempty"`
}

func (m *Cluster) Reset()                    { *m = Cluster{} }
func (m *Cluster) String() string            { return proto.CompactTextString(m) }
func (*Cluster) ProtoMessage()               {}
func (*Cluster) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Cluster) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type Empty struct {
}

func (m *Empty) Reset()                    { *m = Empty{} }
func (m *Empty) String() string            { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()               {}
func (*Empty) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func init() {
	proto.RegisterType((*Result)(nil), "grpcshim.Result")
	proto.RegisterType((*ClusterProperties)(nil), "grpcshim.ClusterProperties")
	proto.RegisterType((*Cluster)(nil), "grpcshim.Cluster")
	proto.RegisterType((*Empty)(nil), "grpcshim.Empty")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for KubeSpawn service

type KubeSpawnClient interface {
	Create(ctx context.Context, in *ClusterProperties, opts ...grpc.CallOption) (*Result, error)
	Start(ctx context.Context, in *Cluster, opts ...grpc.CallOption) (*Result, error)
	GetKubeconfig(ctx context.Context, in *Cluster, opts ...grpc.CallOption) (*Result, error)
	List(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Result, error)
	Destroy(ctx context.Context, in *Cluster, opts ...grpc.CallOption) (*Result, error)
}

type kubeSpawnClient struct {
	cc *grpc.ClientConn
}

func NewKubeSpawnClient(cc *grpc.ClientConn) KubeSpawnClient {
	return &kubeSpawnClient{cc}
}

func (c *kubeSpawnClient) Create(ctx context.Context, in *ClusterProperties, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := grpc.Invoke(ctx, "/grpcshim.KubeSpawn/Create", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kubeSpawnClient) Start(ctx context.Context, in *Cluster, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := grpc.Invoke(ctx, "/grpcshim.KubeSpawn/Start", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kubeSpawnClient) GetKubeconfig(ctx context.Context, in *Cluster, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := grpc.Invoke(ctx, "/grpcshim.KubeSpawn/GetKubeconfig", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kubeSpawnClient) List(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := grpc.Invoke(ctx, "/grpcshim.KubeSpawn/List", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kubeSpawnClient) Destroy(ctx context.Context, in *Cluster, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := grpc.Invoke(ctx, "/grpcshim.KubeSpawn/Destroy", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for KubeSpawn service

type KubeSpawnServer interface {
	Create(context.Context, *ClusterProperties) (*Result, error)
	Start(context.Context, *Cluster) (*Result, error)
	GetKubeconfig(context.Context, *Cluster) (*Result, error)
	List(context.Context, *Empty) (*Result, error)
	Destroy(context.Context, *Cluster) (*Result, error)
}

func RegisterKubeSpawnServer(s *grpc.Server, srv KubeSpawnServer) {
	s.RegisterService(&_KubeSpawn_serviceDesc, srv)
}

func _KubeSpawn_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClusterProperties)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KubeSpawnServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpcshim.KubeSpawn/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KubeSpawnServer).Create(ctx, req.(*ClusterProperties))
	}
	return interceptor(ctx, in, info, handler)
}

func _KubeSpawn_Start_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Cluster)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KubeSpawnServer).Start(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpcshim.KubeSpawn/Start",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KubeSpawnServer).Start(ctx, req.(*Cluster))
	}
	return interceptor(ctx, in, info, handler)
}

func _KubeSpawn_GetKubeconfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Cluster)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KubeSpawnServer).GetKubeconfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpcshim.KubeSpawn/GetKubeconfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KubeSpawnServer).GetKubeconfig(ctx, req.(*Cluster))
	}
	return interceptor(ctx, in, info, handler)
}

func _KubeSpawn_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KubeSpawnServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpcshim.KubeSpawn/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KubeSpawnServer).List(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _KubeSpawn_Destroy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Cluster)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KubeSpawnServer).Destroy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpcshim.KubeSpawn/Destroy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KubeSpawnServer).Destroy(ctx, req.(*Cluster))
	}
	return interceptor(ctx, in, info, handler)
}

var _KubeSpawn_serviceDesc = grpc.ServiceDesc{
	ServiceName: "grpcshim.KubeSpawn",
	HandlerType: (*KubeSpawnServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _KubeSpawn_Create_Handler,
		},
		{
			MethodName: "Start",
			Handler:    _KubeSpawn_Start_Handler,
		},
		{
			MethodName: "GetKubeconfig",
			Handler:    _KubeSpawn_GetKubeconfig_Handler,
		},
		{
			MethodName: "List",
			Handler:    _KubeSpawn_List_Handler,
		},
		{
			MethodName: "Destroy",
			Handler:    _KubeSpawn_Destroy_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "kube_spawn.proto",
}

func init() { proto.RegisterFile("kube_spawn.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 304 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0xcd, 0x4a, 0xc3, 0x60,
	0x10, 0x34, 0x35, 0x3f, 0xed, 0x96, 0x62, 0xbb, 0x88, 0x04, 0x45, 0x29, 0x39, 0x15, 0x84, 0x20,
	0x0a, 0x82, 0xe7, 0x2a, 0x22, 0x95, 0x22, 0xe9, 0x03, 0x48, 0x92, 0xae, 0x35, 0xd8, 0xe4, 0x0b,
	0xfb, 0xed, 0x87, 0xf4, 0x0d, 0xbc, 0xf8, 0xce, 0xd2, 0xf4, 0x4f, 0x68, 0x0e, 0x3d, 0xce, 0xcc,
	0xee, 0x30, 0x3b, 0x2c, 0x74, 0xbf, 0x4c, 0x42, 0xef, 0xba, 0x8c, 0xbf, 0x8b, 0xb0, 0x64, 0x25,
	0x0a, 0x9b, 0x33, 0x2e, 0x53, 0xfd, 0x99, 0xe5, 0xc1, 0x8f, 0x05, 0x6e, 0x44, 0xda, 0xcc, 0x05,
	0x7d, 0xf0, 0xb4, 0x49, 0x53, 0xd2, 0xda, 0xb7, 0xfa, 0xd6, 0xa0, 0x19, 0x6d, 0x20, 0x9e, 0x82,
	0x43, 0xcc, 0x8a, 0xfd, 0x46, 0xdf, 0x1a, 0xb4, 0xa2, 0x15, 0xc0, 0x33, 0x70, 0xb5, 0x4c, 0x95,
	0x11, 0xff, 0xb8, 0xa2, 0xd7, 0x68, 0xcd, 0x13, 0xb3, 0x6f, 0x6f, 0x79, 0x62, 0xc6, 0x2b, 0x80,
	0x65, 0x90, 0x54, 0x15, 0x1f, 0xd9, 0xcc, 0x77, 0x2a, 0xed, 0x1f, 0x13, 0xbc, 0x40, 0x6f, 0x38,
	0x37, 0x5a, 0x88, 0xdf, 0x58, 0x95, 0xc4, 0x92, 0x91, 0x46, 0x04, 0x7b, 0x1c, 0xe7, 0x54, 0x25,
	0x6a, 0x45, 0x76, 0x11, 0xe7, 0x84, 0x7d, 0x68, 0x8f, 0x4d, 0x9e, 0x10, 0x8f, 0xd5, 0x94, 0x74,
	0x15, 0xaa, 0x13, 0xb5, 0x8b, 0x1d, 0x15, 0x5c, 0x82, 0xb7, 0xb6, 0xaa, 0x33, 0x08, 0x3c, 0x70,
	0x9e, 0xf2, 0x52, 0x16, 0xb7, 0xbf, 0x0d, 0x68, 0x8d, 0x4c, 0x42, 0x93, 0x65, 0x37, 0xf8, 0x00,
	0xee, 0x90, 0x29, 0x16, 0xc2, 0x8b, 0x70, 0x53, 0x50, 0xb8, 0x17, 0xe9, 0xbc, 0xbb, 0x13, 0x57,
	0xcd, 0x05, 0x47, 0x18, 0x82, 0x33, 0x91, 0x98, 0x05, 0x7b, 0x7b, 0x9b, 0xb5, 0xf3, 0xf7, 0xd0,
	0x79, 0x26, 0x19, 0x6d, 0x8f, 0x3f, 0x74, 0xef, 0x1a, 0xec, 0xd7, 0x4c, 0x0b, 0x9e, 0xec, 0xb4,
	0xea, 0x92, 0xda, 0xe1, 0x1b, 0xf0, 0x1e, 0x49, 0x0b, 0xab, 0xc5, 0x81, 0xf6, 0x89, 0x5b, 0xbd,
	0xc7, 0xdd, 0x5f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x24, 0x80, 0x24, 0xff, 0x32, 0x02, 0x00, 0x00,
}