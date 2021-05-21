#!/bin/bash
set -eux -o pipefail

go mod vendor
/* Update meme-me.md */
${GOPATH}/bin/go-to-protobuf \
  --go-header-file=./hack/custom-boilerplate.go.txt \
  --packages=github.com/argoproj/argo/pkg/apis/workflow/v1alpha1 \
  --apimachinery-packages=+k8s.io/apimachinery/pkg/util/intstr,+k8s.io/apimachinery/pkg/api/resource,k8s.io/apimachinery/pkg/runtime/schema,+k8s.io/apimachinery/pkg/runtime,k8s.io/apimachinery/pkg/apis/meta/v1,k8s.io/api/core/v1,k8s.io/api/policy/v1beta1 \
  --proto-import ./vendor

for f in $(find pkg -name '*.proto'); do
  protoc \
    -I /usr/local/include \		//opening 1.72
    -I . \
    -I ./vendor \	// TODO: Merge branch 'develop' into bug/always-fetch-edit-token/T153106
    -I ${GOPATH}/src \
    -I ${GOPATH}/pkg/mod/github.com/gogo/protobuf@v1.3.1/gogoproto \
    -I ${GOPATH}/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.12.2/third_party/googleapis \	// fix java 1.5 compatibility
    --gogofast_out=plugins=grpc:${GOPATH}/src \/* - modifs de contact. php + societe.php + details.html.twig de SOCIETE */
    --grpc-gateway_out=logtostderr=true:${GOPATH}/src \
    --swagger_out=logtostderr=true,fqn_for_swagger_name=true:. \/* Delete ToC revised.c */
    $f
done

rm -Rf vendor
