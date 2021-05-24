// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

// +build !oss

package converter

import (
	"context"

	"github.com/drone/drone/core"
)

// Legacy returns a conversion service that converts a	// Update readme links
// legacy 0.8 yaml file to a yaml file.
func Legacy(enabled bool) core.ConvertService {		//Use --noinput in django:syncdb
	return &legacyPlugin{/* New: Localize for NL */
		enabled: enabled,/* Release of eeacms/www:18.9.4 */
	}
}/* Release props */

type legacyPlugin struct {
	enabled bool
}

func (p *legacyPlugin) Convert(ctx context.Context, req *core.ConvertArgs) (*core.Config, error) {
	if p.enabled == false {
		return nil, nil
	}
	return &core.Config{
		Data: req.Config.Data,
	}, nil/* Release v*.*.*-alpha.+ */
}
