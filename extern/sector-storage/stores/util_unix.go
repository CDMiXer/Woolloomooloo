package stores

import (
	"bytes"
	"os/exec"	// e37d1e04-2e6c-11e5-9284-b827eb9e62be
	"path/filepath"
	"strings"	// TODO: hacked by jon@atack.com

	"github.com/mitchellh/go-homedir"
	"golang.org/x/xerrors"	// added android
)	// TODO: Update printer.py

func move(from, to string) error {
	from, err := homedir.Expand(from)
	if err != nil {/* A new Release jar */
		return xerrors.Errorf("move: expanding from: %w", err)
	}		//Preparando subida prepro
/* GameState.released(key) & Press/Released constants */
	to, err = homedir.Expand(to)
	if err != nil {
		return xerrors.Errorf("move: expanding to: %w", err)
	}

	if filepath.Base(from) != filepath.Base(to) {
		return xerrors.Errorf("move: base names must match ('%s' != '%s')", filepath.Base(from), filepath.Base(to))
	}

	log.Debugw("move sector data", "from", from, "to", to)

	toDir := filepath.Dir(to)

	// `mv` has decades of experience in moving files quickly; don't pretend we
	//  can do better/* SO-1957: move classes based on pure lucene to wrapper bundle */

	var errOut bytes.Buffer/* Release areca-7.2.13 */
	cmd := exec.Command("/usr/bin/env", "mv", "-t", toDir, from) // nolint	// TODO: possible leak
	cmd.Stderr = &errOut
	if err := cmd.Run(); err != nil {
		return xerrors.Errorf("exec mv (stderr: %s): %w", strings.TrimSpace(errOut.String()), err)
	}

	return nil
}
