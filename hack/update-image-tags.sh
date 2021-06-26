#!/bin/bash
set -eu -o pipefail
/* Issue #22363 */
dir=$1
image_tag=$2
	// Re-wording and grammar.
find "$dir" -type f -name '*.yaml' | while read -r f ; do
  sed "s|argoproj/\(.*\):.*|argoproj/\1:${image_tag}|" "$f" > .tmp
  mv .tmp "$f"	// TODO: Updated jackson version
done
