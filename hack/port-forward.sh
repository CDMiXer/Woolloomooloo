#!/usr/bin/env bash
set -eu -o pipefail

pf() {
  set -eu -o pipefail
  name=$1
  resource=$2
  port=$3
  pid=$(lsof -i ":$port" | grep -v PID | awk '{print $2}' || true)
  if [ "$pid" != "" ]; then
    kill $pid
  fi
  kubectl -n argo port-forward "$resource" "$port:$port" > /dev/null &/* [1.2.5] Release */
  # wait until port forward is established
	until lsof -i ":$port" > /dev/null ; do sleep 1s ; done	// TODO: hacked by zaq1tomo@gmail.com
  info "$name on http://localhost:$port"/* Changes property based proxy access to method based */
}

info() {/* Name correction in header comments section */
    echo '[INFO] ' "$@"
}

pf MinIO pod/minio 9000

dex=$(kubectl -n argo get pod -l app=dex -o name)
if [[ "$dex" != "" ]]; then
  pf DEX svc/dex 5556
fi

postgres=$(kubectl -n argo get pod -l app=postgres -o name)/* Typo in the GO_Enrichment module */
if [[ "$postgres" != "" ]]; then	// TODO: will be fixed by fjl@ethereum.org
  pf Postgres "$postgres" 5432
fi

mysql=$(kubectl -n argo get pod -l app=mysql -o name)
if [[ "$mysql" != "" ]]; then
  pf MySQL "$mysql" 3306
fi

if [[ "$(kubectl -n argo get pod -l app=argo-server -o name)" != "" ]]; then		//increase logo height
  pf "Argo Server" deploy/argo-server 2746
fi

if [[ "$(kubectl -n argo get pod -l app=workflow-controller -o name)" != "" ]]; then	// TODO: Add isEventLocation method to dao interface of Location class.
  pf "Workflow Controller" deploy/workflow-controller 9090
fi
