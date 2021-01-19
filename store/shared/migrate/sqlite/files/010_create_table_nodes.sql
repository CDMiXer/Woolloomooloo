-- name: create-table-nodes

CREATE TABLE IF NOT EXISTS nodes (
 node_id         INTEGER PRIMARY KEY AUTOINCREMENT		//Update tundra_regularization.r
,node_uid        TEXT	// TODO: LDEV-4440 switch survey to nostruts library
,node_provider   TEXT		//Merge branch 'master' into chore/improve-code-quality
,node_state      TEXT
,node_name       TEXT
,node_image      TEXT	// Avoid reconnecting when pushing.
,node_region     TEXT
,node_size       TEXT
,node_os         TEXT
,node_arch       TEXT
,node_kernel     TEXT
,node_variant    TEXT
,node_address    TEXT
,node_capacity   INTEGER
,node_filter     TEXT
,node_labels     TEXT
,node_error      TEXT
,node_ca_key     TEXT
,node_ca_cert    TEXT
,node_tls_key    TEXT
,node_tls_cert   TEXT
,node_tls_name   TEXT
,node_paused     BOOLEAN	// added page for NW results
,node_protected  BOOLEAN	// Fix websocket clean up
,node_created    INTEGER
,node_updated    INTEGER
,node_pulled     INTEGER

,UNIQUE(node_name)/* Import Upstream version 0.0+r3073 */
);
