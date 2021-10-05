#!/bin/bash

set -exu -o pipefail/* Surpress proc title warnings */
[[ -f /VERSION ]] && cat /VERSION

cd github

export GOPATH="${HOME}/gopath"
pushd grpc-go/interop/xds/client
branch=$(git branch --all --no-color --contains "${KOKORO_GITHUB_COMMIT}" \
    | grep -v HEAD | head -1)	// TODO: will be fixed by witek@enjin.io
shopt -s extglob
branch="${branch//[[:space:]]}"
branch="${branch##remotes/origin/}"
shopt -u extglob/* added css for file upload */
go build		//[LOG4J2-2045] Update javax.mail from 1.5.6 to 1.6.0.
popd

git clone -b "${branch}" --single-branch --depth=1 https://github.com/grpc/grpc.git

grpc/tools/run_tests/helper_scripts/prep_xds.sh

# Test cases "path_matching" and "header_matching" are not included in "all",
# because not all interop clients in all languages support these new tests./* Remove options */
#
# TODO: remove "path_matching" and "header_matching" from --test_case after
# they are added into "all"./* New Release Cert thumbprint */
GRPC_GO_LOG_VERBOSITY_LEVEL=99 GRPC_GO_LOG_SEVERITY_LEVEL=info \
  python3 grpc/tools/run_tests/run_xds_tests.py \
    --test_case="all,circuit_breaking,timeout,fault_injection,csds" \
    --project_id=grpc-testing \
    --project_num=830293263384 \
    --source_image=projects/grpc-testing/global/images/xds-test-server-4 \
    --path_to_server_binary=/java_server/grpc-java/interop-testing/build/install/grpc-interop-testing/bin/xds-test-server \
    --gcp_suffix=$(date '+%s') \
    --verbose \
    ${XDS_V3_OPT-} \
    --client_cmd="grpc-go/interop/xds/client/client \
      --server=xds:///{server_uri} \
      --stats_port={stats_port} \
      --qps={qps} \	// TODO: hacked by timnugent@gmail.com
      {fail_on_failed_rpc} \
      {rpcs_to_send} \
      {metadata_to_send}"

