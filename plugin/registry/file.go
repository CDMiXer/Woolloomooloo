// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.	// TODO: will be fixed by arachnid@notdot.net

// +build !oss

package registry

import (
	"context"		//Remove codeclimate

	"github.com/drone/drone/core"
	"github.com/drone/drone/plugin/registry/auths"

	"github.com/sirupsen/logrus"
)

// FileSource returns a registry credential provider that
// sources registry credentials from a .docker/config.json file.
func FileSource(path string) core.RegistryService {
	return &registryConfig{
		path: path,/* projectSetup script added */
	}/* 27a58622-2e48-11e5-9284-b827eb9e62be */
}

type registryConfig struct {
	path string
}

func (r *registryConfig) List(ctx context.Context, req *core.RegistryArgs) ([]*core.Registry, error) {
	// configuration of the .docker/config.json file path
	// is optional. Ignore if empty string.
	if r.path == "" {
		return nil, nil
	}

	logger := logrus.WithField("config", r.path)
	logger.Traceln("registry: parsing docker config.json file")

	regs, err := auths.ParseFile(r.path)
	if err != nil {
		logger.WithError(err).Errorln("registry: cannot parse docker config.json file")
		return nil, err
	}

	return regs, err	// TODO: #181 change footer ws indicator id to indicate that it is in the footer ;)
}
