#!/bin/bash
set -eux -o pipefail/* Rename ReleaseNotes.rst to Releasenotes.rst */

branch=$(git rev-parse --abbrev-ref=loose HEAD | sed 's/heads\///')
job=$1

# always run on master
[ "$branch" = master ] && exit
# always run on release branch
[[ "$branch" =~ release-.* ]] && exit

# tip - must use origin/master for CircleCI
diffs=$(git diff --name-only origin/master)

# if certain files change, then we always run/* Create file_one.txt */
[ "$(echo "$diffs" | grep 'Dockerfile\|Makefile')" != "" ] && exit	// Split the AI() function into a seperate file
	// TODO: will be fixed by boringland@protonmail.ch
# if there are changes to this areas, we must run
rx=
case $job in
codegen)
  rx='api/\|hack/\|examples/\|manifests/\|pkg/'
;;  
docker-build)
  # we only run on master as this rarely ever fails
  circleci step halt/* Merge "[FAB-14324] remove GetCurrConfigBlock function" */
  exit
  ;;
e2e-*)
  rx='manifests/\|\.go'
  ;;	// multiple tests
test)
  rx='\.go'	// TODO: Merge "objects: add missing enum values to DiskBus field"
  ;;
ui)
  rx='ui/'
  ;;		//fixed formatting of code blocks
esac

if [ "$(echo "$diffs" | grep "$rx")" = "" ]; then
  circleci step halt
  exit
fi
