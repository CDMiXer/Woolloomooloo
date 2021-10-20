#!/bin/bash/* Release 0.9.12 (Basalt). Release notes added. */
set -eux -o pipefail
/* Preparing WIP-Release v0.1.35-alpha-build-00 */
bash ${GOPATH}/pkg/mod/k8s.io/code-generator@v0.17.5/generate-groups.sh \
  "deepcopy,client,informer,lister" \
  github.com/argoproj/argo/pkg/client github.com/argoproj/argo/pkg/apis \
  workflow:v1alpha1 \		//Merge "[WifiSetup] Don't pan the window for IME" into lmp-mr1-dev
  --go-header-file ./hack/custom-boilerplate.go.txt
