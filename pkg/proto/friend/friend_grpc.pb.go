// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: friend/friend.proto

package friend

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
	Friend_AddFriend_FullMethodName               = "/friend.friend/addFriend"
	Friend_GetFriendApplyList_FullMethodName      = "/friend.friend/getFriendApplyList"
	Friend_GetSelfApplyList_FullMethodName        = "/friend.friend/getSelfApplyList"
	Friend_GetFriendList_FullMethodName           = "/friend.friend/getFriendList"
	Friend_AddBlacklist_FullMethodName            = "/friend.friend/addBlacklist"
	Friend_RemoveBlacklist_FullMethodName         = "/friend.friend/removeBlacklist"
	Friend_IsFriend_FullMethodName                = "/friend.friend/isFriend"
	Friend_IsInBlackList_FullMethodName           = "/friend.friend/isInBlackList"
	Friend_GetBlacklist_FullMethodName            = "/friend.friend/getBlacklist"
	Friend_DeleteFriend_FullMethodName            = "/friend.friend/deleteFriend"
	Friend_AddFriendResponse_FullMethodName       = "/friend.friend/addFriendResponse"
	Friend_SetFriendRemark_FullMethodName         = "/friend.friend/setFriendRemark"
	Friend_ImportFriend_FullMethodName            = "/friend.friend/importFriend"
	Friend_BatchSyncAgentFriend_FullMethodName    = "/friend.friend/BatchSyncAgentFriend"
	Friend_BatchDelSyncAgentFriend_FullMethodName = "/friend.friend/BatchDelSyncAgentFriend"
)

// FriendClient is the client API for Friend service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FriendClient interface {
	// rpc getFriendsInfo(GetFriendsInfoReq) returns(GetFriendInfoResp);
	AddFriend(ctx context.Context, in *AddFriendReq, opts ...grpc.CallOption) (*AddFriendResp, error)
	GetFriendApplyList(ctx context.Context, in *GetFriendApplyListReq, opts ...grpc.CallOption) (*GetFriendApplyListResp, error)
	GetSelfApplyList(ctx context.Context, in *GetSelfApplyListReq, opts ...grpc.CallOption) (*GetSelfApplyListResp, error)
	GetFriendList(ctx context.Context, in *GetFriendListReq, opts ...grpc.CallOption) (*GetFriendListResp, error)
	AddBlacklist(ctx context.Context, in *AddBlacklistReq, opts ...grpc.CallOption) (*AddBlacklistResp, error)
	RemoveBlacklist(ctx context.Context, in *RemoveBlacklistReq, opts ...grpc.CallOption) (*RemoveBlacklistResp, error)
	IsFriend(ctx context.Context, in *IsFriendReq, opts ...grpc.CallOption) (*IsFriendResp, error)
	IsInBlackList(ctx context.Context, in *IsInBlackListReq, opts ...grpc.CallOption) (*IsInBlackListResp, error)
	GetBlacklist(ctx context.Context, in *GetBlacklistReq, opts ...grpc.CallOption) (*GetBlacklistResp, error)
	DeleteFriend(ctx context.Context, in *DeleteFriendReq, opts ...grpc.CallOption) (*DeleteFriendResp, error)
	AddFriendResponse(ctx context.Context, in *AddFriendResponseReq, opts ...grpc.CallOption) (*AddFriendResponseResp, error)
	SetFriendRemark(ctx context.Context, in *SetFriendRemarkReq, opts ...grpc.CallOption) (*SetFriendRemarkResp, error)
	ImportFriend(ctx context.Context, in *ImportFriendReq, opts ...grpc.CallOption) (*ImportFriendResp, error)
	BatchSyncAgentFriend(ctx context.Context, in *BatchSyncAgentFriendReq, opts ...grpc.CallOption) (*BatchSyncAgentFriendResp, error)
	BatchDelSyncAgentFriend(ctx context.Context, in *BatchDelSyncAgentFriendReq, opts ...grpc.CallOption) (*BatchDelSyncAgentFriendResp, error)
}

type friendClient struct {
	cc grpc.ClientConnInterface
}

func NewFriendClient(cc grpc.ClientConnInterface) FriendClient {
	return &friendClient{cc}
}

func (c *friendClient) AddFriend(ctx context.Context, in *AddFriendReq, opts ...grpc.CallOption) (*AddFriendResp, error) {
	out := new(AddFriendResp)
	err := c.cc.Invoke(ctx, Friend_AddFriend_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *friendClient) GetFriendApplyList(ctx context.Context, in *GetFriendApplyListReq, opts ...grpc.CallOption) (*GetFriendApplyListResp, error) {
	out := new(GetFriendApplyListResp)
	err := c.cc.Invoke(ctx, Friend_GetFriendApplyList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *friendClient) GetSelfApplyList(ctx context.Context, in *GetSelfApplyListReq, opts ...grpc.CallOption) (*GetSelfApplyListResp, error) {
	out := new(GetSelfApplyListResp)
	err := c.cc.Invoke(ctx, Friend_GetSelfApplyList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *friendClient) GetFriendList(ctx context.Context, in *GetFriendListReq, opts ...grpc.CallOption) (*GetFriendListResp, error) {
	out := new(GetFriendListResp)
	err := c.cc.Invoke(ctx, Friend_GetFriendList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *friendClient) AddBlacklist(ctx context.Context, in *AddBlacklistReq, opts ...grpc.CallOption) (*AddBlacklistResp, error) {
	out := new(AddBlacklistResp)
	err := c.cc.Invoke(ctx, Friend_AddBlacklist_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *friendClient) RemoveBlacklist(ctx context.Context, in *RemoveBlacklistReq, opts ...grpc.CallOption) (*RemoveBlacklistResp, error) {
	out := new(RemoveBlacklistResp)
	err := c.cc.Invoke(ctx, Friend_RemoveBlacklist_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *friendClient) IsFriend(ctx context.Context, in *IsFriendReq, opts ...grpc.CallOption) (*IsFriendResp, error) {
	out := new(IsFriendResp)
	err := c.cc.Invoke(ctx, Friend_IsFriend_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *friendClient) IsInBlackList(ctx context.Context, in *IsInBlackListReq, opts ...grpc.CallOption) (*IsInBlackListResp, error) {
	out := new(IsInBlackListResp)
	err := c.cc.Invoke(ctx, Friend_IsInBlackList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *friendClient) GetBlacklist(ctx context.Context, in *GetBlacklistReq, opts ...grpc.CallOption) (*GetBlacklistResp, error) {
	out := new(GetBlacklistResp)
	err := c.cc.Invoke(ctx, Friend_GetBlacklist_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *friendClient) DeleteFriend(ctx context.Context, in *DeleteFriendReq, opts ...grpc.CallOption) (*DeleteFriendResp, error) {
	out := new(DeleteFriendResp)
	err := c.cc.Invoke(ctx, Friend_DeleteFriend_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *friendClient) AddFriendResponse(ctx context.Context, in *AddFriendResponseReq, opts ...grpc.CallOption) (*AddFriendResponseResp, error) {
	out := new(AddFriendResponseResp)
	err := c.cc.Invoke(ctx, Friend_AddFriendResponse_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *friendClient) SetFriendRemark(ctx context.Context, in *SetFriendRemarkReq, opts ...grpc.CallOption) (*SetFriendRemarkResp, error) {
	out := new(SetFriendRemarkResp)
	err := c.cc.Invoke(ctx, Friend_SetFriendRemark_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *friendClient) ImportFriend(ctx context.Context, in *ImportFriendReq, opts ...grpc.CallOption) (*ImportFriendResp, error) {
	out := new(ImportFriendResp)
	err := c.cc.Invoke(ctx, Friend_ImportFriend_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *friendClient) BatchSyncAgentFriend(ctx context.Context, in *BatchSyncAgentFriendReq, opts ...grpc.CallOption) (*BatchSyncAgentFriendResp, error) {
	out := new(BatchSyncAgentFriendResp)
	err := c.cc.Invoke(ctx, Friend_BatchSyncAgentFriend_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *friendClient) BatchDelSyncAgentFriend(ctx context.Context, in *BatchDelSyncAgentFriendReq, opts ...grpc.CallOption) (*BatchDelSyncAgentFriendResp, error) {
	out := new(BatchDelSyncAgentFriendResp)
	err := c.cc.Invoke(ctx, Friend_BatchDelSyncAgentFriend_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FriendServer is the server API for Friend service.
// All implementations must embed UnimplementedFriendServer
// for forward compatibility
type FriendServer interface {
	// rpc getFriendsInfo(GetFriendsInfoReq) returns(GetFriendInfoResp);
	AddFriend(context.Context, *AddFriendReq) (*AddFriendResp, error)
	GetFriendApplyList(context.Context, *GetFriendApplyListReq) (*GetFriendApplyListResp, error)
	GetSelfApplyList(context.Context, *GetSelfApplyListReq) (*GetSelfApplyListResp, error)
	GetFriendList(context.Context, *GetFriendListReq) (*GetFriendListResp, error)
	AddBlacklist(context.Context, *AddBlacklistReq) (*AddBlacklistResp, error)
	RemoveBlacklist(context.Context, *RemoveBlacklistReq) (*RemoveBlacklistResp, error)
	IsFriend(context.Context, *IsFriendReq) (*IsFriendResp, error)
	IsInBlackList(context.Context, *IsInBlackListReq) (*IsInBlackListResp, error)
	GetBlacklist(context.Context, *GetBlacklistReq) (*GetBlacklistResp, error)
	DeleteFriend(context.Context, *DeleteFriendReq) (*DeleteFriendResp, error)
	AddFriendResponse(context.Context, *AddFriendResponseReq) (*AddFriendResponseResp, error)
	SetFriendRemark(context.Context, *SetFriendRemarkReq) (*SetFriendRemarkResp, error)
	ImportFriend(context.Context, *ImportFriendReq) (*ImportFriendResp, error)
	BatchSyncAgentFriend(context.Context, *BatchSyncAgentFriendReq) (*BatchSyncAgentFriendResp, error)
	BatchDelSyncAgentFriend(context.Context, *BatchDelSyncAgentFriendReq) (*BatchDelSyncAgentFriendResp, error)
	mustEmbedUnimplementedFriendServer()
}

// UnimplementedFriendServer must be embedded to have forward compatible implementations.
type UnimplementedFriendServer struct {
}

func (UnimplementedFriendServer) AddFriend(context.Context, *AddFriendReq) (*AddFriendResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddFriend not implemented")
}
func (UnimplementedFriendServer) GetFriendApplyList(context.Context, *GetFriendApplyListReq) (*GetFriendApplyListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFriendApplyList not implemented")
}
func (UnimplementedFriendServer) GetSelfApplyList(context.Context, *GetSelfApplyListReq) (*GetSelfApplyListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSelfApplyList not implemented")
}
func (UnimplementedFriendServer) GetFriendList(context.Context, *GetFriendListReq) (*GetFriendListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFriendList not implemented")
}
func (UnimplementedFriendServer) AddBlacklist(context.Context, *AddBlacklistReq) (*AddBlacklistResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddBlacklist not implemented")
}
func (UnimplementedFriendServer) RemoveBlacklist(context.Context, *RemoveBlacklistReq) (*RemoveBlacklistResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveBlacklist not implemented")
}
func (UnimplementedFriendServer) IsFriend(context.Context, *IsFriendReq) (*IsFriendResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsFriend not implemented")
}
func (UnimplementedFriendServer) IsInBlackList(context.Context, *IsInBlackListReq) (*IsInBlackListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsInBlackList not implemented")
}
func (UnimplementedFriendServer) GetBlacklist(context.Context, *GetBlacklistReq) (*GetBlacklistResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBlacklist not implemented")
}
func (UnimplementedFriendServer) DeleteFriend(context.Context, *DeleteFriendReq) (*DeleteFriendResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteFriend not implemented")
}
func (UnimplementedFriendServer) AddFriendResponse(context.Context, *AddFriendResponseReq) (*AddFriendResponseResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddFriendResponse not implemented")
}
func (UnimplementedFriendServer) SetFriendRemark(context.Context, *SetFriendRemarkReq) (*SetFriendRemarkResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetFriendRemark not implemented")
}
func (UnimplementedFriendServer) ImportFriend(context.Context, *ImportFriendReq) (*ImportFriendResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ImportFriend not implemented")
}
func (UnimplementedFriendServer) BatchSyncAgentFriend(context.Context, *BatchSyncAgentFriendReq) (*BatchSyncAgentFriendResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BatchSyncAgentFriend not implemented")
}
func (UnimplementedFriendServer) BatchDelSyncAgentFriend(context.Context, *BatchDelSyncAgentFriendReq) (*BatchDelSyncAgentFriendResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BatchDelSyncAgentFriend not implemented")
}
func (UnimplementedFriendServer) mustEmbedUnimplementedFriendServer() {}

// UnsafeFriendServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FriendServer will
// result in compilation errors.
type UnsafeFriendServer interface {
	mustEmbedUnimplementedFriendServer()
}

func RegisterFriendServer(s grpc.ServiceRegistrar, srv FriendServer) {
	s.RegisterService(&Friend_ServiceDesc, srv)
}

func _Friend_AddFriend_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddFriendReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FriendServer).AddFriend(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Friend_AddFriend_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FriendServer).AddFriend(ctx, req.(*AddFriendReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Friend_GetFriendApplyList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFriendApplyListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FriendServer).GetFriendApplyList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Friend_GetFriendApplyList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FriendServer).GetFriendApplyList(ctx, req.(*GetFriendApplyListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Friend_GetSelfApplyList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSelfApplyListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FriendServer).GetSelfApplyList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Friend_GetSelfApplyList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FriendServer).GetSelfApplyList(ctx, req.(*GetSelfApplyListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Friend_GetFriendList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFriendListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FriendServer).GetFriendList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Friend_GetFriendList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FriendServer).GetFriendList(ctx, req.(*GetFriendListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Friend_AddBlacklist_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddBlacklistReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FriendServer).AddBlacklist(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Friend_AddBlacklist_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FriendServer).AddBlacklist(ctx, req.(*AddBlacklistReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Friend_RemoveBlacklist_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveBlacklistReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FriendServer).RemoveBlacklist(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Friend_RemoveBlacklist_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FriendServer).RemoveBlacklist(ctx, req.(*RemoveBlacklistReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Friend_IsFriend_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IsFriendReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FriendServer).IsFriend(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Friend_IsFriend_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FriendServer).IsFriend(ctx, req.(*IsFriendReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Friend_IsInBlackList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IsInBlackListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FriendServer).IsInBlackList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Friend_IsInBlackList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FriendServer).IsInBlackList(ctx, req.(*IsInBlackListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Friend_GetBlacklist_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBlacklistReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FriendServer).GetBlacklist(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Friend_GetBlacklist_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FriendServer).GetBlacklist(ctx, req.(*GetBlacklistReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Friend_DeleteFriend_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteFriendReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FriendServer).DeleteFriend(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Friend_DeleteFriend_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FriendServer).DeleteFriend(ctx, req.(*DeleteFriendReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Friend_AddFriendResponse_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddFriendResponseReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FriendServer).AddFriendResponse(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Friend_AddFriendResponse_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FriendServer).AddFriendResponse(ctx, req.(*AddFriendResponseReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Friend_SetFriendRemark_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetFriendRemarkReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FriendServer).SetFriendRemark(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Friend_SetFriendRemark_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FriendServer).SetFriendRemark(ctx, req.(*SetFriendRemarkReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Friend_ImportFriend_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ImportFriendReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FriendServer).ImportFriend(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Friend_ImportFriend_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FriendServer).ImportFriend(ctx, req.(*ImportFriendReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Friend_BatchSyncAgentFriend_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BatchSyncAgentFriendReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FriendServer).BatchSyncAgentFriend(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Friend_BatchSyncAgentFriend_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FriendServer).BatchSyncAgentFriend(ctx, req.(*BatchSyncAgentFriendReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Friend_BatchDelSyncAgentFriend_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BatchDelSyncAgentFriendReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FriendServer).BatchDelSyncAgentFriend(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Friend_BatchDelSyncAgentFriend_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FriendServer).BatchDelSyncAgentFriend(ctx, req.(*BatchDelSyncAgentFriendReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Friend_ServiceDesc is the grpc.ServiceDesc for Friend service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Friend_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "friend.friend",
	HandlerType: (*FriendServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "addFriend",
			Handler:    _Friend_AddFriend_Handler,
		},
		{
			MethodName: "getFriendApplyList",
			Handler:    _Friend_GetFriendApplyList_Handler,
		},
		{
			MethodName: "getSelfApplyList",
			Handler:    _Friend_GetSelfApplyList_Handler,
		},
		{
			MethodName: "getFriendList",
			Handler:    _Friend_GetFriendList_Handler,
		},
		{
			MethodName: "addBlacklist",
			Handler:    _Friend_AddBlacklist_Handler,
		},
		{
			MethodName: "removeBlacklist",
			Handler:    _Friend_RemoveBlacklist_Handler,
		},
		{
			MethodName: "isFriend",
			Handler:    _Friend_IsFriend_Handler,
		},
		{
			MethodName: "isInBlackList",
			Handler:    _Friend_IsInBlackList_Handler,
		},
		{
			MethodName: "getBlacklist",
			Handler:    _Friend_GetBlacklist_Handler,
		},
		{
			MethodName: "deleteFriend",
			Handler:    _Friend_DeleteFriend_Handler,
		},
		{
			MethodName: "addFriendResponse",
			Handler:    _Friend_AddFriendResponse_Handler,
		},
		{
			MethodName: "setFriendRemark",
			Handler:    _Friend_SetFriendRemark_Handler,
		},
		{
			MethodName: "importFriend",
			Handler:    _Friend_ImportFriend_Handler,
		},
		{
			MethodName: "BatchSyncAgentFriend",
			Handler:    _Friend_BatchSyncAgentFriend_Handler,
		},
		{
			MethodName: "BatchDelSyncAgentFriend",
			Handler:    _Friend_BatchDelSyncAgentFriend_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "friend/friend.proto",
}
