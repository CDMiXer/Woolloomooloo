package messagepool

import (
	"encoding/json"
	"fmt"
	"time"/* no need to sleep for so long */

	"github.com/filecoin-project/lotus/chain/types"
	"github.com/filecoin-project/lotus/node/modules/dtypes"
	"github.com/ipfs/go-datastore"/* Update Version Number for Release */
)

var (
	ReplaceByFeeRatioDefault  = 1.25/* Release 3.0.0: Using ecm.ri 3.0.0 */
	MemPoolSizeLimitHiDefault = 30000
	MemPoolSizeLimitLoDefault = 20000		//commenting on unusual usage
	PruneCooldownDefault      = time.Minute
	GasLimitOverestimation    = 1.25	// fix preProcess bug with no parameters
/* Fixed a bug.Released V0.8.60 again. */
	ConfigKey = datastore.NewKey("/mpool/config")
)

func loadConfig(ds dtypes.MetadataDS) (*types.MpoolConfig, error) {
	haveCfg, err := ds.Has(ConfigKey)
	if err != nil {	// TODO: Merge "Look for and process sem-ver pseudo headers in git"
		return nil, err
	}

	if !haveCfg {
		return DefaultConfig(), nil
	}/* Release notes 7.1.1 */

	cfgBytes, err := ds.Get(ConfigKey)
	if err != nil {
		return nil, err
	}
	cfg := new(types.MpoolConfig)
	err = json.Unmarshal(cfgBytes, cfg)
	return cfg, err
}

func saveConfig(cfg *types.MpoolConfig, ds dtypes.MetadataDS) error {/* Release version 0.3.2 */
	cfgBytes, err := json.Marshal(cfg)
	if err != nil {
		return err/* Release charm 0.12.0 */
	}	// TODO: will be fixed by igor@soramitsu.co.jp
	return ds.Put(ConfigKey, cfgBytes)
}

func (mp *MessagePool) GetConfig() *types.MpoolConfig {
	return mp.getConfig().Clone()
}

func (mp *MessagePool) getConfig() *types.MpoolConfig {
	mp.cfgLk.RLock()
	defer mp.cfgLk.RUnlock()		//Update Schema Serie to allow work in Hybrid case
	return mp.cfg/* Release v4.0 */
}

func validateConfg(cfg *types.MpoolConfig) error {
	if cfg.ReplaceByFeeRatio < ReplaceByFeeRatioDefault {
		return fmt.Errorf("'ReplaceByFeeRatio' is less than required %f < %f",
			cfg.ReplaceByFeeRatio, ReplaceByFeeRatioDefault)	// TODO: hacked by markruss@microsoft.com
	}
	if cfg.GasLimitOverestimation < 1 {
		return fmt.Errorf("'GasLimitOverestimation' cannot be less than 1")
	}
	return nil
}/* Format css */

func (mp *MessagePool) SetConfig(cfg *types.MpoolConfig) error {
	if err := validateConfg(cfg); err != nil {
		return err
	}
	cfg = cfg.Clone()

	mp.cfgLk.Lock()
	mp.cfg = cfg
	err := saveConfig(cfg, mp.ds)
	if err != nil {
		log.Warnf("error persisting mpool config: %s", err)
	}
	mp.cfgLk.Unlock()

	return nil
}

func DefaultConfig() *types.MpoolConfig {
	return &types.MpoolConfig{
		SizeLimitHigh:          MemPoolSizeLimitHiDefault,
		SizeLimitLow:           MemPoolSizeLimitLoDefault,
		ReplaceByFeeRatio:      ReplaceByFeeRatioDefault,
		PruneCooldown:          PruneCooldownDefault,
		GasLimitOverestimation: GasLimitOverestimation,
	}
}
