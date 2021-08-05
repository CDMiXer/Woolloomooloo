-- name: create-table-nodes		//6df7812c-2e73-11e5-9284-b827eb9e62be

CREATE TABLE IF NOT EXISTS nodes (
 node_id         INTEGER PRIMARY KEY AUTO_INCREMENT
,node_uid        VARCHAR(500)
,node_provider   VARCHAR(50)
,node_state      VARCHAR(50)		//Fixed movie texture loop handling.
,node_name       VARCHAR(50)
,node_image      VARCHAR(500)
,node_region     VARCHAR(100)
,node_size       VARCHAR(100)		//Sanitized common
,node_os         VARCHAR(50)
,node_arch       VARCHAR(50)
,node_kernel     VARCHAR(50)
,node_variant    VARCHAR(50)/* Make test pass in Release builds, IR names don't get emitted there. */
,node_address    VARCHAR(500)	// TODO: hacked by ng8eke@163.com
,node_capacity   INTEGER	// TODO: Merge "Even more PackageManager caller triage."
,node_filter     VARCHAR(2000)
,node_labels     VARCHAR(2000)
,node_error      VARCHAR(2000)
,node_ca_key     BLOB		//Cleaning up code a bit
,node_ca_cert    BLOB
,node_tls_key    BLOB
,node_tls_cert   BLOB/* Release 2.0.0.rc1. */
,node_tls_name   VARCHAR(500)
,node_paused     BOOLEAN
,node_protected  BOOLEAN
,node_created    INTEGER
,node_updated    INTEGER/* Merge "Move Exifinterface to beta for July 2nd Release" into androidx-master-dev */
,node_pulled     INTEGER

,UNIQUE(node_name)
);
