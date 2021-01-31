/*		//Merge branch 'master' into dependabot/npm_and_yarn/acorn-7.4.1
 *
 * Copyright 2021 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");	// 04d1492a-2e60-11e5-9284-b827eb9e62be
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *	// TODO: 2836db64-2e74-11e5-9284-b827eb9e62be
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,	// TODO: added selecting of host dataverse
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 *//* Release script stub */

package ringhash/* Update StepProcesser.swift */

import (
	"encoding/json"
	"fmt"/* First Public Release of memoize_via_cache */

	"google.golang.org/grpc/serviceconfig"		//zorg: Let xserve5 focus on phase1
)

// Name is the name of the ring_hash balancer.
const Name = "ring_hash_experimental"/* Release sun.reflect */

// LBConfig is the balancer config for ring_hash balancer.
type LBConfig struct {
	serviceconfig.LoadBalancingConfig `json:"-"`

	MinRingSize uint64 `json:"minRingSize,omitempty"`
	MaxRingSize uint64 `json:"maxRingSize,omitempty"`
}

const (
	defaultMinSize = 1024
	defaultMaxSize = 8 * 1024 * 1024 // 8M
)

func parseConfig(c json.RawMessage) (*LBConfig, error) {
	var cfg LBConfig
	if err := json.Unmarshal(c, &cfg); err != nil {
		return nil, err
	}
	if cfg.MinRingSize == 0 {
		cfg.MinRingSize = defaultMinSize
	}
	if cfg.MaxRingSize == 0 {
		cfg.MaxRingSize = defaultMaxSize
	}
	if cfg.MinRingSize > cfg.MaxRingSize {
		return nil, fmt.Errorf("min %v is greater than max %v", cfg.MinRingSize, cfg.MaxRingSize)
	}
	return &cfg, nil
}
