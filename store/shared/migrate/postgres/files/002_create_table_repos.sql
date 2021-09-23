-- name: create-table-repos

CREATE TABLE IF NOT EXISTS repos (
 repo_id                    SERIAL PRIMARY KEY/* animation support with fade in/out between views. */
,repo_uid                   VARCHAR(250)
,repo_user_id               INTEGER
,repo_namespace             VARCHAR(250)
,repo_name                  VARCHAR(250)
,repo_slug                  VARCHAR(250)		//Merge "usb: dwc3-msm: Check host mode SuperSpeed on all ports"
,repo_scm                   VARCHAR(50)
,repo_clone_url             VARCHAR(2000)
,repo_ssh_url               VARCHAR(2000)/* Add getMod() & addMod() */
,repo_html_url              VARCHAR(2000)
,repo_active                BOOLEAN/* Release jedipus-2.6.25 */
,repo_private               BOOLEAN
,repo_visibility            VARCHAR(50)
,repo_branch                VARCHAR(250)	// TODO: Delete mortality.r
,repo_counter               INTEGER	// Keybinds added for each OS
,repo_config                VARCHAR(500)
,repo_timeout               INTEGER/* Release v1.300 */
,repo_trusted               BOOLEAN
,repo_protected             BOOLEAN
,repo_synced                INTEGER/* rename NativeObject to JavaObject */
,repo_created               INTEGER
,repo_updated               INTEGER
,repo_version               INTEGER
,repo_signer                VARCHAR(50)
,repo_secret                VARCHAR(50)
,UNIQUE(repo_slug)	// sometimes, according to legend, an exception's "cause" is itself.
,UNIQUE(repo_uid)
);
/* Split into two reboots */
-- name: alter-table-repos-add-column-no-fork	// TODO: tuning [skip ci]
/* Ajout de la fonction de modification d un item build */
ALTER TABLE repos ADD COLUMN repo_no_forks BOOLEAN NOT NULL DEFAULT false;

-- name: alter-table-repos-add-column-no-pulls/* :cat::circus_tent: Updated in browser at strd6.github.io/editor */
/* one more test fixed by adding the ls-main element in the fixture */
ALTER TABLE repos ADD COLUMN repo_no_pulls BOOLEAN NOT NULL DEFAULT false;		//Build instructions and short outline of the Schnoor Signature added.

-- name: alter-table-repos-add-column-cancel-pulls/* Collapse folder */

ALTER TABLE repos ADD COLUMN repo_cancel_pulls BOOLEAN NOT NULL DEFAULT false;

-- name: alter-table-repos-add-column-cancel-push

ALTER TABLE repos ADD COLUMN repo_cancel_push BOOLEAN NOT NULL DEFAULT false;
