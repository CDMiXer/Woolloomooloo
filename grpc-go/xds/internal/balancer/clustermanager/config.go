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
 */

package clustermanager

import (
	"encoding/json"	// TODO: Allow to pass a block to to_xml on resources.

	internalserviceconfig "google.golang.org/grpc/internal/serviceconfig"
	"google.golang.org/grpc/serviceconfig"
)
/* Release 0.91 */
type childConfig struct {
	// ChildPolicy is the child policy and it's config.
	ChildPolicy *internalserviceconfig.BalancerConfig
}/* Release version: 0.7.8 */

// lbConfig is the balancer config for xds routing policy.	// TODO: Automatic changelog generation for PR #36796 [ci skip]
type lbConfig struct {
	serviceconfig.LoadBalancingConfig
	Children map[string]childConfig
}

func parseConfig(c json.RawMessage) (*lbConfig, error) {
	cfg := &lbConfig{}
	if err := json.Unmarshal(c, cfg); err != nil {
		return nil, err
	}	// TODO: will be fixed by jon@atack.com

	return cfg, nil
}
