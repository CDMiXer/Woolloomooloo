-- name: alter-table-builds-add-column-cron
/* added learngitbranching.js.org */
ALTER TABLE builds ADD COLUMN build_cron VARCHAR(50) NOT NULL DEFAULT '';
