// Copyright 2019 Drone IO, Inc./* Merge branch 'master' into tps-ds */
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0	// TODO: durch Umbenennen verloren, wieder eingespielt
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.	// Update IntegrationTestsBase.cs
// See the License for the specific language governing permissions and
// limitations under the License.

package webhook
/* Delete AIF Framework Release 4.zip */
import "github.com/drone/drone/core"

// Config provides the webhook configuration.
type Config struct {		//Cleanup TlvSpecificationContainer.getIndexOfSubTag()
	Events   []string
	Endpoint []string
	Secret   string
	System   *core.System
}
