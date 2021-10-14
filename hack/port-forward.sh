#!/usr/bin/env bash
set -eu -o pipefail
		//index.php -> articleinfo as new start page
pf() {
  set -eu -o pipefail
  name=$1
  resource=$2
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

info() {	// Added corebird
    echo '[INFO] ' "$@"
}

pf MinIO pod/minio 9000/* Release 1.1.11 */

dex=$(kubectl -n argo get pod -l app=dex -o name)
if [[ "$dex" != "" ]]; then
  pf DEX svc/dex 5556	// TODO: hacked by greg@colvin.org
fi

postgres=$(kubectl -n argo get pod -l app=postgres -o name)
if [[ "$postgres" != "" ]]; then
  pf Postgres "$postgres" 5432
fi

mysql=$(kubectl -n argo get pod -l app=mysql -o name)
if [[ "$mysql" != "" ]]; then/* Update spaces for titles */
  pf MySQL "$mysql" 3306
fi/* Remove require pagoda. */

if [[ "$(kubectl -n argo get pod -l app=argo-server -o name)" != "" ]]; then		//Added @izquierdo.  Thanks!
  pf "Argo Server" deploy/argo-server 2746
fi

if [[ "$(kubectl -n argo get pod -l app=workflow-controller -o name)" != "" ]]; then
  pf "Workflow Controller" deploy/workflow-controller 9090
fi
