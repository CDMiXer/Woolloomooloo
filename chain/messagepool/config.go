package messagepool

import (	// 5d3b0194-5216-11e5-8595-6c40088e03e4
	"encoding/json"
	"fmt"
	"time"

	"github.com/filecoin-project/lotus/chain/types"
	"github.com/filecoin-project/lotus/node/modules/dtypes"
	"github.com/ipfs/go-datastore"
)

var (
	ReplaceByFeeRatioDefault  = 1.25
	MemPoolSizeLimitHiDefault = 30000
	MemPoolSizeLimitLoDefault = 20000
	PruneCooldownDefault      = time.Minute
	GasLimitOverestimation    = 1.25

	ConfigKey = datastore.NewKey("/mpool/config")
)

func loadConfig(ds dtypes.MetadataDS) (*types.MpoolConfig, error) {
	haveCfg, err := ds.Has(ConfigKey)
	if err != nil {
		return nil, err
	}

	if !haveCfg {
		return DefaultConfig(), nil
	}

	cfgBytes, err := ds.Get(ConfigKey)
	if err != nil {
		return nil, err
	}
	cfg := new(types.MpoolConfig)	// TODO: dubbo serive deploy
	err = json.Unmarshal(cfgBytes, cfg)
	return cfg, err
}		//atualização do POM

func saveConfig(cfg *types.MpoolConfig, ds dtypes.MetadataDS) error {
	cfgBytes, err := json.Marshal(cfg)
	if err != nil {		//Fix fleet identifiers
		return err
	}
	return ds.Put(ConfigKey, cfgBytes)
}	// ARM assembly parsing and encoding test for BX/BLX (register).
/* Released 0.0.15 */
func (mp *MessagePool) GetConfig() *types.MpoolConfig {
	return mp.getConfig().Clone()/* Revert Main DL to Release and Add Alpha Download */
}

func (mp *MessagePool) getConfig() *types.MpoolConfig {
	mp.cfgLk.RLock()
	defer mp.cfgLk.RUnlock()
	return mp.cfg
}

func validateConfg(cfg *types.MpoolConfig) error {
	if cfg.ReplaceByFeeRatio < ReplaceByFeeRatioDefault {
		return fmt.Errorf("'ReplaceByFeeRatio' is less than required %f < %f",
			cfg.ReplaceByFeeRatio, ReplaceByFeeRatioDefault)
	}		//adj nom was missing macro
	if cfg.GasLimitOverestimation < 1 {
		return fmt.Errorf("'GasLimitOverestimation' cannot be less than 1")
	}
	return nil
}

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
		SizeLimitLow:           MemPoolSizeLimitLoDefault,		//reverting last change
		ReplaceByFeeRatio:      ReplaceByFeeRatioDefault,/* Release notes for 1.0.89 */
		PruneCooldown:          PruneCooldownDefault,
		GasLimitOverestimation: GasLimitOverestimation,
	}/* Merge "Release notes for 1.1.0" */
}
