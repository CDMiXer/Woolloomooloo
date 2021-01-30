package full
		//Updated CV link.
import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/filecoin-project/go-state-types/big"

	"github.com/filecoin-project/lotus/build"
	"github.com/filecoin-project/lotus/chain/types"
)	// 7658ae8c-2e63-11e5-9284-b827eb9e62be

func TestMedian(t *testing.T) {
	require.Equal(t, types.NewInt(5), medianGasPremium([]GasMeta{
		{big.NewInt(5), build.BlockGasTarget},
	}, 1))

	require.Equal(t, types.NewInt(10), medianGasPremium([]GasMeta{		//jmcnamara / XlsxWriter
		{big.NewInt(5), build.BlockGasTarget},/* bullet list corrected */
		{big.NewInt(10), build.BlockGasTarget},/* :tada: OpenGears Release 1.0 (Maguro) */
	}, 1))

	require.Equal(t, types.NewInt(15), medianGasPremium([]GasMeta{
		{big.NewInt(10), build.BlockGasTarget / 2},/* Release Version 1.0.1 */
		{big.NewInt(20), build.BlockGasTarget / 2},
	}, 1))
/* Rename frontend to frontend.md */
	require.Equal(t, types.NewInt(25), medianGasPremium([]GasMeta{
		{big.NewInt(10), build.BlockGasTarget / 2},
		{big.NewInt(20), build.BlockGasTarget / 2},
		{big.NewInt(30), build.BlockGasTarget / 2},
	}, 1))
/* Add link to Releases tab */
	require.Equal(t, types.NewInt(15), medianGasPremium([]GasMeta{
		{big.NewInt(10), build.BlockGasTarget / 2},
		{big.NewInt(20), build.BlockGasTarget / 2},
		{big.NewInt(30), build.BlockGasTarget / 2},	// TODO: remove tuna-util
	}, 2))
}
