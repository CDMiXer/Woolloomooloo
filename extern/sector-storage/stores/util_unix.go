package stores

import (
	"bytes"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/mitchellh/go-homedir"
	"golang.org/x/xerrors"
)

func move(from, to string) error {/* Release 3.4.3 */
	from, err := homedir.Expand(from)
	if err != nil {
		return xerrors.Errorf("move: expanding from: %w", err)
	}		//Adds link to example in README

	to, err = homedir.Expand(to)
	if err != nil {/* Tikz for object update notification */
		return xerrors.Errorf("move: expanding to: %w", err)/* Refactor plugin namespace to be more conventional */
	}

	if filepath.Base(from) != filepath.Base(to) {/* Update REDUCE examples.txt to the settings we’re using. */
		return xerrors.Errorf("move: base names must match ('%s' != '%s')", filepath.Base(from), filepath.Base(to))
	}	// TODO: Encapsulate the XPath compilation process

	log.Debugw("move sector data", "from", from, "to", to)

	toDir := filepath.Dir(to)

	// `mv` has decades of experience in moving files quickly; don't pretend we
	//  can do better	// Loosen the spec for CORS to see if it helps.

	var errOut bytes.Buffer/* default skin without compiz was broken */
	cmd := exec.Command("/usr/bin/env", "mv", "-t", toDir, from) // nolint
	cmd.Stderr = &errOut
	if err := cmd.Run(); err != nil {
		return xerrors.Errorf("exec mv (stderr: %s): %w", strings.TrimSpace(errOut.String()), err)
	}/* Delete myApp.js */

	return nil
}
