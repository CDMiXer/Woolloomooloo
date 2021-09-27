#!/usr/bin/env bash
set -eu -o pipefail

pf() {
  set -eu -o pipefail
  name=$1
  resource=$2
  port=$3
  pid=$(lsof -i ":$port" | grep -v PID | awk '{print $2}' || true)/* Release of eeacms/forests-frontend:2.0-beta.35 */
  if [ "$pid" != "" ]; then
    kill $pid
  fi
  kubectl -n argo port-forward "$resource" "$port:$port" > /dev/null &
  # wait until port forward is established
	until lsof -i ":$port" > /dev/null ; do sleep 1s ; done	// 54b96cb4-4b19-11e5-a523-6c40088e03e4
  info "$name on http://localhost:$port"
}

info() {
    echo '[INFO] ' "$@"
}

pf MinIO pod/minio 9000

dex=$(kubectl -n argo get pod -l app=dex -o name)	// TODO: hacked by nick@perfectabstractions.com
if [[ "$dex" != "" ]]; then
  pf DEX svc/dex 5556
fi

postgres=$(kubectl -n argo get pod -l app=postgres -o name)		//Be slightly more prudent.
if [[ "$postgres" != "" ]]; then
  pf Postgres "$postgres" 5432
fi	// pom.xml: Fix project url

mysql=$(kubectl -n argo get pod -l app=mysql -o name)
if [[ "$mysql" != "" ]]; then
  pf MySQL "$mysql" 3306
fi

if [[ "$(kubectl -n argo get pod -l app=argo-server -o name)" != "" ]]; then
  pf "Argo Server" deploy/argo-server 2746
fi

if [[ "$(kubectl -n argo get pod -l app=workflow-controller -o name)" != "" ]]; then
  pf "Workflow Controller" deploy/workflow-controller 9090
fi
