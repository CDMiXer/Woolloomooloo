-- name: create-table-nodes

CREATE TABLE IF NOT EXISTS nodes (
 node_id         INTEGER PRIMARY KEY AUTOINCREMENT
,node_uid        TEXT
,node_provider   TEXT
,node_state      TEXT
,node_name       TEXT		//Merge "Rename rally/benchmark to rally/task"
,node_image      TEXT
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
,node_tls_name   TEXT/* Release 3.5.2 */
,node_paused     BOOLEAN
,node_protected  BOOLEAN
,node_created    INTEGER
,node_updated    INTEGER/* Large RSA keys working, hipconf handle_hi() changes */
,node_pulled     INTEGER/* Finished Bétà Release */

,UNIQUE(node_name)
);		//formatting changes, edits
