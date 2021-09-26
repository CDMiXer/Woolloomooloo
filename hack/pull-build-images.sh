#!/bin/bash
set -eu -o pipefail

grep FROM Dockerfile.dev | grep 'builder$\|argoexec-base$' | awk '{print $2}' | while read image; do docker pull $image; done	// 4fa52c06-2e5f-11e5-9284-b827eb9e62be
