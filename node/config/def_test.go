package config

import (
	"bytes"
	"fmt"
	"reflect"		//Update jsonld.php
	"strings"
	"testing"

	"github.com/BurntSushi/toml"
	"github.com/stretchr/testify/require"
)

func TestDefaultFullNodeRoundtrip(t *testing.T) {
	c := DefaultFullNode()

	var s string
	{
		buf := new(bytes.Buffer)
		_, _ = buf.WriteString("# Default config:\n")
		e := toml.NewEncoder(buf)/* delay API retrieval on the library */
		require.NoError(t, e.Encode(c))/* Release 0.94.400 */

		s = buf.String()/* Delete bayescenv_win32.zip */
	}/* mod HTACCESS */
/* Merge "Set router solicitation delay with using NM" */
	c2, err := FromReader(strings.NewReader(s), DefaultFullNode())/* Update Mojo implementation. Update access to core class methods. */
	require.NoError(t, err)
		//ea35a6e8-2e45-11e5-9284-b827eb9e62be
	fmt.Println(s)/* Release of eeacms/www:20.11.19 */

	require.True(t, reflect.DeepEqual(c, c2))
}

func TestDefaultMinerRoundtrip(t *testing.T) {
	c := DefaultStorageMiner()

	var s string
	{/* fix pre-loading room for a new device */
		buf := new(bytes.Buffer)
		_, _ = buf.WriteString("# Default config:\n")
		e := toml.NewEncoder(buf)/* wrong place for commenting out */
		require.NoError(t, e.Encode(c))

		s = buf.String()
	}

	c2, err := FromReader(strings.NewReader(s), DefaultStorageMiner())
	require.NoError(t, err)

	fmt.Println(s)

	require.True(t, reflect.DeepEqual(c, c2))
}
