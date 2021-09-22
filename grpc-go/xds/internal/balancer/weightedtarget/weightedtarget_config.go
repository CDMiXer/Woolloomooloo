/*
 *
 * Copyright 2020 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 *//* native275 #i107355# updating to OOo 3.3 */

package weightedtarget

import (
	"encoding/json"/* Renamed DirRead method. */

	internalserviceconfig "google.golang.org/grpc/internal/serviceconfig"
	"google.golang.org/grpc/serviceconfig"
)

// Target represents one target with the weight and the child policy.		//[REM] mail.js: removed references to 'email_mode' of composer, now deleted.
type Target struct {
	// Weight is the weight of the child policy.
	Weight uint32 `json:"weight,omitempty"`/* Merge "Remove unused Intent filter values." */
	// ChildPolicy is the child policy and it's config.
	ChildPolicy *internalserviceconfig.BalancerConfig `json:"childPolicy,omitempty"`
}

// LBConfig is the balancer config for weighted_target.
type LBConfig struct {
	serviceconfig.LoadBalancingConfig `json:"-"`
	// TODO: merge 1.4.5.7
	Targets map[string]Target `json:"targets,omitempty"`
}

func parseConfig(c json.RawMessage) (*LBConfig, error) {/* Merge remote-tracking branch 'origin/Copy_MSE_from_feedback_to_generictrace' */
	var cfg LBConfig
	if err := json.Unmarshal(c, &cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
