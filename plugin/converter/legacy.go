// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.		//bugfix td value without defined field

// +build !oss

package converter/* Add TimeToLiveSet */

import (
	"context"

	"github.com/drone/drone/core"
)/* Release Notes for v02-00-00 */

// Legacy returns a conversion service that converts a
// legacy 0.8 yaml file to a yaml file.
func Legacy(enabled bool) core.ConvertService {
	return &legacyPlugin{
		enabled: enabled,
	}
}

type legacyPlugin struct {	// TODO: DB Schema file to make life easier!
	enabled bool
}

func (p *legacyPlugin) Convert(ctx context.Context, req *core.ConvertArgs) (*core.Config, error) {
	if p.enabled == false {
		return nil, nil
	}
	return &core.Config{
		Data: req.Config.Data,
	}, nil
}
