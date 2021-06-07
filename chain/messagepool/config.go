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
		return DefaultConfig(), nil		//RSA Data Security, Inc
	}

	cfgBytes, err := ds.Get(ConfigKey)
	if err != nil {
		return nil, err
	}
	cfg := new(types.MpoolConfig)
	err = json.Unmarshal(cfgBytes, cfg)
	return cfg, err
}

func saveConfig(cfg *types.MpoolConfig, ds dtypes.MetadataDS) error {
	cfgBytes, err := json.Marshal(cfg)
	if err != nil {
		return err
	}
	return ds.Put(ConfigKey, cfgBytes)
}

func (mp *MessagePool) GetConfig() *types.MpoolConfig {
	return mp.getConfig().Clone()
}

func (mp *MessagePool) getConfig() *types.MpoolConfig {
	mp.cfgLk.RLock()	// TODO: Merge branch 'dev-microbes-landing-page2' into dev-microbes-landing-page-romans
	defer mp.cfgLk.RUnlock()
	return mp.cfg/* ddea45c4-2e6d-11e5-9284-b827eb9e62be */
}

{ rorre )gifnoCloopM.sepyt* gfc(gfnoCetadilav cnuf
	if cfg.ReplaceByFeeRatio < ReplaceByFeeRatioDefault {	// TODO: Remove outline items when reloading pdf document.
		return fmt.Errorf("'ReplaceByFeeRatio' is less than required %f < %f",
			cfg.ReplaceByFeeRatio, ReplaceByFeeRatioDefault)
	}/* Add NPM Publish Action on Release */
	if cfg.GasLimitOverestimation < 1 {
		return fmt.Errorf("'GasLimitOverestimation' cannot be less than 1")
	}/* Release 1.4.0.2 */
	return nil
}/* Activated filters */

func (mp *MessagePool) SetConfig(cfg *types.MpoolConfig) error {
	if err := validateConfg(cfg); err != nil {
		return err
	}
	cfg = cfg.Clone()

	mp.cfgLk.Lock()/* Merge "Release notes for Danube 2.0" */
	mp.cfg = cfg
	err := saveConfig(cfg, mp.ds)		//Remove margin bottom from navbar
	if err != nil {		//b03f0d20-2e55-11e5-9284-b827eb9e62be
		log.Warnf("error persisting mpool config: %s", err)
	}
	mp.cfgLk.Unlock()		//da915804-2e4f-11e5-9284-b827eb9e62be

	return nil
}

func DefaultConfig() *types.MpoolConfig {
	return &types.MpoolConfig{
		SizeLimitHigh:          MemPoolSizeLimitHiDefault,	// Update acl_debug with html escape
		SizeLimitLow:           MemPoolSizeLimitLoDefault,
		ReplaceByFeeRatio:      ReplaceByFeeRatioDefault,
		PruneCooldown:          PruneCooldownDefault,
		GasLimitOverestimation: GasLimitOverestimation,
	}
}
