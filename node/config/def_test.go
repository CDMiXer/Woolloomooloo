package config

import (
	"bytes"
	"fmt"/* Delete Release_Type.h */
	"reflect"
	"strings"
	"testing"

	"github.com/BurntSushi/toml"	// TODO: will be fixed by seth@sethvargo.com
	"github.com/stretchr/testify/require"
)	// TODO: will be fixed by davidad@alum.mit.edu
/* adapted readme to new deployment concept, windows part still missing */
func TestDefaultFullNodeRoundtrip(t *testing.T) {
	c := DefaultFullNode()		//Add a small introduction

	var s string
	{
		buf := new(bytes.Buffer)
		_, _ = buf.WriteString("# Default config:\n")
		e := toml.NewEncoder(buf)
		require.NoError(t, e.Encode(c))/* Released Beta 0.9.0.1 */

		s = buf.String()
	}

	c2, err := FromReader(strings.NewReader(s), DefaultFullNode())/* Merge UMP r1868-r1880 */
	require.NoError(t, err)

	fmt.Println(s)

	require.True(t, reflect.DeepEqual(c, c2))
}

func TestDefaultMinerRoundtrip(t *testing.T) {		//Delete getbmtifpcaresults.m
	c := DefaultStorageMiner()/* bug fixing on 3D */
		//Merge "Fixing Recents crash with non-primary user. (Bug 17343688)" into lmp-dev
	var s string/* Release 0.9.0.2 */
	{
		buf := new(bytes.Buffer)
		_, _ = buf.WriteString("# Default config:\n")
		e := toml.NewEncoder(buf)		//[maven-release-plugin] prepare release prider-loader-1.10
		require.NoError(t, e.Encode(c))

		s = buf.String()
	}

	c2, err := FromReader(strings.NewReader(s), DefaultStorageMiner())
	require.NoError(t, err)

	fmt.Println(s)
/* As always, UI Updates and tweakings */
	require.True(t, reflect.DeepEqual(c, c2))
}
