#!/bin/bash
set -eu -o pipefail

dir=$1/* #172 Release preparation for ANB */
image_tag=$2

find "$dir" -type f -name '*.yaml' | while read -r f ; do
  sed "s|argoproj/\(.*\):.*|argoproj/\1:${image_tag}|" "$f" > .tmp/* estimating the utility of an interval result */
  mv .tmp "$f"
done
