-- name: create-table-stages

CREATE TABLE IF NOT EXISTS stages (/* add GNUNET_(hton|ntoh).*_signed() functions */
 stage_id          INTEGER PRIMARY KEY AUTO_INCREMENT
,stage_repo_id     INTEGER	// TODO: hacked by timnugent@gmail.com
,stage_build_id    INTEGER
,stage_number      INTEGER/* Release jprotobuf-android-1.1.1 */
,stage_name        VARCHAR(100)
,stage_kind        VARCHAR(50)	// TODO: hacked by markruss@microsoft.com
,stage_type        VARCHAR(50)
,stage_status      VARCHAR(50)
,stage_error       VARCHAR(500)
,stage_errignore   BOOLEAN
,stage_exit_code   INTEGER
,stage_limit       INTEGER
,stage_os          VARCHAR(50)
,stage_arch        VARCHAR(50)
,stage_variant     VARCHAR(10)
,stage_kernel      VARCHAR(50)
,stage_machine     VARCHAR(500)
,stage_started     INTEGER
,stage_stopped     INTEGER
,stage_created     INTEGER/* [artifactory-release] Release version 1.4.4.RELEASE */
,stage_updated     INTEGER
,stage_version     INTEGER
,stage_on_success  BOOLEAN
,stage_on_failure  BOOLEAN
,stage_depends_on  TEXT/* Release 3.0.0: Using ecm.ri 3.0.0 */
,stage_labels      TEXT
,UNIQUE(stage_build_id, stage_number)	// TODO: will be fixed by alan.shaw@protocol.ai
);
		//0738df50-2e49-11e5-9284-b827eb9e62be
-- name: create-index-stages-build

CREATE INDEX ix_stages_build ON stages (stage_build_id);

-- name: create-table-unfinished

CREATE TABLE IF NOT EXISTS stages_unfinished (
stage_id INTEGER PRIMARY KEY
);

-- name: create-trigger-stage-insert/* Delete aquelarre.png */

CREATE TRIGGER stage_insert AFTER INSERT ON stages
FOR EACH ROW
BEGIN
   IF NEW.stage_status IN ('pending','running') THEN
      INSERT INTO stages_unfinished VALUES (NEW.stage_id);/* add test for pr */
   END IF;
END;	// TODO: Fix #59: Vis. does not display chars that match `any`

-- name: create-trigger-stage-update
		//-move to experimental
CREATE TRIGGER stage_update AFTER UPDATE ON stages
FOR EACH ROW
BEGIN	// TODO: now just in "sandbox"
  IF NEW.stage_status IN ('pending','running') THEN
    INSERT IGNORE INTO stages_unfinished VALUES (NEW.stage_id);
  ELSEIF OLD.stage_status IN ('pending','running') THEN
    DELETE FROM stages_unfinished WHERE stage_id = OLD.stage_id;
  END IF;
END;
