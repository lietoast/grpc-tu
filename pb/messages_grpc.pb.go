// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.9
// source: messages.proto

package pb

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
	EmployeeService_GetByNo_FullMethodName  = "/pb.EmployeeService/GetByNo"
	EmployeeService_GetAll_FullMethodName   = "/pb.EmployeeService/GetAll"
	EmployeeService_AddPhoto_FullMethodName = "/pb.EmployeeService/AddPhoto"
	EmployeeService_Save_FullMethodName     = "/pb.EmployeeService/Save"
	EmployeeService_SaveAll_FullMethodName  = "/pb.EmployeeService/SaveAll"
)

// EmployeeServiceClient is the client API for EmployeeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EmployeeServiceClient interface {
	GetByNo(ctx context.Context, in *GetByNoRequest, opts ...grpc.CallOption) (*EmployeeResponse, error)
	GetAll(ctx context.Context, in *GetAllRequest, opts ...grpc.CallOption) (EmployeeService_GetAllClient, error)
	AddPhoto(ctx context.Context, opts ...grpc.CallOption) (EmployeeService_AddPhotoClient, error)
	Save(ctx context.Context, in *EmployeeRequest, opts ...grpc.CallOption) (*EmployeeResponse, error)
	SaveAll(ctx context.Context, opts ...grpc.CallOption) (EmployeeService_SaveAllClient, error)
}

type employeeServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewEmployeeServiceClient(cc grpc.ClientConnInterface) EmployeeServiceClient {
	return &employeeServiceClient{cc}
}

func (c *employeeServiceClient) GetByNo(ctx context.Context, in *GetByNoRequest, opts ...grpc.CallOption) (*EmployeeResponse, error) {
	out := new(EmployeeResponse)
	err := c.cc.Invoke(ctx, EmployeeService_GetByNo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *employeeServiceClient) GetAll(ctx context.Context, in *GetAllRequest, opts ...grpc.CallOption) (EmployeeService_GetAllClient, error) {
	stream, err := c.cc.NewStream(ctx, &EmployeeService_ServiceDesc.Streams[0], EmployeeService_GetAll_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &employeeServiceGetAllClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type EmployeeService_GetAllClient interface {
	Recv() (*EmployeeResponse, error)
	grpc.ClientStream
}

type employeeServiceGetAllClient struct {
	grpc.ClientStream
}

func (x *employeeServiceGetAllClient) Recv() (*EmployeeResponse, error) {
	m := new(EmployeeResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *employeeServiceClient) AddPhoto(ctx context.Context, opts ...grpc.CallOption) (EmployeeService_AddPhotoClient, error) {
	stream, err := c.cc.NewStream(ctx, &EmployeeService_ServiceDesc.Streams[1], EmployeeService_AddPhoto_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &employeeServiceAddPhotoClient{stream}
	return x, nil
}

type EmployeeService_AddPhotoClient interface {
	Send(*AddPhotoRequest) error
	CloseAndRecv() (*AddPhotoResponse, error)
	grpc.ClientStream
}

type employeeServiceAddPhotoClient struct {
	grpc.ClientStream
}

func (x *employeeServiceAddPhotoClient) Send(m *AddPhotoRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *employeeServiceAddPhotoClient) CloseAndRecv() (*AddPhotoResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(AddPhotoResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *employeeServiceClient) Save(ctx context.Context, in *EmployeeRequest, opts ...grpc.CallOption) (*EmployeeResponse, error) {
	out := new(EmployeeResponse)
	err := c.cc.Invoke(ctx, EmployeeService_Save_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *employeeServiceClient) SaveAll(ctx context.Context, opts ...grpc.CallOption) (EmployeeService_SaveAllClient, error) {
	stream, err := c.cc.NewStream(ctx, &EmployeeService_ServiceDesc.Streams[2], EmployeeService_SaveAll_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &employeeServiceSaveAllClient{stream}
	return x, nil
}

type EmployeeService_SaveAllClient interface {
	Send(*EmployeeRequest) error
	Recv() (*EmployeeResponse, error)
	grpc.ClientStream
}

type employeeServiceSaveAllClient struct {
	grpc.ClientStream
}

func (x *employeeServiceSaveAllClient) Send(m *EmployeeRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *employeeServiceSaveAllClient) Recv() (*EmployeeResponse, error) {
	m := new(EmployeeResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// EmployeeServiceServer is the server API for EmployeeService service.
// All implementations must embed UnimplementedEmployeeServiceServer
// for forward compatibility
type EmployeeServiceServer interface {
	GetByNo(context.Context, *GetByNoRequest) (*EmployeeResponse, error)
	GetAll(*GetAllRequest, EmployeeService_GetAllServer) error
	AddPhoto(EmployeeService_AddPhotoServer) error
	Save(context.Context, *EmployeeRequest) (*EmployeeResponse, error)
	SaveAll(EmployeeService_SaveAllServer) error
	mustEmbedUnimplementedEmployeeServiceServer()
}

// UnimplementedEmployeeServiceServer must be embedded to have forward compatible implementations.
type UnimplementedEmployeeServiceServer struct {
}

func (UnimplementedEmployeeServiceServer) GetByNo(context.Context, *GetByNoRequest) (*EmployeeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetByNo not implemented")
}
func (UnimplementedEmployeeServiceServer) GetAll(*GetAllRequest, EmployeeService_GetAllServer) error {
	return status.Errorf(codes.Unimplemented, "method GetAll not implemented")
}
func (UnimplementedEmployeeServiceServer) AddPhoto(EmployeeService_AddPhotoServer) error {
	return status.Errorf(codes.Unimplemented, "method AddPhoto not implemented")
}
func (UnimplementedEmployeeServiceServer) Save(context.Context, *EmployeeRequest) (*EmployeeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Save not implemented")
}
func (UnimplementedEmployeeServiceServer) SaveAll(EmployeeService_SaveAllServer) error {
	return status.Errorf(codes.Unimplemented, "method SaveAll not implemented")
}
func (UnimplementedEmployeeServiceServer) mustEmbedUnimplementedEmployeeServiceServer() {}

// UnsafeEmployeeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EmployeeServiceServer will
// result in compilation errors.
type UnsafeEmployeeServiceServer interface {
	mustEmbedUnimplementedEmployeeServiceServer()
}

func RegisterEmployeeServiceServer(s grpc.ServiceRegistrar, srv EmployeeServiceServer) {
	s.RegisterService(&EmployeeService_ServiceDesc, srv)
}

func _EmployeeService_GetByNo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetByNoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmployeeServiceServer).GetByNo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EmployeeService_GetByNo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmployeeServiceServer).GetByNo(ctx, req.(*GetByNoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EmployeeService_GetAll_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetAllRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(EmployeeServiceServer).GetAll(m, &employeeServiceGetAllServer{stream})
}

type EmployeeService_GetAllServer interface {
	Send(*EmployeeResponse) error
	grpc.ServerStream
}

type employeeServiceGetAllServer struct {
	grpc.ServerStream
}

func (x *employeeServiceGetAllServer) Send(m *EmployeeResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _EmployeeService_AddPhoto_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(EmployeeServiceServer).AddPhoto(&employeeServiceAddPhotoServer{stream})
}

type EmployeeService_AddPhotoServer interface {
	SendAndClose(*AddPhotoResponse) error
	Recv() (*AddPhotoRequest, error)
	grpc.ServerStream
}

type employeeServiceAddPhotoServer struct {
	grpc.ServerStream
}

func (x *employeeServiceAddPhotoServer) SendAndClose(m *AddPhotoResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *employeeServiceAddPhotoServer) Recv() (*AddPhotoRequest, error) {
	m := new(AddPhotoRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _EmployeeService_Save_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmployeeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmployeeServiceServer).Save(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EmployeeService_Save_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmployeeServiceServer).Save(ctx, req.(*EmployeeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EmployeeService_SaveAll_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(EmployeeServiceServer).SaveAll(&employeeServiceSaveAllServer{stream})
}

type EmployeeService_SaveAllServer interface {
	Send(*EmployeeResponse) error
	Recv() (*EmployeeRequest, error)
	grpc.ServerStream
}

type employeeServiceSaveAllServer struct {
	grpc.ServerStream
}

func (x *employeeServiceSaveAllServer) Send(m *EmployeeResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *employeeServiceSaveAllServer) Recv() (*EmployeeRequest, error) {
	m := new(EmployeeRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// EmployeeService_ServiceDesc is the grpc.ServiceDesc for EmployeeService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var EmployeeService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.EmployeeService",
	HandlerType: (*EmployeeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetByNo",
			Handler:    _EmployeeService_GetByNo_Handler,
		},
		{
			MethodName: "Save",
			Handler:    _EmployeeService_Save_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetAll",
			Handler:       _EmployeeService_GetAll_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "AddPhoto",
			Handler:       _EmployeeService_AddPhoto_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "SaveAll",
			Handler:       _EmployeeService_SaveAll_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "messages.proto",
}
