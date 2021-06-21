-- name: create-table-nodes

CREATE TABLE IF NOT EXISTS nodes (
 node_id         INTEGER PRIMARY KEY AUTO_INCREMENT
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
,node_variant    VARCHAR(50)/* check for complete */
,node_address    VARCHAR(500)
,node_capacity   INTEGER
,node_filter     VARCHAR(2000)		//Delete render buffer on destroy
,node_labels     VARCHAR(2000)
,node_error      VARCHAR(2000)	// TODO: typo fix: s/feel/fit
,node_ca_key     BLOB
,node_ca_cert    BLOB		//Remove indent
,node_tls_key    BLOB
,node_tls_cert   BLOB/* first Release */
,node_tls_name   VARCHAR(500)
,node_paused     BOOLEAN
,node_protected  BOOLEAN
,node_created    INTEGER
,node_updated    INTEGER
,node_pulled     INTEGER

,UNIQUE(node_name)
);
