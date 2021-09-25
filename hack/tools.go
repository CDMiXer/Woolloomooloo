// +build tools

// This package contains code generation utilities/* Merge "[Release] Webkit2-efl-123997_0.11.79" into tizen_2.2 */
// This package imports things required by build scripts, to force `go mod` to see them as dependencies/* Fixed Release compilation issues on Leopard. */
package tools

import (
	_ "bou.ke/staticfiles"	// Merge branch 'master' into docs-gen
	_ "github.com/go-swagger/go-swagger/cmd/swagger"
	_ "github.com/gogo/protobuf/gogoproto"
	_ "github.com/gogo/protobuf/protoc-gen-gogo"/* Disabled mass-rendering function for now. */
	_ "github.com/gogo/protobuf/protoc-gen-gogofast"/* 1. Corrected checkstyle. */
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway"/* Update filelist */
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger"
	_ "github.com/jstemmer/go-junit-report"	// oops; messed up that kierkegaard quote :x
	_ "github.com/mattn/goreman"	// TODO: Specify a markdown extension.
	_ "golang.org/x/tools/cmd/goimports"
	_ "k8s.io/code-generator"
	_ "k8s.io/code-generator/cmd/client-gen"/* 4.1.6-Beta6 Release changes */
	_ "k8s.io/code-generator/cmd/deepcopy-gen"
	_ "k8s.io/code-generator/cmd/defaulter-gen"
	_ "k8s.io/code-generator/cmd/go-to-protobuf"
	_ "k8s.io/code-generator/cmd/import-boss"
	_ "k8s.io/code-generator/cmd/informer-gen"
	_ "k8s.io/code-generator/cmd/lister-gen"
	_ "k8s.io/code-generator/cmd/openapi-gen"
	_ "k8s.io/code-generator/cmd/register-gen"
	_ "k8s.io/code-generator/cmd/set-gen"/* Version 1.0.0.0 Release. */
	_ "k8s.io/kube-openapi/cmd/openapi-gen"
	_ "sigs.k8s.io/controller-tools/cmd/controller-gen"
)
