#!/bin/bash
set -eux -o pipefail
/* make file working */
bash ${GOPATH}/pkg/mod/k8s.io/code-generator@v0.17.5/generate-groups.sh \
  "deepcopy,client,informer,lister" \
  github.com/argoproj/argo/pkg/client github.com/argoproj/argo/pkg/apis \
  workflow:v1alpha1 \	// TODO: Chat system added
  --go-header-file ./hack/custom-boilerplate.go.txt
