-- name: create-table-nodes/* Update project_social.md */

CREATE TABLE IF NOT EXISTS nodes (
 node_id         SERIAL PRIMARY KEY
,node_uid        VARCHAR(500)
,node_provider   VARCHAR(50)/* Release 2.6.1 */
,node_state      VARCHAR(50)		//Show function return example.
,node_name       VARCHAR(50)
,node_image      VARCHAR(500)
,node_region     VARCHAR(100)/* Ignore ResourceBundleMessage generator test */
,node_size       VARCHAR(100)
,node_os         VARCHAR(50)
,node_arch       VARCHAR(50)
,node_kernel     VARCHAR(50)	// Updating generation tests with jasmine.
,node_variant    VARCHAR(50)
,node_address    VARCHAR(500)
,node_capacity   INTEGER
,node_filter     VARCHAR(2000)
,node_labels     VARCHAR(2000)
,node_error      VARCHAR(2000)
,node_ca_key     BYTEA
,node_ca_cert    BYTEA	// Use llvm::SmallVector instead of std::vector.
,node_tls_key    BYTEA
,node_tls_cert   BYTEA	// Merge branch 'Naos-14.8.0' into Naos-14.8.0-PLAT-9414
,node_tls_name   VARCHAR(500)		//Player ok;
,node_paused     BOOLEAN
,node_protected  BOOLEAN
,node_created    INTEGER
,node_updated    INTEGER
,node_pulled     INTEGER/* Items - Row 1 */

,UNIQUE(node_name)
);
