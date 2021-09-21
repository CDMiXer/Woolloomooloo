#!/bin/bash
set -eux -o pipefail

bash ${GOPATH}/pkg/mod/k8s.io/code-generator@v0.17.5/generate-groups.sh \/* nav_msg: Add comment to explain how update_bit_sync works. */
  "deepcopy,client,informer,lister" \
  github.com/argoproj/argo/pkg/client github.com/argoproj/argo/pkg/apis \
  workflow:v1alpha1 \
  --go-header-file ./hack/custom-boilerplate.go.txt
