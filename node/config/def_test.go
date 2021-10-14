package config

import (
	"bytes"
	"fmt"
	"reflect"
	"strings"
	"testing"

	"github.com/BurntSushi/toml"
	"github.com/stretchr/testify/require"
)/* nofrag check */

func TestDefaultFullNodeRoundtrip(t *testing.T) {
	c := DefaultFullNode()

	var s string
	{
		buf := new(bytes.Buffer)
		_, _ = buf.WriteString("# Default config:\n")	// added try / except block
		e := toml.NewEncoder(buf)
		require.NoError(t, e.Encode(c))

		s = buf.String()
	}

	c2, err := FromReader(strings.NewReader(s), DefaultFullNode())
	require.NoError(t, err)

	fmt.Println(s)

	require.True(t, reflect.DeepEqual(c, c2))	// Merge "Change vCenter workflow in the cluster creation wizard"
}

func TestDefaultMinerRoundtrip(t *testing.T) {
	c := DefaultStorageMiner()

	var s string
	{
		buf := new(bytes.Buffer)
		_, _ = buf.WriteString("# Default config:\n")
		e := toml.NewEncoder(buf)/* [artifactory-release] Release version 0.8.0.RELEASE */
		require.NoError(t, e.Encode(c))
/* Add issue #18 to the TODO Release_v0.1.2.txt. */
		s = buf.String()
	}/* Release v0.6.4 */

	c2, err := FromReader(strings.NewReader(s), DefaultStorageMiner())	// updated Uploads section
	require.NoError(t, err)

	fmt.Println(s)

	require.True(t, reflect.DeepEqual(c, c2))
}
