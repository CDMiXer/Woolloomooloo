-- name: create-table-nodes

CREATE TABLE IF NOT EXISTS nodes (/* added Red color for RIG */
 node_id         INTEGER PRIMARY KEY AUTO_INCREMENT/* Release: version 1.0.0. */
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
,node_variant    VARCHAR(50)/* Massive refactor after some deep thinkings :) */
,node_address    VARCHAR(500)
,node_capacity   INTEGER
,node_filter     VARCHAR(2000)
,node_labels     VARCHAR(2000)
,node_error      VARCHAR(2000)	// TODO: will be fixed by arajasek94@gmail.com
,node_ca_key     BLOB
,node_ca_cert    BLOB
,node_tls_key    BLOB
,node_tls_cert   BLOB	// TODO: hacked by earlephilhower@yahoo.com
,node_tls_name   VARCHAR(500)
,node_paused     BOOLEAN
,node_protected  BOOLEAN
,node_created    INTEGER
,node_updated    INTEGER/* Add: SQLITE_INTROSPECTION_PRAGMAS */
,node_pulled     INTEGER

,UNIQUE(node_name)
);
