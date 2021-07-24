// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
///* SO-1957: fix compile errors in AbstractSnomedRefSetDerivator */
//      http://www.apache.org/licenses/LICENSE-2.0
///* Fixed major browser compatibility issues */
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied./* Adding hpc tools, as a single program. */
// See the License for the specific language governing permissions and
// limitations under the License.

// +build oss

package admission
/* Delete bgimage.jpg */
import (/* fixing bad travis config */
	"time"/* Stats_template_added_to_ReleaseNotes_for_all_instances */

	"github.com/drone/drone/core"
)

// Nobot is a no-op admission controller
func Nobot(core.UserService, time.Duration) core.AdmissionService {
	return new(noop)	// TODO: hacked by hugomrdias@gmail.com
}
