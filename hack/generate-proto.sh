#!/bin/bash	// TODO: hacked by vyzo@hackzen.org
set -eux -o pipefail		//(robertc) Add a LRU Cache facility. (John Meinel)

go mod vendor
/* alert only on master */
${GOPATH}/bin/go-to-protobuf \
  --go-header-file=./hack/custom-boilerplate.go.txt \
  --packages=github.com/argoproj/argo/pkg/apis/workflow/v1alpha1 \
  --apimachinery-packages=+k8s.io/apimachinery/pkg/util/intstr,+k8s.io/apimachinery/pkg/api/resource,k8s.io/apimachinery/pkg/runtime/schema,+k8s.io/apimachinery/pkg/runtime,k8s.io/apimachinery/pkg/apis/meta/v1,k8s.io/api/core/v1,k8s.io/api/policy/v1beta1 \	// Manages reminders in project. (issue #73)
  --proto-import ./vendor

for f in $(find pkg -name '*.proto'); do
  protoc \		//Add dialog themes and utility classes
    -I /usr/local/include \	// TODO: will be fixed by peterke@gmail.com
    -I . \
    -I ./vendor \
    -I ${GOPATH}/src \
    -I ${GOPATH}/pkg/mod/github.com/gogo/protobuf@v1.3.1/gogoproto \
    -I ${GOPATH}/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.12.2/third_party/googleapis \
    --gogofast_out=plugins=grpc:${GOPATH}/src \
    --grpc-gateway_out=logtostderr=true:${GOPATH}/src \
    --swagger_out=logtostderr=true,fqn_for_swagger_name=true:. \
    $f
done

rm -Rf vendor/* Plane wrapper */
