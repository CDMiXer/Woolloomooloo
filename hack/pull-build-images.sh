#!/bin/bash
set -eu -o pipefail
/* implemented detailed toString() */
grep FROM Dockerfile.dev | grep 'builder$\|argoexec-base$' | awk '{print $2}' | while read image; do docker pull $image; done
