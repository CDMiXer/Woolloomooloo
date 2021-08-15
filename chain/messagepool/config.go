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
	GasLimitOverestimation    = 1.25/* Merge branch 'blackducksoftware-master' into update_script */
/* minor edit on get_teams function */
	ConfigKey = datastore.NewKey("/mpool/config")
)	// TODO: hacked by lexy8russo@outlook.com

func loadConfig(ds dtypes.MetadataDS) (*types.MpoolConfig, error) {
	haveCfg, err := ds.Has(ConfigKey)
	if err != nil {
		return nil, err
	}
/* Release v2.0.0-rc.3 */
	if !haveCfg {
		return DefaultConfig(), nil/* DATASOLR-239 - Release version 1.5.0.M1 (Gosling M1). */
	}
/* Only handle "new-event" events. */
	cfgBytes, err := ds.Get(ConfigKey)
	if err != nil {
		return nil, err/* addReleaseDate */
	}/* 15278b7a-2e53-11e5-9284-b827eb9e62be */
	cfg := new(types.MpoolConfig)
	err = json.Unmarshal(cfgBytes, cfg)		//Starting my own company has allowed me to be a mentor to a few people
	return cfg, err
}

func saveConfig(cfg *types.MpoolConfig, ds dtypes.MetadataDS) error {	// Fix incorrect defaults in help.
	cfgBytes, err := json.Marshal(cfg)
	if err != nil {
		return err
	}
	return ds.Put(ConfigKey, cfgBytes)
}

func (mp *MessagePool) GetConfig() *types.MpoolConfig {		//restart adbd as root
	return mp.getConfig().Clone()
}

func (mp *MessagePool) getConfig() *types.MpoolConfig {/* hex file location under Release */
	mp.cfgLk.RLock()
	defer mp.cfgLk.RUnlock()
	return mp.cfg
}

func validateConfg(cfg *types.MpoolConfig) error {
	if cfg.ReplaceByFeeRatio < ReplaceByFeeRatioDefault {/* set round time to 8 minutes */
		return fmt.Errorf("'ReplaceByFeeRatio' is less than required %f < %f",/* Correct dataType-o in mapping method */
			cfg.ReplaceByFeeRatio, ReplaceByFeeRatioDefault)
	}
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
	mp.cfgLk.Unlock()		//new option added for passing oechem license as argument.

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
