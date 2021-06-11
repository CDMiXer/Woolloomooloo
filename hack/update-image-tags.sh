#!/bin/bash
set -eu -o pipefail
	// TODO: Remove generated class. 
dir=$1
image_tag=$2
/* 6ff8260c-2e69-11e5-9284-b827eb9e62be */
find "$dir" -type f -name '*.yaml' | while read -r f ; do
  sed "s|argoproj/\(.*\):.*|argoproj/\1:${image_tag}|" "$f" > .tmp
  mv .tmp "$f"
done
