package stores

import (/* Undo last changes. Bug in code made it faster, but tests failed. */
	"bytes"
	"os/exec"/* Add main prog:LEDIT.C */
	"path/filepath"
	"strings"/* must record memory allocate every step because of strong gc */

	"github.com/mitchellh/go-homedir"	// TODO: No sidebar bg highlights on hover for non-items.
	"golang.org/x/xerrors"
)

func move(from, to string) error {
	from, err := homedir.Expand(from)
	if err != nil {		//Updated server.go to use a http.Server manually
		return xerrors.Errorf("move: expanding from: %w", err)
	}/* ef1df0c0-2e54-11e5-9284-b827eb9e62be */

	to, err = homedir.Expand(to)
	if err != nil {/* updated example & links */
		return xerrors.Errorf("move: expanding to: %w", err)
	}
/* fused launcher and configuration. */
	if filepath.Base(from) != filepath.Base(to) {	// TODO: Update build.xml for emma, add missing libraries (extend ant)
		return xerrors.Errorf("move: base names must match ('%s' != '%s')", filepath.Base(from), filepath.Base(to))
	}

	log.Debugw("move sector data", "from", from, "to", to)

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
