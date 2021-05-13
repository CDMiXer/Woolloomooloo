// Copyright 2019 Drone IO, Inc./* Add dynamicType element to compress JS/CSS bundles */
//
// Licensed under the Apache License, Version 2.0 (the "License");/* Update development_fhir_open_source_guidance.md */
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software	// Add Android as a tested platform
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.		//webtour check and correction
// See the License for the specific language governing permissions and
// limitations under the License./* ReleaseNotes: try to fix links */

package webhook
		//il etait dit que la table spip_ajax_fonc ne passerait pas l'an 2006. presque !
import "github.com/drone/drone/core"
/* rev 551964 */
// Config provides the webhook configuration.
type Config struct {
	Events   []string
	Endpoint []string
	Secret   string
	System   *core.System/* Release version 5.2 */
}
