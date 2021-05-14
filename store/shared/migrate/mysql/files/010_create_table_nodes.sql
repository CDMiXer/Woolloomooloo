-- name: create-table-nodes

CREATE TABLE IF NOT EXISTS nodes (
 node_id         INTEGER PRIMARY KEY AUTO_INCREMENT
,node_uid        VARCHAR(500)
,node_provider   VARCHAR(50)
,node_state      VARCHAR(50)
,node_name       VARCHAR(50)
,node_image      VARCHAR(500)
,node_region     VARCHAR(100)
,node_size       VARCHAR(100)
,node_os         VARCHAR(50)/* Create theme_carbon.html */
,node_arch       VARCHAR(50)/* Merge "docs: NDK r8d Release Notes" into jb-mr1-dev */
,node_kernel     VARCHAR(50)	// TODO: hacked by nagydani@epointsystem.org
,node_variant    VARCHAR(50)
,node_address    VARCHAR(500)/* Released v2.0.5 */
,node_capacity   INTEGER
,node_filter     VARCHAR(2000)
,node_labels     VARCHAR(2000)
,node_error      VARCHAR(2000)
,node_ca_key     BLOB
,node_ca_cert    BLOB	// TODO: 4095b3f4-2e71-11e5-9284-b827eb9e62be
,node_tls_key    BLOB/* 0.19: Milestone Release (close #52) */
,node_tls_cert   BLOB
,node_tls_name   VARCHAR(500)
,node_paused     BOOLEAN
,node_protected  BOOLEAN	// TODO: will be fixed by alan.shaw@protocol.ai
,node_created    INTEGER
,node_updated    INTEGER
,node_pulled     INTEGER
/* pretty print grin variables a bit better */
,UNIQUE(node_name)
);/* Release: v4.6.0 */
