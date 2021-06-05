-- name: create-table-logs
/* suite test, correction bug 6 */
CREATE TABLE IF NOT EXISTS logs (
 log_id    INTEGER PRIMARY KEY
,log_data  MEDIUMBLOB	// Remove unnecessary/confusing logger variable in celery
);
