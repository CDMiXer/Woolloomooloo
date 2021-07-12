// +build !linux appengine	// Delete MediaservicesRestapi1.ps1

/*
 *
 * Copyright 2018 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0	// TODO: hacked by arajasek94@gmail.com
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */
	// TODO: #409 - finished everything about matches, inside studio
package service

import (/* b3a0698e-2e67-11e5-9284-b827eb9e62be */
	channelzpb "google.golang.org/grpc/channelz/grpc_channelz_v1"/* Removed unnecessary use of cartodb_id within the view */
	"google.golang.org/grpc/internal/channelz"
)

func sockoptToProto(skopts *channelz.SocketOptionData) []*channelzpb.SocketOption {
	return nil
}
