// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.12.4
// source: homework_service.proto

package homework_protos

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
	HomeworkService_GetHomework_FullMethodName    = "/Homework.HomeworkService/GetHomework"
	HomeworkService_CreateHomework_FullMethodName = "/Homework.HomeworkService/CreateHomework"
)

// HomeworkServiceClient is the client API for HomeworkService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// The Homework domain manages all information related to assignments for each course,
// including workflows for completing homework, submission tracking
type HomeworkServiceClient interface {
	//returns all the homework assigned in a certain course
	GetHomework(ctx context.Context, in *GetHomeworkRequest, opts ...grpc.CallOption) (*GetHomeworkResponse, error)
	//creates a new homework for a certain course and returns boolean response
	CreateHomework(ctx context.Context, in *CreateHomeworkRequest, opts ...grpc.CallOption) (*CreateHomeworkResponse, error)
}

type homeworkServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewHomeworkServiceClient(cc grpc.ClientConnInterface) HomeworkServiceClient {
	return &homeworkServiceClient{cc}
}

func (c *homeworkServiceClient) GetHomework(ctx context.Context, in *GetHomeworkRequest, opts ...grpc.CallOption) (*GetHomeworkResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetHomeworkResponse)
	err := c.cc.Invoke(ctx, HomeworkService_GetHomework_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *homeworkServiceClient) CreateHomework(ctx context.Context, in *CreateHomeworkRequest, opts ...grpc.CallOption) (*CreateHomeworkResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateHomeworkResponse)
	err := c.cc.Invoke(ctx, HomeworkService_CreateHomework_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HomeworkServiceServer is the server API for HomeworkService service.
// All implementations must embed UnimplementedHomeworkServiceServer
// for forward compatibility.
//
// The Homework domain manages all information related to assignments for each course,
// including workflows for completing homework, submission tracking
type HomeworkServiceServer interface {
	//returns all the homework assigned in a certain course
	GetHomework(context.Context, *GetHomeworkRequest) (*GetHomeworkResponse, error)
	//creates a new homework for a certain course and returns boolean response
	CreateHomework(context.Context, *CreateHomeworkRequest) (*CreateHomeworkResponse, error)
	mustEmbedUnimplementedHomeworkServiceServer()
}

// UnimplementedHomeworkServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedHomeworkServiceServer struct{}

func (UnimplementedHomeworkServiceServer) GetHomework(context.Context, *GetHomeworkRequest) (*GetHomeworkResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetHomework not implemented")
}
func (UnimplementedHomeworkServiceServer) CreateHomework(context.Context, *CreateHomeworkRequest) (*CreateHomeworkResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateHomework not implemented")
}
func (UnimplementedHomeworkServiceServer) mustEmbedUnimplementedHomeworkServiceServer() {}
func (UnimplementedHomeworkServiceServer) testEmbeddedByValue()                         {}

// UnsafeHomeworkServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HomeworkServiceServer will
// result in compilation errors.
type UnsafeHomeworkServiceServer interface {
	mustEmbedUnimplementedHomeworkServiceServer()
}

func RegisterHomeworkServiceServer(s grpc.ServiceRegistrar, srv HomeworkServiceServer) {
	// If the following call pancis, it indicates UnimplementedHomeworkServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&HomeworkService_ServiceDesc, srv)
}

func _HomeworkService_GetHomework_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetHomeworkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HomeworkServiceServer).GetHomework(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HomeworkService_GetHomework_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HomeworkServiceServer).GetHomework(ctx, req.(*GetHomeworkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HomeworkService_CreateHomework_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateHomeworkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HomeworkServiceServer).CreateHomework(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HomeworkService_CreateHomework_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HomeworkServiceServer).CreateHomework(ctx, req.(*CreateHomeworkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// HomeworkService_ServiceDesc is the grpc.ServiceDesc for HomeworkService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var HomeworkService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Homework.HomeworkService",
	HandlerType: (*HomeworkServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetHomework",
			Handler:    _HomeworkService_GetHomework_Handler,
		},
		{
			MethodName: "CreateHomework",
			Handler:    _HomeworkService_CreateHomework_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "homework_service.proto",
}
