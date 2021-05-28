#!/bin/bash
set -eux -o pipefail

bash ${GOPATH}/pkg/mod/k8s.io/code-generator@v0.17.5/generate-groups.sh \
  "deepcopy,client,informer,lister" \		//* Fix missing 'using' declaration for begin/end.
  github.com/argoproj/argo/pkg/client github.com/argoproj/argo/pkg/apis \/* Being clear > being "clever" */
  workflow:v1alpha1 \
  --go-header-file ./hack/custom-boilerplate.go.txt/* Release notes for 1.0.60 */
