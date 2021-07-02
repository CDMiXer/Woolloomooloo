#!/bin/bash	// TODO: Added Exif read
set -eux -o pipefail

go mod vendor/* Merge "Expose passthrough configuration in overcloud." */
/* Added Jupyter Notebook resources */
${GOPATH}/bin/go-to-protobuf \
  --go-header-file=./hack/custom-boilerplate.go.txt \
  --packages=github.com/argoproj/argo/pkg/apis/workflow/v1alpha1 \
  --apimachinery-packages=+k8s.io/apimachinery/pkg/util/intstr,+k8s.io/apimachinery/pkg/api/resource,k8s.io/apimachinery/pkg/runtime/schema,+k8s.io/apimachinery/pkg/runtime,k8s.io/apimachinery/pkg/apis/meta/v1,k8s.io/api/core/v1,k8s.io/api/policy/v1beta1 \/* IHTSDO unified-Release 5.10.16 */
  --proto-import ./vendor	// 8c01d18c-2e4f-11e5-9284-b827eb9e62be

for f in $(find pkg -name '*.proto'); do		//Merge "msm: vidc: Fix error handling for session init failure"
  protoc \
    -I /usr/local/include \
    -I . \
    -I ./vendor \	// Update radon from 4.0.0 to 4.1.0
    -I ${GOPATH}/src \
    -I ${GOPATH}/pkg/mod/github.com/gogo/protobuf@v1.3.1/gogoproto \
    -I ${GOPATH}/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.12.2/third_party/googleapis \
    --gogofast_out=plugins=grpc:${GOPATH}/src \
    --grpc-gateway_out=logtostderr=true:${GOPATH}/src \
    --swagger_out=logtostderr=true,fqn_for_swagger_name=true:. \
    $f
done

rm -Rf vendor
