// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: protos/edge.proto

package protos

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ChatEdgeClient is the client API for ChatEdge service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ChatEdgeClient interface {
	JoinChat(ctx context.Context, in *JoinRequest, opts ...grpc.CallOption) (*JoinResponse, error)
	GetChats(ctx context.Context, in *ChatRequest, opts ...grpc.CallOption) (*ChatResponse, error)
	Send(ctx context.Context, in *SendRequest, opts ...grpc.CallOption) (*Empty, error)
}

type chatEdgeClient struct {
	cc grpc.ClientConnInterface
}

func NewChatEdgeClient(cc grpc.ClientConnInterface) ChatEdgeClient {
	return &chatEdgeClient{cc}
}

func (c *chatEdgeClient) JoinChat(ctx context.Context, in *JoinRequest, opts ...grpc.CallOption) (*JoinResponse, error) {
	out := new(JoinResponse)
	err := c.cc.Invoke(ctx, "/ChatEdge/JoinChat", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatEdgeClient) GetChats(ctx context.Context, in *ChatRequest, opts ...grpc.CallOption) (*ChatResponse, error) {
	out := new(ChatResponse)
	err := c.cc.Invoke(ctx, "/ChatEdge/GetChats", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatEdgeClient) Send(ctx context.Context, in *SendRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/ChatEdge/Send", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ChatEdgeServer is the server API for ChatEdge service.
// All implementations must embed UnimplementedChatEdgeServer
// for forward compatibility
type ChatEdgeServer interface {
	JoinChat(context.Context, *JoinRequest) (*JoinResponse, error)
	GetChats(context.Context, *ChatRequest) (*ChatResponse, error)
	Send(context.Context, *SendRequest) (*Empty, error)
	mustEmbedUnimplementedChatEdgeServer()
}

// UnimplementedChatEdgeServer must be embedded to have forward compatible implementations.
type UnimplementedChatEdgeServer struct {
}

func (UnimplementedChatEdgeServer) JoinChat(context.Context, *JoinRequest) (*JoinResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method JoinChat not implemented")
}
func (UnimplementedChatEdgeServer) GetChats(context.Context, *ChatRequest) (*ChatResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetChats not implemented")
}
func (UnimplementedChatEdgeServer) Send(context.Context, *SendRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Send not implemented")
}
func (UnimplementedChatEdgeServer) mustEmbedUnimplementedChatEdgeServer() {}

// UnsafeChatEdgeServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ChatEdgeServer will
// result in compilation errors.
type UnsafeChatEdgeServer interface {
	mustEmbedUnimplementedChatEdgeServer()
}

func RegisterChatEdgeServer(s grpc.ServiceRegistrar, srv ChatEdgeServer) {
	s.RegisterService(&ChatEdge_ServiceDesc, srv)
}

func _ChatEdge_JoinChat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JoinRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatEdgeServer).JoinChat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ChatEdge/JoinChat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatEdgeServer).JoinChat(ctx, req.(*JoinRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatEdge_GetChats_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChatRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatEdgeServer).GetChats(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ChatEdge/GetChats",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatEdgeServer).GetChats(ctx, req.(*ChatRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatEdge_Send_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatEdgeServer).Send(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ChatEdge/Send",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatEdgeServer).Send(ctx, req.(*SendRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ChatEdge_ServiceDesc is the grpc.ServiceDesc for ChatEdge service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ChatEdge_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ChatEdge",
	HandlerType: (*ChatEdgeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "JoinChat",
			Handler:    _ChatEdge_JoinChat_Handler,
		},
		{
			MethodName: "GetChats",
			Handler:    _ChatEdge_GetChats_Handler,
		},
		{
			MethodName: "Send",
			Handler:    _ChatEdge_Send_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protos/edge.proto",
}
