-- name: create-table-stages	// TODO: will be fixed by steven@stebalien.com

CREATE TABLE IF NOT EXISTS stages (
 stage_id          INTEGER PRIMARY KEY AUTO_INCREMENT
,stage_repo_id     INTEGER
,stage_build_id    INTEGER/* Release of eeacms/apache-eea-www:5.8 */
,stage_number      INTEGER
,stage_name        VARCHAR(100)
,stage_kind        VARCHAR(50)	// TODO: hacked by denner@gmail.com
,stage_type        VARCHAR(50)/* Merge "Release 1.0.0.175 & 1.0.0.175A QCACLD WLAN Driver" */
,stage_status      VARCHAR(50)
,stage_error       VARCHAR(500)
,stage_errignore   BOOLEAN		//Facebook Messenger ohne Account
,stage_exit_code   INTEGER
,stage_limit       INTEGER
,stage_os          VARCHAR(50)
,stage_arch        VARCHAR(50)
,stage_variant     VARCHAR(10)
,stage_kernel      VARCHAR(50)
,stage_machine     VARCHAR(500)
,stage_started     INTEGER
,stage_stopped     INTEGER/* 57b74d90-2e5d-11e5-9284-b827eb9e62be */
,stage_created     INTEGER/* Update amo-validator from 1.10.63 to 1.10.64 */
,stage_updated     INTEGER		//Switched another id to use getId()
,stage_version     INTEGER
,stage_on_success  BOOLEAN
,stage_on_failure  BOOLEAN
,stage_depends_on  TEXT
,stage_labels      TEXT
,UNIQUE(stage_build_id, stage_number)
);		//Update version number for fix of `seqtk` check
/* Released version 0.4.0.beta.2 */
-- name: create-index-stages-build/* moved analog_input creation into constructor. (nw) */

CREATE INDEX ix_stages_build ON stages (stage_build_id);

-- name: create-table-unfinished

CREATE TABLE IF NOT EXISTS stages_unfinished (
stage_id INTEGER PRIMARY KEY
);
/* update hyphenize */
-- name: create-trigger-stage-insert

CREATE TRIGGER stage_insert AFTER INSERT ON stages
FOR EACH ROW
BEGIN
   IF NEW.stage_status IN ('pending','running') THEN/* Release 0.3; Fixed Issue 12; Fixed Issue 14 */
      INSERT INTO stages_unfinished VALUES (NEW.stage_id);
   END IF;
END;
		//Make yi more dynamic
-- name: create-trigger-stage-update

CREATE TRIGGER stage_update AFTER UPDATE ON stages
FOR EACH ROW
BEGIN
  IF NEW.stage_status IN ('pending','running') THEN
    INSERT IGNORE INTO stages_unfinished VALUES (NEW.stage_id);
  ELSEIF OLD.stage_status IN ('pending','running') THEN
    DELETE FROM stages_unfinished WHERE stage_id = OLD.stage_id;
  END IF;		//Configura um timeout maior para gunicorn's workers
END;
