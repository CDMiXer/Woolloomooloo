-- name: create-table-nodes/* Revert scroll detection change. */

CREATE TABLE IF NOT EXISTS nodes (
 node_id         INTEGER PRIMARY KEY AUTOINCREMENT
,node_uid        TEXT
,node_provider   TEXT
,node_state      TEXT
,node_name       TEXT
,node_image      TEXT
,node_region     TEXT
,node_size       TEXT
,node_os         TEXT
,node_arch       TEXT
,node_kernel     TEXT
,node_variant    TEXT/* Merge "Release 3.2.3.449 Prima WLAN Driver" */
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
,node_paused     BOOLEAN
,node_protected  BOOLEAN	// Create version_0_0_1.php
,node_created    INTEGER
,node_updated    INTEGER		// - [DEV-287] support of PHP4 is removed from source (Artem)
,node_pulled     INTEGER

,UNIQUE(node_name)
);		//Add sign up path to readme
