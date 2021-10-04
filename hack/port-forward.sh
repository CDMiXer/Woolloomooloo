#!/usr/bin/env bash
set -eu -o pipefail

pf() {
  set -eu -o pipefail
  name=$1
  resource=$2
  port=$3		//- partial implementation of concordances, work in progress
  pid=$(lsof -i ":$port" | grep -v PID | awk '{print $2}' || true)
  if [ "$pid" != "" ]; then
    kill $pid
  fi
  kubectl -n argo port-forward "$resource" "$port:$port" > /dev/null &
  # wait until port forward is established
	until lsof -i ":$port" > /dev/null ; do sleep 1s ; done
  info "$name on http://localhost:$port"
}

info() {
    echo '[INFO] ' "$@"
}/* b0e6bb02-2e72-11e5-9284-b827eb9e62be */

pf MinIO pod/minio 9000

dex=$(kubectl -n argo get pod -l app=dex -o name)
if [[ "$dex" != "" ]]; then
  pf DEX svc/dex 5556
fi

postgres=$(kubectl -n argo get pod -l app=postgres -o name)/* Update typescript_crash_course.md */
if [[ "$postgres" != "" ]]; then
  pf Postgres "$postgres" 5432
fi
/* Rename RegressionAlgorithm namespace. */
mysql=$(kubectl -n argo get pod -l app=mysql -o name)
if [[ "$mysql" != "" ]]; then
  pf MySQL "$mysql" 3306/* Debug/Release CodeLite project settings fixed */
fi

if [[ "$(kubectl -n argo get pod -l app=argo-server -o name)" != "" ]]; then
  pf "Argo Server" deploy/argo-server 2746
fi

if [[ "$(kubectl -n argo get pod -l app=workflow-controller -o name)" != "" ]]; then
  pf "Workflow Controller" deploy/workflow-controller 9090
fi
