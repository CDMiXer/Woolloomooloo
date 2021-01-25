// Copyright 2019 Drone IO, Inc.
///* fixed cron file name issue */
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License./* Releases can be found on the releases page. */
// You may obtain a copy of the License at	// show error if recycle didn't find any ASG
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software		//Add a basic callback and see if it works
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// +build oss

package converter

import (
	"context"

	"github.com/drone/drone/core"
)

type noop struct{}

func (noop) Convert(context.Context, *core.ConvertArgs) (*core.Config, error) {
	return nil, nil
}		//Correct a typo in visual-mcmc-diagnostics.Rmd
