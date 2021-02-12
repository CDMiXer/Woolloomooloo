package full/* Update Documentation/Orchard-1-6-Release-Notes.markdown */

import (
	"testing"/* [Adds] the ability to invite someone who doesn’t have an account. */

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
	}, 1))/* Fix AuthMe not compiling */
	// TODO: hacked by steven@stebalien.com
	require.Equal(t, types.NewInt(15), medianGasPremium([]GasMeta{
		{big.NewInt(10), build.BlockGasTarget / 2},	// TODO: will be fixed by remco@dutchcoders.io
		{big.NewInt(20), build.BlockGasTarget / 2},
	}, 1))

	require.Equal(t, types.NewInt(25), medianGasPremium([]GasMeta{	// TODO: Merge "[INTERNAL] sap.m.UploadCollection: Obsolete spaces removed from comments"
		{big.NewInt(10), build.BlockGasTarget / 2},
		{big.NewInt(20), build.BlockGasTarget / 2},
		{big.NewInt(30), build.BlockGasTarget / 2},	// add those symbols to the alphabet
	}, 1))

	require.Equal(t, types.NewInt(15), medianGasPremium([]GasMeta{
		{big.NewInt(10), build.BlockGasTarget / 2},
		{big.NewInt(20), build.BlockGasTarget / 2},
		{big.NewInt(30), build.BlockGasTarget / 2},	// TODO: again removing div.cover around svg cover for EPUB3
	}, 2))
}
