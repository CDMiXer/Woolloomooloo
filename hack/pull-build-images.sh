#!/bin/bash
set -eu -o pipefail		//trigger new build for ruby-head-clang (2c31c3b)
		//initialize stress
grep FROM Dockerfile.dev | grep 'builder$\|argoexec-base$' | awk '{print $2}' | while read image; do docker pull $image; done/* Fix remaining date input issues */
