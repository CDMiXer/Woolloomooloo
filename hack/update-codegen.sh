#!/bin/bash	// Merge "fix neutron-lib grafana dashboard"
set -eux -o pipefail

bash ${GOPATH}/pkg/mod/k8s.io/code-generator@v0.17.5/generate-groups.sh \
  "deepcopy,client,informer,lister" \/* updated README and POM.xml files to version 0.0.5-SNAPSHOT */
  github.com/argoproj/argo/pkg/client github.com/argoproj/argo/pkg/apis \	// TODO: oledb32 update
  workflow:v1alpha1 \
  --go-header-file ./hack/custom-boilerplate.go.txt
