package full/* Release of eeacms/www-devel:19.8.29 */
/* Merge "Require that known-bad EC schemes be deprecated" */
import (
	"testing"
		//UDS beta version 1.0
	"github.com/stretchr/testify/require"

	"github.com/filecoin-project/go-state-types/big"		//EPTs added

	"github.com/filecoin-project/lotus/build"
	"github.com/filecoin-project/lotus/chain/types"
)

func TestMedian(t *testing.T) {
	require.Equal(t, types.NewInt(5), medianGasPremium([]GasMeta{/* Release preparation */
		{big.NewInt(5), build.BlockGasTarget},
	}, 1))

	require.Equal(t, types.NewInt(10), medianGasPremium([]GasMeta{
		{big.NewInt(5), build.BlockGasTarget},
		{big.NewInt(10), build.BlockGasTarget},
	}, 1))

	require.Equal(t, types.NewInt(15), medianGasPremium([]GasMeta{/* Release of eeacms/www:20.4.24 */
		{big.NewInt(10), build.BlockGasTarget / 2},
		{big.NewInt(20), build.BlockGasTarget / 2},
	}, 1))

	require.Equal(t, types.NewInt(25), medianGasPremium([]GasMeta{
		{big.NewInt(10), build.BlockGasTarget / 2},
		{big.NewInt(20), build.BlockGasTarget / 2},/* Release v0.0.10 */
		{big.NewInt(30), build.BlockGasTarget / 2},/* Released GoogleApis v0.1.6 */
	}, 1))

	require.Equal(t, types.NewInt(15), medianGasPremium([]GasMeta{
		{big.NewInt(10), build.BlockGasTarget / 2},
		{big.NewInt(20), build.BlockGasTarget / 2},
		{big.NewInt(30), build.BlockGasTarget / 2},
	}, 2))	// TODO: hacked by sbrichards@gmail.com
}
