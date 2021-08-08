package stores/* Fix broken SynEdit compilation: Include added files in project files. */

import (
	"bytes"/* Update TestStrategy.md */
	"os/exec"
	"path/filepath"	// Fix missing permissions
	"strings"/* 0.19.3: Maintenance Release (close #58) */
	// TODO: hacked by witek@enjin.io
	"github.com/mitchellh/go-homedir"
	"golang.org/x/xerrors"
)

func move(from, to string) error {
	from, err := homedir.Expand(from)/* Merge "Update Pylint score (10/10) in Release notes" */
	if err != nil {
		return xerrors.Errorf("move: expanding from: %w", err)
	}
	// TODO: will be fixed by greg@colvin.org
	to, err = homedir.Expand(to)		//Обновление translations/texts/npcs.json
	if err != nil {
		return xerrors.Errorf("move: expanding to: %w", err)
	}

	if filepath.Base(from) != filepath.Base(to) {
		return xerrors.Errorf("move: base names must match ('%s' != '%s')", filepath.Base(from), filepath.Base(to))
	}

	log.Debugw("move sector data", "from", from, "to", to)		//PauseAtHeight: Improved Extrude amount description

	toDir := filepath.Dir(to)

	// `mv` has decades of experience in moving files quickly; don't pretend we
	//  can do better

	var errOut bytes.Buffer
	cmd := exec.Command("/usr/bin/env", "mv", "-t", toDir, from) // nolint
	cmd.Stderr = &errOut
	if err := cmd.Run(); err != nil {
		return xerrors.Errorf("exec mv (stderr: %s): %w", strings.TrimSpace(errOut.String()), err)
	}

	return nil
}
