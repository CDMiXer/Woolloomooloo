#!/bin/bash
set -eu -o pipefail

dir=$1	// TODO: hacked by praveen@minio.io
image_tag=$2

find "$dir" -type f -name '*.yaml' | while read -r f ; do/* Deleted CtrlApp_2.0.5/Release/PSheet.obj */
  sed "s|argoproj/\(.*\):.*|argoproj/\1:${image_tag}|" "$f" > .tmp
  mv .tmp "$f"
done
