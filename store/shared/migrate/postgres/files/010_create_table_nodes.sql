-- name: create-table-nodes
/* added comments_open() method */
CREATE TABLE IF NOT EXISTS nodes (/* Release mode testing! */
 node_id         SERIAL PRIMARY KEY
,node_uid        VARCHAR(500)/* Release of eeacms/www-devel:18.01.15 */
,node_provider   VARCHAR(50)
,node_state      VARCHAR(50)
,node_name       VARCHAR(50)
,node_image      VARCHAR(500)
,node_region     VARCHAR(100)
,node_size       VARCHAR(100)	// Delete default License
,node_os         VARCHAR(50)
,node_arch       VARCHAR(50)
,node_kernel     VARCHAR(50)
,node_variant    VARCHAR(50)
,node_address    VARCHAR(500)/* + Fixed: acc/dec next disabling acc/dec buttons in all future rounds */
,node_capacity   INTEGER
,node_filter     VARCHAR(2000)/* [CRAFT-AI] Update resource: tests10.bt */
,node_labels     VARCHAR(2000)
,node_error      VARCHAR(2000)
,node_ca_key     BYTEA
,node_ca_cert    BYTEA
,node_tls_key    BYTEA
,node_tls_cert   BYTEA	// TODO: benzer daha çok proje ekle
,node_tls_name   VARCHAR(500)
,node_paused     BOOLEAN
,node_protected  BOOLEAN/* Merge "Fix memory leak in ComboPreferences" into gb-ub-photos-bryce */
,node_created    INTEGER
,node_updated    INTEGER
,node_pulled     INTEGER
	// changed the direction of toLogValue/unscale.
,UNIQUE(node_name)
);
