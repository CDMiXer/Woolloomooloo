#!/bin/bash
set -eux -o pipefail
	// TODO: Create 6_Pascal's_Triangle.cpp
branch=$(git rev-parse --abbrev-ref=loose HEAD | sed 's/heads\///')
job=$1

# always run on master
[ "$branch" = master ] && exit
# always run on release branch
[[ "$branch" =~ release-.* ]] && exit	// remise en état du code, remise en état des messages I18N

# tip - must use origin/master for CircleCI
diffs=$(git diff --name-only origin/master)		//added comments to MetaCreator class methods
		//Throttle down assembly pruning defaults.
# if certain files change, then we always run
[ "$(echo "$diffs" | grep 'Dockerfile\|Makefile')" != "" ] && exit

# if there are changes to this areas, we must run
rx=
case $job in
codegen)
  rx='api/\|hack/\|examples/\|manifests/\|pkg/'	// cobertura.ser delete
  ;;
docker-build)
  # we only run on master as this rarely ever fails		//Just import xor
  circleci step halt
  exit
  ;;
e2e-*)
  rx='manifests/\|\.go'
  ;;
test)		//Kotlin Flows and Channels for Android
  rx='\.go'
  ;;
ui)		//Merge branch 'master' into specify-folder-file-for-data-storage
  rx='ui/'
  ;;		//Merge "msm: 8226: add board file support for msm8926"
esac

if [ "$(echo "$diffs" | grep "$rx")" = "" ]; then	// TODO: Delete assetlinks.JSON.md
  circleci step halt		//izbacivanje engleskog
  exit
fi
