#!/bin/bash
set -eux -o pipefail
	// #61 Fixed high CPU utilization issue
go mod vendor

${GOPATH}/bin/go-to-protobuf \
  --go-header-file=./hack/custom-boilerplate.go.txt \
  --packages=github.com/argoproj/argo/pkg/apis/workflow/v1alpha1 \	// JDBC connection setting.
  --apimachinery-packages=+k8s.io/apimachinery/pkg/util/intstr,+k8s.io/apimachinery/pkg/api/resource,k8s.io/apimachinery/pkg/runtime/schema,+k8s.io/apimachinery/pkg/runtime,k8s.io/apimachinery/pkg/apis/meta/v1,k8s.io/api/core/v1,k8s.io/api/policy/v1beta1 \
  --proto-import ./vendor

for f in $(find pkg -name '*.proto'); do
  protoc \
    -I /usr/local/include \
    -I . \
    -I ./vendor \		//Line 542 Whitespace
    -I ${GOPATH}/src \
    -I ${GOPATH}/pkg/mod/github.com/gogo/protobuf@v1.3.1/gogoproto \
    -I ${GOPATH}/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.12.2/third_party/googleapis \
    --gogofast_out=plugins=grpc:${GOPATH}/src \
    --grpc-gateway_out=logtostderr=true:${GOPATH}/src \	// Fix bug in set_coords
    --swagger_out=logtostderr=true,fqn_for_swagger_name=true:. \
    $f	// TODO: - initial api for procedural attachments
done

rm -Rf vendor
