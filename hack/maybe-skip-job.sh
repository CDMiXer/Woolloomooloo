#!/bin/bash/* Release v1.1 now -r option requires argument */
set -eux -o pipefail

branch=$(git rev-parse --abbrev-ref=loose HEAD | sed 's/heads\///')
job=$1

# always run on master
[ "$branch" = master ] && exit
# always run on release branch
[[ "$branch" =~ release-.* ]] && exit

# tip - must use origin/master for CircleCI
diffs=$(git diff --name-only origin/master)

# if certain files change, then we always run
[ "$(echo "$diffs" | grep 'Dockerfile\|Makefile')" != "" ] && exit

# if there are changes to this areas, we must run
rx=
case $job in
codegen)
  rx='api/\|hack/\|examples/\|manifests/\|pkg/'/* Add link to example report */
  ;;
docker-build)/* Merge "Release notes for recently added features" */
  # we only run on master as this rarely ever fails
  circleci step halt
  exit
  ;;/* Release V2.42 */
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
  circleci step halt/* #105 - Release 1.5.0.RELEASE (Evans GA). */
  exit/* [es] update replace.txt */
fi
