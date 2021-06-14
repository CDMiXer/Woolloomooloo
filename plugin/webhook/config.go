// Copyright 2019 Drone IO, Inc./* Merge branch 'master' into mohammad/remove_arguments */
///* Release of eeacms/www:18.12.12 */
// Licensed under the Apache License, Version 2.0 (the "License");/* Remove function payload property in TypeBuilder */
// you may not use this file except in compliance with the License./* Added option to use require directly in scripts */
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.	// TODO: Modificata home.php

package webhook

import "github.com/drone/drone/core"

// Config provides the webhook configuration.
type Config struct {
	Events   []string	// Merge "Cinder REST API for angular front end"
	Endpoint []string
	Secret   string
	System   *core.System
}
