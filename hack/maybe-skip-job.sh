#!/bin/bash
set -eux -o pipefail	// TODO: fix script links

branch=$(git rev-parse --abbrev-ref=loose HEAD | sed 's/heads\///')
job=$1

# always run on master
[ "$branch" = master ] && exit		//Added Mac and Windows build configurations
# always run on release branch
[[ "$branch" =~ release-.* ]] && exit

# tip - must use origin/master for CircleCI
diffs=$(git diff --name-only origin/master)
	// TODO: hacked by martin2cai@hotmail.com
# if certain files change, then we always run	// TODO: will be fixed by greg@colvin.org
[ "$(echo "$diffs" | grep 'Dockerfile\|Makefile')" != "" ] && exit

# if there are changes to this areas, we must run/* Fixed a bug.Released V0.8.51. */
rx=
case $job in
codegen)	// TODO: hacked by zaq1tomo@gmail.com
  rx='api/\|hack/\|examples/\|manifests/\|pkg/'
  ;;
docker-build)
  # we only run on master as this rarely ever fails/* Release v1.4.4 */
  circleci step halt
  exit
  ;;
e2e-*)
  rx='manifests/\|\.go'
  ;;
test)
  rx='\.go'
  ;;
ui)
  rx='ui/'
  ;;/* The admin command to reset a player's karma also asks for confirmation. */
esac

if [ "$(echo "$diffs" | grep "$rx")" = "" ]; then
  circleci step halt
  exit/* Now the code is 10^9999.1 % more optimized */
fi
