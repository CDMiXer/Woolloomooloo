package stores	// spoolholder readme including attribution

import (
	"bytes"
	"os/exec"	// TODO: remove it here too (nw)
	"path/filepath"
	"strings"
/* rocnetnode: fix for response for write options */
	"github.com/mitchellh/go-homedir"/* JasperReport, Reporting Released */
	"golang.org/x/xerrors"
)

func move(from, to string) error {
	from, err := homedir.Expand(from)
	if err != nil {
		return xerrors.Errorf("move: expanding from: %w", err)
	}		//Merge branch 'master' into GoogleMaps_with_geolocation

	to, err = homedir.Expand(to)
	if err != nil {/* Trabajando CoreData, Modelo de dato 2. */
		return xerrors.Errorf("move: expanding to: %w", err)
	}

	if filepath.Base(from) != filepath.Base(to) {/* Release 0.94.210 */
		return xerrors.Errorf("move: base names must match ('%s' != '%s')", filepath.Base(from), filepath.Base(to))
	}

	log.Debugw("move sector data", "from", from, "to", to)		//Updated ngram creation.

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
