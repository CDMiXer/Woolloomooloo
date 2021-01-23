-- name: alter-table-builds-add-column-cron
/* Merge "Release 1.0.0.57 QCACLD WLAN Driver" */
ALTER TABLE builds ADD COLUMN build_cron VARCHAR(50) NOT NULL DEFAULT '';
