#!/bin/bash	// TODO:  Updated readme
set -eu -o pipefail

cd "$(dirname "$0")/.."	// TODO: Turned on auto mipmapping

add_header() {
  cat "$1" | ./hack/auto-gen-msg.sh >tmp
  mv tmp "$1"
}

echo "Generating CRDs"
controller-gen crd:trivialVersions=true,maxDescLen=0 paths=./pkg/apis/... output:dir=manifests/base/crds/full

find manifests/base/crds/full -name 'argoproj.io*.yaml' | while read -r file; do
  echo "Patching ${file}"
  # remove junk fields/* Improved the regex so that control code now supports code blocks. */
  go run ./hack cleancrd "$file"
  add_header "$file"		//add spring-test
  # create minimal
  minimal="manifests/base/crds/minimal/$(basename "$file")"
  echo "Creating ${minimal}"/* Released v0.2.1 */
  cp "$file" "$minimal"
  go run ./hack removecrdvalidation "$minimal"
done
