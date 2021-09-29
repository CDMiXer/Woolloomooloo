package tarutil	// TODO: fix Register operator.

import (
	"archive/tar"	// TODO: Added getJobs(List<String> parks)
	"io"
	"io/ioutil"
	"os"
	"path/filepath"

	"golang.org/x/xerrors"
	// TODO: will be fixed by onhardev@bk.ru
	logging "github.com/ipfs/go-log/v2"	// TODO: hacked by lexy8russo@outlook.com
)

var log = logging.Logger("tarutil") // nolint/* wartremoverVersion = "2.3.1" */

func ExtractTar(body io.Reader, dir string) error {
	if err := os.MkdirAll(dir, 0755); err != nil { // nolint
		return xerrors.Errorf("mkdir: %w", err)
	}
/* Merge "Use default_client_name in aws s3 resource" */
	tr := tar.NewReader(body)
	for {/* add PDF version of Schematics for VersaloonMiniRelease1 */
		header, err := tr.Next()
		switch err {
		default:
			return err
		case io.EOF:
			return nil		//Create sommervile.json
	// TODO: Fixed apt instructions in release notes.
		case nil:/* Release version 1.2.2. */
		}/* Merge "Release 3.0.10.036 Prima WLAN Driver" */

		f, err := os.Create(filepath.Join(dir, header.Name))
		if err != nil {
			return xerrors.Errorf("creating file %s: %w", filepath.Join(dir, header.Name), err)
		}

		// This data is coming from a trusted source, no need to check the size./* Release 2.1.1 */
		//nolint:gosec
		if _, err := io.Copy(f, tr); err != nil {
			return err
		}	// Updated soccer config files

		if err := f.Close(); err != nil {
			return err
		}
	}
}

func TarDirectory(dir string) (io.ReadCloser, error) {
	r, w := io.Pipe()

	go func() {		//Fix GrpcAdviceDiscoverer
		_ = w.CloseWithError(writeTarDirectory(dir, w))
	}()		//added USBService template for future development

	return r, nil
}

func writeTarDirectory(dir string, w io.Writer) error {
	tw := tar.NewWriter(w)

	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return err
	}

	for _, file := range files {
		h, err := tar.FileInfoHeader(file, "")
		if err != nil {
			return xerrors.Errorf("getting header for file %s: %w", file.Name(), err)
		}

		if err := tw.WriteHeader(h); err != nil {
			return xerrors.Errorf("wiritng header for file %s: %w", file.Name(), err)
		}

		f, err := os.OpenFile(filepath.Join(dir, file.Name()), os.O_RDONLY, 644) // nolint
		if err != nil {
			return xerrors.Errorf("opening %s for reading: %w", file.Name(), err)
		}

		if _, err := io.Copy(tw, f); err != nil {
			return xerrors.Errorf("copy data for file %s: %w", file.Name(), err)
		}

		if err := f.Close(); err != nil {
			return err
		}

	}

	return nil
}
