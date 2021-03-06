// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package api

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

// ApiClient is the client API for Api service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ApiClient interface {
	Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error)
	SendMessage(ctx context.Context, in *SendMessageRequest, opts ...grpc.CallOption) (*SendMessageResponse, error)
	GetNodeID(ctx context.Context, in *GetNodeIDRequest, opts ...grpc.CallOption) (*GetNodeIDResponse, error)
	SetNickname(ctx context.Context, in *SetNicknameRequest, opts ...grpc.CallOption) (*SetNicknameResponse, error)
	GetNickname(ctx context.Context, in *GetNicknameRequest, opts ...grpc.CallOption) (*GetNicknameResponse, error)
	JoinRoom(ctx context.Context, in *JoinRoomRequest, opts ...grpc.CallOption) (*JoinRoomResponse, error)
	GetRoomParticipants(ctx context.Context, in *GetRoomParticipantsRequest, opts ...grpc.CallOption) (*GetRoomParticipantsResponse, error)
	SubscribeToEvents(ctx context.Context, in *SubscribeToEventsRequest, opts ...grpc.CallOption) (Api_SubscribeToEventsClient, error)
}

type apiClient struct {
	cc grpc.ClientConnInterface
}

func NewApiClient(cc grpc.ClientConnInterface) ApiClient {
	return &apiClient{cc}
}

func (c *apiClient) Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error) {
	out := new(PingResponse)
	err := c.cc.Invoke(ctx, "/api.Api/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiClient) SendMessage(ctx context.Context, in *SendMessageRequest, opts ...grpc.CallOption) (*SendMessageResponse, error) {
	out := new(SendMessageResponse)
	err := c.cc.Invoke(ctx, "/api.Api/SendMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiClient) GetNodeID(ctx context.Context, in *GetNodeIDRequest, opts ...grpc.CallOption) (*GetNodeIDResponse, error) {
	out := new(GetNodeIDResponse)
	err := c.cc.Invoke(ctx, "/api.Api/GetNodeID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiClient) SetNickname(ctx context.Context, in *SetNicknameRequest, opts ...grpc.CallOption) (*SetNicknameResponse, error) {
	out := new(SetNicknameResponse)
	err := c.cc.Invoke(ctx, "/api.Api/SetNickname", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiClient) GetNickname(ctx context.Context, in *GetNicknameRequest, opts ...grpc.CallOption) (*GetNicknameResponse, error) {
	out := new(GetNicknameResponse)
	err := c.cc.Invoke(ctx, "/api.Api/GetNickname", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiClient) JoinRoom(ctx context.Context, in *JoinRoomRequest, opts ...grpc.CallOption) (*JoinRoomResponse, error) {
	out := new(JoinRoomResponse)
	err := c.cc.Invoke(ctx, "/api.Api/JoinRoom", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiClient) GetRoomParticipants(ctx context.Context, in *GetRoomParticipantsRequest, opts ...grpc.CallOption) (*GetRoomParticipantsResponse, error) {
	out := new(GetRoomParticipantsResponse)
	err := c.cc.Invoke(ctx, "/api.Api/GetRoomParticipants", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiClient) SubscribeToEvents(ctx context.Context, in *SubscribeToEventsRequest, opts ...grpc.CallOption) (Api_SubscribeToEventsClient, error) {
	stream, err := c.cc.NewStream(ctx, &Api_ServiceDesc.Streams[0], "/api.Api/SubscribeToEvents", opts...)
	if err != nil {
		return nil, err
	}
	x := &apiSubscribeToEventsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Api_SubscribeToEventsClient interface {
	Recv() (*Event, error)
	grpc.ClientStream
}

type apiSubscribeToEventsClient struct {
	grpc.ClientStream
}

func (x *apiSubscribeToEventsClient) Recv() (*Event, error) {
	m := new(Event)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ApiServer is the server API for Api service.
// All implementations must embed UnimplementedApiServer
// for forward compatibility
type ApiServer interface {
	Ping(context.Context, *PingRequest) (*PingResponse, error)
	SendMessage(context.Context, *SendMessageRequest) (*SendMessageResponse, error)
	GetNodeID(context.Context, *GetNodeIDRequest) (*GetNodeIDResponse, error)
	SetNickname(context.Context, *SetNicknameRequest) (*SetNicknameResponse, error)
	GetNickname(context.Context, *GetNicknameRequest) (*GetNicknameResponse, error)
	JoinRoom(context.Context, *JoinRoomRequest) (*JoinRoomResponse, error)
	GetRoomParticipants(context.Context, *GetRoomParticipantsRequest) (*GetRoomParticipantsResponse, error)
	SubscribeToEvents(*SubscribeToEventsRequest, Api_SubscribeToEventsServer) error
	mustEmbedUnimplementedApiServer()
}

// UnimplementedApiServer must be embedded to have forward compatible implementations.
type UnimplementedApiServer struct {
}

func (UnimplementedApiServer) Ping(context.Context, *PingRequest) (*PingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (UnimplementedApiServer) SendMessage(context.Context, *SendMessageRequest) (*SendMessageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendMessage not implemented")
}
func (UnimplementedApiServer) GetNodeID(context.Context, *GetNodeIDRequest) (*GetNodeIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNodeID not implemented")
}
func (UnimplementedApiServer) SetNickname(context.Context, *SetNicknameRequest) (*SetNicknameResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetNickname not implemented")
}
func (UnimplementedApiServer) GetNickname(context.Context, *GetNicknameRequest) (*GetNicknameResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNickname not implemented")
}
func (UnimplementedApiServer) JoinRoom(context.Context, *JoinRoomRequest) (*JoinRoomResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method JoinRoom not implemented")
}
func (UnimplementedApiServer) GetRoomParticipants(context.Context, *GetRoomParticipantsRequest) (*GetRoomParticipantsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRoomParticipants not implemented")
}
func (UnimplementedApiServer) SubscribeToEvents(*SubscribeToEventsRequest, Api_SubscribeToEventsServer) error {
	return status.Errorf(codes.Unimplemented, "method SubscribeToEvents not implemented")
}
func (UnimplementedApiServer) mustEmbedUnimplementedApiServer() {}

// UnsafeApiServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ApiServer will
// result in compilation errors.
type UnsafeApiServer interface {
	mustEmbedUnimplementedApiServer()
}

func RegisterApiServer(s grpc.ServiceRegistrar, srv ApiServer) {
	s.RegisterService(&Api_ServiceDesc, srv)
}

func _Api_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Api/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServer).Ping(ctx, req.(*PingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Api_SendMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServer).SendMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Api/SendMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServer).SendMessage(ctx, req.(*SendMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Api_GetNodeID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetNodeIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServer).GetNodeID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Api/GetNodeID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServer).GetNodeID(ctx, req.(*GetNodeIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Api_SetNickname_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetNicknameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServer).SetNickname(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Api/SetNickname",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServer).SetNickname(ctx, req.(*SetNicknameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Api_GetNickname_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetNicknameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServer).GetNickname(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Api/GetNickname",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServer).GetNickname(ctx, req.(*GetNicknameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Api_JoinRoom_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JoinRoomRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServer).JoinRoom(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Api/JoinRoom",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServer).JoinRoom(ctx, req.(*JoinRoomRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Api_GetRoomParticipants_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRoomParticipantsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServer).GetRoomParticipants(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Api/GetRoomParticipants",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServer).GetRoomParticipants(ctx, req.(*GetRoomParticipantsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Api_SubscribeToEvents_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SubscribeToEventsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ApiServer).SubscribeToEvents(m, &apiSubscribeToEventsServer{stream})
}

type Api_SubscribeToEventsServer interface {
	Send(*Event) error
	grpc.ServerStream
}

type apiSubscribeToEventsServer struct {
	grpc.ServerStream
}

func (x *apiSubscribeToEventsServer) Send(m *Event) error {
	return x.ServerStream.SendMsg(m)
}

// Api_ServiceDesc is the grpc.ServiceDesc for Api service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Api_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.Api",
	HandlerType: (*ApiServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _Api_Ping_Handler,
		},
		{
			MethodName: "SendMessage",
			Handler:    _Api_SendMessage_Handler,
		},
		{
			MethodName: "GetNodeID",
			Handler:    _Api_GetNodeID_Handler,
		},
		{
			MethodName: "SetNickname",
			Handler:    _Api_SetNickname_Handler,
		},
		{
			MethodName: "GetNickname",
			Handler:    _Api_GetNickname_Handler,
		},
		{
			MethodName: "JoinRoom",
			Handler:    _Api_JoinRoom_Handler,
		},
		{
			MethodName: "GetRoomParticipants",
			Handler:    _Api_GetRoomParticipants_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SubscribeToEvents",
			Handler:       _Api_SubscribeToEvents_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "proto/api.proto",
}
