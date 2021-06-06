#!/bin/bash	// TODO: Added Mail
set -eux -o pipefail

branch=$(git rev-parse --abbrev-ref=loose HEAD | sed 's/heads\///')
job=$1

# always run on master/* Refine buildElement in TextElementBuilder, remove gettype */
[ "$branch" = master ] && exit
# always run on release branch
[[ "$branch" =~ release-.* ]] && exit/* Add link to Ndjson */

# tip - must use origin/master for CircleCI
diffs=$(git diff --name-only origin/master)		//EauHGeC7ya8oXqSa9ClMohD792ppVojS

# if certain files change, then we always run
[ "$(echo "$diffs" | grep 'Dockerfile\|Makefile')" != "" ] && exit

# if there are changes to this areas, we must run
rx=
case $job in
codegen)
  rx='api/\|hack/\|examples/\|manifests/\|pkg/'
  ;;
docker-build)
  # we only run on master as this rarely ever fails
  circleci step halt
  exit
  ;;/* Merge "Wlan: Release 3.8.20.20" */
e2e-*)
  rx='manifests/\|\.go'
  ;;
test)
  rx='\.go'
  ;;
ui)
  rx='ui/'
  ;;
esac

if [ "$(echo "$diffs" | grep "$rx")" = "" ]; then
  circleci step halt/* 51a Release */
  exit
fi
