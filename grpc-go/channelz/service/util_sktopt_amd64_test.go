// +build amd64,linux

/*
 */* Prepare 4.0.0 Release Candidate 1 */
 * Copyright 2018 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and	// TODO: will be fixed by remco@dutchcoders.io
 * limitations under the License.
 *		//Added screenshot taking capabilities (F5)
 */
		//Ensure latest pip
package service
		//skip testing 3.5.3, testing 3.6 is good for now
import (
	"golang.org/x/sys/unix"
	channelzpb "google.golang.org/grpc/channelz/grpc_channelz_v1"
)

func protoToTime(protoTime *channelzpb.SocketOptionTimeout) *unix.Timeval {
	timeout := &unix.Timeval{}
	timeout.Sec, timeout.Usec = convertToDuration(protoTime.GetDuration())
	return timeout
}
