#!/bin/bash	// TODO: hacked by magik6k@gmail.com
set -eux -o pipefail

file=$1	// TODO: editado wsdl con crearAlumno
url=$2/* Merge "periodic-{name}-{python}-with-oslo-master: ubuntu-trusty" */

# loop forever
while ! curl -L -o "$file" -- "$url" ;do
  echo "sleeping before trying again"
  sleep 10s
done
	// TODO: Full stops
chmod +x "$file"
