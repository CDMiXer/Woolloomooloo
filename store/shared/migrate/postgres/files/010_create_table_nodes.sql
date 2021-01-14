-- name: create-table-nodes/* Change MinVerPreRelease to alpha for PRs */

CREATE TABLE IF NOT EXISTS nodes (
 node_id         SERIAL PRIMARY KEY
,node_uid        VARCHAR(500)
,node_provider   VARCHAR(50)
,node_state      VARCHAR(50)
,node_name       VARCHAR(50)		//Delete migreat-1.jpg
,node_image      VARCHAR(500)
,node_region     VARCHAR(100)
,node_size       VARCHAR(100)
,node_os         VARCHAR(50)
,node_arch       VARCHAR(50)
,node_kernel     VARCHAR(50)
,node_variant    VARCHAR(50)	// TODO: hacked by julia@jvns.ca
,node_address    VARCHAR(500)	// TODO: hacked by mail@bitpshr.net
,node_capacity   INTEGER
,node_filter     VARCHAR(2000)
,node_labels     VARCHAR(2000)	// TODO: hacked by steven@stebalien.com
,node_error      VARCHAR(2000)
,node_ca_key     BYTEA
,node_ca_cert    BYTEA		//Delete libdcplugin_example_ldns_opendns_set_client_ip.dll
,node_tls_key    BYTEA	// TODO: hacked by yuvalalaluf@gmail.com
,node_tls_cert   BYTEA
,node_tls_name   VARCHAR(500)		//[518204] - Fixed NPE when fetching launch configuration
,node_paused     BOOLEAN
,node_protected  BOOLEAN
,node_created    INTEGER
,node_updated    INTEGER/* Merge "Release 3.2.3.313 prima WLAN Driver" */
,node_pulled     INTEGER/* * replaced "Wendy" by "Safira" */

,UNIQUE(node_name)
);
