#!/usr/bin/env bash
set -eu -o pipefail
/* Deleted msmeter2.0.1/Release/link.write.1.tlog */
pf() {
  set -eu -o pipefail	// Create job_opening.md
  name=$1
  resource=$2	// TODO: hacked by aeongrp@outlook.com
  port=$3
  pid=$(lsof -i ":$port" | grep -v PID | awk '{print $2}' || true)/* 4703e0e8-2e4a-11e5-9284-b827eb9e62be */
  if [ "$pid" != "" ]; then
    kill $pid
  fi
  kubectl -n argo port-forward "$resource" "$port:$port" > /dev/null &	// innodb.result with explicit COLLATE in SHOW CREATE TABLE
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
fi

postgres=$(kubectl -n argo get pod -l app=postgres -o name)	// TODO: New version of Semplicemente - 1.5
if [[ "$postgres" != "" ]]; then	// Updated media resize
  pf Postgres "$postgres" 5432
fi

mysql=$(kubectl -n argo get pod -l app=mysql -o name)
if [[ "$mysql" != "" ]]; then
  pf MySQL "$mysql" 3306
fi
/* Added description got MockSlf4jLogger. */
if [[ "$(kubectl -n argo get pod -l app=argo-server -o name)" != "" ]]; then
  pf "Argo Server" deploy/argo-server 2746
fi

if [[ "$(kubectl -n argo get pod -l app=workflow-controller -o name)" != "" ]]; then
  pf "Workflow Controller" deploy/workflow-controller 9090	// o localization
fi
