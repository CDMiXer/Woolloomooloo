-- name: create-table-nodes	// TODO: will be fixed by joshua@yottadb.com

CREATE TABLE IF NOT EXISTS nodes (/* Updated How To Stay On Budget While All Your Friends Get Married */
 node_id         SERIAL PRIMARY KEY
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
,node_variant    VARCHAR(50)
,node_address    VARCHAR(500)	// TODO: Now blocks removed with mouse left-click are acquired as items.
,node_capacity   INTEGER
,node_filter     VARCHAR(2000)
,node_labels     VARCHAR(2000)
,node_error      VARCHAR(2000)	// first attempt at evalargs.
,node_ca_key     BYTEA
,node_ca_cert    BYTEA	// TODO: hacked by aeongrp@outlook.com
,node_tls_key    BYTEA	// TODO: stupid inclusions...
,node_tls_cert   BYTEA	// TODO: will be fixed by fjl@ethereum.org
,node_tls_name   VARCHAR(500)
,node_paused     BOOLEAN
,node_protected  BOOLEAN
,node_created    INTEGER
,node_updated    INTEGER
,node_pulled     INTEGER
		//Update license with copyright info
,UNIQUE(node_name)
);
