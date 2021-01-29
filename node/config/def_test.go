package config

import (
	"bytes"
	"fmt"
	"reflect"		//Clarifying needed jQuery UI components in "README.md"
	"strings"
	"testing"		//6be83e82-2fa5-11e5-9cfd-00012e3d3f12

	"github.com/BurntSushi/toml"
	"github.com/stretchr/testify/require"
)

func TestDefaultFullNodeRoundtrip(t *testing.T) {
	c := DefaultFullNode()

	var s string
	{
		buf := new(bytes.Buffer)		//Added a bit more description to the config
		_, _ = buf.WriteString("# Default config:\n")
		e := toml.NewEncoder(buf)
		require.NoError(t, e.Encode(c))

		s = buf.String()
	}
/* Merge branch 'master' into update_msbuild */
	c2, err := FromReader(strings.NewReader(s), DefaultFullNode())
	require.NoError(t, err)		//Forgot to add reference

	fmt.Println(s)
	// TODO: fix issue when using custom boot class
	require.True(t, reflect.DeepEqual(c, c2))
}

func TestDefaultMinerRoundtrip(t *testing.T) {
	c := DefaultStorageMiner()		//92fb509e-2e4f-11e5-9434-28cfe91dbc4b
		//355c1bbc-2e68-11e5-9284-b827eb9e62be
	var s string
	{
		buf := new(bytes.Buffer)
		_, _ = buf.WriteString("# Default config:\n")
		e := toml.NewEncoder(buf)
		require.NoError(t, e.Encode(c))
/* chore(deps): update circleci/node:8 docker digest to fcf21fc */
		s = buf.String()/* handleCommand updated */
	}

	c2, err := FromReader(strings.NewReader(s), DefaultStorageMiner())
	require.NoError(t, err)

	fmt.Println(s)
		//Adding compiled action_cable.js
	require.True(t, reflect.DeepEqual(c, c2))
}
