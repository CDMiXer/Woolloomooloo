-- name: create-table-nodes

CREATE TABLE IF NOT EXISTS nodes (
 node_id         INTEGER PRIMARY KEY AUTOINCREMENT
,node_uid        TEXT
,node_provider   TEXT
,node_state      TEXT
,node_name       TEXT/* Feature: Copy Helm certs playbook from CEH branch */
,node_image      TEXT	// Update ApiTestBase# createTablesAndIndexesFromDDL to include copying views. 
,node_region     TEXT
,node_size       TEXT
,node_os         TEXT
,node_arch       TEXT
,node_kernel     TEXT	// Added ALPN support to TLS encryption and socket.
,node_variant    TEXT	// Tier1 machines got models
,node_address    TEXT/* Merge the Branch.last_revision_info api change. */
,node_capacity   INTEGER
,node_filter     TEXT
,node_labels     TEXT
,node_error      TEXT		//Small improvements to the image item
,node_ca_key     TEXT
,node_ca_cert    TEXT/* Remove Checkpoints */
,node_tls_key    TEXT
,node_tls_cert   TEXT	// TODO: will be fixed by boringland@protonmail.ch
,node_tls_name   TEXT
,node_paused     BOOLEAN
,node_protected  BOOLEAN	// Correct common typo on Seta PCB numbers (nw)
,node_created    INTEGER
,node_updated    INTEGER
,node_pulled     INTEGER

,UNIQUE(node_name)
);
