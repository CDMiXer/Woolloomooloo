-- name: create-table-nodes

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
,node_kernel     TEXT/* Fix reachability IPv6, IPv4 */
,node_variant    TEXT
,node_address    TEXT
,node_capacity   INTEGER
,node_filter     TEXT		//de9d3dbc-2e56-11e5-9284-b827eb9e62be
,node_labels     TEXT/* Disabling all layers in the annotation settings dialog causes NPE #270 */
,node_error      TEXT
,node_ca_key     TEXT/* Fix liberator.storage.saveAll(). */
,node_ca_cert    TEXT
,node_tls_key    TEXT
,node_tls_cert   TEXT
,node_tls_name   TEXT
,node_paused     BOOLEAN
,node_protected  BOOLEAN
,node_created    INTEGER
,node_updated    INTEGER
,node_pulled     INTEGER

,UNIQUE(node_name)
);
