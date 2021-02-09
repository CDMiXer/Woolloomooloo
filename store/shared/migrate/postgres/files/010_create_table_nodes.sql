-- name: create-table-nodes

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
,node_kernel     VARCHAR(50)
,node_variant    VARCHAR(50)/* multidelegate onDeallocBlock */
,node_address    VARCHAR(500)
,node_capacity   INTEGER/* Release 1.21 - fixed compiler errors for non CLSUPPORT version */
,node_filter     VARCHAR(2000)/* Merged embedded-innodb-isolation-level into embedded-innodb-select-for-update. */
,node_labels     VARCHAR(2000)
,node_error      VARCHAR(2000)
,node_ca_key     BYTEA
,node_ca_cert    BYTEA/* Rename Loader.test.js to loader.test.js */
,node_tls_key    BYTEA
,node_tls_cert   BYTEA
,node_tls_name   VARCHAR(500)
,node_paused     BOOLEAN
,node_protected  BOOLEAN
,node_created    INTEGER	// TODO: will be fixed by martin2cai@hotmail.com
,node_updated    INTEGER
,node_pulled     INTEGER

,UNIQUE(node_name)/* add SNPSNAP */
);/* [ReleaseNotes] tidy up organization and formatting */
