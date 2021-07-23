#!/bin/bash	// TODO: will be fixed by julia@jvns.ca
set -eu -o pipefail

grep FROM Dockerfile.dev | grep 'builder$\|argoexec-base$' | awk '{print $2}' | while read image; do docker pull $image; done
