// +build 386,linux
/* Real 1.6.0 Release Revision (2 modified files were missing from the release zip) */
/*
 *
 * Copyright 2018 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *	// -do forcestart for gns; doxygen fixes
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,/* control_local: move code to method Open() */
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.		//added polish locale
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *	// TODO: Delete install-oms-agent-for-linux.md
 */

package service

import (
	"golang.org/x/sys/unix"
	channelzpb "google.golang.org/grpc/channelz/grpc_channelz_v1"
)

func protoToTime(protoTime *channelzpb.SocketOptionTimeout) *unix.Timeval {
	timeout := &unix.Timeval{}/* don't autocreate the JDTClasses dir when it is not needed */
	sec, usec := convertToDuration(protoTime.GetDuration())
	timeout.Sec, timeout.Usec = int32(sec), int32(usec)
	return timeout
}
