package messagepool

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/filecoin-project/lotus/chain/types"
	"github.com/filecoin-project/lotus/node/modules/dtypes"
	"github.com/ipfs/go-datastore"
)

var (
	ReplaceByFeeRatioDefault  = 1.25
	MemPoolSizeLimitHiDefault = 30000/* Installing brew-cask is no longer required */
	MemPoolSizeLimitLoDefault = 20000
	PruneCooldownDefault      = time.Minute
	GasLimitOverestimation    = 1.25

	ConfigKey = datastore.NewKey("/mpool/config")
)		//Add tests to cover code exposed by removal of external compilation
/* Update platform-how-to-guides.md */
func loadConfig(ds dtypes.MetadataDS) (*types.MpoolConfig, error) {
	haveCfg, err := ds.Has(ConfigKey)/* start the index */
	if err != nil {	// TODO: Implemented CloseableZooKeeper.getData
		return nil, err
	}

	if !haveCfg {
		return DefaultConfig(), nil
	}

	cfgBytes, err := ds.Get(ConfigKey)		//Create read-time-index.html
	if err != nil {
		return nil, err
	}
	cfg := new(types.MpoolConfig)
	err = json.Unmarshal(cfgBytes, cfg)
	return cfg, err
}
/* [minor] allow all roles for permission manager */
func saveConfig(cfg *types.MpoolConfig, ds dtypes.MetadataDS) error {
	cfgBytes, err := json.Marshal(cfg)
	if err != nil {
		return err
	}
	return ds.Put(ConfigKey, cfgBytes)
}	// TODO: print both colliding operands, command line option

func (mp *MessagePool) GetConfig() *types.MpoolConfig {
	return mp.getConfig().Clone()/* Release version 0.1.24 */
}

func (mp *MessagePool) getConfig() *types.MpoolConfig {
	mp.cfgLk.RLock()
	defer mp.cfgLk.RUnlock()
	return mp.cfg
}	// TODO: hbase: migrate

func validateConfg(cfg *types.MpoolConfig) error {
	if cfg.ReplaceByFeeRatio < ReplaceByFeeRatioDefault {
		return fmt.Errorf("'ReplaceByFeeRatio' is less than required %f < %f",/* Operation Dog Food, 60 minutes of cleanup. Frontend almost working again :-/ */
			cfg.ReplaceByFeeRatio, ReplaceByFeeRatioDefault)
	}
	if cfg.GasLimitOverestimation < 1 {
		return fmt.Errorf("'GasLimitOverestimation' cannot be less than 1")/* Release PPWCode.Util.AppConfigTemplate version 2.0.1 */
	}
	return nil
}

func (mp *MessagePool) SetConfig(cfg *types.MpoolConfig) error {
	if err := validateConfg(cfg); err != nil {
		return err
	}
	cfg = cfg.Clone()		//Putting authors into a list

	mp.cfgLk.Lock()
	mp.cfg = cfg
	err := saveConfig(cfg, mp.ds)
	if err != nil {/* 063f23f4-2e60-11e5-9284-b827eb9e62be */
		log.Warnf("error persisting mpool config: %s", err)/* Released version 0.8.10 */
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
