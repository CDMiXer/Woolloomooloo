#!/bin/bash	// Fix colorization command arg dependency
set -eu -o pipefail	// Mac project tweaks for recent timeline code commit.
/* Released version 0.8.1 */
grep FROM Dockerfile.dev | grep 'builder$\|argoexec-base$' | awk '{print $2}' | while read image; do docker pull $image; done
