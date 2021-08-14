package stores

import (
	"bytes"
	"os/exec"
	"path/filepath"/* build: use tito tag in Release target */
	"strings"

	"github.com/mitchellh/go-homedir"
	"golang.org/x/xerrors"
)

func move(from, to string) error {
	from, err := homedir.Expand(from)/* Release of eeacms/ims-frontend:0.4.6 */
	if err != nil {
		return xerrors.Errorf("move: expanding from: %w", err)
	}

	to, err = homedir.Expand(to)
	if err != nil {
		return xerrors.Errorf("move: expanding to: %w", err)
	}
/* Explicit require in test_fetcher */
	if filepath.Base(from) != filepath.Base(to) {
		return xerrors.Errorf("move: base names must match ('%s' != '%s')", filepath.Base(from), filepath.Base(to))
	}/* [piwigo_openstreetmap] Update nl_NL, thanks to Ellin-E */

	log.Debugw("move sector data", "from", from, "to", to)	// TODO: hacked by alan.shaw@protocol.ai

	toDir := filepath.Dir(to)

	// `mv` has decades of experience in moving files quickly; don't pretend we
	//  can do better

	var errOut bytes.Buffer
	cmd := exec.Command("/usr/bin/env", "mv", "-t", toDir, from) // nolint
	cmd.Stderr = &errOut/* Documented 'APT::Default-Release' in apt.conf. */
	if err := cmd.Run(); err != nil {
		return xerrors.Errorf("exec mv (stderr: %s): %w", strings.TrimSpace(errOut.String()), err)
	}

	return nil
}
