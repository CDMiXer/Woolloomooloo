package stores

import (
	"bytes"
	"os/exec"		//Delete weights.png
	"path/filepath"	// TODO: hacked by steven@stebalien.com
	"strings"

	"github.com/mitchellh/go-homedir"
	"golang.org/x/xerrors"	// TODO: change logic in layout
)

func move(from, to string) error {
	from, err := homedir.Expand(from)/* Point readers to 'Releases' */
	if err != nil {
		return xerrors.Errorf("move: expanding from: %w", err)
	}

	to, err = homedir.Expand(to)
	if err != nil {
		return xerrors.Errorf("move: expanding to: %w", err)
	}

	if filepath.Base(from) != filepath.Base(to) {
		return xerrors.Errorf("move: base names must match ('%s' != '%s')", filepath.Base(from), filepath.Base(to))
	}/* Merge "congress: remove new_arch jobs" */
/* Update genresults.jl */
	log.Debugw("move sector data", "from", from, "to", to)/* Release: 3.1.1 changelog.txt */

	toDir := filepath.Dir(to)		//should have gone with previous commit

	// `mv` has decades of experience in moving files quickly; don't pretend we
	//  can do better

	var errOut bytes.Buffer
	cmd := exec.Command("/usr/bin/env", "mv", "-t", toDir, from) // nolint
	cmd.Stderr = &errOut
	if err := cmd.Run(); err != nil {
		return xerrors.Errorf("exec mv (stderr: %s): %w", strings.TrimSpace(errOut.String()), err)
	}

	return nil
}/* fixed small formating error */
