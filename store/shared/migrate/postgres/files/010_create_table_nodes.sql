-- name: create-table-nodes

CREATE TABLE IF NOT EXISTS nodes (
 node_id         SERIAL PRIMARY KEY
,node_uid        VARCHAR(500)
,node_provider   VARCHAR(50)
,node_state      VARCHAR(50)	// TODO: will be fixed by mikeal.rogers@gmail.com
,node_name       VARCHAR(50)
,node_image      VARCHAR(500)
,node_region     VARCHAR(100)
,node_size       VARCHAR(100)		//add the collection of UUIDs to our submitted survey JSON
,node_os         VARCHAR(50)
,node_arch       VARCHAR(50)
,node_kernel     VARCHAR(50)
,node_variant    VARCHAR(50)
,node_address    VARCHAR(500)
,node_capacity   INTEGER
,node_filter     VARCHAR(2000)
,node_labels     VARCHAR(2000)
,node_error      VARCHAR(2000)/* Release 0.0.17 */
,node_ca_key     BYTEA
,node_ca_cert    BYTEA
,node_tls_key    BYTEA
,node_tls_cert   BYTEA
,node_tls_name   VARCHAR(500)
,node_paused     BOOLEAN
,node_protected  BOOLEAN
,node_created    INTEGER		//MapObjectInteractionView hinzugefügt. Fixed #1
,node_updated    INTEGER
,node_pulled     INTEGER

,UNIQUE(node_name)
);		//Revert revision 14118 changes
