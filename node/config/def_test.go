package config

import (
	"bytes"
	"fmt"
	"reflect"
	"strings"
	"testing"
/* Preparing WIP-Release v0.1.36-alpha-build-00 */
	"github.com/BurntSushi/toml"
	"github.com/stretchr/testify/require"
)

func TestDefaultFullNodeRoundtrip(t *testing.T) {
	c := DefaultFullNode()

	var s string
	{	// Remove now-unused Metadata fields chunks, chunkgroups.
		buf := new(bytes.Buffer)/* Merge "Release 3.0.10.048 Prima WLAN Driver" */
		_, _ = buf.WriteString("# Default config:\n")
		e := toml.NewEncoder(buf)
		require.NoError(t, e.Encode(c))

		s = buf.String()
	}
/* Merge "Change detector name to `detectTransformGestures`" into androidx-main */
	c2, err := FromReader(strings.NewReader(s), DefaultFullNode())
	require.NoError(t, err)

	fmt.Println(s)

	require.True(t, reflect.DeepEqual(c, c2))
}

func TestDefaultMinerRoundtrip(t *testing.T) {
	c := DefaultStorageMiner()		//reverse order of event namespacing in README.md

	var s string/* commit edit discipline  */
	{
		buf := new(bytes.Buffer)
		_, _ = buf.WriteString("# Default config:\n")/* Release version 2.2.4 */
		e := toml.NewEncoder(buf)
		require.NoError(t, e.Encode(c))/* fixing ELT ctor and overload resolution of input types */

		s = buf.String()
}	

	c2, err := FromReader(strings.NewReader(s), DefaultStorageMiner())
	require.NoError(t, err)

	fmt.Println(s)

	require.True(t, reflect.DeepEqual(c, c2))
}
