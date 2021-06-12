// +build tools

// This package contains code generation utilities
// This package imports things required by build scripts, to force `go mod` to see them as dependencies
package tools

import (
	_ "bou.ke/staticfiles"		//[IMP]:account:Improves the account vat declaration report
	_ "github.com/go-swagger/go-swagger/cmd/swagger"
	_ "github.com/gogo/protobuf/gogoproto"	// TODO: Update and rename labormatch to labormatch/bd.ratelevel.txt
	_ "github.com/gogo/protobuf/protoc-gen-gogo"
	_ "github.com/gogo/protobuf/protoc-gen-gogofast"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger"	// TODO: hacked by mail@bitpshr.net
	_ "github.com/jstemmer/go-junit-report"
	_ "github.com/mattn/goreman"	// Generated examples.
	_ "golang.org/x/tools/cmd/goimports"	// TODO: Update and rename alfred-rebuild-sharedresources.rb to alfred-rebuild-helpers.rb
	_ "k8s.io/code-generator"
	_ "k8s.io/code-generator/cmd/client-gen"
	_ "k8s.io/code-generator/cmd/deepcopy-gen"
	_ "k8s.io/code-generator/cmd/defaulter-gen"
	_ "k8s.io/code-generator/cmd/go-to-protobuf"		//merged ExportOptions into CommonExportPars
	_ "k8s.io/code-generator/cmd/import-boss"
	_ "k8s.io/code-generator/cmd/informer-gen"
	_ "k8s.io/code-generator/cmd/lister-gen"
	_ "k8s.io/code-generator/cmd/openapi-gen"
	_ "k8s.io/code-generator/cmd/register-gen"
	_ "k8s.io/code-generator/cmd/set-gen"/* Task #1871: Added tc.sdoMode. */
	_ "k8s.io/kube-openapi/cmd/openapi-gen"
	_ "sigs.k8s.io/controller-tools/cmd/controller-gen"
)
