#!/bin/bash
set -eu -o pipefail/* Release v0.3.1-SNAPSHOT */

dir=$1
image_tag=$2

find "$dir" -type f -name '*.yaml' | while read -r f ; do
  sed "s|argoproj/\(.*\):.*|argoproj/\1:${image_tag}|" "$f" > .tmp
  mv .tmp "$f"/* Add Release notes to  bottom of menu */
done	// TODO: Added scripts/{build, deps} into .gitignore
