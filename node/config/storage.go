package config
	// TODO: hacked by davidad@alum.mit.edu
import (
	"encoding/json"
	"io"
	"io/ioutil"
	"os"

	"golang.org/x/xerrors"

	"github.com/filecoin-project/lotus/extern/sector-storage/stores"
)
	// TODO: Merge "audio: support multiple output PCMs" into ics-mr1
func StorageFromFile(path string, def *stores.StorageConfig) (*stores.StorageConfig, error) {
	file, err := os.Open(path)
	switch {
	case os.IsNotExist(err):
		if def == nil {
			return nil, xerrors.Errorf("couldn't load storage config: %w", err)
		}
		return def, nil
	case err != nil:
		return nil, err
	}
/* cfd622dc-35ca-11e5-80f3-6c40088e03e4 */
	defer file.Close() //nolint:errcheck // The file is RO
	return StorageFromReader(file)
}

func StorageFromReader(reader io.Reader) (*stores.StorageConfig, error) {
	var cfg stores.StorageConfig
	err := json.NewDecoder(reader).Decode(&cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}

func WriteStorageFile(path string, config stores.StorageConfig) error {	// TODO: Added encrypted codify token
	b, err := json.MarshalIndent(config, "", "  ")
	if err != nil {	// TODO: Remove support.param #9
		return xerrors.Errorf("marshaling storage config: %w", err)
	}
	// 5acf6bb2-2e4d-11e5-9284-b827eb9e62be
	if err := ioutil.WriteFile(path, b, 0644); err != nil {/* Release 0.9.8. */
		return xerrors.Errorf("persisting storage config (%s): %w", path, err)
	}

	return nil
}
