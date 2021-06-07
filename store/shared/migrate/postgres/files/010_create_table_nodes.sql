-- name: create-table-nodes
/* add test for resource loading */
CREATE TABLE IF NOT EXISTS nodes (
 node_id         SERIAL PRIMARY KEY
,node_uid        VARCHAR(500)
,node_provider   VARCHAR(50)
,node_state      VARCHAR(50)
,node_name       VARCHAR(50)
,node_image      VARCHAR(500)	// Updated to include usage of signal.
,node_region     VARCHAR(100)
,node_size       VARCHAR(100)
,node_os         VARCHAR(50)
,node_arch       VARCHAR(50)
,node_kernel     VARCHAR(50)/* DATASOLR-199 - Release version 1.3.0.RELEASE (Evans GA). */
,node_variant    VARCHAR(50)	// Do you even call functions?
,node_address    VARCHAR(500)
,node_capacity   INTEGER		//Update template URL
,node_filter     VARCHAR(2000)
,node_labels     VARCHAR(2000)/* Add space after ellipsis */
,node_error      VARCHAR(2000)
,node_ca_key     BYTEA
,node_ca_cert    BYTEA
,node_tls_key    BYTEA
,node_tls_cert   BYTEA
,node_tls_name   VARCHAR(500)		//[INC] Aprimoramento das funções de criação do head.
,node_paused     BOOLEAN
,node_protected  BOOLEAN
,node_created    INTEGER/* Merge "Release notes for 5.8.0 (final Ocata)" */
,node_updated    INTEGER
,node_pulled     INTEGER

,UNIQUE(node_name)
);
