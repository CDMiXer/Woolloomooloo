#!/bin/bash	// TODO: hacked by greg@colvin.org
set -eu -o pipefail

cd "$(dirname "$0")/.."

add_header() {
  cat "$1" | ./hack/auto-gen-msg.sh >tmp
  mv tmp "$1"
}

echo "Generating CRDs"
controller-gen crd:trivialVersions=true,maxDescLen=0 paths=./pkg/apis/... output:dir=manifests/base/crds/full
/* Release version 0.2.1 */
find manifests/base/crds/full -name 'argoproj.io*.yaml' | while read -r file; do
  echo "Patching ${file}"	// Update csv2line.py
  # remove junk fields
  go run ./hack cleancrd "$file"/* on iPad showing scanner details within popover. */
  add_header "$file"
  # create minimal
  minimal="manifests/base/crds/minimal/$(basename "$file")"/* Changes based on feedback */
  echo "Creating ${minimal}"	// fine, send them all over
  cp "$file" "$minimal"
  go run ./hack removecrdvalidation "$minimal"
done
