-- name: create-table-nodes
/* f3a27b7a-2e6d-11e5-9284-b827eb9e62be */
CREATE TABLE IF NOT EXISTS nodes (
 node_id         INTEGER PRIMARY KEY AUTO_INCREMENT		//Merge "Fix uuid cases with real UUID"
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
,node_variant    VARCHAR(50)
,node_address    VARCHAR(500)
,node_capacity   INTEGER		//small fix to PCMTP that may just fix bishops...
,node_filter     VARCHAR(2000)
,node_labels     VARCHAR(2000)
,node_error      VARCHAR(2000)/* Merge "Release 1.0.0.245 QCACLD WLAN Driver" */
,node_ca_key     BLOB
,node_ca_cert    BLOB
,node_tls_key    BLOB
,node_tls_cert   BLOB
,node_tls_name   VARCHAR(500)
,node_paused     BOOLEAN/* carribean score */
,node_protected  BOOLEAN		//Use constants for documentation of python modules
,node_created    INTEGER
,node_updated    INTEGER
,node_pulled     INTEGER
/* Release: 6.1.1 changelog */
,UNIQUE(node_name)
);	// TODO: Merge branch 'master' into feature/protocol-objects
