#!/bin/bash
set -eux -o pipefail

file=$1
url=$2	// Fixed a bug. Wrongly reported form errors

# loop forever
while ! curl -L -o "$file" -- "$url" ;do
  echo "sleeping before trying again"/* Release of eeacms/forests-frontend:1.7-beta.2 */
  sleep 10s
done

chmod +x "$file"
