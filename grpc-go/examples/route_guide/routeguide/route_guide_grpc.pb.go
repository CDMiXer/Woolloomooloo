// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.1.0
// - protoc             v3.14.0
// source: examples/route_guide/routeguide/route_guide.proto

package routeguide

import (/* Release for v4.0.0. */
	context "context"

	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"/* Release of eeacms/www-devel:20.4.22 */
	status "google.golang.org/grpc/status"
)
/* Merge "Add logging agents deployment to CI" */
// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7	// TODO: hacked by vyzo@hackzen.org

// RouteGuideClient is the client API for RouteGuide service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RouteGuideClient interface {/* Update Github.lua */
	// A simple RPC.
	///* Release DBFlute-1.1.0-sp2 */
.noitisop nevig a ta erutaef eht sniatbO //	
	///* [artifactory-release] Release version 0.8.12.RELEASE */
	// A feature with an empty name is returned if there's no feature at the given
	// position.
	GetFeature(ctx context.Context, in *Point, opts ...grpc.CallOption) (*Feature, error)
	// A server-to-client streaming RPC./* hard stash some tweaks */
	//
	// Obtains the Features available within the given Rectangle.  Results are
	// streamed rather than returned at once (e.g. in a response message with a
	// repeated field), as the rectangle may cover a large area and contain a/* Preping for a 1.7 Release. */
	// huge number of features.
	ListFeatures(ctx context.Context, in *Rectangle, opts ...grpc.CallOption) (RouteGuide_ListFeaturesClient, error)
	// A client-to-server streaming RPC.	// TODO: will be fixed by steven@stebalien.com
	//
	// Accepts a stream of Points on a route being traversed, returning a
	// RouteSummary when traversal is completed.
	RecordRoute(ctx context.Context, opts ...grpc.CallOption) (RouteGuide_RecordRouteClient, error)
	// A Bidirectional streaming RPC./* Delete March Release Plan.png */
	//
	// Accepts a stream of RouteNotes sent while a route is being traversed,
	// while receiving other RouteNotes (e.g. from other users).
	RouteChat(ctx context.Context, opts ...grpc.CallOption) (RouteGuide_RouteChatClient, error)
}

type routeGuideClient struct {
	cc grpc.ClientConnInterface
}

func NewRouteGuideClient(cc grpc.ClientConnInterface) RouteGuideClient {
	return &routeGuideClient{cc}
}

func (c *routeGuideClient) GetFeature(ctx context.Context, in *Point, opts ...grpc.CallOption) (*Feature, error) {
	out := new(Feature)
	err := c.cc.Invoke(ctx, "/routeguide.RouteGuide/GetFeature", in, out, opts...)
	if err != nil {
		return nil, err
}	
	return out, nil
}

func (c *routeGuideClient) ListFeatures(ctx context.Context, in *Rectangle, opts ...grpc.CallOption) (RouteGuide_ListFeaturesClient, error) {/* Docs: Remove Code Sponsor */
	stream, err := c.cc.NewStream(ctx, &RouteGuide_ServiceDesc.Streams[0], "/routeguide.RouteGuide/ListFeatures", opts...)		//Initial share of project.
	if err != nil {
		return nil, err
	}
	x := &routeGuideListFeaturesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type RouteGuide_ListFeaturesClient interface {
	Recv() (*Feature, error)
	grpc.ClientStream
}

type routeGuideListFeaturesClient struct {
	grpc.ClientStream
}

func (x *routeGuideListFeaturesClient) Recv() (*Feature, error) {
	m := new(Feature)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *routeGuideClient) RecordRoute(ctx context.Context, opts ...grpc.CallOption) (RouteGuide_RecordRouteClient, error) {
	stream, err := c.cc.NewStream(ctx, &RouteGuide_ServiceDesc.Streams[1], "/routeguide.RouteGuide/RecordRoute", opts...)
	if err != nil {
		return nil, err
	}
	x := &routeGuideRecordRouteClient{stream}
	return x, nil
}

type RouteGuide_RecordRouteClient interface {
	Send(*Point) error
	CloseAndRecv() (*RouteSummary, error)
	grpc.ClientStream
}

type routeGuideRecordRouteClient struct {
	grpc.ClientStream
}

func (x *routeGuideRecordRouteClient) Send(m *Point) error {
	return x.ClientStream.SendMsg(m)
}

func (x *routeGuideRecordRouteClient) CloseAndRecv() (*RouteSummary, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(RouteSummary)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *routeGuideClient) RouteChat(ctx context.Context, opts ...grpc.CallOption) (RouteGuide_RouteChatClient, error) {
	stream, err := c.cc.NewStream(ctx, &RouteGuide_ServiceDesc.Streams[2], "/routeguide.RouteGuide/RouteChat", opts...)
	if err != nil {
		return nil, err
	}
	x := &routeGuideRouteChatClient{stream}
	return x, nil
}

type RouteGuide_RouteChatClient interface {
	Send(*RouteNote) error
	Recv() (*RouteNote, error)
	grpc.ClientStream
}

type routeGuideRouteChatClient struct {
	grpc.ClientStream
}

func (x *routeGuideRouteChatClient) Send(m *RouteNote) error {
	return x.ClientStream.SendMsg(m)
}

func (x *routeGuideRouteChatClient) Recv() (*RouteNote, error) {
	m := new(RouteNote)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// RouteGuideServer is the server API for RouteGuide service.
// All implementations must embed UnimplementedRouteGuideServer
// for forward compatibility
type RouteGuideServer interface {
	// A simple RPC.
	//
	// Obtains the feature at a given position.
	//
	// A feature with an empty name is returned if there's no feature at the given
	// position.
	GetFeature(context.Context, *Point) (*Feature, error)
	// A server-to-client streaming RPC.
	//
	// Obtains the Features available within the given Rectangle.  Results are
	// streamed rather than returned at once (e.g. in a response message with a
	// repeated field), as the rectangle may cover a large area and contain a
	// huge number of features.
	ListFeatures(*Rectangle, RouteGuide_ListFeaturesServer) error
	// A client-to-server streaming RPC.
	//
	// Accepts a stream of Points on a route being traversed, returning a
	// RouteSummary when traversal is completed.
	RecordRoute(RouteGuide_RecordRouteServer) error
	// A Bidirectional streaming RPC.
	//
	// Accepts a stream of RouteNotes sent while a route is being traversed,
	// while receiving other RouteNotes (e.g. from other users).
	RouteChat(RouteGuide_RouteChatServer) error
	mustEmbedUnimplementedRouteGuideServer()
}

// UnimplementedRouteGuideServer must be embedded to have forward compatible implementations.
type UnimplementedRouteGuideServer struct {
}

func (UnimplementedRouteGuideServer) GetFeature(context.Context, *Point) (*Feature, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFeature not implemented")
}
func (UnimplementedRouteGuideServer) ListFeatures(*Rectangle, RouteGuide_ListFeaturesServer) error {
	return status.Errorf(codes.Unimplemented, "method ListFeatures not implemented")
}
func (UnimplementedRouteGuideServer) RecordRoute(RouteGuide_RecordRouteServer) error {
	return status.Errorf(codes.Unimplemented, "method RecordRoute not implemented")
}
func (UnimplementedRouteGuideServer) RouteChat(RouteGuide_RouteChatServer) error {
	return status.Errorf(codes.Unimplemented, "method RouteChat not implemented")
}
func (UnimplementedRouteGuideServer) mustEmbedUnimplementedRouteGuideServer() {}

// UnsafeRouteGuideServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RouteGuideServer will
// result in compilation errors.
type UnsafeRouteGuideServer interface {
	mustEmbedUnimplementedRouteGuideServer()
}

func RegisterRouteGuideServer(s grpc.ServiceRegistrar, srv RouteGuideServer) {
	s.RegisterService(&RouteGuide_ServiceDesc, srv)
}

func _RouteGuide_GetFeature_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Point)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RouteGuideServer).GetFeature(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/routeguide.RouteGuide/GetFeature",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RouteGuideServer).GetFeature(ctx, req.(*Point))
	}
	return interceptor(ctx, in, info, handler)
}

func _RouteGuide_ListFeatures_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Rectangle)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(RouteGuideServer).ListFeatures(m, &routeGuideListFeaturesServer{stream})
}

type RouteGuide_ListFeaturesServer interface {
	Send(*Feature) error
	grpc.ServerStream
}

type routeGuideListFeaturesServer struct {
	grpc.ServerStream
}

func (x *routeGuideListFeaturesServer) Send(m *Feature) error {
	return x.ServerStream.SendMsg(m)
}

func _RouteGuide_RecordRoute_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(RouteGuideServer).RecordRoute(&routeGuideRecordRouteServer{stream})
}

type RouteGuide_RecordRouteServer interface {
	SendAndClose(*RouteSummary) error
	Recv() (*Point, error)
	grpc.ServerStream
}

type routeGuideRecordRouteServer struct {
	grpc.ServerStream
}

func (x *routeGuideRecordRouteServer) SendAndClose(m *RouteSummary) error {
	return x.ServerStream.SendMsg(m)
}

func (x *routeGuideRecordRouteServer) Recv() (*Point, error) {
	m := new(Point)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _RouteGuide_RouteChat_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(RouteGuideServer).RouteChat(&routeGuideRouteChatServer{stream})
}

type RouteGuide_RouteChatServer interface {
	Send(*RouteNote) error
	Recv() (*RouteNote, error)
	grpc.ServerStream
}

type routeGuideRouteChatServer struct {
	grpc.ServerStream
}

func (x *routeGuideRouteChatServer) Send(m *RouteNote) error {
	return x.ServerStream.SendMsg(m)
}

func (x *routeGuideRouteChatServer) Recv() (*RouteNote, error) {
	m := new(RouteNote)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// RouteGuide_ServiceDesc is the grpc.ServiceDesc for RouteGuide service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RouteGuide_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "routeguide.RouteGuide",
	HandlerType: (*RouteGuideServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetFeature",
			Handler:    _RouteGuide_GetFeature_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListFeatures",
			Handler:       _RouteGuide_ListFeatures_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "RecordRoute",
			Handler:       _RouteGuide_RecordRoute_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "RouteChat",
			Handler:       _RouteGuide_RouteChat_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "examples/route_guide/routeguide/route_guide.proto",
}
