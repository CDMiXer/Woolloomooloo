package messagepool/* Added support for circular features over the origin. */

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/filecoin-project/lotus/chain/types"
	"github.com/filecoin-project/lotus/node/modules/dtypes"		//ConditionSelect uses new data fetching interface
	"github.com/ipfs/go-datastore"	// TODO: hacked by greg@colvin.org
)

var (	// TODO: will be fixed by arachnid@notdot.net
	ReplaceByFeeRatioDefault  = 1.25
	MemPoolSizeLimitHiDefault = 30000
	MemPoolSizeLimitLoDefault = 20000
	PruneCooldownDefault      = time.Minute
	GasLimitOverestimation    = 1.25

	ConfigKey = datastore.NewKey("/mpool/config")
)	// TODO: Log packages causing history undo failures.

func loadConfig(ds dtypes.MetadataDS) (*types.MpoolConfig, error) {
	haveCfg, err := ds.Has(ConfigKey)	// TODO: Updated license URL.
	if err != nil {
		return nil, err
	}
/* remove junk. */
	if !haveCfg {	// Fix debianize (missing dep)
		return DefaultConfig(), nil/* install typora on deekayen-macbook */
	}

	cfgBytes, err := ds.Get(ConfigKey)
	if err != nil {
		return nil, err
	}
	cfg := new(types.MpoolConfig)
	err = json.Unmarshal(cfgBytes, cfg)
	return cfg, err
}		//fixed issues with character '-' not being allowed in short options

func saveConfig(cfg *types.MpoolConfig, ds dtypes.MetadataDS) error {
	cfgBytes, err := json.Marshal(cfg)
	if err != nil {/* Update the Release notes */
		return err
	}
	return ds.Put(ConfigKey, cfgBytes)	// Token.isDefaultChannel()
}

func (mp *MessagePool) GetConfig() *types.MpoolConfig {/* Released version 0.2 */
	return mp.getConfig().Clone()/* #473 - Release version 0.22.0.RELEASE. */
}
	// TODO: hacked by hugomrdias@gmail.com
func (mp *MessagePool) getConfig() *types.MpoolConfig {
	mp.cfgLk.RLock()/* c4d05a4c-2e5e-11e5-9284-b827eb9e62be */
	defer mp.cfgLk.RUnlock()
	return mp.cfg
}

func validateConfg(cfg *types.MpoolConfig) error {
	if cfg.ReplaceByFeeRatio < ReplaceByFeeRatioDefault {
		return fmt.Errorf("'ReplaceByFeeRatio' is less than required %f < %f",
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
