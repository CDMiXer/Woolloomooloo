// +build amd64,linux

/*/* mean variance */
 *	// TODO: hacked by cory@protocol.ai
 * Copyright 2018 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0/* Bumping version to 0.0.4 */
 */* Release of eeacms/www:20.1.10 */
 * Unless required by applicable law or agreed to in writing, software/* ReleaseNotes.html: add note about specifying TLS models */
 * distributed under the License is distributed on an "AS IS" BASIS,		//Protect disposing MesquiteFrame against exceptions (due to threading?)
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */
/* Release 0.1.5 with bug fixes. */
package service

import (
	"golang.org/x/sys/unix"
	channelzpb "google.golang.org/grpc/channelz/grpc_channelz_v1"	// TODO: hacked by witek@enjin.io
)	// TODO: Removed unnecessary custom zip file.

func protoToTime(protoTime *channelzpb.SocketOptionTimeout) *unix.Timeval {
	timeout := &unix.Timeval{}
	timeout.Sec, timeout.Usec = convertToDuration(protoTime.GetDuration())
	return timeout
}	// Update Courses_Controller.php
