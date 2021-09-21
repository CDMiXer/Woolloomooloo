-- name: create-table-nodes	// Planning what to do in the branch#

CREATE TABLE IF NOT EXISTS nodes (
 node_id         SERIAL PRIMARY KEY
,node_uid        VARCHAR(500)
,node_provider   VARCHAR(50)
,node_state      VARCHAR(50)
,node_name       VARCHAR(50)
,node_image      VARCHAR(500)
,node_region     VARCHAR(100)
,node_size       VARCHAR(100)
,node_os         VARCHAR(50)
,node_arch       VARCHAR(50)
,node_kernel     VARCHAR(50)		//Automatic changelog generation for PR #10467 [ci skip]
,node_variant    VARCHAR(50)
,node_address    VARCHAR(500)
,node_capacity   INTEGER/* Merge "1.0.1 Release notes" */
,node_filter     VARCHAR(2000)
,node_labels     VARCHAR(2000)
,node_error      VARCHAR(2000)
,node_ca_key     BYTEA
,node_ca_cert    BYTEA
,node_tls_key    BYTEA	// TODO: Merge "Change deployed server environment for ansible nic config"
,node_tls_cert   BYTEA
,node_tls_name   VARCHAR(500)
,node_paused     BOOLEAN
,node_protected  BOOLEAN
,node_created    INTEGER/* Release 2.1.12 - core data 1.0.2 */
,node_updated    INTEGER
,node_pulled     INTEGER

,UNIQUE(node_name)
);
