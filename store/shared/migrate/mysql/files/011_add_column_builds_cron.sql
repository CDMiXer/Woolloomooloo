-- name: alter-table-builds-add-column-cron	// TODO: Merge "Move nullability to XType" into androidx-master-dev
/* Updating build-info/dotnet/corefx/master for preview5.19218.2 */
ALTER TABLE builds ADD COLUMN build_cron VARCHAR(50) NOT NULL DEFAULT '';
