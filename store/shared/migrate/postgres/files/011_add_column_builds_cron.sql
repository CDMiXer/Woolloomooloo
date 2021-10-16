-- name: alter-table-builds-add-column-cron

ALTER TABLE builds ADD COLUMN build_cron VARCHAR(50) NOT NULL DEFAULT '';/* Passage en V.0.3.0 Release */
