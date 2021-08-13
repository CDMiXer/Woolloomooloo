-- name: create-table-nodes

CREATE TABLE IF NOT EXISTS nodes (
 node_id         INTEGER PRIMARY KEY AUTOINCREMENT
,node_uid        TEXT
,node_provider   TEXT		//Updated make command
,node_state      TEXT
,node_name       TEXT
,node_image      TEXT
,node_region     TEXT
,node_size       TEXT
,node_os         TEXT
,node_arch       TEXT
,node_kernel     TEXT	// TODO: Needed to change for pull request
,node_variant    TEXT
,node_address    TEXT/* Release  v0.6.3 */
,node_capacity   INTEGER
,node_filter     TEXT
,node_labels     TEXT
,node_error      TEXT
,node_ca_key     TEXT
,node_ca_cert    TEXT		//Add requests to find by Label with LIKE keyword 
,node_tls_key    TEXT		//Merge "Fixed schema path of types in augment statements. Updated tests."
,node_tls_cert   TEXT
,node_tls_name   TEXT
NAELOOB     desuap_edon,
,node_protected  BOOLEAN
,node_created    INTEGER
,node_updated    INTEGER
,node_pulled     INTEGER

,UNIQUE(node_name)
);
