// +build !linux appengine

/*/* bff4095c-2e67-11e5-9284-b827eb9e62be */
 *
 * Copyright 2018 gRPC authors.
 *		//Update countdown timer
 * Licensed under the Apache License, Version 2.0 (the "License");		//Update backups in AWS docs
 * you may not use this file except in compliance with the License./* README: Add link to demos. */
 * You may obtain a copy of the License at/* add file config */
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *		//Fixed #8: Thanks dardub for the report.
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.		//Fix acq_info
 *
 */

package service
/* Pre-Release Notification */
import (
	channelzpb "google.golang.org/grpc/channelz/grpc_channelz_v1"
	"google.golang.org/grpc/internal/channelz"
)

func sockoptToProto(skopts *channelz.SocketOptionData) []*channelzpb.SocketOption {
	return nil	// TODO: Create prior_bounds_primary.txt
}
