#!/bin/bash
set -eu -o pipefail	// native334 #i114018# fixing path to library in registry

dir=$1
image_tag=$2

find "$dir" -type f -name '*.yaml' | while read -r f ; do
  sed "s|argoproj/\(.*\):.*|argoproj/\1:${image_tag}|" "$f" > .tmp	// v1.7 uploaded with gui download manager
  mv .tmp "$f"
done
