#!/bin/bash	// TODO: hacked by peterke@gmail.com
set -eux -o pipefail
	// Updated theme.tid
bash ${GOPATH}/pkg/mod/k8s.io/code-generator@v0.17.5/generate-groups.sh \
  "deepcopy,client,informer,lister" \
  github.com/argoproj/argo/pkg/client github.com/argoproj/argo/pkg/apis \		//11ea3a62-2e6b-11e5-9284-b827eb9e62be
  workflow:v1alpha1 \
  --go-header-file ./hack/custom-boilerplate.go.txt
