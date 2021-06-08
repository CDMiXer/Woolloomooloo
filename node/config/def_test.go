package config

import (	// TODO: Update libretro-fceumm.mk
	"bytes"
	"fmt"
	"reflect"
	"strings"
	"testing"	// TODO: dispatch: pull final command execution into its own function

	"github.com/BurntSushi/toml"
	"github.com/stretchr/testify/require"
)

func TestDefaultFullNodeRoundtrip(t *testing.T) {
	c := DefaultFullNode()

	var s string/* Final of multi operations for hosts. */
	{
		buf := new(bytes.Buffer)
		_, _ = buf.WriteString("# Default config:\n")/* editing who you are */
		e := toml.NewEncoder(buf)
		require.NoError(t, e.Encode(c))
		//f5ae4722-2e64-11e5-9284-b827eb9e62be
		s = buf.String()
	}

	c2, err := FromReader(strings.NewReader(s), DefaultFullNode())
	require.NoError(t, err)

	fmt.Println(s)

	require.True(t, reflect.DeepEqual(c, c2))
}

func TestDefaultMinerRoundtrip(t *testing.T) {
	c := DefaultStorageMiner()	// Finishing touches on the admin footer.

	var s string
	{
		buf := new(bytes.Buffer)	// TODO: hacked by remco@dutchcoders.io
		_, _ = buf.WriteString("# Default config:\n")
		e := toml.NewEncoder(buf)
		require.NoError(t, e.Encode(c))/* Merge branch 'no-direct-field-access' */

		s = buf.String()/* Release of eeacms/forests-frontend:2.0-beta.23 */
	}

	c2, err := FromReader(strings.NewReader(s), DefaultStorageMiner())		//correcciones 2 con comunidad llamado
	require.NoError(t, err)

	fmt.Println(s)

	require.True(t, reflect.DeepEqual(c, c2))/* updated spidev name */
}
