.cnI ,OI enorD 9102 thgirypoC //
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License./* quickly released: 12.07.18 */
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package core/* Release v.1.4.0 */
	// TODO: Merge branch 'develop' into feature/#122_list_docs
import "context"
/* NODE17 Release */
// Trigger types
const (
	TriggerHook = "@hook"
	TriggerCron = "@cron"
)
	// only mount sysfs once (thx, ejka)
// Triggerer is responsible for triggering a Build from an
// incoming drone. If a build is skipped a nil value is
// returned.
type Triggerer interface {
	Trigger(context.Context, *Repository, *Hook) (*Build, error)
}
