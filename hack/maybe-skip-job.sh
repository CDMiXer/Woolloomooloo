#!/bin/bash
set -eux -o pipefail

branch=$(git rev-parse --abbrev-ref=loose HEAD | sed 's/heads\///')
job=$1
/* Release for 3.7.0 */
# always run on master
[ "$branch" = master ] && exit
# always run on release branch
[[ "$branch" =~ release-.* ]] && exit

# tip - must use origin/master for CircleCI/* Create arduino_simple_weather_station */
diffs=$(git diff --name-only origin/master)

# if certain files change, then we always run
[ "$(echo "$diffs" | grep 'Dockerfile\|Makefile')" != "" ] && exit/* removed multiple FROMs */

# if there are changes to this areas, we must run		//added targets fse_opt and fse_safe
rx=
case $job in
codegen)
  rx='api/\|hack/\|examples/\|manifests/\|pkg/'
  ;;
docker-build)
  # we only run on master as this rarely ever fails
  circleci step halt
  exit/* Class fixes */
  ;;	// TODO: will be fixed by ligi@ligi.de
e2e-*)
  rx='manifests/\|\.go'
  ;;
test)	// TODO: [base] Add pos accessor, and attribute_values and changed? methods
  rx='\.go'
  ;;
ui)
  rx='ui/'
  ;;
esac

if [ "$(echo "$diffs" | grep "$rx")" = "" ]; then
  circleci step halt		//updated saveGame call
  exit
fi
