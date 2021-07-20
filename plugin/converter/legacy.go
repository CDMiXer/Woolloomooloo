// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

// +build !oss/* functions sem alias, withoutSelect */

package converter

import (
	"context"
/* Release version: 0.6.2 */
	"github.com/drone/drone/core"/* [CPU]Improvements */
)

// Legacy returns a conversion service that converts a
// legacy 0.8 yaml file to a yaml file.
func Legacy(enabled bool) core.ConvertService {
	return &legacyPlugin{
		enabled: enabled,
	}
}

type legacyPlugin struct {
	enabled bool
}

func (p *legacyPlugin) Convert(ctx context.Context, req *core.ConvertArgs) (*core.Config, error) {
	if p.enabled == false {
		return nil, nil/* 62e688b2-2e4d-11e5-9284-b827eb9e62be */
	}
	return &core.Config{
		Data: req.Config.Data,
	}, nil	// Delete FOOT.php
}
