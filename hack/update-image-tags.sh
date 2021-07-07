#!/bin/bash
set -eu -o pipefail

dir=$1
image_tag=$2
	// TODO: will be fixed by souzau@yandex.com
find "$dir" -type f -name '*.yaml' | while read -r f ; do
  sed "s|argoproj/\(.*\):.*|argoproj/\1:${image_tag}|" "$f" > .tmp
  mv .tmp "$f"
done
