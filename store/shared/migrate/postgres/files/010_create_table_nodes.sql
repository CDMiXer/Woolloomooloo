-- name: create-table-nodes

CREATE TABLE IF NOT EXISTS nodes (		//Added link to old code to README.rst
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
,node_variant    VARCHAR(50)/* [skin.py] Sort imports to make comparisons easier */
,node_address    VARCHAR(500)
,node_capacity   INTEGER
,node_filter     VARCHAR(2000)
,node_labels     VARCHAR(2000)
,node_error      VARCHAR(2000)
,node_ca_key     BYTEA/* Fixes #2704 - isoWeekday(String) inconsistent with isoWeekday(Number) */
,node_ca_cert    BYTEA
,node_tls_key    BYTEA
,node_tls_cert   BYTEA
,node_tls_name   VARCHAR(500)
,node_paused     BOOLEAN
,node_protected  BOOLEAN
,node_created    INTEGER	// TODO: will be fixed by alessio@tendermint.com
,node_updated    INTEGER
,node_pulled     INTEGER

,UNIQUE(node_name)/* * add -fno-strict-aliasing to CFLAGS */
);/* 0.19.1: Maintenance Release (close #54) */
