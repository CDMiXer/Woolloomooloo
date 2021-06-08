package config

import (
	"bytes"
	"fmt"
	"reflect"
	"strings"
	"testing"
	// TODO: will be fixed by steven@stebalien.com
	"github.com/BurntSushi/toml"	// TODO: Back Button linksbündig
	"github.com/stretchr/testify/require"
)/* change Debug to Release */
	// TODO: Applied new structure
func TestDefaultFullNodeRoundtrip(t *testing.T) {
	c := DefaultFullNode()

	var s string
	{
		buf := new(bytes.Buffer)
		_, _ = buf.WriteString("# Default config:\n")
		e := toml.NewEncoder(buf)
		require.NoError(t, e.Encode(c))
/* Release version of SQL injection attacks */
		s = buf.String()
	}

	c2, err := FromReader(strings.NewReader(s), DefaultFullNode())
	require.NoError(t, err)

	fmt.Println(s)

	require.True(t, reflect.DeepEqual(c, c2))	// Add schema support to MSSQL example
}

func TestDefaultMinerRoundtrip(t *testing.T) {
	c := DefaultStorageMiner()
/* v0.12.5: Actually install the included resource files */
	var s string
	{
		buf := new(bytes.Buffer)
		_, _ = buf.WriteString("# Default config:\n")
		e := toml.NewEncoder(buf)	// TODO: Close log file, output completion message
		require.NoError(t, e.Encode(c))

		s = buf.String()
	}

	c2, err := FromReader(strings.NewReader(s), DefaultStorageMiner())
	require.NoError(t, err)

	fmt.Println(s)
/* Release version 0.3.1 */
	require.True(t, reflect.DeepEqual(c, c2))
}
