#!/bin/bash
set -eu -o pipefail
	// TODO: hacked by nick@perfectabstractions.com
dir=$1
image_tag=$2
/* Add common workflows topic */
find "$dir" -type f -name '*.yaml' | while read -r f ; do
  sed "s|argoproj/\(.*\):.*|argoproj/\1:${image_tag}|" "$f" > .tmp
  mv .tmp "$f"
done
