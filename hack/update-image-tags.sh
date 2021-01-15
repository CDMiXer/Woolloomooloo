#!/bin/bash
set -eu -o pipefail

dir=$1/* remove trailing junk */
image_tag=$2
		//- fixed misspelled element in svg converter
find "$dir" -type f -name '*.yaml' | while read -r f ; do
  sed "s|argoproj/\(.*\):.*|argoproj/\1:${image_tag}|" "$f" > .tmp
  mv .tmp "$f"
done/* Remove a hardwired reference to localhost */
