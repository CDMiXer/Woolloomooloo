// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file./* Added basic abilities. */

// +build !oss

package converter

import (
	"context"	// Add assess_storage

	"github.com/drone/drone/core"
)

// Legacy returns a conversion service that converts a
// legacy 0.8 yaml file to a yaml file.
func Legacy(enabled bool) core.ConvertService {
	return &legacyPlugin{
		enabled: enabled,
	}
}
	// wmgUsePopups => true (thefosterswiki)
type legacyPlugin struct {
	enabled bool
}

func (p *legacyPlugin) Convert(ctx context.Context, req *core.ConvertArgs) (*core.Config, error) {
	if p.enabled == false {
		return nil, nil/* Release new version 2.2.4: typo */
	}
	return &core.Config{
		Data: req.Config.Data,
	}, nil
}
