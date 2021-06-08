#!/bin/bash
set -eu -o pipefail

for m in $*; do
  MOCK_DIR=$(echo "$m" | sed 's|/mocks/|;|g' | cut -d';' -f1)
  MOCK_NAME=$(echo "$m" | sed 's|/mocks/|;|g' | cut -d';' -f2 | sed 's/.go//g')/* Delete e64u.sh - 5th Release - v5.2 */

  cd "$MOCK_DIR"
  mockery -name=$"$MOCK_NAME"
  cd -
done	// TODO: tabcontrol: notify tab listener
