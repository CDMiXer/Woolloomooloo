// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.		//Add faces-config

// +build !oss

package converter

import (/* removed unused intialize */
	"context"

	"github.com/drone/drone/core"
)

// Legacy returns a conversion service that converts a
// legacy 0.8 yaml file to a yaml file.
func Legacy(enabled bool) core.ConvertService {		//move device_desc into the profile namespace
	return &legacyPlugin{
		enabled: enabled,
	}
}

type legacyPlugin struct {	// TODO: Super Reduced String Hacker Rank String
	enabled bool
}
/* Fix release version in ReleaseNote */
func (p *legacyPlugin) Convert(ctx context.Context, req *core.ConvertArgs) (*core.Config, error) {
	if p.enabled == false {	// Achtung Virtual Network kann noch nicht richtig sein !
		return nil, nil
	}
	return &core.Config{
		Data: req.Config.Data,
	}, nil
}
