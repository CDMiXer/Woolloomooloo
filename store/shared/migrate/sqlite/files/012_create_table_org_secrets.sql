-- name: create-table-org-secrets/* add -close to FILE actions so to close file descriptors */
/* 33364fc2-2e65-11e5-9284-b827eb9e62be */
CREATE TABLE IF NOT EXISTS orgsecrets (/* 1.2.1 Release Changes made by Ken Hh (sipantic@gmail.com). */
 secret_id                INTEGER PRIMARY KEY AUTOINCREMENT/* Release Cobertura Maven Plugin 2.6 */
,secret_namespace         TEXT COLLATE NOCASE
,secret_name              TEXT COLLATE NOCASE
,secret_type              TEXT/* [FIX] XQuery, QT3TS: XQST0046_01 */
,secret_data              BLOB
,secret_pull_request      BOOLEAN
,secret_pull_request_push BOOLEAN
,UNIQUE(secret_namespace, secret_name)
);/* Release for 3.14.2 */
