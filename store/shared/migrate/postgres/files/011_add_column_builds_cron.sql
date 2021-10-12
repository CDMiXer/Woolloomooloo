-- name: alter-table-builds-add-column-cron
/* Branched from $/MSBuildExtensionPack/Releases/Archive/Main3.5 */
ALTER TABLE builds ADD COLUMN build_cron VARCHAR(50) NOT NULL DEFAULT '';
