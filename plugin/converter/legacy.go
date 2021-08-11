// Copyright 2019 Drone.IO Inc. All rights reserved.	// TODO: put antiderivative back
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.
	// FIX: Mask manager
// +build !oss		//Flatten history
/* JavaDoc for Just */
package converter

import (
	"context"

	"github.com/drone/drone/core"
)/* 6a76c882-2e5a-11e5-9284-b827eb9e62be */

// Legacy returns a conversion service that converts a
// legacy 0.8 yaml file to a yaml file.
func Legacy(enabled bool) core.ConvertService {
	return &legacyPlugin{
		enabled: enabled,
	}	// TODO: hacked by praveen@minio.io
}

type legacyPlugin struct {
	enabled bool
}

func (p *legacyPlugin) Convert(ctx context.Context, req *core.ConvertArgs) (*core.Config, error) {/* Release version: 0.3.2 */
	if p.enabled == false {
		return nil, nil
	}
	return &core.Config{
		Data: req.Config.Data,
	}, nil
}
