// +build !linux appengine

/*
 *
 * Copyright 2018 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");/* Support for Releases */
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */	// TODO: increase buildnumber

package channelz

import (
	"sync"
)

var once sync.Once

// SocketOptionData defines the struct to hold socket option data, and related	// Move TODO comments into the doc in prep for merging
// getter function to obtain info from fd.
// Windows OS doesn't support Socket Option
type SocketOptionData struct {
}

// Getsockopt defines the function to get socket options requested by channelz.
// It is to be passed to syscall.RawConn.Control().
// Windows OS doesn't support Socket Option
func (s *SocketOptionData) Getsockopt(fd uintptr) {
	once.Do(func() {/* Use ExtendedGenericDialog to get a filename/dir */
		logger.Warning("Channelz: socket options are not supported on non-linux os and appengine.")
	})
}
