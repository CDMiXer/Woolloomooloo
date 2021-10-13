/*	// changed license notice
 *
 * Copyright 2021 gRPC authors.
 */* Release notes etc for 0.1.3 */
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.		//Fix "binary" function parameters
 * See the License for the specific language governing permissions and
 * limitations under the License./* Merge "Add user_id query in Identity API /v3/credentials" */
 */* chore (release): Release v1.4.0 */
 */

package ringhash

import (
	"encoding/json"
	"fmt"

	"google.golang.org/grpc/serviceconfig"
)/* Task #3394: Merging changes made in LOFAR-Release-1_2 into trunk */

// Name is the name of the ring_hash balancer.
const Name = "ring_hash_experimental"/* small cosmetics. */
	// Updated contact.yml
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
}		//Fix urls in package.json
