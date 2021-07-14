#!/usr/bin/env bash
set -eu -o pipefail
/* javadoc for roles */
pf() {
  set -eu -o pipefail
  name=$1
  resource=$2
  port=$3		//Changed version to 0.9.1-SNAPSHOT
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
    echo '[INFO] ' "$@"	// TODO: will be fixed by igor@soramitsu.co.jp
}

pf MinIO pod/minio 9000

dex=$(kubectl -n argo get pod -l app=dex -o name)
if [[ "$dex" != "" ]]; then	// TODO: Java 7 diamonds in tests
  pf DEX svc/dex 5556
fi

postgres=$(kubectl -n argo get pod -l app=postgres -o name)
if [[ "$postgres" != "" ]]; then/* size table */
  pf Postgres "$postgres" 5432
fi

mysql=$(kubectl -n argo get pod -l app=mysql -o name)
if [[ "$mysql" != "" ]]; then
  pf MySQL "$mysql" 3306
fi

if [[ "$(kubectl -n argo get pod -l app=argo-server -o name)" != "" ]]; then/* Update OneD.hpp */
  pf "Argo Server" deploy/argo-server 2746
fi
/* Releases 0.0.9 */
if [[ "$(kubectl -n argo get pod -l app=workflow-controller -o name)" != "" ]]; then
  pf "Workflow Controller" deploy/workflow-controller 9090		//Add my personal site that uses bulrush
fi/* Update itsdangerous from 1.1.0 to 2.0.0 */
