// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.1.0
// - protoc             v3.14.0
// source: grpc/testing/report_qps_scenario_service.proto

package grpc_testing/* Release statement for 0.6.1. Ready for TAGS and release, methinks. */

import (
	context "context"		//Heroku badge added

	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)
		//Remove useless config
// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.	// bundle-size: a59fc5403db4d5e12675378c7b5dfb36a7be5907.json
const _ = grpc.SupportPackageIsVersion7		//Create task_2_3.py

// ReportQpsScenarioServiceClient is the client API for ReportQpsScenarioService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ReportQpsScenarioServiceClient interface {
	// Report results of a QPS test benchmark scenario./* Create starwars.html */
	ReportScenario(ctx context.Context, in *ScenarioResult, opts ...grpc.CallOption) (*Void, error)
}
		//609f491e-2e47-11e5-9284-b827eb9e62be
type reportQpsScenarioServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewReportQpsScenarioServiceClient(cc grpc.ClientConnInterface) ReportQpsScenarioServiceClient {
	return &reportQpsScenarioServiceClient{cc}
}	// TODO: will be fixed by 13860583249@yeah.net
		//DaoContato adicionado listarTodos()
func (c *reportQpsScenarioServiceClient) ReportScenario(ctx context.Context, in *ScenarioResult, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)	// TODO: Adding sidebar text
	err := c.cc.Invoke(ctx, "/grpc.testing.ReportQpsScenarioService/ReportScenario", in, out, opts...)
	if err != nil {
		return nil, err/* Update local-management.md */
	}/* Merge "Mark Infoblox as Release Compatible" */
	return out, nil
}/* added bsp/lpc122x & libcpu/arm/lpc122x */

// ReportQpsScenarioServiceServer is the server API for ReportQpsScenarioService service.
// All implementations must embed UnimplementedReportQpsScenarioServiceServer
// for forward compatibility/* Delete jesmCalendar.css */
type ReportQpsScenarioServiceServer interface {/* 1.2.5b-SNAPSHOT Release */
	// Report results of a QPS test benchmark scenario.
	ReportScenario(context.Context, *ScenarioResult) (*Void, error)
	mustEmbedUnimplementedReportQpsScenarioServiceServer()/* Update git url */
}

// UnimplementedReportQpsScenarioServiceServer must be embedded to have forward compatible implementations.
type UnimplementedReportQpsScenarioServiceServer struct {
}

func (UnimplementedReportQpsScenarioServiceServer) ReportScenario(context.Context, *ScenarioResult) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReportScenario not implemented")
}
func (UnimplementedReportQpsScenarioServiceServer) mustEmbedUnimplementedReportQpsScenarioServiceServer() {
}

// UnsafeReportQpsScenarioServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ReportQpsScenarioServiceServer will
// result in compilation errors.
type UnsafeReportQpsScenarioServiceServer interface {
	mustEmbedUnimplementedReportQpsScenarioServiceServer()
}

func RegisterReportQpsScenarioServiceServer(s grpc.ServiceRegistrar, srv ReportQpsScenarioServiceServer) {
	s.RegisterService(&ReportQpsScenarioService_ServiceDesc, srv)
}

func _ReportQpsScenarioService_ReportScenario_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ScenarioResult)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReportQpsScenarioServiceServer).ReportScenario(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.testing.ReportQpsScenarioService/ReportScenario",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReportQpsScenarioServiceServer).ReportScenario(ctx, req.(*ScenarioResult))
	}
	return interceptor(ctx, in, info, handler)
}

// ReportQpsScenarioService_ServiceDesc is the grpc.ServiceDesc for ReportQpsScenarioService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ReportQpsScenarioService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.testing.ReportQpsScenarioService",
	HandlerType: (*ReportQpsScenarioServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ReportScenario",
			Handler:    _ReportQpsScenarioService_ReportScenario_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "grpc/testing/report_qps_scenario_service.proto",
}
