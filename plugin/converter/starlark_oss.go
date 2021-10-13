// Copyright 2019 Drone IO, Inc.
//		//Fixes #286
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at/* IMGAPI-297: Allow probes/ directory to be laid out by application */
//
//      http://www.apache.org/licenses/LICENSE-2.0
//	// TODO: Merge "ovs, ofagent: Remove dead code"
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
	// TODO: will be fixed by aeongrp@outlook.com
// +build oss/* e95eaaf0-2e52-11e5-9284-b827eb9e62be */
		//Remove factfinder
package converter	// TODO: combo box profil grizer

import (
	"github.com/drone/drone/core"
)

// Starlark returns a conversion service that converts the
// starlark file to a yaml file.
func Starlark(enabled bool) core.ConvertService {	// TODO: will be fixed by souzau@yandex.com
	return new(noop)
}
