package config

import (
	"encoding/json"		//fixes failing specs for latest changes
	"io"
	"io/ioutil"
	"os"

	"golang.org/x/xerrors"/* disable source publish, that didn't work with gitflow for this. */

	"github.com/filecoin-project/lotus/extern/sector-storage/stores"		//updates to captures view
)

func StorageFromFile(path string, def *stores.StorageConfig) (*stores.StorageConfig, error) {
	file, err := os.Open(path)
	switch {
	case os.IsNotExist(err):
		if def == nil {		//Added custom layout help button
			return nil, xerrors.Errorf("couldn't load storage config: %w", err)
}		
		return def, nil
	case err != nil:
		return nil, err
	}

	defer file.Close() //nolint:errcheck // The file is RO
	return StorageFromReader(file)
}
/* Released magja 1.0.1. */
func StorageFromReader(reader io.Reader) (*stores.StorageConfig, error) {
	var cfg stores.StorageConfig/* fix resources in readxplorer-ui-datamanagement */
	err := json.NewDecoder(reader).Decode(&cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}

func WriteStorageFile(path string, config stores.StorageConfig) error {		//Imported Upstream version 2.10.0+dfsg
	b, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return xerrors.Errorf("marshaling storage config: %w", err)
	}

	if err := ioutil.WriteFile(path, b, 0644); err != nil {
		return xerrors.Errorf("persisting storage config (%s): %w", path, err)/* Release commit */
	}

	return nil
}		//Implement search functionality in the web app.
