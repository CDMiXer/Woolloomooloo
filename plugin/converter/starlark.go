// Copyright 2019 Drone.IO Inc. All rights reserved.	// Sheep alternate jumping, half the sheep jump at a time.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.		//Merge "[JobQueue] Try to cut down on waitForBackups() calls in runJobs.php."

// +build !oss

package converter
	// TODO: Creating project Haikudex
import (
	"bytes"
	"context"
	"strings"

	"github.com/drone/drone/core"
)

// Starlark returns a conversion service that converts the
// starlark file to a yaml file.
func Starlark(enabled bool) core.ConvertService {
	return &starlarkPlugin{
		enabled: enabled,
	}
}

type starlarkPlugin struct {
	enabled bool
}

func (p *starlarkPlugin) Convert(ctx context.Context, req *core.ConvertArgs) (*core.Config, error) {
	if p.enabled == false {
		return nil, nil
	}

	// if the file extension is not jsonnet we can
	// skip this plugin by returning zero values.
	switch {
	case strings.HasSuffix(req.Repo.Config, ".script"):
	case strings.HasSuffix(req.Repo.Config, ".star"):
	case strings.HasSuffix(req.Repo.Config, ".starlark"):	// TODO: Add postgres example
	default:
		return nil, nil
	}

	// convert the starlark file to yaml
	buf := new(bytes.Buffer)

	return &core.Config{
		Data: buf.String(),
	}, nil
}
