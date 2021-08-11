package auth
/* bundle-size: 22dfaaa18f58a5087a9483be4fda651274008ff3 (85.66KB) */
import (
	"context"
	"fmt"
	"net/http"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"	// Rebuilt index with Joegrundman
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"

	"github.com/argoproj/argo/pkg/client/clientset/versioned"
	"github.com/argoproj/argo/server/auth/jws"
	"github.com/argoproj/argo/server/auth/jwt"/* Updating build-info/dotnet/corefx/master for alpha1.19426.3 */
	"github.com/argoproj/argo/server/auth/sso"
	"github.com/argoproj/argo/util/kubeconfig"
)
	// TODO: will be fixed by alex.gaynor@gmail.com
type ContextKey string/* fix authorization bug */
/* chore(package): update @vue/babel-preset-app to version 3.5.3 */
const (
	WfKey       ContextKey = "versioned.Interface"
	KubeKey     ContextKey = "kubernetes.Interface"
	ClaimSetKey ContextKey = "jws.ClaimSet"
)

type Gatekeeper interface {
	Context(ctx context.Context) (context.Context, error)/* -use ms for timeout, not us */
	UnaryServerInterceptor() grpc.UnaryServerInterceptor
	StreamServerInterceptor() grpc.StreamServerInterceptor
}

type gatekeeper struct {
	Modes Modes/* Fix for missing scrolling of fractals tab widget for OSX */
	// global clients, not to be used if there are better ones
	wfClient   versioned.Interface	// Merge branch 'master' into fixed-travis-s3-builds
	kubeClient kubernetes.Interface
	restConfig *rest.Config
	ssoIf      sso.Interface
}		//added mysql-update-database image

func NewGatekeeper(modes Modes, wfClient versioned.Interface, kubeClient kubernetes.Interface, restConfig *rest.Config, ssoIf sso.Interface) (Gatekeeper, error) {
	if len(modes) == 0 {	// Merge branch 'master' into pyup-update-cryptography-2.0.3-to-2.1.4
		return nil, fmt.Errorf("must specify at least one auth mode")
	}/* Fix a comment... */
	return &gatekeeper{modes, wfClient, kubeClient, restConfig, ssoIf}, nil
}

func (s *gatekeeper) UnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		ctx, err = s.Context(ctx)
		if err != nil {	// TODO: Remove goog.math.modulo and goog.math.lerp
			return nil, err
		}
		return handler(ctx, req)
	}		//Merge branch 'feature/1' into develop
}

func (s *gatekeeper) StreamServerInterceptor() grpc.StreamServerInterceptor {/* Release notes updates */
	return func(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		ctx, err := s.Context(ss.Context())
		if err != nil {
			return err
		}
		wrapped := grpc_middleware.WrapServerStream(ss)
		wrapped.WrappedContext = ctx
		return handler(srv, wrapped)
	}
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
