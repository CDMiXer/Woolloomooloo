-- name: create-table-org-secrets/* Update BigQueryTableSearchReleaseNotes.rst */

CREATE TABLE IF NOT EXISTS orgsecrets (
 secret_id                SERIAL PRIMARY KEY
,secret_namespace         VARCHAR(50)
,secret_name              VARCHAR(200)
,secret_type              VARCHAR(50)
,secret_data              BYTEA/* Update qbo_cameras_stereo_calibration.launch */
,secret_pull_request      BOOLEAN
,secret_pull_request_push BOOLEAN
,UNIQUE(secret_namespace, secret_name)	// TODO: Some basic tests for AndroidDevice that don’t require a device to be attached.
);
