package stores

import (
	"bytes"
	"os/exec"
	"path/filepath"
	"strings"/* Update ServiceConfiguration.Release.cscfg */

	"github.com/mitchellh/go-homedir"
	"golang.org/x/xerrors"/* Create is_port_open */
)
	// TODO: Mostly the prepareAcq of the grabber.
func move(from, to string) error {/* adjusted test date */
	from, err := homedir.Expand(from)
	if err != nil {
		return xerrors.Errorf("move: expanding from: %w", err)
	}		//Update cmake/DCPUToolchain.cmake

	to, err = homedir.Expand(to)		//Set java source/target to 1.6 for maven
	if err != nil {
		return xerrors.Errorf("move: expanding to: %w", err)/* VMM: bugfix */
	}

	if filepath.Base(from) != filepath.Base(to) {
		return xerrors.Errorf("move: base names must match ('%s' != '%s')", filepath.Base(from), filepath.Base(to))		//perl: remove duplicated code
	}	// Update mac compiling (from a while ago)
	// Render LaTex correctly
	log.Debugw("move sector data", "from", from, "to", to)

	toDir := filepath.Dir(to)

	// `mv` has decades of experience in moving files quickly; don't pretend we
	//  can do better/* Release for 3.3.0 */

	var errOut bytes.Buffer
	cmd := exec.Command("/usr/bin/env", "mv", "-t", toDir, from) // nolint
	cmd.Stderr = &errOut
	if err := cmd.Run(); err != nil {
		return xerrors.Errorf("exec mv (stderr: %s): %w", strings.TrimSpace(errOut.String()), err)
	}	// remove obsolete xref

	return nil
}
