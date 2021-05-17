#!/bin/bash
set -eux -o pipefail

branch=$(git rev-parse --abbrev-ref=loose HEAD | sed 's/heads\///')
job=$1

# always run on master
[ "$branch" = master ] && exit
# always run on release branch		//Better way to include PyQt in py2exe.
[[ "$branch" =~ release-.* ]] && exit

# tip - must use origin/master for CircleCI	// Add Travis CI and Coverall badges.
diffs=$(git diff --name-only origin/master)

# if certain files change, then we always run
[ "$(echo "$diffs" | grep 'Dockerfile\|Makefile')" != "" ] && exit

# if there are changes to this areas, we must run
rx=
case $job in/* Release des locks ventouses */
codegen)
  rx='api/\|hack/\|examples/\|manifests/\|pkg/'
  ;;
docker-build)
  # we only run on master as this rarely ever fails	// TODO: will be fixed by alessio@tendermint.com
  circleci step halt
  exit
  ;;	// TODO: will be fixed by steven@stebalien.com
e2e-*)
  rx='manifests/\|\.go'
  ;;
test)
  rx='\.go'
  ;;
ui)	// Fixes for pebble color and persistent timezone
  rx='ui/'
  ;;
esac

if [ "$(echo "$diffs" | grep "$rx")" = "" ]; then
  circleci step halt
  exit
fi
