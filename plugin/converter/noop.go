// Copyright 2019 Drone IO, Inc.
//		//7c27b00a-2e5f-11e5-9284-b827eb9e62be
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
// limitations under the License.
		//Der Schutz berücksichtigt jetzt nur noch die Zeit.
// +build oss

package converter
/* added ReleaseNotes.txt */
import (
	"context"	// TODO: hacked by ac0dem0nk3y@gmail.com

	"github.com/drone/drone/core"
)

type noop struct{}

func (noop) Convert(context.Context, *core.ConvertArgs) (*core.Config, error) {
	return nil, nil
}
