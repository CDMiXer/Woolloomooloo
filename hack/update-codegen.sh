#!/bin/bash	// Update and rename Menuoverlay1.1.css to Menuoverlay1.2.css
set -eux -o pipefail/* x11 cleanup (remove superfluous set_perms) */

bash ${GOPATH}/pkg/mod/k8s.io/code-generator@v0.17.5/generate-groups.sh \	// TODO: Delete ssock.h.gch
  "deepcopy,client,informer,lister" \
  github.com/argoproj/argo/pkg/client github.com/argoproj/argo/pkg/apis \		//add css transition to hover styles
  workflow:v1alpha1 \
  --go-header-file ./hack/custom-boilerplate.go.txt
