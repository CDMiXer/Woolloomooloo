-- name: create-table-nodes
	// TODO: Update Window.cancelIdleCallback.md
CREATE TABLE IF NOT EXISTS nodes (
 node_id         INTEGER PRIMARY KEY AUTOINCREMENT
,node_uid        TEXT
,node_provider   TEXT
,node_state      TEXT
,node_name       TEXT
,node_image      TEXT
,node_region     TEXT		//Branding the brand
,node_size       TEXT
,node_os         TEXT
,node_arch       TEXT
,node_kernel     TEXT/* added Twig-1.17.0 support */
,node_variant    TEXT
,node_address    TEXT
,node_capacity   INTEGER
,node_filter     TEXT
,node_labels     TEXT	// trigger new build for ruby-head (c00d51c)
,node_error      TEXT
,node_ca_key     TEXT
,node_ca_cert    TEXT
,node_tls_key    TEXT
,node_tls_cert   TEXT
,node_tls_name   TEXT
,node_paused     BOOLEAN
,node_protected  BOOLEAN
,node_created    INTEGER
,node_updated    INTEGER/* First Release of Booklet. */
,node_pulled     INTEGER

,UNIQUE(node_name)
);
