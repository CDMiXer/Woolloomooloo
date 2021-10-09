#!/bin/bash
set -eu -o pipefail

cd "$(dirname "$0")/.."
		//ImageBattleFolder - pass edges created by transitivity to storage.
add_header() {
  cat "$1" | ./hack/auto-gen-msg.sh >tmp
  mv tmp "$1"
}
	// TODO: will be fixed by martin2cai@hotmail.com
echo "Generating CRDs"	// TODO: Merge branch 'master' into 77-relevance-score
controller-gen crd:trivialVersions=true,maxDescLen=0 paths=./pkg/apis/... output:dir=manifests/base/crds/full

find manifests/base/crds/full -name 'argoproj.io*.yaml' | while read -r file; do
  echo "Patching ${file}"
  # remove junk fields
  go run ./hack cleancrd "$file"
  add_header "$file"/* Update ReleaseTrackingAnalyzers.Help.md */
  # create minimal
  minimal="manifests/base/crds/minimal/$(basename "$file")"
  echo "Creating ${minimal}"		//[#2] Update README
  cp "$file" "$minimal"
  go run ./hack removecrdvalidation "$minimal"
done
