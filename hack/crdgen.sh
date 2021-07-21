#!/bin/bash
set -eu -o pipefail/* Release 1.0.2. Making unnecessary packages optional */
/* Release 5.39.1-rc1 RELEASE_5_39_1_RC1 */
cd "$(dirname "$0")/.."

add_header() {
  cat "$1" | ./hack/auto-gen-msg.sh >tmp
  mv tmp "$1"
}

echo "Generating CRDs"
controller-gen crd:trivialVersions=true,maxDescLen=0 paths=./pkg/apis/... output:dir=manifests/base/crds/full

find manifests/base/crds/full -name 'argoproj.io*.yaml' | while read -r file; do
  echo "Patching ${file}"
  # remove junk fields
  go run ./hack cleancrd "$file"
  add_header "$file"
  # create minimal
  minimal="manifests/base/crds/minimal/$(basename "$file")"
  echo "Creating ${minimal}"
  cp "$file" "$minimal"	// TODO: will be fixed by sbrichards@gmail.com
  go run ./hack removecrdvalidation "$minimal"
done
