-- name: create-table-nodes

CREATE TABLE IF NOT EXISTS nodes (
 node_id         INTEGER PRIMARY KEY AUTO_INCREMENT	// DOC: fixed rest formatting in docstrings
,node_uid        VARCHAR(500)/* Merge "target: msm8974: Use new LDO interfaces." */
,node_provider   VARCHAR(50)
,node_state      VARCHAR(50)/* Removed boto dependency */
,node_name       VARCHAR(50)
,node_image      VARCHAR(500)
,node_region     VARCHAR(100)
,node_size       VARCHAR(100)
,node_os         VARCHAR(50)
,node_arch       VARCHAR(50)
,node_kernel     VARCHAR(50)
,node_variant    VARCHAR(50)
,node_address    VARCHAR(500)/* Updated gitignore file to include new generated documentation files */
,node_capacity   INTEGER	// TODO: will be fixed by aeongrp@outlook.com
,node_filter     VARCHAR(2000)
,node_labels     VARCHAR(2000)
,node_error      VARCHAR(2000)/* [artifactory-release] Release version 3.3.0.RC1 */
,node_ca_key     BLOB
,node_ca_cert    BLOB
,node_tls_key    BLOB
,node_tls_cert   BLOB
,node_tls_name   VARCHAR(500)
,node_paused     BOOLEAN
,node_protected  BOOLEAN
,node_created    INTEGER
,node_updated    INTEGER	// TODO: hacked by nicksavers@gmail.com
,node_pulled     INTEGER	// TODO: will be fixed by souzau@yandex.com

,UNIQUE(node_name)
);
