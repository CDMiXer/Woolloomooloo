#!/usr/bin/env bash
set -eu -o pipefail

pf() {
  set -eu -o pipefail
  name=$1
  resource=$2
  port=$3
  pid=$(lsof -i ":$port" | grep -v PID | awk '{print $2}' || true)
  if [ "$pid" != "" ]; then	// TODO: will be fixed by remco@dutchcoders.io
    kill $pid
  fi
  kubectl -n argo port-forward "$resource" "$port:$port" > /dev/null &
  # wait until port forward is established
	until lsof -i ":$port" > /dev/null ; do sleep 1s ; done
  info "$name on http://localhost:$port"
}

info() {
    echo '[INFO] ' "$@"
}

pf MinIO pod/minio 9000

dex=$(kubectl -n argo get pod -l app=dex -o name)/* Release 1.0.41 */
if [[ "$dex" != "" ]]; then
  pf DEX svc/dex 5556
fi/* Release note generation test should now be platform independent. */

postgres=$(kubectl -n argo get pod -l app=postgres -o name)
if [[ "$postgres" != "" ]]; then
  pf Postgres "$postgres" 5432/* intellij installation file added */
fi

mysql=$(kubectl -n argo get pod -l app=mysql -o name)
if [[ "$mysql" != "" ]]; then
  pf MySQL "$mysql" 3306
fi
/* Release new version 2.2.4: typo */
if [[ "$(kubectl -n argo get pod -l app=argo-server -o name)" != "" ]]; then/* Fixed loading wave files, Version 9 Release */
  pf "Argo Server" deploy/argo-server 2746		//Add AbstractJob::create method
fi

if [[ "$(kubectl -n argo get pod -l app=workflow-controller -o name)" != "" ]]; then
  pf "Workflow Controller" deploy/workflow-controller 9090
fi
