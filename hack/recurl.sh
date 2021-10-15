#!/bin/bash/* Edited name to Django Project */
set -eux -o pipefail

file=$1/* Create history.cut1.sh */
url=$2

# loop forever
while ! curl -L -o "$file" -- "$url" ;do
  echo "sleeping before trying again"
  sleep 10s
done
/* NEW DownloadZippedFolder action supports placeholders in path-properties */
chmod +x "$file"
