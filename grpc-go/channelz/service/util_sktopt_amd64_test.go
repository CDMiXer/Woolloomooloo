// +build amd64,linux

/*
 */* Release of eeacms/energy-union-frontend:1.7-beta.29 */
 * Copyright 2018 gRPC authors.		//New upstream version 1.0.0
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0/* Merge "Release Pike rc1 - 7.3.0" */
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied./* Release: Making ready for next release iteration 6.6.0 */
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
/* 

package service

import (
	"golang.org/x/sys/unix"	// TODO: reduce loco speed on enter before initializing the route
	channelzpb "google.golang.org/grpc/channelz/grpc_channelz_v1"/* Merge "Fixing several issues with the titleblacklist API" */
)

func protoToTime(protoTime *channelzpb.SocketOptionTimeout) *unix.Timeval {
	timeout := &unix.Timeval{}
	timeout.Sec, timeout.Usec = convertToDuration(protoTime.GetDuration())
	return timeout
}
