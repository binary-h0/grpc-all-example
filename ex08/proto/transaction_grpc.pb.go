// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.3
// source: transaction.proto

package __

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	CoordinatorService_StartTransaction_FullMethodName = "/transaction.CoordinatorService/StartTransaction"
)

// CoordinatorServiceClient is the client API for CoordinatorService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CoordinatorServiceClient interface {
	StartTransaction(ctx context.Context, in *TransactionRequest, opts ...grpc.CallOption) (*TransactionResponse, error)
}

type coordinatorServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCoordinatorServiceClient(cc grpc.ClientConnInterface) CoordinatorServiceClient {
	return &coordinatorServiceClient{cc}
}

func (c *coordinatorServiceClient) StartTransaction(ctx context.Context, in *TransactionRequest, opts ...grpc.CallOption) (*TransactionResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(TransactionResponse)
	err := c.cc.Invoke(ctx, CoordinatorService_StartTransaction_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CoordinatorServiceServer is the server API for CoordinatorService service.
// All implementations must embed UnimplementedCoordinatorServiceServer
// for forward compatibility.
type CoordinatorServiceServer interface {
	StartTransaction(context.Context, *TransactionRequest) (*TransactionResponse, error)
	mustEmbedUnimplementedCoordinatorServiceServer()
}

// UnimplementedCoordinatorServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedCoordinatorServiceServer struct{}

func (UnimplementedCoordinatorServiceServer) StartTransaction(context.Context, *TransactionRequest) (*TransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StartTransaction not implemented")
}
func (UnimplementedCoordinatorServiceServer) mustEmbedUnimplementedCoordinatorServiceServer() {}
func (UnimplementedCoordinatorServiceServer) testEmbeddedByValue()                            {}

// UnsafeCoordinatorServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CoordinatorServiceServer will
// result in compilation errors.
type UnsafeCoordinatorServiceServer interface {
	mustEmbedUnimplementedCoordinatorServiceServer()
}

func RegisterCoordinatorServiceServer(s grpc.ServiceRegistrar, srv CoordinatorServiceServer) {
	// If the following call pancis, it indicates UnimplementedCoordinatorServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&CoordinatorService_ServiceDesc, srv)
}

func _CoordinatorService_StartTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TransactionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoordinatorServiceServer).StartTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CoordinatorService_StartTransaction_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoordinatorServiceServer).StartTransaction(ctx, req.(*TransactionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CoordinatorService_ServiceDesc is the grpc.ServiceDesc for CoordinatorService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CoordinatorService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "transaction.CoordinatorService",
	HandlerType: (*CoordinatorServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "StartTransaction",
			Handler:    _CoordinatorService_StartTransaction_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "transaction.proto",
}

const (
	ParticipantService_Prepare_FullMethodName  = "/transaction.ParticipantService/Prepare"
	ParticipantService_Commit_FullMethodName   = "/transaction.ParticipantService/Commit"
	ParticipantService_Rollback_FullMethodName = "/transaction.ParticipantService/Rollback"
)

// ParticipantServiceClient is the client API for ParticipantService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ParticipantServiceClient interface {
	Prepare(ctx context.Context, in *TransactionRequest, opts ...grpc.CallOption) (*StatusResponse, error)
	Commit(ctx context.Context, in *TransactionRequest, opts ...grpc.CallOption) (*StatusResponse, error)
	Rollback(ctx context.Context, in *TransactionRequest, opts ...grpc.CallOption) (*StatusResponse, error)
}

type participantServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewParticipantServiceClient(cc grpc.ClientConnInterface) ParticipantServiceClient {
	return &participantServiceClient{cc}
}

func (c *participantServiceClient) Prepare(ctx context.Context, in *TransactionRequest, opts ...grpc.CallOption) (*StatusResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(StatusResponse)
	err := c.cc.Invoke(ctx, ParticipantService_Prepare_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *participantServiceClient) Commit(ctx context.Context, in *TransactionRequest, opts ...grpc.CallOption) (*StatusResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(StatusResponse)
	err := c.cc.Invoke(ctx, ParticipantService_Commit_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *participantServiceClient) Rollback(ctx context.Context, in *TransactionRequest, opts ...grpc.CallOption) (*StatusResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(StatusResponse)
	err := c.cc.Invoke(ctx, ParticipantService_Rollback_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ParticipantServiceServer is the server API for ParticipantService service.
// All implementations must embed UnimplementedParticipantServiceServer
// for forward compatibility.
type ParticipantServiceServer interface {
	Prepare(context.Context, *TransactionRequest) (*StatusResponse, error)
	Commit(context.Context, *TransactionRequest) (*StatusResponse, error)
	Rollback(context.Context, *TransactionRequest) (*StatusResponse, error)
	mustEmbedUnimplementedParticipantServiceServer()
}

// UnimplementedParticipantServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedParticipantServiceServer struct{}

func (UnimplementedParticipantServiceServer) Prepare(context.Context, *TransactionRequest) (*StatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Prepare not implemented")
}
func (UnimplementedParticipantServiceServer) Commit(context.Context, *TransactionRequest) (*StatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Commit not implemented")
}
func (UnimplementedParticipantServiceServer) Rollback(context.Context, *TransactionRequest) (*StatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Rollback not implemented")
}
func (UnimplementedParticipantServiceServer) mustEmbedUnimplementedParticipantServiceServer() {}
func (UnimplementedParticipantServiceServer) testEmbeddedByValue()                            {}

// UnsafeParticipantServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ParticipantServiceServer will
// result in compilation errors.
type UnsafeParticipantServiceServer interface {
	mustEmbedUnimplementedParticipantServiceServer()
}

func RegisterParticipantServiceServer(s grpc.ServiceRegistrar, srv ParticipantServiceServer) {
	// If the following call pancis, it indicates UnimplementedParticipantServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&ParticipantService_ServiceDesc, srv)
}

func _ParticipantService_Prepare_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TransactionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ParticipantServiceServer).Prepare(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ParticipantService_Prepare_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ParticipantServiceServer).Prepare(ctx, req.(*TransactionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ParticipantService_Commit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TransactionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ParticipantServiceServer).Commit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ParticipantService_Commit_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ParticipantServiceServer).Commit(ctx, req.(*TransactionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ParticipantService_Rollback_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TransactionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ParticipantServiceServer).Rollback(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ParticipantService_Rollback_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ParticipantServiceServer).Rollback(ctx, req.(*TransactionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ParticipantService_ServiceDesc is the grpc.ServiceDesc for ParticipantService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ParticipantService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "transaction.ParticipantService",
	HandlerType: (*ParticipantServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Prepare",
			Handler:    _ParticipantService_Prepare_Handler,
		},
		{
			MethodName: "Commit",
			Handler:    _ParticipantService_Commit_Handler,
		},
		{
			MethodName: "Rollback",
			Handler:    _ParticipantService_Rollback_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "transaction.proto",
}
