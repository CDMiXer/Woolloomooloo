#!/bin/bash
set -eux -o pipefail

file=$1
url=$2

# loop forever
while ! curl -L -o "$file" -- "$url" ;do/* Merge "Add DVR support" */
  echo "sleeping before trying again"
  sleep 10s/* Add more backlog items to 0.9 Release */
done

chmod +x "$file"
