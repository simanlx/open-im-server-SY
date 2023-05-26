// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: msg/msg.proto

package msg

import (
	sdk_ws "Open_IM/pkg/proto/sdk_ws"
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
	Msg_GetMaxAndMinSeq_FullMethodName                  = "/msg.msg/GetMaxAndMinSeq"
	Msg_PullMessageBySeqList_FullMethodName             = "/msg.msg/PullMessageBySeqList"
	Msg_SendMsg_FullMethodName                          = "/msg.msg/SendMsg"
	Msg_DelMsgList_FullMethodName                       = "/msg.msg/DelMsgList"
	Msg_DelSuperGroupMsg_FullMethodName                 = "/msg.msg/DelSuperGroupMsg"
	Msg_ClearMsg_FullMethodName                         = "/msg.msg/ClearMsg"
	Msg_SetMsgMinSeq_FullMethodName                     = "/msg.msg/SetMsgMinSeq"
	Msg_SetSendMsgStatus_FullMethodName                 = "/msg.msg/SetSendMsgStatus"
	Msg_GetSendMsgStatus_FullMethodName                 = "/msg.msg/GetSendMsgStatus"
	Msg_GetSuperGroupMsg_FullMethodName                 = "/msg.msg/GetSuperGroupMsg"
	Msg_GetWriteDiffMsg_FullMethodName                  = "/msg.msg/GetWriteDiffMsg"
	Msg_SetMessageReactionExtensions_FullMethodName     = "/msg.msg/SetMessageReactionExtensions"
	Msg_GetMessageListReactionExtensions_FullMethodName = "/msg.msg/GetMessageListReactionExtensions"
	Msg_AddMessageReactionExtensions_FullMethodName     = "/msg.msg/AddMessageReactionExtensions"
	Msg_DeleteMessageReactionExtensions_FullMethodName  = "/msg.msg/DeleteMessageReactionExtensions"
	Msg_MsgCollect_FullMethodName                       = "/msg.msg/MsgCollect"
	Msg_MsgCollectDel_FullMethodName                    = "/msg.msg/MsgCollectDel"
	Msg_MsgCollectList_FullMethodName                   = "/msg.msg/MsgCollectList"
)

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MsgClient interface {
	GetMaxAndMinSeq(ctx context.Context, in *sdk_ws.GetMaxAndMinSeqReq, opts ...grpc.CallOption) (*sdk_ws.GetMaxAndMinSeqResp, error)
	PullMessageBySeqList(ctx context.Context, in *sdk_ws.PullMessageBySeqListReq, opts ...grpc.CallOption) (*sdk_ws.PullMessageBySeqListResp, error)
	SendMsg(ctx context.Context, in *SendMsgReq, opts ...grpc.CallOption) (*SendMsgResp, error)
	DelMsgList(ctx context.Context, in *sdk_ws.DelMsgListReq, opts ...grpc.CallOption) (*sdk_ws.DelMsgListResp, error)
	DelSuperGroupMsg(ctx context.Context, in *DelSuperGroupMsgReq, opts ...grpc.CallOption) (*DelSuperGroupMsgResp, error)
	ClearMsg(ctx context.Context, in *ClearMsgReq, opts ...grpc.CallOption) (*ClearMsgResp, error)
	SetMsgMinSeq(ctx context.Context, in *SetMsgMinSeqReq, opts ...grpc.CallOption) (*SetMsgMinSeqResp, error)
	SetSendMsgStatus(ctx context.Context, in *SetSendMsgStatusReq, opts ...grpc.CallOption) (*SetSendMsgStatusResp, error)
	GetSendMsgStatus(ctx context.Context, in *GetSendMsgStatusReq, opts ...grpc.CallOption) (*GetSendMsgStatusResp, error)
	GetSuperGroupMsg(ctx context.Context, in *GetSuperGroupMsgReq, opts ...grpc.CallOption) (*GetSuperGroupMsgResp, error)
	GetWriteDiffMsg(ctx context.Context, in *GetWriteDiffMsgReq, opts ...grpc.CallOption) (*GetWriteDiffMsgResp, error)
	// modify msg
	SetMessageReactionExtensions(ctx context.Context, in *SetMessageReactionExtensionsReq, opts ...grpc.CallOption) (*SetMessageReactionExtensionsResp, error)
	GetMessageListReactionExtensions(ctx context.Context, in *GetMessageListReactionExtensionsReq, opts ...grpc.CallOption) (*GetMessageListReactionExtensionsResp, error)
	AddMessageReactionExtensions(ctx context.Context, in *AddMessageReactionExtensionsReq, opts ...grpc.CallOption) (*AddMessageReactionExtensionsResp, error)
	DeleteMessageReactionExtensions(ctx context.Context, in *DeleteMessageListReactionExtensionsReq, opts ...grpc.CallOption) (*DeleteMessageListReactionExtensionsResp, error)
	//消息收藏
	MsgCollect(ctx context.Context, in *MsgCollectReq, opts ...grpc.CallOption) (*MsgCollectResp, error)
	//删除消息收藏
	MsgCollectDel(ctx context.Context, in *MsgCollectDelReq, opts ...grpc.CallOption) (*MsgCollectDelResp, error)
	//消息收藏列表
	MsgCollectList(ctx context.Context, in *MsgCollectListReq, opts ...grpc.CallOption) (*MsgCollectListResp, error)
}

type msgClient struct {
	cc grpc.ClientConnInterface
}

func NewMsgClient(cc grpc.ClientConnInterface) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) GetMaxAndMinSeq(ctx context.Context, in *sdk_ws.GetMaxAndMinSeqReq, opts ...grpc.CallOption) (*sdk_ws.GetMaxAndMinSeqResp, error) {
	out := new(sdk_ws.GetMaxAndMinSeqResp)
	err := c.cc.Invoke(ctx, Msg_GetMaxAndMinSeq_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) PullMessageBySeqList(ctx context.Context, in *sdk_ws.PullMessageBySeqListReq, opts ...grpc.CallOption) (*sdk_ws.PullMessageBySeqListResp, error) {
	out := new(sdk_ws.PullMessageBySeqListResp)
	err := c.cc.Invoke(ctx, Msg_PullMessageBySeqList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) SendMsg(ctx context.Context, in *SendMsgReq, opts ...grpc.CallOption) (*SendMsgResp, error) {
	out := new(SendMsgResp)
	err := c.cc.Invoke(ctx, Msg_SendMsg_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) DelMsgList(ctx context.Context, in *sdk_ws.DelMsgListReq, opts ...grpc.CallOption) (*sdk_ws.DelMsgListResp, error) {
	out := new(sdk_ws.DelMsgListResp)
	err := c.cc.Invoke(ctx, Msg_DelMsgList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) DelSuperGroupMsg(ctx context.Context, in *DelSuperGroupMsgReq, opts ...grpc.CallOption) (*DelSuperGroupMsgResp, error) {
	out := new(DelSuperGroupMsgResp)
	err := c.cc.Invoke(ctx, Msg_DelSuperGroupMsg_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) ClearMsg(ctx context.Context, in *ClearMsgReq, opts ...grpc.CallOption) (*ClearMsgResp, error) {
	out := new(ClearMsgResp)
	err := c.cc.Invoke(ctx, Msg_ClearMsg_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) SetMsgMinSeq(ctx context.Context, in *SetMsgMinSeqReq, opts ...grpc.CallOption) (*SetMsgMinSeqResp, error) {
	out := new(SetMsgMinSeqResp)
	err := c.cc.Invoke(ctx, Msg_SetMsgMinSeq_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) SetSendMsgStatus(ctx context.Context, in *SetSendMsgStatusReq, opts ...grpc.CallOption) (*SetSendMsgStatusResp, error) {
	out := new(SetSendMsgStatusResp)
	err := c.cc.Invoke(ctx, Msg_SetSendMsgStatus_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) GetSendMsgStatus(ctx context.Context, in *GetSendMsgStatusReq, opts ...grpc.CallOption) (*GetSendMsgStatusResp, error) {
	out := new(GetSendMsgStatusResp)
	err := c.cc.Invoke(ctx, Msg_GetSendMsgStatus_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) GetSuperGroupMsg(ctx context.Context, in *GetSuperGroupMsgReq, opts ...grpc.CallOption) (*GetSuperGroupMsgResp, error) {
	out := new(GetSuperGroupMsgResp)
	err := c.cc.Invoke(ctx, Msg_GetSuperGroupMsg_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) GetWriteDiffMsg(ctx context.Context, in *GetWriteDiffMsgReq, opts ...grpc.CallOption) (*GetWriteDiffMsgResp, error) {
	out := new(GetWriteDiffMsgResp)
	err := c.cc.Invoke(ctx, Msg_GetWriteDiffMsg_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) SetMessageReactionExtensions(ctx context.Context, in *SetMessageReactionExtensionsReq, opts ...grpc.CallOption) (*SetMessageReactionExtensionsResp, error) {
	out := new(SetMessageReactionExtensionsResp)
	err := c.cc.Invoke(ctx, Msg_SetMessageReactionExtensions_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) GetMessageListReactionExtensions(ctx context.Context, in *GetMessageListReactionExtensionsReq, opts ...grpc.CallOption) (*GetMessageListReactionExtensionsResp, error) {
	out := new(GetMessageListReactionExtensionsResp)
	err := c.cc.Invoke(ctx, Msg_GetMessageListReactionExtensions_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) AddMessageReactionExtensions(ctx context.Context, in *AddMessageReactionExtensionsReq, opts ...grpc.CallOption) (*AddMessageReactionExtensionsResp, error) {
	out := new(AddMessageReactionExtensionsResp)
	err := c.cc.Invoke(ctx, Msg_AddMessageReactionExtensions_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) DeleteMessageReactionExtensions(ctx context.Context, in *DeleteMessageListReactionExtensionsReq, opts ...grpc.CallOption) (*DeleteMessageListReactionExtensionsResp, error) {
	out := new(DeleteMessageListReactionExtensionsResp)
	err := c.cc.Invoke(ctx, Msg_DeleteMessageReactionExtensions_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) MsgCollect(ctx context.Context, in *MsgCollectReq, opts ...grpc.CallOption) (*MsgCollectResp, error) {
	out := new(MsgCollectResp)
	err := c.cc.Invoke(ctx, Msg_MsgCollect_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) MsgCollectDel(ctx context.Context, in *MsgCollectDelReq, opts ...grpc.CallOption) (*MsgCollectDelResp, error) {
	out := new(MsgCollectDelResp)
	err := c.cc.Invoke(ctx, Msg_MsgCollectDel_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) MsgCollectList(ctx context.Context, in *MsgCollectListReq, opts ...grpc.CallOption) (*MsgCollectListResp, error) {
	out := new(MsgCollectListResp)
	err := c.cc.Invoke(ctx, Msg_MsgCollectList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
// All implementations must embed UnimplementedMsgServer
// for forward compatibility
type MsgServer interface {
	GetMaxAndMinSeq(context.Context, *sdk_ws.GetMaxAndMinSeqReq) (*sdk_ws.GetMaxAndMinSeqResp, error)
	PullMessageBySeqList(context.Context, *sdk_ws.PullMessageBySeqListReq) (*sdk_ws.PullMessageBySeqListResp, error)
	SendMsg(context.Context, *SendMsgReq) (*SendMsgResp, error)
	DelMsgList(context.Context, *sdk_ws.DelMsgListReq) (*sdk_ws.DelMsgListResp, error)
	DelSuperGroupMsg(context.Context, *DelSuperGroupMsgReq) (*DelSuperGroupMsgResp, error)
	ClearMsg(context.Context, *ClearMsgReq) (*ClearMsgResp, error)
	SetMsgMinSeq(context.Context, *SetMsgMinSeqReq) (*SetMsgMinSeqResp, error)
	SetSendMsgStatus(context.Context, *SetSendMsgStatusReq) (*SetSendMsgStatusResp, error)
	GetSendMsgStatus(context.Context, *GetSendMsgStatusReq) (*GetSendMsgStatusResp, error)
	GetSuperGroupMsg(context.Context, *GetSuperGroupMsgReq) (*GetSuperGroupMsgResp, error)
	GetWriteDiffMsg(context.Context, *GetWriteDiffMsgReq) (*GetWriteDiffMsgResp, error)
	// modify msg
	SetMessageReactionExtensions(context.Context, *SetMessageReactionExtensionsReq) (*SetMessageReactionExtensionsResp, error)
	GetMessageListReactionExtensions(context.Context, *GetMessageListReactionExtensionsReq) (*GetMessageListReactionExtensionsResp, error)
	AddMessageReactionExtensions(context.Context, *AddMessageReactionExtensionsReq) (*AddMessageReactionExtensionsResp, error)
	DeleteMessageReactionExtensions(context.Context, *DeleteMessageListReactionExtensionsReq) (*DeleteMessageListReactionExtensionsResp, error)
	//消息收藏
	MsgCollect(context.Context, *MsgCollectReq) (*MsgCollectResp, error)
	//删除消息收藏
	MsgCollectDel(context.Context, *MsgCollectDelReq) (*MsgCollectDelResp, error)
	//消息收藏列表
	MsgCollectList(context.Context, *MsgCollectListReq) (*MsgCollectListResp, error)
	mustEmbedUnimplementedMsgServer()
}

// UnimplementedMsgServer must be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (UnimplementedMsgServer) GetMaxAndMinSeq(context.Context, *sdk_ws.GetMaxAndMinSeqReq) (*sdk_ws.GetMaxAndMinSeqResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMaxAndMinSeq not implemented")
}
func (UnimplementedMsgServer) PullMessageBySeqList(context.Context, *sdk_ws.PullMessageBySeqListReq) (*sdk_ws.PullMessageBySeqListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PullMessageBySeqList not implemented")
}
func (UnimplementedMsgServer) SendMsg(context.Context, *SendMsgReq) (*SendMsgResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendMsg not implemented")
}
func (UnimplementedMsgServer) DelMsgList(context.Context, *sdk_ws.DelMsgListReq) (*sdk_ws.DelMsgListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DelMsgList not implemented")
}
func (UnimplementedMsgServer) DelSuperGroupMsg(context.Context, *DelSuperGroupMsgReq) (*DelSuperGroupMsgResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DelSuperGroupMsg not implemented")
}
func (UnimplementedMsgServer) ClearMsg(context.Context, *ClearMsgReq) (*ClearMsgResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ClearMsg not implemented")
}
func (UnimplementedMsgServer) SetMsgMinSeq(context.Context, *SetMsgMinSeqReq) (*SetMsgMinSeqResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetMsgMinSeq not implemented")
}
func (UnimplementedMsgServer) SetSendMsgStatus(context.Context, *SetSendMsgStatusReq) (*SetSendMsgStatusResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetSendMsgStatus not implemented")
}
func (UnimplementedMsgServer) GetSendMsgStatus(context.Context, *GetSendMsgStatusReq) (*GetSendMsgStatusResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSendMsgStatus not implemented")
}
func (UnimplementedMsgServer) GetSuperGroupMsg(context.Context, *GetSuperGroupMsgReq) (*GetSuperGroupMsgResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSuperGroupMsg not implemented")
}
func (UnimplementedMsgServer) GetWriteDiffMsg(context.Context, *GetWriteDiffMsgReq) (*GetWriteDiffMsgResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetWriteDiffMsg not implemented")
}
func (UnimplementedMsgServer) SetMessageReactionExtensions(context.Context, *SetMessageReactionExtensionsReq) (*SetMessageReactionExtensionsResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetMessageReactionExtensions not implemented")
}
func (UnimplementedMsgServer) GetMessageListReactionExtensions(context.Context, *GetMessageListReactionExtensionsReq) (*GetMessageListReactionExtensionsResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMessageListReactionExtensions not implemented")
}
func (UnimplementedMsgServer) AddMessageReactionExtensions(context.Context, *AddMessageReactionExtensionsReq) (*AddMessageReactionExtensionsResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddMessageReactionExtensions not implemented")
}
func (UnimplementedMsgServer) DeleteMessageReactionExtensions(context.Context, *DeleteMessageListReactionExtensionsReq) (*DeleteMessageListReactionExtensionsResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteMessageReactionExtensions not implemented")
}
func (UnimplementedMsgServer) MsgCollect(context.Context, *MsgCollectReq) (*MsgCollectResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MsgCollect not implemented")
}
func (UnimplementedMsgServer) MsgCollectDel(context.Context, *MsgCollectDelReq) (*MsgCollectDelResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MsgCollectDel not implemented")
}
func (UnimplementedMsgServer) MsgCollectList(context.Context, *MsgCollectListReq) (*MsgCollectListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MsgCollectList not implemented")
}
func (UnimplementedMsgServer) mustEmbedUnimplementedMsgServer() {}

// UnsafeMsgServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MsgServer will
// result in compilation errors.
type UnsafeMsgServer interface {
	mustEmbedUnimplementedMsgServer()
}

func RegisterMsgServer(s grpc.ServiceRegistrar, srv MsgServer) {
	s.RegisterService(&Msg_ServiceDesc, srv)
}

func _Msg_GetMaxAndMinSeq_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(sdk_ws.GetMaxAndMinSeqReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).GetMaxAndMinSeq(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_GetMaxAndMinSeq_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).GetMaxAndMinSeq(ctx, req.(*sdk_ws.GetMaxAndMinSeqReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_PullMessageBySeqList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(sdk_ws.PullMessageBySeqListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).PullMessageBySeqList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_PullMessageBySeqList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).PullMessageBySeqList(ctx, req.(*sdk_ws.PullMessageBySeqListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_SendMsg_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendMsgReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).SendMsg(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_SendMsg_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).SendMsg(ctx, req.(*SendMsgReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_DelMsgList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(sdk_ws.DelMsgListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).DelMsgList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_DelMsgList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).DelMsgList(ctx, req.(*sdk_ws.DelMsgListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_DelSuperGroupMsg_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DelSuperGroupMsgReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).DelSuperGroupMsg(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_DelSuperGroupMsg_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).DelSuperGroupMsg(ctx, req.(*DelSuperGroupMsgReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_ClearMsg_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClearMsgReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).ClearMsg(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_ClearMsg_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).ClearMsg(ctx, req.(*ClearMsgReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_SetMsgMinSeq_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetMsgMinSeqReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).SetMsgMinSeq(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_SetMsgMinSeq_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).SetMsgMinSeq(ctx, req.(*SetMsgMinSeqReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_SetSendMsgStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetSendMsgStatusReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).SetSendMsgStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_SetSendMsgStatus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).SetSendMsgStatus(ctx, req.(*SetSendMsgStatusReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_GetSendMsgStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSendMsgStatusReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).GetSendMsgStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_GetSendMsgStatus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).GetSendMsgStatus(ctx, req.(*GetSendMsgStatusReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_GetSuperGroupMsg_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSuperGroupMsgReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).GetSuperGroupMsg(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_GetSuperGroupMsg_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).GetSuperGroupMsg(ctx, req.(*GetSuperGroupMsgReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_GetWriteDiffMsg_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetWriteDiffMsgReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).GetWriteDiffMsg(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_GetWriteDiffMsg_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).GetWriteDiffMsg(ctx, req.(*GetWriteDiffMsgReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_SetMessageReactionExtensions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetMessageReactionExtensionsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).SetMessageReactionExtensions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_SetMessageReactionExtensions_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).SetMessageReactionExtensions(ctx, req.(*SetMessageReactionExtensionsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_GetMessageListReactionExtensions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMessageListReactionExtensionsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).GetMessageListReactionExtensions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_GetMessageListReactionExtensions_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).GetMessageListReactionExtensions(ctx, req.(*GetMessageListReactionExtensionsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_AddMessageReactionExtensions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddMessageReactionExtensionsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).AddMessageReactionExtensions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_AddMessageReactionExtensions_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).AddMessageReactionExtensions(ctx, req.(*AddMessageReactionExtensionsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_DeleteMessageReactionExtensions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteMessageListReactionExtensionsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).DeleteMessageReactionExtensions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_DeleteMessageReactionExtensions_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).DeleteMessageReactionExtensions(ctx, req.(*DeleteMessageListReactionExtensionsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_MsgCollect_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCollectReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).MsgCollect(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_MsgCollect_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).MsgCollect(ctx, req.(*MsgCollectReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_MsgCollectDel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCollectDelReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).MsgCollectDel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_MsgCollectDel_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).MsgCollectDel(ctx, req.(*MsgCollectDelReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_MsgCollectList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCollectListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).MsgCollectList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_MsgCollectList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).MsgCollectList(ctx, req.(*MsgCollectListReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Msg_ServiceDesc is the grpc.ServiceDesc for Msg service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Msg_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "msg.msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetMaxAndMinSeq",
			Handler:    _Msg_GetMaxAndMinSeq_Handler,
		},
		{
			MethodName: "PullMessageBySeqList",
			Handler:    _Msg_PullMessageBySeqList_Handler,
		},
		{
			MethodName: "SendMsg",
			Handler:    _Msg_SendMsg_Handler,
		},
		{
			MethodName: "DelMsgList",
			Handler:    _Msg_DelMsgList_Handler,
		},
		{
			MethodName: "DelSuperGroupMsg",
			Handler:    _Msg_DelSuperGroupMsg_Handler,
		},
		{
			MethodName: "ClearMsg",
			Handler:    _Msg_ClearMsg_Handler,
		},
		{
			MethodName: "SetMsgMinSeq",
			Handler:    _Msg_SetMsgMinSeq_Handler,
		},
		{
			MethodName: "SetSendMsgStatus",
			Handler:    _Msg_SetSendMsgStatus_Handler,
		},
		{
			MethodName: "GetSendMsgStatus",
			Handler:    _Msg_GetSendMsgStatus_Handler,
		},
		{
			MethodName: "GetSuperGroupMsg",
			Handler:    _Msg_GetSuperGroupMsg_Handler,
		},
		{
			MethodName: "GetWriteDiffMsg",
			Handler:    _Msg_GetWriteDiffMsg_Handler,
		},
		{
			MethodName: "SetMessageReactionExtensions",
			Handler:    _Msg_SetMessageReactionExtensions_Handler,
		},
		{
			MethodName: "GetMessageListReactionExtensions",
			Handler:    _Msg_GetMessageListReactionExtensions_Handler,
		},
		{
			MethodName: "AddMessageReactionExtensions",
			Handler:    _Msg_AddMessageReactionExtensions_Handler,
		},
		{
			MethodName: "DeleteMessageReactionExtensions",
			Handler:    _Msg_DeleteMessageReactionExtensions_Handler,
		},
		{
			MethodName: "MsgCollect",
			Handler:    _Msg_MsgCollect_Handler,
		},
		{
			MethodName: "MsgCollectDel",
			Handler:    _Msg_MsgCollectDel_Handler,
		},
		{
			MethodName: "MsgCollectList",
			Handler:    _Msg_MsgCollectList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "msg/msg.proto",
}
