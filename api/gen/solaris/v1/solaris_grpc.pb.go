// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.3
// source: solaris.proto

package solaris

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
	Service_CreateLog_FullMethodName     = "/solaris.v1.Service/CreateLog"
	Service_UpdateLog_FullMethodName     = "/solaris.v1.Service/UpdateLog"
	Service_QueryLogs_FullMethodName     = "/solaris.v1.Service/QueryLogs"
	Service_DeleteLogs_FullMethodName    = "/solaris.v1.Service/DeleteLogs"
	Service_AppendRecords_FullMethodName = "/solaris.v1.Service/AppendRecords"
	Service_QueryRecords_FullMethodName  = "/solaris.v1.Service/QueryRecords"
)

// ServiceClient is the client API for Service service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ServiceClient interface {
	// CreateLog creates then new log
	CreateLog(ctx context.Context, in *Log, opts ...grpc.CallOption) (*Log, error)
	// UpdateLog changes the log settings (tags)
	UpdateLog(ctx context.Context, in *Log, opts ...grpc.CallOption) (*Log, error)
	// QueryLogs requests list of logs by the query request ordered by the log IDs ascending order
	QueryLogs(ctx context.Context, in *QueryLogsRequest, opts ...grpc.CallOption) (*QueryLogsResult, error)
	// DeleteLogs removes one or more logs
	DeleteLogs(ctx context.Context, in *DeleteLogsRequest, opts ...grpc.CallOption) (*DeleteLogsResult, error)
	// AppendRecords appends a bunch of records to the log
	AppendRecords(ctx context.Context, in *AppendRecordsRequest, opts ...grpc.CallOption) (*AppendRecordsResult, error)
	// QueryRecords read records from one or many logs, merging them together into the result set
	// sorted in ascending or descending order by the records IDs (timestamps)
	QueryRecords(ctx context.Context, in *QueryRecordsRequest, opts ...grpc.CallOption) (*QueryRecordsResult, error)
}

type serviceClient struct {
	cc grpc.ClientConnInterface
}

func NewServiceClient(cc grpc.ClientConnInterface) ServiceClient {
	return &serviceClient{cc}
}

func (c *serviceClient) CreateLog(ctx context.Context, in *Log, opts ...grpc.CallOption) (*Log, error) {
	out := new(Log)
	err := c.cc.Invoke(ctx, Service_CreateLog_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) UpdateLog(ctx context.Context, in *Log, opts ...grpc.CallOption) (*Log, error) {
	out := new(Log)
	err := c.cc.Invoke(ctx, Service_UpdateLog_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) QueryLogs(ctx context.Context, in *QueryLogsRequest, opts ...grpc.CallOption) (*QueryLogsResult, error) {
	out := new(QueryLogsResult)
	err := c.cc.Invoke(ctx, Service_QueryLogs_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) DeleteLogs(ctx context.Context, in *DeleteLogsRequest, opts ...grpc.CallOption) (*DeleteLogsResult, error) {
	out := new(DeleteLogsResult)
	err := c.cc.Invoke(ctx, Service_DeleteLogs_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) AppendRecords(ctx context.Context, in *AppendRecordsRequest, opts ...grpc.CallOption) (*AppendRecordsResult, error) {
	out := new(AppendRecordsResult)
	err := c.cc.Invoke(ctx, Service_AppendRecords_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) QueryRecords(ctx context.Context, in *QueryRecordsRequest, opts ...grpc.CallOption) (*QueryRecordsResult, error) {
	out := new(QueryRecordsResult)
	err := c.cc.Invoke(ctx, Service_QueryRecords_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServiceServer is the server API for Service service.
// All implementations must embed UnimplementedServiceServer
// for forward compatibility
type ServiceServer interface {
	// CreateLog creates then new log
	CreateLog(context.Context, *Log) (*Log, error)
	// UpdateLog changes the log settings (tags)
	UpdateLog(context.Context, *Log) (*Log, error)
	// QueryLogs requests list of logs by the query request ordered by the log IDs ascending order
	QueryLogs(context.Context, *QueryLogsRequest) (*QueryLogsResult, error)
	// DeleteLogs removes one or more logs
	DeleteLogs(context.Context, *DeleteLogsRequest) (*DeleteLogsResult, error)
	// AppendRecords appends a bunch of records to the log
	AppendRecords(context.Context, *AppendRecordsRequest) (*AppendRecordsResult, error)
	// QueryRecords read records from one or many logs, merging them together into the result set
	// sorted in ascending or descending order by the records IDs (timestamps)
	QueryRecords(context.Context, *QueryRecordsRequest) (*QueryRecordsResult, error)
	mustEmbedUnimplementedServiceServer()
}

// UnimplementedServiceServer must be embedded to have forward compatible implementations.
type UnimplementedServiceServer struct {
}

func (UnimplementedServiceServer) CreateLog(context.Context, *Log) (*Log, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateLog not implemented")
}
func (UnimplementedServiceServer) UpdateLog(context.Context, *Log) (*Log, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateLog not implemented")
}
func (UnimplementedServiceServer) QueryLogs(context.Context, *QueryLogsRequest) (*QueryLogsResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryLogs not implemented")
}
func (UnimplementedServiceServer) DeleteLogs(context.Context, *DeleteLogsRequest) (*DeleteLogsResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteLogs not implemented")
}
func (UnimplementedServiceServer) AppendRecords(context.Context, *AppendRecordsRequest) (*AppendRecordsResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AppendRecords not implemented")
}
func (UnimplementedServiceServer) QueryRecords(context.Context, *QueryRecordsRequest) (*QueryRecordsResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryRecords not implemented")
}
func (UnimplementedServiceServer) mustEmbedUnimplementedServiceServer() {}

// UnsafeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ServiceServer will
// result in compilation errors.
type UnsafeServiceServer interface {
	mustEmbedUnimplementedServiceServer()
}

func RegisterServiceServer(s grpc.ServiceRegistrar, srv ServiceServer) {
	s.RegisterService(&Service_ServiceDesc, srv)
}

func _Service_CreateLog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Log)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).CreateLog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Service_CreateLog_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).CreateLog(ctx, req.(*Log))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_UpdateLog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Log)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).UpdateLog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Service_UpdateLog_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).UpdateLog(ctx, req.(*Log))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_QueryLogs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryLogsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).QueryLogs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Service_QueryLogs_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).QueryLogs(ctx, req.(*QueryLogsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_DeleteLogs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteLogsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).DeleteLogs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Service_DeleteLogs_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).DeleteLogs(ctx, req.(*DeleteLogsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_AppendRecords_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AppendRecordsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).AppendRecords(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Service_AppendRecords_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).AppendRecords(ctx, req.(*AppendRecordsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_QueryRecords_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryRecordsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).QueryRecords(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Service_QueryRecords_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).QueryRecords(ctx, req.(*QueryRecordsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Service_ServiceDesc is the grpc.ServiceDesc for Service service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Service_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "solaris.v1.Service",
	HandlerType: (*ServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateLog",
			Handler:    _Service_CreateLog_Handler,
		},
		{
			MethodName: "UpdateLog",
			Handler:    _Service_UpdateLog_Handler,
		},
		{
			MethodName: "QueryLogs",
			Handler:    _Service_QueryLogs_Handler,
		},
		{
			MethodName: "DeleteLogs",
			Handler:    _Service_DeleteLogs_Handler,
		},
		{
			MethodName: "AppendRecords",
			Handler:    _Service_AppendRecords_Handler,
		},
		{
			MethodName: "QueryRecords",
			Handler:    _Service_QueryRecords_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "solaris.proto",
}
