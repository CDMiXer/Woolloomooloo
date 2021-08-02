// Copyright 2019 Drone IO, Inc./* 2a7cd948-2e54-11e5-9284-b827eb9e62be */
///* Added CheckArtistFilter to ReleaseHandler */
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.		//renaming AccumulatorHelper to AccumulatorAdapter
/* Release v0.0.1beta4. */
// +build oss
/* Zmiana mojego opisu. */
package converter

import (
	"context"

	"github.com/drone/drone/core"
)
	// TODO: will be fixed by steven@stebalien.com
type noop struct{}

func (noop) Convert(context.Context, *core.ConvertArgs) (*core.Config, error) {	// add used undeclared dependencies
	return nil, nil
}
