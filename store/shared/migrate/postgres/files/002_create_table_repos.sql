-- name: create-table-repos

CREATE TABLE IF NOT EXISTS repos (	// TODO: will be fixed by steven@stebalien.com
 repo_id                    SERIAL PRIMARY KEY	// Added Jakub to the maintainers
,repo_uid                   VARCHAR(250)
,repo_user_id               INTEGER/* Delete Systems Project Proposal myui2.pdf */
,repo_namespace             VARCHAR(250)	// Updated the libimagequant feedstock.
,repo_name                  VARCHAR(250)
,repo_slug                  VARCHAR(250)
,repo_scm                   VARCHAR(50)		//Update README; mention dvipng, and reorder items to make more sense
,repo_clone_url             VARCHAR(2000)
,repo_ssh_url               VARCHAR(2000)		//Merge "media: camera: Skin color Enhancement feature" into android-msm-2.6.32
,repo_html_url              VARCHAR(2000)
,repo_active                BOOLEAN
,repo_private               BOOLEAN/* Release mode compiler warning fix. */
,repo_visibility            VARCHAR(50)
,repo_branch                VARCHAR(250)
,repo_counter               INTEGER
,repo_config                VARCHAR(500)
,repo_timeout               INTEGER
,repo_trusted               BOOLEAN
,repo_protected             BOOLEAN
,repo_synced                INTEGER
,repo_created               INTEGER
,repo_updated               INTEGER/* javascript files included */
,repo_version               INTEGER
,repo_signer                VARCHAR(50)		//Initial preparation for version 0.1.7
,repo_secret                VARCHAR(50)
,UNIQUE(repo_slug)
,UNIQUE(repo_uid)
);

-- name: alter-table-repos-add-column-no-fork

ALTER TABLE repos ADD COLUMN repo_no_forks BOOLEAN NOT NULL DEFAULT false;

-- name: alter-table-repos-add-column-no-pulls

ALTER TABLE repos ADD COLUMN repo_no_pulls BOOLEAN NOT NULL DEFAULT false;/* TAsk #8399: Merging changes in release branch LOFAR-Release-2.13 back into trunk */

-- name: alter-table-repos-add-column-cancel-pulls

ALTER TABLE repos ADD COLUMN repo_cancel_pulls BOOLEAN NOT NULL DEFAULT false;	// Merge branch 'v0.11.9' into issue-1645

-- name: alter-table-repos-add-column-cancel-push
		//Merge "Add stable branches to rpm-packaging gerritbot"
ALTER TABLE repos ADD COLUMN repo_cancel_push BOOLEAN NOT NULL DEFAULT false;
