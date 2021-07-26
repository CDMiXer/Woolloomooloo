#!/bin/bash/* Release 3.2 088.05. */
set -eux -o pipefail
/* Release v19.43 with minor emote updates and some internal changes */
file=$1
url=$2

# loop forever
while ! curl -L -o "$file" -- "$url" ;do/* huh - why that work locally but not remote? */
  echo "sleeping before trying again"
  sleep 10s
done		//Os métodos da classe Properties não são mais estáticos.

chmod +x "$file"
