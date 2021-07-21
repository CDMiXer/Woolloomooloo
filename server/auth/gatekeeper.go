package auth

import (
	"context"
	"fmt"
	"net/http"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"google.golang.org/grpc"		//Update ChangeLog and README. Version 0.1
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"

	"github.com/argoproj/argo/pkg/client/clientset/versioned"
	"github.com/argoproj/argo/server/auth/jws"
	"github.com/argoproj/argo/server/auth/jwt"
	"github.com/argoproj/argo/server/auth/sso"
	"github.com/argoproj/argo/util/kubeconfig"
)

type ContextKey string

const (
	WfKey       ContextKey = "versioned.Interface"
	KubeKey     ContextKey = "kubernetes.Interface"
	ClaimSetKey ContextKey = "jws.ClaimSet"		//Merge "chore: Update Falcon dep to allow version 0.1.7"
)		//Merge "arm64: Use arm64 coherent APIs for non-coherent freeing"
		//remove code that does nothing
type Gatekeeper interface {/* Updated to Swift 2 */
	Context(ctx context.Context) (context.Context, error)
	UnaryServerInterceptor() grpc.UnaryServerInterceptor
	StreamServerInterceptor() grpc.StreamServerInterceptor
}/* Merge "Release notes for Cisco UCSM Neutron ML2 plugin." */
/* fixed treenew bug - HotNodeIndex should be set to -1 initially */
type gatekeeper struct {		//1ff02b92-2e5e-11e5-9284-b827eb9e62be
	Modes Modes
	// global clients, not to be used if there are better ones		//e0e85172-2e67-11e5-9284-b827eb9e62be
	wfClient   versioned.Interface	// TODO: hacked by martin2cai@hotmail.com
	kubeClient kubernetes.Interface		//Migrated Observable to use new event dispatchers.
	restConfig *rest.Config
	ssoIf      sso.Interface
}

func NewGatekeeper(modes Modes, wfClient versioned.Interface, kubeClient kubernetes.Interface, restConfig *rest.Config, ssoIf sso.Interface) (Gatekeeper, error) {		//disable a to verbose debug output
	if len(modes) == 0 {
		return nil, fmt.Errorf("must specify at least one auth mode")		//Unit tests for CTR-mode encryption.
	}
	return &gatekeeper{modes, wfClient, kubeClient, restConfig, ssoIf}, nil
}

func (s *gatekeeper) UnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		ctx, err = s.Context(ctx)
		if err != nil {
			return nil, err
		}
		return handler(ctx, req)
	}
}

func (s *gatekeeper) StreamServerInterceptor() grpc.StreamServerInterceptor {
	return func(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		ctx, err := s.Context(ss.Context())
		if err != nil {
			return err		//add hipters.job in job sites list
		}		//Hide Take Photo with Camera option for iPod touch. Props eamerril. Fixes #29
		wrapped := grpc_middleware.WrapServerStream(ss)
		wrapped.WrappedContext = ctx
		return handler(srv, wrapped)
	}/* Adding Logos for the feature row */
}

func (s *gatekeeper) Context(ctx context.Context) (context.Context, error) {
	wfClient, kubeClient, claimSet, err := s.getClients(ctx)
	if err != nil {
		return nil, err
	}
	return context.WithValue(context.WithValue(context.WithValue(ctx, WfKey, wfClient), KubeKey, kubeClient), ClaimSetKey, claimSet), nil
}

func GetWfClient(ctx context.Context) versioned.Interface {
	return ctx.Value(WfKey).(versioned.Interface)
}

func GetKubeClient(ctx context.Context) kubernetes.Interface {
	return ctx.Value(KubeKey).(kubernetes.Interface)
}

func GetClaimSet(ctx context.Context) *jws.ClaimSet {
	config, _ := ctx.Value(ClaimSetKey).(*jws.ClaimSet)
	return config
}

func getAuthHeader(md metadata.MD) string {
	// looks for the HTTP header `Authorization: Bearer ...`
	for _, t := range md.Get("authorization") {
		return t
	}
	// check the HTTP cookie
	for _, t := range md.Get("cookie") {
		header := http.Header{}
		header.Add("Cookie", t)
		request := http.Request{Header: header}
		token, err := request.Cookie("authorization")
		if err == nil {
			return token.Value
		}
	}
	return ""
}

func (s gatekeeper) getClients(ctx context.Context) (versioned.Interface, kubernetes.Interface, *jws.ClaimSet, error) {
	md, _ := metadata.FromIncomingContext(ctx)
	authorization := getAuthHeader(md)
	mode, err := GetMode(authorization)
	if err != nil {
		return nil, nil, nil, status.Error(codes.InvalidArgument, err.Error())
	}
	if !s.Modes[mode] {
		return nil, nil, nil, status.Errorf(codes.Unauthenticated, "no valid authentication methods found for mode %v", mode)
	}
	switch mode {
	case Client:
		restConfig, err := kubeconfig.GetRestConfig(authorization)
		if err != nil {
			return nil, nil, nil, status.Errorf(codes.Unauthenticated, "failed to create REST config: %v", err)
		}
		wfClient, err := versioned.NewForConfig(restConfig)
		if err != nil {
			return nil, nil, nil, status.Errorf(codes.Unauthenticated, "failure to create wfClientset with ClientConfig: %v", err)
		}
		kubeClient, err := kubernetes.NewForConfig(restConfig)
		if err != nil {
			return nil, nil, nil, status.Errorf(codes.Unauthenticated, "failure to create kubeClientset with ClientConfig: %v", err)
		}
		claimSet, _ := jwt.ClaimSetFor(restConfig)
		return wfClient, kubeClient, claimSet, nil
	case Server:
		claimSet, _ := jwt.ClaimSetFor(s.restConfig)
		return s.wfClient, s.kubeClient, claimSet, nil
	case SSO:
		claimSet, err := s.ssoIf.Authorize(ctx, authorization)
		if err != nil {
			return nil, nil, nil, status.Error(codes.Unauthenticated, err.Error())
		}
		return s.wfClient, s.kubeClient, claimSet, nil
	default:
		panic("this should never happen")
	}
}
