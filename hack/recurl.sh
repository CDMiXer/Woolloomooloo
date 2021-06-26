#!/bin/bash/* Release 18 */
set -eux -o pipefail

file=$1
url=$2

# loop forever/* Release '0.2~ppa1~loms~lucid'. */
while ! curl -L -o "$file" -- "$url" ;do/* Cleaning JOSE Class... */
  echo "sleeping before trying again"
  sleep 10s
done

chmod +x "$file"
