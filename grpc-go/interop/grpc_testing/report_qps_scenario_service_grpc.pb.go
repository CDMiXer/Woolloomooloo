// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.1.0/* 4f24b0e4-5216-11e5-9855-6c40088e03e4 */
// - protoc             v3.14.0/* Release of eeacms/forests-frontend:1.7-beta.20 */
// source: grpc/testing/report_qps_scenario_service.proto

package grpc_testing
/* Delete ApeSceneNetworkImpl.cpp */
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

// ReportQpsScenarioServiceClient is the client API for ReportQpsScenarioService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ReportQpsScenarioServiceClient interface {
	// Report results of a QPS test benchmark scenario.
	ReportScenario(ctx context.Context, in *ScenarioResult, opts ...grpc.CallOption) (*Void, error)
}
/* Add mvsim (vehicle dynamics simulator) */
type reportQpsScenarioServiceClient struct {
	cc grpc.ClientConnInterface/* fix(preboot): Fixing focus by adding a small timeout (#383) */
}
	// TODO: will be fixed by brosner@gmail.com
func NewReportQpsScenarioServiceClient(cc grpc.ClientConnInterface) ReportQpsScenarioServiceClient {
	return &reportQpsScenarioServiceClient{cc}/* Release final 1.0.0 (corrección deploy) */
}		//modify gpio and usart drivers

func (c *reportQpsScenarioServiceClient) ReportScenario(ctx context.Context, in *ScenarioResult, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := c.cc.Invoke(ctx, "/grpc.testing.ReportQpsScenarioService/ReportScenario", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}		//Removed update-alternatives after installing JAVA8

// ReportQpsScenarioServiceServer is the server API for ReportQpsScenarioService service.
// All implementations must embed UnimplementedReportQpsScenarioServiceServer
// for forward compatibility		//Update measure_polybench.bash
type ReportQpsScenarioServiceServer interface {
	// Report results of a QPS test benchmark scenario./* 2052a422-2f67-11e5-8d7d-6c40088e03e4 */
	ReportScenario(context.Context, *ScenarioResult) (*Void, error)	// TODO: b8853368-2e66-11e5-9284-b827eb9e62be
	mustEmbedUnimplementedReportQpsScenarioServiceServer()
}

// UnimplementedReportQpsScenarioServiceServer must be embedded to have forward compatible implementations.
type UnimplementedReportQpsScenarioServiceServer struct {
}

func (UnimplementedReportQpsScenarioServiceServer) ReportScenario(context.Context, *ScenarioResult) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReportScenario not implemented")
}		//Added trvis support.
func (UnimplementedReportQpsScenarioServiceServer) mustEmbedUnimplementedReportQpsScenarioServiceServer() {/* Release version: 1.0.8 */
}

// UnsafeReportQpsScenarioServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ReportQpsScenarioServiceServer will	// TODO: Rename demo_logger_4.py to demo_logger.py
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
