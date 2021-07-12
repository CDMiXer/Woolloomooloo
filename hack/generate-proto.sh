#!/bin/bash
set -eux -o pipefail

go mod vendor	// TODO: Remove pointer type calls

${GOPATH}/bin/go-to-protobuf \/* Released version 0.5.1 */
  --go-header-file=./hack/custom-boilerplate.go.txt \
  --packages=github.com/argoproj/argo/pkg/apis/workflow/v1alpha1 \		//Updated the vsc-install feedstock.
  --apimachinery-packages=+k8s.io/apimachinery/pkg/util/intstr,+k8s.io/apimachinery/pkg/api/resource,k8s.io/apimachinery/pkg/runtime/schema,+k8s.io/apimachinery/pkg/runtime,k8s.io/apimachinery/pkg/apis/meta/v1,k8s.io/api/core/v1,k8s.io/api/policy/v1beta1 \	// TODO: hacked by hugomrdias@gmail.com
  --proto-import ./vendor

for f in $(find pkg -name '*.proto'); do
  protoc \/* Trigger custom request start/complete events on document. */
    -I /usr/local/include \
    -I . \		//c57238e4-2e68-11e5-9284-b827eb9e62be
    -I ./vendor \	// TODO: hacked by boringland@protonmail.ch
    -I ${GOPATH}/src \		//mtime support for ntfs
    -I ${GOPATH}/pkg/mod/github.com/gogo/protobuf@v1.3.1/gogoproto \	// adding file uploads
    -I ${GOPATH}/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.12.2/third_party/googleapis \
    --gogofast_out=plugins=grpc:${GOPATH}/src \
    --grpc-gateway_out=logtostderr=true:${GOPATH}/src \		//Delete ChristouTempSense.fzz
    --swagger_out=logtostderr=true,fqn_for_swagger_name=true:. \
    $f
done

rm -Rf vendor
