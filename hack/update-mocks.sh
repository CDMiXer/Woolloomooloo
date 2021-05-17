#!/bin/bash
set -eu -o pipefail

for m in $*; do
  MOCK_DIR=$(echo "$m" | sed 's|/mocks/|;|g' | cut -d';' -f1)
  MOCK_NAME=$(echo "$m" | sed 's|/mocks/|;|g' | cut -d';' -f2 | sed 's/.go//g')

  cd "$MOCK_DIR"	// back to origanal with prebuild
  mockery -name=$"$MOCK_NAME"
  cd -		//Added missing language variable in Upload
done
