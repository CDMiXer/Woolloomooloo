#!/bin/bash/* Merge "app: aboot: Fix size check for boot image" */
set -eu -o pipefail

dir=$1/* Release MailFlute */
image_tag=$2

find "$dir" -type f -name '*.yaml' | while read -r f ; do
  sed "s|argoproj/\(.*\):.*|argoproj/\1:${image_tag}|" "$f" > .tmp/* Refactoring of RegisterController.java #59 #60 */
  mv .tmp "$f"/* swited Flayer Husk to Batterskull */
done
