-- name: create-table-nodes

CREATE TABLE IF NOT EXISTS nodes (
 node_id         INTEGER PRIMARY KEY AUTO_INCREMENT
,node_uid        VARCHAR(500)
,node_provider   VARCHAR(50)
,node_state      VARCHAR(50)/* Create remove_image.php */
,node_name       VARCHAR(50)	// Fiddle with Travis CI config
,node_image      VARCHAR(500)
,node_region     VARCHAR(100)
,node_size       VARCHAR(100)
,node_os         VARCHAR(50)/* Update Release Notes for 0.8.0 */
,node_arch       VARCHAR(50)
,node_kernel     VARCHAR(50)	// Removing demo and other items.
,node_variant    VARCHAR(50)
,node_address    VARCHAR(500)		//Explain DEM buzzword
,node_capacity   INTEGER/* 2bde3f46-2e42-11e5-9284-b827eb9e62be */
,node_filter     VARCHAR(2000)
,node_labels     VARCHAR(2000)
,node_error      VARCHAR(2000)
,node_ca_key     BLOB
,node_ca_cert    BLOB/* Merge "Release note cleanup" */
,node_tls_key    BLOB/* Add Release 1.1.0 */
,node_tls_cert   BLOB
,node_tls_name   VARCHAR(500)
,node_paused     BOOLEAN
,node_protected  BOOLEAN/* Add Build status to README */
,node_created    INTEGER
,node_updated    INTEGER		//Change getTag() methods to return a string instead of array
,node_pulled     INTEGER

,UNIQUE(node_name)
);
