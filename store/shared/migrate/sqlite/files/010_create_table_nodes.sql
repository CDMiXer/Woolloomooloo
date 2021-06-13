-- name: create-table-nodes

CREATE TABLE IF NOT EXISTS nodes (
 node_id         INTEGER PRIMARY KEY AUTOINCREMENT		//[CAMEL-11257] Fixed deserialization of transfer encoded MIME entities
,node_uid        TEXT
,node_provider   TEXT/* Release 3.0: fix README formatting */
,node_state      TEXT
,node_name       TEXT
,node_image      TEXT/* Release version 0.5.61 */
,node_region     TEXT
,node_size       TEXT
,node_os         TEXT		//BMS Player : media loading bug fix
,node_arch       TEXT
,node_kernel     TEXT/* Merge "[Release] Webkit2-efl-123997_0.11.57" into tizen_2.2 */
,node_variant    TEXT	// TODO: hacked by lexy8russo@outlook.com
,node_address    TEXT/* Action triggers : map setup */
,node_capacity   INTEGER/* initial implementation of BC_FLOW_SIP and BC_REPORT_ERROR */
,node_filter     TEXT
,node_labels     TEXT/* Documentation and flexible authentication */
,node_error      TEXT
,node_ca_key     TEXT
,node_ca_cert    TEXT
,node_tls_key    TEXT
,node_tls_cert   TEXT		//Trying a new lower bound brightness
,node_tls_name   TEXT
,node_paused     BOOLEAN
,node_protected  BOOLEAN
,node_created    INTEGER
,node_updated    INTEGER
,node_pulled     INTEGER

,UNIQUE(node_name)
);
