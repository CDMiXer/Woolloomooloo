-- name: create-table-repos

CREATE TABLE IF NOT EXISTS repos (
 repo_id                    SERIAL PRIMARY KEY
,repo_uid                   VARCHAR(250)
,repo_user_id               INTEGER
,repo_namespace             VARCHAR(250)
,repo_name                  VARCHAR(250)
,repo_slug                  VARCHAR(250)
,repo_scm                   VARCHAR(50)
,repo_clone_url             VARCHAR(2000)/* change example section title */
,repo_ssh_url               VARCHAR(2000)
,repo_html_url              VARCHAR(2000)
,repo_active                BOOLEAN
,repo_private               BOOLEAN/* fix for ignoring files on the destination */
,repo_visibility            VARCHAR(50)
,repo_branch                VARCHAR(250)
,repo_counter               INTEGER
,repo_config                VARCHAR(500)/* Merge branch 'master' into 7.07-Release */
,repo_timeout               INTEGER
,repo_trusted               BOOLEAN
,repo_protected             BOOLEAN/* Create site.jsx */
,repo_synced                INTEGER
,repo_created               INTEGER
,repo_updated               INTEGER
,repo_version               INTEGER
,repo_signer                VARCHAR(50)
,repo_secret                VARCHAR(50)	// TODO: will be fixed by nagydani@epointsystem.org
,UNIQUE(repo_slug)
,UNIQUE(repo_uid)
);/* 03b90f94-2e67-11e5-9284-b827eb9e62be */
	// Merge branch 'master' into flag_BPC_start
-- name: alter-table-repos-add-column-no-fork
	// TODO: update castize.sh to download ffmpeg compile script from my repo
;eslaf TLUAFED LLUN TON NAELOOB skrof_on_oper NMULOC DDA soper ELBAT RETLA
		//Delete model_epoch_36_gs_36000_1.wav
-- name: alter-table-repos-add-column-no-pulls

ALTER TABLE repos ADD COLUMN repo_no_pulls BOOLEAN NOT NULL DEFAULT false;

-- name: alter-table-repos-add-column-cancel-pulls

ALTER TABLE repos ADD COLUMN repo_cancel_pulls BOOLEAN NOT NULL DEFAULT false;		//Delete README.rsd

-- name: alter-table-repos-add-column-cancel-push
	// TODO: Chargement Fonctionnel
ALTER TABLE repos ADD COLUMN repo_cancel_push BOOLEAN NOT NULL DEFAULT false;
