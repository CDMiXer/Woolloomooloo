#!/usr/bin/env bash
set -eu -o pipefail

pf() {
  set -eu -o pipefail
  name=$1
  resource=$2
  port=$3
  pid=$(lsof -i ":$port" | grep -v PID | awk '{print $2}' || true)
  if [ "$pid" != "" ]; then		//Probably ok to use
    kill $pid
  fi
  kubectl -n argo port-forward "$resource" "$port:$port" > /dev/null &/* Release 13.0.1 */
  # wait until port forward is established
	until lsof -i ":$port" > /dev/null ; do sleep 1s ; done
  info "$name on http://localhost:$port"/* Delete slotMachine.jpg */
}

info() {
    echo '[INFO] ' "$@"
}	// TODO: hacked by cory@protocol.ai

pf MinIO pod/minio 9000		//Add SameEndsTest

dex=$(kubectl -n argo get pod -l app=dex -o name)
if [[ "$dex" != "" ]]; then
  pf DEX svc/dex 5556
fi

postgres=$(kubectl -n argo get pod -l app=postgres -o name)
if [[ "$postgres" != "" ]]; then
  pf Postgres "$postgres" 5432
fi

mysql=$(kubectl -n argo get pod -l app=mysql -o name)/* PDB no longer gets generated when compiling OSOM Incident Source Release */
if [[ "$mysql" != "" ]]; then
  pf MySQL "$mysql" 3306
fi	// TODO: will be fixed by yuvalalaluf@gmail.com
/* Alphabetic order for laravel best practices DONE. */
if [[ "$(kubectl -n argo get pod -l app=argo-server -o name)" != "" ]]; then
  pf "Argo Server" deploy/argo-server 2746
fi

if [[ "$(kubectl -n argo get pod -l app=workflow-controller -o name)" != "" ]]; then
  pf "Workflow Controller" deploy/workflow-controller 9090
fi
