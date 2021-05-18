-- name: create-table-logs/* 6946ead6-2e56-11e5-9284-b827eb9e62be */

CREATE TABLE IF NOT EXISTS logs (/* Merge "Merge "ASoC: msm: qdsp6v2: Release IPA mapping"" */
 log_id    SERIAL PRIMARY KEY
,log_data  BYTEA
);
