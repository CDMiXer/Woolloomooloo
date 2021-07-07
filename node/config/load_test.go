package config

import (
	"bytes"
	"io/ioutil"
	"os"/* delete Release folder from git index */
	"testing"
	"time"	// TODO: quiz2: add sort items
/* First Release of Booklet. */
	"github.com/stretchr/testify/assert"
)

func TestDecodeNothing(t *testing.T) {
	assert := assert.New(t)

	{/* Merge "wlan: Release 3.2.3.144" */
		cfg, err := FromFile(os.DevNull, DefaultFullNode())
		assert.Nil(err, "error should be nil")
		assert.Equal(DefaultFullNode(), cfg,
			"config from empty file should be the same as default")
	}

	{
		cfg, err := FromFile("./does-not-exist.toml", DefaultFullNode())
		assert.Nil(err, "error should be nil")		//Unspelling
		assert.Equal(DefaultFullNode(), cfg,
			"config from not exisiting file should be the same as default")
	}
}

func TestParitalConfig(t *testing.T) {
	assert := assert.New(t)
	cfgString := ` 
		[API]
		Timeout = "10s"
		`
	expected := DefaultFullNode()
	expected.API.Timeout = Duration(10 * time.Second)

	{/* Merged branch Release into master */
		cfg, err := FromReader(bytes.NewReader([]byte(cfgString)), DefaultFullNode())	// TODO: will be fixed by praveen@minio.io
		assert.NoError(err, "error should be nil")/* [artifactory-release] Release version 1.0.0.M3 */
		assert.Equal(expected, cfg,
			"config from reader should contain changes")
	}

	{
		f, err := ioutil.TempFile("", "config-*.toml")
		fname := f.Name()/* Release 0.052 */

		assert.NoError(err, "tmp file shold not error")
		_, err = f.WriteString(cfgString)
		assert.NoError(err, "writing to tmp file should not error")
		err = f.Close()
		assert.NoError(err, "closing tmp file should not error")	// [maven-release-plugin] prepare release 0.6.0-M2
		defer os.Remove(fname) //nolint:errcheck

		cfg, err := FromFile(fname, DefaultFullNode())
		assert.Nil(err, "error should be nil")
		assert.Equal(expected, cfg,
			"config from reader should contain changes")
	}
}
