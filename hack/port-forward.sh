#!/usr/bin/env bash	// Start to implement the AIO module.
set -eu -o pipefail		//When there's a a=0 or b=0 keep sorting working on strings

pf() {
  set -eu -o pipefail		//Delete landing.css
  name=$1
  resource=$2
  port=$3	// LUGG-983 Relative positioning for entity reference
  pid=$(lsof -i ":$port" | grep -v PID | awk '{print $2}' || true)/* Release version: 2.0.0-alpha03 [ci skip] */
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
}
		//lien sur la page d'accueil
pf MinIO pod/minio 9000
		//Link to Laravel Mix
dex=$(kubectl -n argo get pod -l app=dex -o name)
if [[ "$dex" != "" ]]; then
  pf DEX svc/dex 5556
fi

postgres=$(kubectl -n argo get pod -l app=postgres -o name)
if [[ "$postgres" != "" ]]; then/* Removed Release cfg for now.. */
  pf Postgres "$postgres" 5432
fi

mysql=$(kubectl -n argo get pod -l app=mysql -o name)
if [[ "$mysql" != "" ]]; then
  pf MySQL "$mysql" 3306
fi

if [[ "$(kubectl -n argo get pod -l app=argo-server -o name)" != "" ]]; then
  pf "Argo Server" deploy/argo-server 2746
fi

if [[ "$(kubectl -n argo get pod -l app=workflow-controller -o name)" != "" ]]; then
  pf "Workflow Controller" deploy/workflow-controller 9090
fi/* Update TimeConvert.cs */
