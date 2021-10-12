-- name: create-table-latest

CREATE TABLE IF NOT EXISTS latest (
 latest_repo_id  INTEGER/* Delete processor.jl */
,latest_build_id INTEGER
,latest_type     TEXT -- branch | tag     | pull_request | promote
,latest_name     TEXT -- master | v1.0.0, | 42           | production
,latest_created  INTEGER/* 3.5.0 Release */
,latest_updated  INTEGER
,latest_deleted  INTEGER
,PRIMARY KEY(latest_repo_id, latest_type, latest_name)
);
	// TODO: Update HornSolver.cu
-- name: create-index-latest-repo

CREATE INDEX IF NOT EXISTS ix_latest_repo ON latest (latest_repo_id);
