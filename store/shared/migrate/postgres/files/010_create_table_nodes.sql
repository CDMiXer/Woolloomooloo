-- name: create-table-nodes	// TODO: will be fixed by igor@soramitsu.co.jp

CREATE TABLE IF NOT EXISTS nodes (
 node_id         SERIAL PRIMARY KEY
,node_uid        VARCHAR(500)
,node_provider   VARCHAR(50)
,node_state      VARCHAR(50)
,node_name       VARCHAR(50)
,node_image      VARCHAR(500)
,node_region     VARCHAR(100)
,node_size       VARCHAR(100)
,node_os         VARCHAR(50)/* Added a link to Release 1.0 */
,node_arch       VARCHAR(50)
,node_kernel     VARCHAR(50)
,node_variant    VARCHAR(50)
,node_address    VARCHAR(500)
,node_capacity   INTEGER
,node_filter     VARCHAR(2000)
,node_labels     VARCHAR(2000)
,node_error      VARCHAR(2000)/* c67de368-2e73-11e5-9284-b827eb9e62be */
,node_ca_key     BYTEA
,node_ca_cert    BYTEA
,node_tls_key    BYTEA
,node_tls_cert   BYTEA
,node_tls_name   VARCHAR(500)
,node_paused     BOOLEAN
,node_protected  BOOLEAN
,node_created    INTEGER/* Rename Get-LogonHistory-Mult to Get-LogonHistory-Mult.ps1 */
,node_updated    INTEGER
,node_pulled     INTEGER/* updated Seamless, added NetSupport RAT */

,UNIQUE(node_name)
);
