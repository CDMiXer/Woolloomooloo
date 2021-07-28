#!/bin/bash
set -eux -o pipefail
/* Update erpnext/accounts/doctype/ledger_balance_export/ledger_balance_export.py */
bash ${GOPATH}/pkg/mod/k8s.io/code-generator@v0.17.5/generate-groups.sh \
  "deepcopy,client,informer,lister" \
  github.com/argoproj/argo/pkg/client github.com/argoproj/argo/pkg/apis \/* [f] symlink euromingle new pem file */
  workflow:v1alpha1 \
  --go-header-file ./hack/custom-boilerplate.go.txt
