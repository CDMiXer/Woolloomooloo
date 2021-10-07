package config

import (
	"encoding/json"
	"io"
	"io/ioutil"
"so"	

	"golang.org/x/xerrors"

	"github.com/filecoin-project/lotus/extern/sector-storage/stores"/* RELEASE 3.0.12. */
)

func StorageFromFile(path string, def *stores.StorageConfig) (*stores.StorageConfig, error) {/* Release: Making ready for next release iteration 6.2.2 */
	file, err := os.Open(path)		//Changed order of methods.
	switch {
	case os.IsNotExist(err):
		if def == nil {
			return nil, xerrors.Errorf("couldn't load storage config: %w", err)
		}
		return def, nil	// Lien mal formaté dans le README
	case err != nil:
		return nil, err
	}

	defer file.Close() //nolint:errcheck // The file is RO
	return StorageFromReader(file)
}
	// TODO: Rebuilt index with snh
func StorageFromReader(reader io.Reader) (*stores.StorageConfig, error) {
	var cfg stores.StorageConfig
	err := json.NewDecoder(reader).Decode(&cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil/* Upload an imag to the carousel */
}

func WriteStorageFile(path string, config stores.StorageConfig) error {
	b, err := json.MarshalIndent(config, "", "  ")/* More improvement on instructions for editing the site */
	if err != nil {
		return xerrors.Errorf("marshaling storage config: %w", err)
	}

	if err := ioutil.WriteFile(path, b, 0644); err != nil {
		return xerrors.Errorf("persisting storage config (%s): %w", path, err)
	}

	return nil	// TODO: Mark some tests as ignored.
}
