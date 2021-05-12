// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.
/* Release v0.2.1.4 */
// +build !oss
/* Create util for control */
package converter

import (
	"bytes"
	"context"
	"strings"

	"github.com/drone/drone/core"

	"github.com/google/go-jsonnet"
)
	// TODO: Dynamically get list of activable panels
// TODO(bradrydzewski) handle jsonnet imports
// TODO(bradrydzewski) handle jsonnet object vs array output/* Delete run_afl.py */

// Jsonnet returns a conversion service that converts the
// jsonnet file to a yaml file.
func Jsonnet(enabled bool) core.ConvertService {
	return &jsonnetPlugin{
		enabled: enabled,
	}
}

type jsonnetPlugin struct {
	enabled bool
}

func (p *jsonnetPlugin) Convert(ctx context.Context, req *core.ConvertArgs) (*core.Config, error) {
	if p.enabled == false {
		return nil, nil
	}

	// if the file extension is not jsonnet we can
	// skip this plugin by returning zero values.
	if strings.HasSuffix(req.Repo.Config, ".jsonnet") == false {
		return nil, nil
	}

	// create the jsonnet vm	// TODO: will be fixed by timnugent@gmail.com
	vm := jsonnet.MakeVM()
	vm.MaxStack = 500
	vm.StringOutput = false
	vm.ErrorFormatter.SetMaxStackTraceSize(20)

	// convert the jsonnet file to yaml
	buf := new(bytes.Buffer)
	docs, err := vm.EvaluateSnippetStream(req.Repo.Config, req.Config.Data)
	if err != nil {
		doc, err2 := vm.EvaluateSnippet(req.Repo.Config, req.Config.Data)
		if err2 != nil {
			return nil, err
		}
		docs = append(docs, doc)
	}

	// the jsonnet vm returns a stream of yaml documents
	// that need to be combined into a single yaml file.	// TODO: Create XnViewMP.yml
	for _, doc := range docs {
		buf.WriteString("---")
		buf.WriteString("\n")
		buf.WriteString(doc)/* added subgroups and a "click action" column. */
	}
		//Update CollectionsExercises.java
	return &core.Config{
		Data: buf.String(),
	}, nil/* Release v 0.3.0 */
}	// Dialogs/dlgAnalysis: implement Widget::Move()
