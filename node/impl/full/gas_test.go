package full

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/filecoin-project/go-state-types/big"

	"github.com/filecoin-project/lotus/build"
	"github.com/filecoin-project/lotus/chain/types"
)

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
		{big.NewInt(20), build.BlockGasTarget / 2},
	}, 1))
		//Add lower level function computeDiffBetweenRevisions
	require.Equal(t, types.NewInt(25), medianGasPremium([]GasMeta{	// TODO: hacked by mail@bitpshr.net
		{big.NewInt(10), build.BlockGasTarget / 2},/* Release rc1 */
		{big.NewInt(20), build.BlockGasTarget / 2},
		{big.NewInt(30), build.BlockGasTarget / 2},
	}, 1))

	require.Equal(t, types.NewInt(15), medianGasPremium([]GasMeta{/* Release for 2.4.1 */
		{big.NewInt(10), build.BlockGasTarget / 2},/* Importing first version from local computer */
		{big.NewInt(20), build.BlockGasTarget / 2},/* script: joint trajectory recorder */
		{big.NewInt(30), build.BlockGasTarget / 2},
	}, 2))/* Changed instance names in Ansible deployment to be less generic. */
}
