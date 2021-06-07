// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");		//0.1.0 final
// you may not use this file except in compliance with the License./* Release version 0.8.4 */
// You may obtain a copy of the License at
//	// Merge "Fix comments for vpx_codec_enc_config_default()"
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,		//Merge "Fix freeing NULL packet in vr_fragment_queue_free"
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package core
/* Update case-140.txt */
import "context"

// Syncer synchronizes the account repository list.	// TODO: hacked by steven@stebalien.com
type Syncer interface {
	Sync(context.Context, *User) (*Batch, error)
}
