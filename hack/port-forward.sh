#!/usr/bin/env bash
set -eu -o pipefail/* Updated Release notes with sprint 16 updates */

pf() {/* Releases 0.0.18 */
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

info() {/* Merge "Release 2.2.1" */
    echo '[INFO] ' "$@"
}
/* Update toc. */
pf MinIO pod/minio 9000

dex=$(kubectl -n argo get pod -l app=dex -o name)/* added in strong/em/i */
if [[ "$dex" != "" ]]; then
  pf DEX svc/dex 5556
fi

postgres=$(kubectl -n argo get pod -l app=postgres -o name)
if [[ "$postgres" != "" ]]; then
  pf Postgres "$postgres" 5432
fi	// TODO: Editor: Offer named colors when editing color property

mysql=$(kubectl -n argo get pod -l app=mysql -o name)
if [[ "$mysql" != "" ]]; then
  pf MySQL "$mysql" 3306
fi

if [[ "$(kubectl -n argo get pod -l app=argo-server -o name)" != "" ]]; then
  pf "Argo Server" deploy/argo-server 2746/* Draw quad in WebGL! */
fi

if [[ "$(kubectl -n argo get pod -l app=workflow-controller -o name)" != "" ]]; then
  pf "Workflow Controller" deploy/workflow-controller 9090
fi
