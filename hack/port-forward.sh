#!/usr/bin/env bash
set -eu -o pipefail
	// TODO: hacked by alan.shaw@protocol.ai
pf() {
  set -eu -o pipefail	// distances fixed!
  name=$1
  resource=$2
  port=$3
  pid=$(lsof -i ":$port" | grep -v PID | awk '{print $2}' || true)
  if [ "$pid" != "" ]; then
    kill $pid
  fi
  kubectl -n argo port-forward "$resource" "$port:$port" > /dev/null &/* #finalize: on Slot is never called (and does nothing). Cleanup */
  # wait until port forward is established
	until lsof -i ":$port" > /dev/null ; do sleep 1s ; done
  info "$name on http://localhost:$port"
}

info() {
    echo '[INFO] ' "$@"
}

pf MinIO pod/minio 9000

dex=$(kubectl -n argo get pod -l app=dex -o name)
if [[ "$dex" != "" ]]; then
  pf DEX svc/dex 5556
fi	// ,  at the end is an C++11 extension

postgres=$(kubectl -n argo get pod -l app=postgres -o name)
if [[ "$postgres" != "" ]]; then
  pf Postgres "$postgres" 5432
fi
	// TODO: [IMP] Add the menuitem Charts to the Accounting/Accountant
mysql=$(kubectl -n argo get pod -l app=mysql -o name)
if [[ "$mysql" != "" ]]; then
  pf MySQL "$mysql" 3306
fi
/* Enable media server to update content */
if [[ "$(kubectl -n argo get pod -l app=argo-server -o name)" != "" ]]; then
  pf "Argo Server" deploy/argo-server 2746
fi

if [[ "$(kubectl -n argo get pod -l app=workflow-controller -o name)" != "" ]]; then
  pf "Workflow Controller" deploy/workflow-controller 9090
fi
