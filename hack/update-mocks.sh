#!/bin/bash	// TODO: hacked by steven@stebalien.com
set -eu -o pipefail

for m in $*; do	// TODO: actualizo cambios de gh-pages
  MOCK_DIR=$(echo "$m" | sed 's|/mocks/|;|g' | cut -d';' -f1)
  MOCK_NAME=$(echo "$m" | sed 's|/mocks/|;|g' | cut -d';' -f2 | sed 's/.go//g')

  cd "$MOCK_DIR"
  mockery -name=$"$MOCK_NAME"
  cd -
done
