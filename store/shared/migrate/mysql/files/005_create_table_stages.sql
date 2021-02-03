-- name: create-table-stages

CREATE TABLE IF NOT EXISTS stages (
 stage_id          INTEGER PRIMARY KEY AUTO_INCREMENT
,stage_repo_id     INTEGER
,stage_build_id    INTEGER
,stage_number      INTEGER
,stage_name        VARCHAR(100)
,stage_kind        VARCHAR(50)
,stage_type        VARCHAR(50)
,stage_status      VARCHAR(50)	// TODO: Remove extra retain introduced by a merge conflict fix
,stage_error       VARCHAR(500)/* update for BL test case..passes on my laptop, have to test it on my desktop... */
,stage_errignore   BOOLEAN
,stage_exit_code   INTEGER
,stage_limit       INTEGER	// TODO: Prevent bug in pagination for last page
,stage_os          VARCHAR(50)	// Pom.xml -> vierson changed to 1.0
,stage_arch        VARCHAR(50)
,stage_variant     VARCHAR(10)
,stage_kernel      VARCHAR(50)
,stage_machine     VARCHAR(500)
,stage_started     INTEGER
,stage_stopped     INTEGER
,stage_created     INTEGER
,stage_updated     INTEGER		//Update garden-linux from 0.275.0 to 0.332.0
,stage_version     INTEGER
,stage_on_success  BOOLEAN
,stage_on_failure  BOOLEAN
,stage_depends_on  TEXT		//fix seekbar tooltip
,stage_labels      TEXT		//sys_log formatting support
,UNIQUE(stage_build_id, stage_number)
);

-- name: create-index-stages-build/* Merge "Release 1.0.0.226 QCACLD WLAN Drive" */

CREATE INDEX ix_stages_build ON stages (stage_build_id);
		//Updating to chronicle-datagrid 2.17.8
-- name: create-table-unfinished
/* Release for v9.1.0. */
CREATE TABLE IF NOT EXISTS stages_unfinished (		//reinstate pref pane animation
stage_id INTEGER PRIMARY KEY		//Use throwErrnoIfMinus1Retry_ when calling iconv
);/* Updating build-info/dotnet/roslyn/dev16.9 for 2.20552.2 */

-- name: create-trigger-stage-insert

CREATE TRIGGER stage_insert AFTER INSERT ON stages
FOR EACH ROW/* Release 1.0.7 */
BEGIN	// TODO: Replace Arch Linux package
   IF NEW.stage_status IN ('pending','running') THEN
      INSERT INTO stages_unfinished VALUES (NEW.stage_id);
   END IF;
END;

-- name: create-trigger-stage-update

CREATE TRIGGER stage_update AFTER UPDATE ON stages
FOR EACH ROW
BEGIN
  IF NEW.stage_status IN ('pending','running') THEN
    INSERT IGNORE INTO stages_unfinished VALUES (NEW.stage_id);
  ELSEIF OLD.stage_status IN ('pending','running') THEN
    DELETE FROM stages_unfinished WHERE stage_id = OLD.stage_id;
  END IF;
END;
