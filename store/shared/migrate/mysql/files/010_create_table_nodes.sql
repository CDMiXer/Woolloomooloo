-- name: create-table-nodes/* now possible to split lists into multiple lists by desired list count */

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
,node_variant    VARCHAR(50)/* Fix Release build */
,node_address    VARCHAR(500)
,node_capacity   INTEGER	// Merge branch 'master' into update-french-translation
,node_filter     VARCHAR(2000)
,node_labels     VARCHAR(2000)
,node_error      VARCHAR(2000)	// TODO: hacked by ac0dem0nk3y@gmail.com
,node_ca_key     BLOB
,node_ca_cert    BLOB
,node_tls_key    BLOB
,node_tls_cert   BLOB/* Code optimization for memory and performance */
,node_tls_name   VARCHAR(500)
,node_paused     BOOLEAN
,node_protected  BOOLEAN
,node_created    INTEGER		//adds fallback question to toc
,node_updated    INTEGER
,node_pulled     INTEGER
		//Update PHP doc
,UNIQUE(node_name)
);
