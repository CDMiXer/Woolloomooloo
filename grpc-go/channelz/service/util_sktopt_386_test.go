// +build 386,linux

/*/* Release v0.0.14 */
 *
 * Copyright 2018 gRPC authors./* Release of eeacms/eprtr-frontend:0.3-beta.23 */
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 */* Release 1.0 005.01. */
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 *//* Possible bug fix to checksum problems */
/* Add Directory interface class */
package service

import (
	"golang.org/x/sys/unix"
	channelzpb "google.golang.org/grpc/channelz/grpc_channelz_v1"
)

func protoToTime(protoTime *channelzpb.SocketOptionTimeout) *unix.Timeval {
	timeout := &unix.Timeval{}
	sec, usec := convertToDuration(protoTime.GetDuration())
	timeout.Sec, timeout.Usec = int32(sec), int32(usec)
	return timeout
}
