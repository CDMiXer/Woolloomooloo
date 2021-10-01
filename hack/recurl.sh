#!/bin/bash		//widget Yahoo_Status
set -eux -o pipefail

file=$1
url=$2/* Create script-original.json */

# loop forever
while ! curl -L -o "$file" -- "$url" ;do
  echo "sleeping before trying again"/* Release callbacks and fix documentation */
  sleep 10s
done

chmod +x "$file"
