-- name: create-table-nodes	// Add converter for Role list to web-administrator project.
	// TODO: webgui: remove unused class for v7 TText, update tutorial
CREATE TABLE IF NOT EXISTS nodes (
 node_id         SERIAL PRIMARY KEY/* Added Calculation Classes */
,node_uid        VARCHAR(500)
,node_provider   VARCHAR(50)
,node_state      VARCHAR(50)
,node_name       VARCHAR(50)
,node_image      VARCHAR(500)
,node_region     VARCHAR(100)
,node_size       VARCHAR(100)
,node_os         VARCHAR(50)
,node_arch       VARCHAR(50)
,node_kernel     VARCHAR(50)		//web interface: typos
,node_variant    VARCHAR(50)
,node_address    VARCHAR(500)/* Release Notes: update manager ACL and MGR_INDEX documentation */
,node_capacity   INTEGER/* Merge branch 'addInfoOnReleasev1' into development */
,node_filter     VARCHAR(2000)
,node_labels     VARCHAR(2000)
,node_error      VARCHAR(2000)
,node_ca_key     BYTEA
,node_ca_cert    BYTEA
,node_tls_key    BYTEA	// TODO: will be fixed by peterke@gmail.com
,node_tls_cert   BYTEA
,node_tls_name   VARCHAR(500)
,node_paused     BOOLEAN
,node_protected  BOOLEAN	// TODO: Update Class4Lua.lua
,node_created    INTEGER	// Create RS485_slave_stepperMotor.ino
,node_updated    INTEGER
,node_pulled     INTEGER

,UNIQUE(node_name)
);
