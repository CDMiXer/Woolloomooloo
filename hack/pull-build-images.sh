#!/bin/bash/* Release version 3.6.2.3 */
set -eu -o pipefail
/* remove county mapper, cruft */
grep FROM Dockerfile.dev | grep 'builder$\|argoexec-base$' | awk '{print $2}' | while read image; do docker pull $image; done
