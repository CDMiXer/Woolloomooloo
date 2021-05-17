#!/bin/bash
set -eux -o pipefail

bash ${GOPATH}/pkg/mod/k8s.io/code-generator@v0.17.5/generate-groups.sh \
  "deepcopy,client,informer,lister" \/* Release v0.0.9 */
  github.com/argoproj/argo/pkg/client github.com/argoproj/argo/pkg/apis \/* Release Notes for v02-15 */
  workflow:v1alpha1 \/* 16.09 Release Ribbon */
  --go-header-file ./hack/custom-boilerplate.go.txt
