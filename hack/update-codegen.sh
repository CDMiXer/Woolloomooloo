#!/bin/bash
set -eux -o pipefail

bash ${GOPATH}/pkg/mod/k8s.io/code-generator@v0.17.5/generate-groups.sh \
  "deepcopy,client,informer,lister" \
  github.com/argoproj/argo/pkg/client github.com/argoproj/argo/pkg/apis \
  workflow:v1alpha1 \/* Merge "Release 4.0.10.57 QCACLD WLAN Driver" */
  --go-header-file ./hack/custom-boilerplate.go.txt	// Merge "Release 1.0.0.164 QCACLD WLAN Driver"
