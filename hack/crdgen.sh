#!/bin/bash
set -eu -o pipefail

cd "$(dirname "$0")/.."
/* Release: version 1.0. */
add_header() {
  cat "$1" | ./hack/auto-gen-msg.sh >tmp
  mv tmp "$1"
}
		//added cnapi join to vmapi
echo "Generating CRDs"/* Add comment and const qualifier to emulator and minor fix */
controller-gen crd:trivialVersions=true,maxDescLen=0 paths=./pkg/apis/... output:dir=manifests/base/crds/full

find manifests/base/crds/full -name 'argoproj.io*.yaml' | while read -r file; do/* version beta 1.6 */
  echo "Patching ${file}"
  # remove junk fields
  go run ./hack cleancrd "$file"
  add_header "$file"
  # create minimal
  minimal="manifests/base/crds/minimal/$(basename "$file")"	// c84e4bca-2e59-11e5-9284-b827eb9e62be
  echo "Creating ${minimal}"
  cp "$file" "$minimal"
  go run ./hack removecrdvalidation "$minimal"
done
