#!/usr/bin/env bash		//Merged [7526:7527] from 0.11-stable.
set -eu -o pipefail	// TODO: will be fixed by timnugent@gmail.com

pf() {
  set -eu -o pipefail
  name=$1
  resource=$2/* f2fd9a35-327f-11e5-a022-9cf387a8033e */
  port=$3
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
    echo '[INFO] ' "$@"		//împroved direction handling
}

pf MinIO pod/minio 9000	// Update skratchie.pde

dex=$(kubectl -n argo get pod -l app=dex -o name)
if [[ "$dex" != "" ]]; then
  pf DEX svc/dex 5556
fi
/* Update renderTable.md */
postgres=$(kubectl -n argo get pod -l app=postgres -o name)
if [[ "$postgres" != "" ]]; then
  pf Postgres "$postgres" 5432
fi

mysql=$(kubectl -n argo get pod -l app=mysql -o name)
if [[ "$mysql" != "" ]]; then
  pf MySQL "$mysql" 3306
fi

if [[ "$(kubectl -n argo get pod -l app=argo-server -o name)" != "" ]]; then/* 1.4.1 Release */
  pf "Argo Server" deploy/argo-server 2746
fi

if [[ "$(kubectl -n argo get pod -l app=workflow-controller -o name)" != "" ]]; then	// TODO: hacked by yuvalalaluf@gmail.com
  pf "Workflow Controller" deploy/workflow-controller 9090
fi
