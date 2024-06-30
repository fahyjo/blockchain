// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v3.12.4
// source: proto/blockchain.proto

package blockchain

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	Node_HandleHandshake_FullMethodName   = "/Node/HandleHandshake"
	Node_HandleTransaction_FullMethodName = "/Node/HandleTransaction"
	Node_HandleBlock_FullMethodName       = "/Node/HandleBlock"
	Node_HandleVote_FullMethodName        = "/Node/HandleVote"
)

// NodeClient is the client API for Node service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NodeClient interface {
	HandleHandshake(ctx context.Context, in *Handshake, opts ...grpc.CallOption) (*Handshake, error)
	HandleTransaction(ctx context.Context, in *Transaction, opts ...grpc.CallOption) (*Ack, error)
	HandleBlock(ctx context.Context, in *Block, opts ...grpc.CallOption) (*Ack, error)
	HandleVote(ctx context.Context, in *Vote, opts ...grpc.CallOption) (*Ack, error)
}

type nodeClient struct {
	cc grpc.ClientConnInterface
}

func NewNodeClient(cc grpc.ClientConnInterface) NodeClient {
	return &nodeClient{cc}
}

func (c *nodeClient) HandleHandshake(ctx context.Context, in *Handshake, opts ...grpc.CallOption) (*Handshake, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Handshake)
	err := c.cc.Invoke(ctx, Node_HandleHandshake_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeClient) HandleTransaction(ctx context.Context, in *Transaction, opts ...grpc.CallOption) (*Ack, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Ack)
	err := c.cc.Invoke(ctx, Node_HandleTransaction_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeClient) HandleBlock(ctx context.Context, in *Block, opts ...grpc.CallOption) (*Ack, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Ack)
	err := c.cc.Invoke(ctx, Node_HandleBlock_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeClient) HandleVote(ctx context.Context, in *Vote, opts ...grpc.CallOption) (*Ack, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Ack)
	err := c.cc.Invoke(ctx, Node_HandleVote_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NodeServer is the server API for Node service.
// All implementations must embed UnimplementedNodeServer
// for forward compatibility
type NodeServer interface {
	HandleHandshake(context.Context, *Handshake) (*Handshake, error)
	HandleTransaction(context.Context, *Transaction) (*Ack, error)
	HandleBlock(context.Context, *Block) (*Ack, error)
	HandleVote(context.Context, *Vote) (*Ack, error)
	mustEmbedUnimplementedNodeServer()
}

// UnimplementedNodeServer must be embedded to have forward compatible implementations.
type UnimplementedNodeServer struct {
}

func (UnimplementedNodeServer) HandleHandshake(context.Context, *Handshake) (*Handshake, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HandleHandshake not implemented")
}
func (UnimplementedNodeServer) HandleTransaction(context.Context, *Transaction) (*Ack, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HandleTransaction not implemented")
}
func (UnimplementedNodeServer) HandleBlock(context.Context, *Block) (*Ack, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HandleBlock not implemented")
}
func (UnimplementedNodeServer) HandleVote(context.Context, *Vote) (*Ack, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HandleVote not implemented")
}
func (UnimplementedNodeServer) mustEmbedUnimplementedNodeServer() {}

// UnsafeNodeServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NodeServer will
// result in compilation errors.
type UnsafeNodeServer interface {
	mustEmbedUnimplementedNodeServer()
}

func RegisterNodeServer(s grpc.ServiceRegistrar, srv NodeServer) {
	s.RegisterService(&Node_ServiceDesc, srv)
}

func _Node_HandleHandshake_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Handshake)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServer).HandleHandshake(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Node_HandleHandshake_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServer).HandleHandshake(ctx, req.(*Handshake))
	}
	return interceptor(ctx, in, info, handler)
}

func _Node_HandleTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Transaction)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServer).HandleTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Node_HandleTransaction_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServer).HandleTransaction(ctx, req.(*Transaction))
	}
	return interceptor(ctx, in, info, handler)
}

func _Node_HandleBlock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Block)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServer).HandleBlock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Node_HandleBlock_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServer).HandleBlock(ctx, req.(*Block))
	}
	return interceptor(ctx, in, info, handler)
}

func _Node_HandleVote_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Vote)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServer).HandleVote(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Node_HandleVote_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServer).HandleVote(ctx, req.(*Vote))
	}
	return interceptor(ctx, in, info, handler)
}

// Node_ServiceDesc is the grpc.ServiceDesc for Node service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Node_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Node",
	HandlerType: (*NodeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "HandleHandshake",
			Handler:    _Node_HandleHandshake_Handler,
		},
		{
			MethodName: "HandleTransaction",
			Handler:    _Node_HandleTransaction_Handler,
		},
		{
			MethodName: "HandleBlock",
			Handler:    _Node_HandleBlock_Handler,
		},
		{
			MethodName: "HandleVote",
			Handler:    _Node_HandleVote_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/blockchain.proto",
}
