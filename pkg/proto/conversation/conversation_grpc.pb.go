// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: conversation/conversation.proto

package conversation

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

const (
	Conversation_ModifyConversationField_FullMethodName = "/conversation.conversation/ModifyConversationField"
)

// ConversationClient is the client API for Conversation service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ConversationClient interface {
	ModifyConversationField(ctx context.Context, in *ModifyConversationFieldReq, opts ...grpc.CallOption) (*ModifyConversationFieldResp, error)
}

type conversationClient struct {
	cc grpc.ClientConnInterface
}

func NewConversationClient(cc grpc.ClientConnInterface) ConversationClient {
	return &conversationClient{cc}
}

func (c *conversationClient) ModifyConversationField(ctx context.Context, in *ModifyConversationFieldReq, opts ...grpc.CallOption) (*ModifyConversationFieldResp, error) {
	out := new(ModifyConversationFieldResp)
	err := c.cc.Invoke(ctx, Conversation_ModifyConversationField_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ConversationServer is the server API for Conversation service.
// All implementations must embed UnimplementedConversationServer
// for forward compatibility
type ConversationServer interface {
	ModifyConversationField(context.Context, *ModifyConversationFieldReq) (*ModifyConversationFieldResp, error)
	mustEmbedUnimplementedConversationServer()
}

// UnimplementedConversationServer must be embedded to have forward compatible implementations.
type UnimplementedConversationServer struct {
}

func (UnimplementedConversationServer) ModifyConversationField(context.Context, *ModifyConversationFieldReq) (*ModifyConversationFieldResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ModifyConversationField not implemented")
}
func (UnimplementedConversationServer) mustEmbedUnimplementedConversationServer() {}

// UnsafeConversationServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ConversationServer will
// result in compilation errors.
type UnsafeConversationServer interface {
	mustEmbedUnimplementedConversationServer()
}

func RegisterConversationServer(s grpc.ServiceRegistrar, srv ConversationServer) {
	s.RegisterService(&Conversation_ServiceDesc, srv)
}

func _Conversation_ModifyConversationField_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ModifyConversationFieldReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConversationServer).ModifyConversationField(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Conversation_ModifyConversationField_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConversationServer).ModifyConversationField(ctx, req.(*ModifyConversationFieldReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Conversation_ServiceDesc is the grpc.ServiceDesc for Conversation service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Conversation_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "conversation.conversation",
	HandlerType: (*ConversationServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ModifyConversationField",
			Handler:    _Conversation_ModifyConversationField_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "conversation/conversation.proto",
}
