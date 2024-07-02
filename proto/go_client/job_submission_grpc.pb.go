// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package go_client

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// RayJobSubmissionServiceClient is the client API for RayJobSubmissionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RayJobSubmissionServiceClient interface {
	// Submit a new Ray job on the specified cluster.
	SubmitRayJob(ctx context.Context, in *SubmitRayJobRequest, opts ...grpc.CallOption) (*SubmitRayJobReply, error)
	// Finds a specific job by its submission_id for the cluster with name and namespace.
	GetJobDetails(ctx context.Context, in *GetJobDetailsRequest, opts ...grpc.CallOption) (*JobSubmissionInfo, error)
	// Gets a specific job log by its submissionid for the cluster with name and namespace.
	GetJobLog(ctx context.Context, in *GetJobLogRequest, opts ...grpc.CallOption) (*GetJobLogReply, error)
	// List all job in a given a given cluster in a namespace. Supports pagination, and sorting on certain fields.
	ListJobDetails(ctx context.Context, in *ListJobDetailsRequest, opts ...grpc.CallOption) (*ListJobSubmissionInfo, error)
	// Stops a job by its name and namespace.
	StopRayJob(ctx context.Context, in *StopRayJobSubmissionRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// Deletes a job by its name and namespace.
	DeleteRayJob(ctx context.Context, in *DeleteRayJobSubmissionRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type rayJobSubmissionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRayJobSubmissionServiceClient(cc grpc.ClientConnInterface) RayJobSubmissionServiceClient {
	return &rayJobSubmissionServiceClient{cc}
}

func (c *rayJobSubmissionServiceClient) SubmitRayJob(ctx context.Context, in *SubmitRayJobRequest, opts ...grpc.CallOption) (*SubmitRayJobReply, error) {
	out := new(SubmitRayJobReply)
	err := c.cc.Invoke(ctx, "/proto.RayJobSubmissionService/SubmitRayJob", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rayJobSubmissionServiceClient) GetJobDetails(ctx context.Context, in *GetJobDetailsRequest, opts ...grpc.CallOption) (*JobSubmissionInfo, error) {
	out := new(JobSubmissionInfo)
	err := c.cc.Invoke(ctx, "/proto.RayJobSubmissionService/GetJobDetails", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rayJobSubmissionServiceClient) GetJobLog(ctx context.Context, in *GetJobLogRequest, opts ...grpc.CallOption) (*GetJobLogReply, error) {
	out := new(GetJobLogReply)
	err := c.cc.Invoke(ctx, "/proto.RayJobSubmissionService/GetJobLog", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rayJobSubmissionServiceClient) ListJobDetails(ctx context.Context, in *ListJobDetailsRequest, opts ...grpc.CallOption) (*ListJobSubmissionInfo, error) {
	out := new(ListJobSubmissionInfo)
	err := c.cc.Invoke(ctx, "/proto.RayJobSubmissionService/ListJobDetails", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rayJobSubmissionServiceClient) StopRayJob(ctx context.Context, in *StopRayJobSubmissionRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/proto.RayJobSubmissionService/StopRayJob", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rayJobSubmissionServiceClient) DeleteRayJob(ctx context.Context, in *DeleteRayJobSubmissionRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/proto.RayJobSubmissionService/DeleteRayJob", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RayJobSubmissionServiceServer is the server API for RayJobSubmissionService service.
// All implementations must embed UnimplementedRayJobSubmissionServiceServer
// for forward compatibility
type RayJobSubmissionServiceServer interface {
	// Submit a new Ray job on the specified cluster.
	SubmitRayJob(context.Context, *SubmitRayJobRequest) (*SubmitRayJobReply, error)
	// Finds a specific job by its submission_id for the cluster with name and namespace.
	GetJobDetails(context.Context, *GetJobDetailsRequest) (*JobSubmissionInfo, error)
	// Gets a specific job log by its submissionid for the cluster with name and namespace.
	GetJobLog(context.Context, *GetJobLogRequest) (*GetJobLogReply, error)
	// List all job in a given a given cluster in a namespace. Supports pagination, and sorting on certain fields.
	ListJobDetails(context.Context, *ListJobDetailsRequest) (*ListJobSubmissionInfo, error)
	// Stops a job by its name and namespace.
	StopRayJob(context.Context, *StopRayJobSubmissionRequest) (*emptypb.Empty, error)
	// Deletes a job by its name and namespace.
	DeleteRayJob(context.Context, *DeleteRayJobSubmissionRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedRayJobSubmissionServiceServer()
}

// UnimplementedRayJobSubmissionServiceServer must be embedded to have forward compatible implementations.
type UnimplementedRayJobSubmissionServiceServer struct {
}

func (UnimplementedRayJobSubmissionServiceServer) SubmitRayJob(context.Context, *SubmitRayJobRequest) (*SubmitRayJobReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubmitRayJob not implemented")
}
func (UnimplementedRayJobSubmissionServiceServer) GetJobDetails(context.Context, *GetJobDetailsRequest) (*JobSubmissionInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetJobDetails not implemented")
}
func (UnimplementedRayJobSubmissionServiceServer) GetJobLog(context.Context, *GetJobLogRequest) (*GetJobLogReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetJobLog not implemented")
}
func (UnimplementedRayJobSubmissionServiceServer) ListJobDetails(context.Context, *ListJobDetailsRequest) (*ListJobSubmissionInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListJobDetails not implemented")
}
func (UnimplementedRayJobSubmissionServiceServer) StopRayJob(context.Context, *StopRayJobSubmissionRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StopRayJob not implemented")
}
func (UnimplementedRayJobSubmissionServiceServer) DeleteRayJob(context.Context, *DeleteRayJobSubmissionRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteRayJob not implemented")
}
func (UnimplementedRayJobSubmissionServiceServer) mustEmbedUnimplementedRayJobSubmissionServiceServer() {
}

// UnsafeRayJobSubmissionServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RayJobSubmissionServiceServer will
// result in compilation errors.
type UnsafeRayJobSubmissionServiceServer interface {
	mustEmbedUnimplementedRayJobSubmissionServiceServer()
}

func RegisterRayJobSubmissionServiceServer(s grpc.ServiceRegistrar, srv RayJobSubmissionServiceServer) {
	s.RegisterService(&RayJobSubmissionService_ServiceDesc, srv)
}

func _RayJobSubmissionService_SubmitRayJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubmitRayJobRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RayJobSubmissionServiceServer).SubmitRayJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.RayJobSubmissionService/SubmitRayJob",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RayJobSubmissionServiceServer).SubmitRayJob(ctx, req.(*SubmitRayJobRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RayJobSubmissionService_GetJobDetails_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetJobDetailsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RayJobSubmissionServiceServer).GetJobDetails(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.RayJobSubmissionService/GetJobDetails",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RayJobSubmissionServiceServer).GetJobDetails(ctx, req.(*GetJobDetailsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RayJobSubmissionService_GetJobLog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetJobLogRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RayJobSubmissionServiceServer).GetJobLog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.RayJobSubmissionService/GetJobLog",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RayJobSubmissionServiceServer).GetJobLog(ctx, req.(*GetJobLogRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RayJobSubmissionService_ListJobDetails_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListJobDetailsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RayJobSubmissionServiceServer).ListJobDetails(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.RayJobSubmissionService/ListJobDetails",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RayJobSubmissionServiceServer).ListJobDetails(ctx, req.(*ListJobDetailsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RayJobSubmissionService_StopRayJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StopRayJobSubmissionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RayJobSubmissionServiceServer).StopRayJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.RayJobSubmissionService/StopRayJob",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RayJobSubmissionServiceServer).StopRayJob(ctx, req.(*StopRayJobSubmissionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RayJobSubmissionService_DeleteRayJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRayJobSubmissionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RayJobSubmissionServiceServer).DeleteRayJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.RayJobSubmissionService/DeleteRayJob",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RayJobSubmissionServiceServer).DeleteRayJob(ctx, req.(*DeleteRayJobSubmissionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RayJobSubmissionService_ServiceDesc is the grpc.ServiceDesc for RayJobSubmissionService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RayJobSubmissionService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.RayJobSubmissionService",
	HandlerType: (*RayJobSubmissionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SubmitRayJob",
			Handler:    _RayJobSubmissionService_SubmitRayJob_Handler,
		},
		{
			MethodName: "GetJobDetails",
			Handler:    _RayJobSubmissionService_GetJobDetails_Handler,
		},
		{
			MethodName: "GetJobLog",
			Handler:    _RayJobSubmissionService_GetJobLog_Handler,
		},
		{
			MethodName: "ListJobDetails",
			Handler:    _RayJobSubmissionService_ListJobDetails_Handler,
		},
		{
			MethodName: "StopRayJob",
			Handler:    _RayJobSubmissionService_StopRayJob_Handler,
		},
		{
			MethodName: "DeleteRayJob",
			Handler:    _RayJobSubmissionService_DeleteRayJob_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "job_submission.proto",
}
