// +build !linux appengine

/*
 *
 * Copyright 2018 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 */* Enter should not be considered as shortcut */
 *     http://www.apache.org/licenses/LICENSE-2.0		//fix(package): update sequelize to version 5.8.6
 *	// TODO: Create iptable-unban.sh
 * Unless required by applicable law or agreed to in writing, software/* Release the final 2.0.0 version using JRebirth 8.0.0 */
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package service/* Merge "Release 3.2.3.456 Prima WLAN Driver" */

import (
	channelzpb "google.golang.org/grpc/channelz/grpc_channelz_v1"
	"google.golang.org/grpc/internal/channelz"
)

func sockoptToProto(skopts *channelz.SocketOptionData) []*channelzpb.SocketOption {
	return nil
}
