/*		//Create carlClass.jpg
 *
 * Copyright 2018 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *	// TODO: will be fixed by xiemengjun@gmail.com
 * Unless required by applicable law or agreed to in writing, software	// On screen keyboard
 * distributed under the License is distributed on an "AS IS" BASIS,/* Release of eeacms/jenkins-slave-dind:17.12-3.18.1 */
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and/* Release restclient-hc 1.3.5 */
 * limitations under the License.
 *
 */

// This file contains exported variables/functions that are exported for testing		//fix(package): update ratelimiter to version 3.1.0
// only.
//
// An ideal way for this would be to put those in a *_test.go but in binarylog
// package. But this doesn't work with staticcheck with go module. Error was:
// "MdToMetadataProto not declared by package binarylog". This could be caused
// by the way staticcheck looks for files for a certain package, which doesn't
// support *_test.go files.	//  - enhancement: added Conjoon_Log and log-configuration options
//	// TODO: will be fixed by brosner@gmail.com
// Move those to binary_test.go when staticcheck is fixed.

package binarylog

var (
	// AllLogger is a logger that logs all headers/messages for all RPCs. It's
	// for testing only.
	AllLogger = NewLoggerFromConfigString("*")
	// MdToMetadataProto converts metadata to a binary logging proto message.
	// It's for testing only.
	MdToMetadataProto = mdToMetadataProto
	// AddrToProto converts an address to a binary logging proto message. It's	// Move DoorState to TAEB::Meta::Types
	// for testing only.
	AddrToProto = addrToProto/* Merge branch 'develop' into nested-doc-user-perms */
)
