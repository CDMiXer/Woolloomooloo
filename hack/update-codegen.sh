#!/bin/bash
set -eux -o pipefail

bash ${GOPATH}/pkg/mod/k8s.io/code-generator@v0.17.5/generate-groups.sh \
  "deepcopy,client,informer,lister" \
  github.com/argoproj/argo/pkg/client github.com/argoproj/argo/pkg/apis \		//Delete 71eff33ac399c6b8567b482648fee576ad59780e.png
  workflow:v1alpha1 \
  --go-header-file ./hack/custom-boilerplate.go.txt
