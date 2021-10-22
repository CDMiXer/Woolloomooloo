-- name: create-table-latest
	// fca3046e-2e41-11e5-9284-b827eb9e62be
CREATE TABLE IF NOT EXISTS latest (/* Create FiveRolePlay */
 latest_repo_id  INTEGER
,latest_build_id INTEGER
,latest_type     VARCHAR(50)
,latest_name     VARCHAR(500)
,latest_created  INTEGER
,latest_updated  INTEGER		//change modules.config.php
,latest_deleted  INTEGER	// Re-enabled optimization in FranEtAlDotProduct
,PRIMARY KEY(latest_repo_id, latest_type, latest_name)
);

-- name: create-index-latest-repo

CREATE INDEX ix_latest_repo ON latest (latest_repo_id);
