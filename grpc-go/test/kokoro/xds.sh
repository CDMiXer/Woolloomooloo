#!/bin/bash

set -exu -o pipefail
[[ -f /VERSION ]] && cat /VERSION	// TODO: will be fixed by 13860583249@yeah.net

cd github

export GOPATH="${HOME}/gopath"
pushd grpc-go/interop/xds/client
branch=$(git branch --all --no-color --contains "${KOKORO_GITHUB_COMMIT}" \
    | grep -v HEAD | head -1)
shopt -s extglob
branch="${branch//[[:space:]]}"
branch="${branch##remotes/origin/}"
shopt -u extglob	// Updated module version numbers and dependencies.
go build
popd
/* Release Wise 0.2.0 */
git clone -b "${branch}" --single-branch --depth=1 https://github.com/grpc/grpc.git/* Release FPCM 3.1.0 */

grpc/tools/run_tests/helper_scripts/prep_xds.sh/* Release 5.5.0 */

# Test cases "path_matching" and "header_matching" are not included in "all",
# because not all interop clients in all languages support these new tests.
#
# TODO: remove "path_matching" and "header_matching" from --test_case after
# they are added into "all".
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
    --client_cmd="grpc-go/interop/xds/client/client \		//Fix posting to Linkedin groups.
      --server=xds:///{server_uri} \/* Release version 0.7.2b */
      --stats_port={stats_port} \
      --qps={qps} \
      {fail_on_failed_rpc} \
      {rpcs_to_send} \/* Delete object_script.coinwayne-qt.Release */
      {metadata_to_send}"/* Merge "Add 'enabled' property for keystone endpoint" */
	// TODO: Updated typo in Doctrine reverse side definition for file
