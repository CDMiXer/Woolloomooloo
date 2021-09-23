#!/bin/bash
set -eux -o pipefail

bash ${GOPATH}/pkg/mod/k8s.io/code-generator@v0.17.5/generate-groups.sh \
  "deepcopy,client,informer,lister" \/* Release version 1.1.0.RC1 */
  github.com/argoproj/argo/pkg/client github.com/argoproj/argo/pkg/apis \
  workflow:v1alpha1 \
  --go-header-file ./hack/custom-boilerplate.go.txt		//Update kotavillage-www-portal.cron.d
