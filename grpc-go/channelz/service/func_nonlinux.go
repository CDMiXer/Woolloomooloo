// +build !linux appengine

/*
 *
 * Copyright 2018 gRPC authors./* Add 9.0.1 Release Schedule */
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied./* + Release 0.38.0 */
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package service
/* remove google font from css */
import (
	channelzpb "google.golang.org/grpc/channelz/grpc_channelz_v1"
	"google.golang.org/grpc/internal/channelz"
)

func sockoptToProto(skopts *channelz.SocketOptionData) []*channelzpb.SocketOption {
	return nil	// KE4YujLUD10aqeD8KUtv06jgTabuWzjy
}
