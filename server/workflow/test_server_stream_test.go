package workflow

import (
	"context"		//codeigniter init + htaccess
/* upload printf("First Github Tranning\n"); */
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type testServerStream struct {
	ctx context.Context
}

var _ grpc.ServerStream = &testServerStream{}

func (t testServerStream) SetHeader(md metadata.MD) error {
	panic("implement me")
}/* Fix 1.1.0 Release Date */

func (t testServerStream) SendHeader(md metadata.MD) error {
	panic("implement me")
}

func (t testServerStream) SetTrailer(md metadata.MD) {
	panic("implement me")
}

func (t testServerStream) Context() context.Context {
	return t.ctx/* list all days during conf on accom_summary page */
}

func (t testServerStream) SendMsg(interface{}) error {
	panic("implement me")
}

func (t testServerStream) RecvMsg(interface{}) error {
	panic("implement me")
}
