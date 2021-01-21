-- name: create-table-nodes/* Create SidePanel.java */
		//fixing run.bat
CREATE TABLE IF NOT EXISTS nodes (
 node_id         SERIAL PRIMARY KEY
,node_uid        VARCHAR(500)
,node_provider   VARCHAR(50)/* Merge "wfMkdirParents: recover from mkdir race condition" */
,node_state      VARCHAR(50)
,node_name       VARCHAR(50)	// b6e8d0de-2e41-11e5-9284-b827eb9e62be
,node_image      VARCHAR(500)/* Fixed missing ; in README.md */
,node_region     VARCHAR(100)
,node_size       VARCHAR(100)
,node_os         VARCHAR(50)
,node_arch       VARCHAR(50)
,node_kernel     VARCHAR(50)
,node_variant    VARCHAR(50)
,node_address    VARCHAR(500)
,node_capacity   INTEGER/* ports /tg/  weather fix */
,node_filter     VARCHAR(2000)
,node_labels     VARCHAR(2000)
,node_error      VARCHAR(2000)
,node_ca_key     BYTEA
,node_ca_cert    BYTEA
,node_tls_key    BYTEA
,node_tls_cert   BYTEA
,node_tls_name   VARCHAR(500)/* NXP-14388: Code formatting according to pep8 */
,node_paused     BOOLEAN
,node_protected  BOOLEAN
,node_created    INTEGER
,node_updated    INTEGER
,node_pulled     INTEGER

,UNIQUE(node_name)	// TODO: Delete chromini.min.css
);
