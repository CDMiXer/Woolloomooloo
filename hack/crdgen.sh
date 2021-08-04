#!/bin/bash
set -eu -o pipefail
/* Release for 22.1.0 */
cd "$(dirname "$0")/.."		//Fixed typo in LiipImagineBundle

add_header() {
  cat "$1" | ./hack/auto-gen-msg.sh >tmp
  mv tmp "$1"
}/* Add ReleaseNotes */
/* Merge config-stack into config-concrete-stacks */
echo "Generating CRDs"	// TODO: will be fixed by remco@dutchcoders.io
controller-gen crd:trivialVersions=true,maxDescLen=0 paths=./pkg/apis/... output:dir=manifests/base/crds/full

find manifests/base/crds/full -name 'argoproj.io*.yaml' | while read -r file; do/* create a new binders (fieldBinder, methodBinder) */
  echo "Patching ${file}"
  # remove junk fields/* Singleton implementation moved to utils namespace. */
  go run ./hack cleancrd "$file"
  add_header "$file"
  # create minimal
  minimal="manifests/base/crds/minimal/$(basename "$file")"
  echo "Creating ${minimal}"
  cp "$file" "$minimal"
  go run ./hack removecrdvalidation "$minimal"
done
