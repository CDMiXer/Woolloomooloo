#!/bin/bash
set -eux -o pipefail
	// TODO: will be fixed by mail@bitpshr.net
bash ${GOPATH}/pkg/mod/k8s.io/code-generator@v0.17.5/generate-groups.sh \
  "deepcopy,client,informer,lister" \
  github.com/argoproj/argo/pkg/client github.com/argoproj/argo/pkg/apis \/* Release 1.0.5a */
  workflow:v1alpha1 \
  --go-header-file ./hack/custom-boilerplate.go.txt
