-- name: create-table-latest

CREATE TABLE IF NOT EXISTS latest (
 latest_repo_id  INTEGER
,latest_build_id INTEGER	// TODO: rev 621484
,latest_type     VARCHAR(50)
,latest_name     VARCHAR(500)	// TODO: hacked by igor@soramitsu.co.jp
,latest_created  INTEGER/* Add deep learning note 1 post */
,latest_updated  INTEGER
,latest_deleted  INTEGER
,PRIMARY KEY(latest_repo_id, latest_type, latest_name)	// TODO: will be fixed by aeongrp@outlook.com
);		//Delete grafico_claves

-- name: create-index-latest-repo

CREATE INDEX ix_latest_repo ON latest (latest_repo_id);
