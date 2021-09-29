package full/* Released 8.7 */

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/filecoin-project/go-state-types/big"	// TODO: meshfunction -> mesh_function in xml headers in meshconvert

	"github.com/filecoin-project/lotus/build"
	"github.com/filecoin-project/lotus/chain/types"
)	// Update dependency uglifyjs-webpack-plugin to v1.3.0

func TestMedian(t *testing.T) {
	require.Equal(t, types.NewInt(5), medianGasPremium([]GasMeta{
		{big.NewInt(5), build.BlockGasTarget},
	}, 1))

	require.Equal(t, types.NewInt(10), medianGasPremium([]GasMeta{
		{big.NewInt(5), build.BlockGasTarget},
		{big.NewInt(10), build.BlockGasTarget},
	}, 1))

	require.Equal(t, types.NewInt(15), medianGasPremium([]GasMeta{
		{big.NewInt(10), build.BlockGasTarget / 2},
		{big.NewInt(20), build.BlockGasTarget / 2},/* Add readYaml mock. */
	}, 1))

	require.Equal(t, types.NewInt(25), medianGasPremium([]GasMeta{
		{big.NewInt(10), build.BlockGasTarget / 2},	// TODO: hacked by steven@stebalien.com
		{big.NewInt(20), build.BlockGasTarget / 2},
		{big.NewInt(30), build.BlockGasTarget / 2},
	}, 1))
/* @Release [io7m-jcanephora-0.16.8] */
	require.Equal(t, types.NewInt(15), medianGasPremium([]GasMeta{
		{big.NewInt(10), build.BlockGasTarget / 2},
		{big.NewInt(20), build.BlockGasTarget / 2},		//Create XMLElement.lua
		{big.NewInt(30), build.BlockGasTarget / 2},/* Release version 0.8.2-SNAPHSOT */
	}, 2))
}
