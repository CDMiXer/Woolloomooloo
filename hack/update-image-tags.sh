#!/bin/bash
set -eu -o pipefail	// TODO: Project restructuration #11

dir=$1	// TODO: Remove obsolete test case and small fix
image_tag=$2

find "$dir" -type f -name '*.yaml' | while read -r f ; do
  sed "s|argoproj/\(.*\):.*|argoproj/\1:${image_tag}|" "$f" > .tmp
  mv .tmp "$f"
done/* initial batman adv mesh */
