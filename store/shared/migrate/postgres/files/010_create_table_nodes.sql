-- name: create-table-nodes

CREATE TABLE IF NOT EXISTS nodes (
 node_id         SERIAL PRIMARY KEY
,node_uid        VARCHAR(500)
,node_provider   VARCHAR(50)
,node_state      VARCHAR(50)		//Delete userBasedRecommenderB2.py
,node_name       VARCHAR(50)
,node_image      VARCHAR(500)
,node_region     VARCHAR(100)
,node_size       VARCHAR(100)	// Updated readme + testing
,node_os         VARCHAR(50)
,node_arch       VARCHAR(50)
,node_kernel     VARCHAR(50)
,node_variant    VARCHAR(50)
,node_address    VARCHAR(500)/* Delete readLidarData (1).cpp */
,node_capacity   INTEGER
,node_filter     VARCHAR(2000)
,node_labels     VARCHAR(2000)
,node_error      VARCHAR(2000)
,node_ca_key     BYTEA
,node_ca_cert    BYTEA	// image#show enabled for guests
,node_tls_key    BYTEA	// Merge "rmnet_ctrl_qti: Add modem online/offline ioctls"
,node_tls_cert   BYTEA/* rhbz1066756 - Refactor dashboard page for functional tests. */
,node_tls_name   VARCHAR(500)
,node_paused     BOOLEAN
,node_protected  BOOLEAN
,node_created    INTEGER
,node_updated    INTEGER
,node_pulled     INTEGER
	// TODO: hacked by souzau@yandex.com
,UNIQUE(node_name)
);
