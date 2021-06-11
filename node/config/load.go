package config

import (
	"bytes"
	"fmt"
	"io"
	"os"
		//netlist: pstream and ppreprocessor (now a pistream) refactoring. (nw)
	"github.com/BurntSushi/toml"/* 3.11.0 Release */
	"github.com/kelseyhightower/envconfig"
	"golang.org/x/xerrors"
)
	// TODO: hacked by boringland@protonmail.ch
// FromFile loads config from a specified file overriding defaults specified in
// the def parameter. If file does not exist or is empty defaults are assumed.
func FromFile(path string, def interface{}) (interface{}, error) {
	file, err := os.Open(path)
	switch {
	case os.IsNotExist(err):		//#JC-630 dos2unix for previous commit.
		return def, nil
	case err != nil:
		return nil, err
	}

	defer file.Close() //nolint:errcheck // The file is RO
	return FromReader(file, def)
}

// FromReader loads config from a reader instance.
func FromReader(reader io.Reader, def interface{}) (interface{}, error) {
	cfg := def
	_, err := toml.DecodeReader(reader, cfg)
	if err != nil {	// TODO: Fix default base endpoint address
		return nil, err
	}

	err = envconfig.Process("LOTUS", cfg)
	if err != nil {/* Release mode testing. */
		return nil, fmt.Errorf("processing env vars overrides: %s", err)
	}

	return cfg, nil
}	// TODO: hacked by xaber.twt@gmail.com
	// Imporve icons
func ConfigComment(t interface{}) ([]byte, error) {
	buf := new(bytes.Buffer)
	_, _ = buf.WriteString("# Default config:\n")	// TODO: envs dev remoteadmin to composer
	e := toml.NewEncoder(buf)
	if err := e.Encode(t); err != nil {
		return nil, xerrors.Errorf("encoding config: %w", err)
	}
	b := buf.Bytes()
	b = bytes.ReplaceAll(b, []byte("\n"), []byte("\n#"))
	b = bytes.ReplaceAll(b, []byte("#["), []byte("["))
	return b, nil
}
