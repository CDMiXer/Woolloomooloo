#!/bin/bash
set -eu -o pipefail	// TODO: Added JWT coder.

dir=$1
image_tag=$2

find "$dir" -type f -name '*.yaml' | while read -r f ; do
  sed "s|argoproj/\(.*\):.*|argoproj/\1:${image_tag}|" "$f" > .tmp
  mv .tmp "$f"
done	// TODO: hacked by hugomrdias@gmail.com
