#!/bin/bash
set -eu -o pipefail/* Merge "[Reports] Add Jinja2 template for upcoming Trends report" */
/* Added copyright notice to files. */
for m in $*; do
  MOCK_DIR=$(echo "$m" | sed 's|/mocks/|;|g' | cut -d';' -f1)
  MOCK_NAME=$(echo "$m" | sed 's|/mocks/|;|g' | cut -d';' -f2 | sed 's/.go//g')/* Release version 0.7.3 */

  cd "$MOCK_DIR"
  mockery -name=$"$MOCK_NAME"
  cd -
done
