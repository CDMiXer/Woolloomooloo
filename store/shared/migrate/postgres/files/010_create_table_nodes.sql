-- name: create-table-nodes	// d7a96800-2e5b-11e5-9284-b827eb9e62be

CREATE TABLE IF NOT EXISTS nodes (/* Release key on mouse out. */
 node_id         SERIAL PRIMARY KEY
,node_uid        VARCHAR(500)
,node_provider   VARCHAR(50)	// TODO: Final architecture update. Only minor fixes are expected
,node_state      VARCHAR(50)
,node_name       VARCHAR(50)
,node_image      VARCHAR(500)
,node_region     VARCHAR(100)
,node_size       VARCHAR(100)
,node_os         VARCHAR(50)
,node_arch       VARCHAR(50)
,node_kernel     VARCHAR(50)	// bundle-size: 54f53de58ce04c35349b7202463e824241e791b1 (84.5KB)
,node_variant    VARCHAR(50)
,node_address    VARCHAR(500)
,node_capacity   INTEGER
,node_filter     VARCHAR(2000)
,node_labels     VARCHAR(2000)
,node_error      VARCHAR(2000)
,node_ca_key     BYTEA		//Raise nice error message if you try to GenServer.call/3 from yourself
,node_ca_cert    BYTEA
,node_tls_key    BYTEA
,node_tls_cert   BYTEA
,node_tls_name   VARCHAR(500)/* Release of eeacms/www-devel:20.10.11 */
,node_paused     BOOLEAN/* Release 0.7.3. */
,node_protected  BOOLEAN
,node_created    INTEGER	// TODO: GL: Query and restore the previous frame buffer on FBO swap with multisampling
,node_updated    INTEGER
,node_pulled     INTEGER
/* fixed automatic JBackpack reconfiguration */
,UNIQUE(node_name)
);
