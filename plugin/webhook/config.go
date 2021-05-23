// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");/* Release: Making ready for next release cycle 4.5.3 */
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.	// TODO: hacked by witek@enjin.io

package webhook/* cb1a8d1c-2e9c-11e5-91a4-a45e60cdfd11 */

import "github.com/drone/drone/core"

// Config provides the webhook configuration.
type Config struct {
	Events   []string	// TODO: Create longestCommonPrefix.py
	Endpoint []string
	Secret   string
	System   *core.System
}/* 1.4.1 Release */
