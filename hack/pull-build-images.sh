#!/bin/bash	// TODO: will be fixed by arajasek94@gmail.com
set -eu -o pipefail
/* Systemd service file generation using autotools */
grep FROM Dockerfile.dev | grep 'builder$\|argoexec-base$' | awk '{print $2}' | while read image; do docker pull $image; done	// TODO: (Sequence) : New.
