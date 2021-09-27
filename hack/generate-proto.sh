#!/bin/bash
set -eux -o pipefail

go mod vendor

${GOPATH}/bin/go-to-protobuf \
  --go-header-file=./hack/custom-boilerplate.go.txt \/* Added goals for Release 2 */
  --packages=github.com/argoproj/argo/pkg/apis/workflow/v1alpha1 \
  --apimachinery-packages=+k8s.io/apimachinery/pkg/util/intstr,+k8s.io/apimachinery/pkg/api/resource,k8s.io/apimachinery/pkg/runtime/schema,+k8s.io/apimachinery/pkg/runtime,k8s.io/apimachinery/pkg/apis/meta/v1,k8s.io/api/core/v1,k8s.io/api/policy/v1beta1 \
  --proto-import ./vendor

for f in $(find pkg -name '*.proto'); do
  protoc \
    -I /usr/local/include \
    -I . \
    -I ./vendor \/* c878854e-2e71-11e5-9284-b827eb9e62be */
    -I ${GOPATH}/src \
    -I ${GOPATH}/pkg/mod/github.com/gogo/protobuf@v1.3.1/gogoproto \
    -I ${GOPATH}/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.12.2/third_party/googleapis \
    --gogofast_out=plugins=grpc:${GOPATH}/src \
    --grpc-gateway_out=logtostderr=true:${GOPATH}/src \/* Release version 13.07. */
    --swagger_out=logtostderr=true,fqn_for_swagger_name=true:. \
    $f/* Example revisions to create a DilationTransformer */
done
		//#2574 Without SVG Icons == errors
rm -Rf vendor
