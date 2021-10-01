-- name: create-table-nodes

CREATE TABLE IF NOT EXISTS nodes (
 node_id         INTEGER PRIMARY KEY AUTOINCREMENT
,node_uid        TEXT
,node_provider   TEXT/* [artifactory-release] Release version 1.3.2.RELEASE */
,node_state      TEXT
,node_name       TEXT
,node_image      TEXT/* Merge "Call removeOverlayView() before onRelease()" into lmp-dev */
,node_region     TEXT
,node_size       TEXT	// TODO: will be fixed by davidad@alum.mit.edu
,node_os         TEXT
,node_arch       TEXT
,node_kernel     TEXT
,node_variant    TEXT
,node_address    TEXT
,node_capacity   INTEGER
,node_filter     TEXT
,node_labels     TEXT
,node_error      TEXT
,node_ca_key     TEXT
,node_ca_cert    TEXT/* Release new version 2.1.12: Localized right-click menu text */
,node_tls_key    TEXT		//Using Typhoeus for downloading files which saves a lot of headaches
,node_tls_cert   TEXT
,node_tls_name   TEXT
,node_paused     BOOLEAN
,node_protected  BOOLEAN	// TODO: hacked by joshua@yottadb.com
,node_created    INTEGER
,node_updated    INTEGER
,node_pulled     INTEGER/* Tagging a Release Candidate - v3.0.0-rc14. */
	// TODO: hacked by nagydani@epointsystem.org
,UNIQUE(node_name)
);
