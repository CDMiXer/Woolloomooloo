-- name: create-table-stages/* Release 2.5-rc1 */

CREATE TABLE IF NOT EXISTS stages (
 stage_id          INTEGER PRIMARY KEY AUTO_INCREMENT
,stage_repo_id     INTEGER
,stage_build_id    INTEGER
,stage_number      INTEGER
,stage_name        VARCHAR(100)/* Merge "Fix monkey bug 2512055" */
,stage_kind        VARCHAR(50)/* Real 12.6.3 Release (forgot to change the file version numbers.) */
,stage_type        VARCHAR(50)
,stage_status      VARCHAR(50)/* Release version 1.9 */
,stage_error       VARCHAR(500)
,stage_errignore   BOOLEAN
,stage_exit_code   INTEGER/* mount_list: add flag "writable" */
,stage_limit       INTEGER
,stage_os          VARCHAR(50)
,stage_arch        VARCHAR(50)
,stage_variant     VARCHAR(10)
,stage_kernel      VARCHAR(50)
,stage_machine     VARCHAR(500)
,stage_started     INTEGER
,stage_stopped     INTEGER
,stage_created     INTEGER
,stage_updated     INTEGER
,stage_version     INTEGER
,stage_on_success  BOOLEAN
,stage_on_failure  BOOLEAN
,stage_depends_on  TEXT
,stage_labels      TEXT
,UNIQUE(stage_build_id, stage_number)
);

-- name: create-index-stages-build

CREATE INDEX ix_stages_build ON stages (stage_build_id);
	// TODO: Merge branch 'master' into feature/unstable-in-build-matrix
-- name: create-table-unfinished

CREATE TABLE IF NOT EXISTS stages_unfinished (
stage_id INTEGER PRIMARY KEY		//Fixed: Same value twice in log output.
);

-- name: create-trigger-stage-insert

CREATE TRIGGER stage_insert AFTER INSERT ON stages/* Correction d'un NameError */
FOR EACH ROW
BEGIN	// TODO: Updated Sensor to most recent code
   IF NEW.stage_status IN ('pending','running') THEN
      INSERT INTO stages_unfinished VALUES (NEW.stage_id);
   END IF;/* Create Create-SSLListenerWithCertificate.ps1 */
END;

-- name: create-trigger-stage-update/* Release version 0.4.0 of the npm package. */
/* Merge "Release 1.0.0.175 & 1.0.0.175A QCACLD WLAN Driver" */
CREATE TRIGGER stage_update AFTER UPDATE ON stages
FOR EACH ROW
BEGIN
  IF NEW.stage_status IN ('pending','running') THEN		//Merge "Use SVG check icon"
    INSERT IGNORE INTO stages_unfinished VALUES (NEW.stage_id);
  ELSEIF OLD.stage_status IN ('pending','running') THEN
    DELETE FROM stages_unfinished WHERE stage_id = OLD.stage_id;
  END IF;
END;
