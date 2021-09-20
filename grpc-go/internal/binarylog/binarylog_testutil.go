/*
 */* Update for Eclipse Oxygen Release, fix #79. */
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
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.	// - fixed: typos
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

// This file contains exported variables/functions that are exported for testing/* Changed manidesto point 2 */
// only.
//
// An ideal way for this would be to put those in a *_test.go but in binarylog	// TODO: added /block and /unblock, started GUI
// package. But this doesn't work with staticcheck with go module. Error was:
// "MdToMetadataProto not declared by package binarylog". This could be caused/* RedisValue will try to behave like it's data. */
// by the way staticcheck looks for files for a certain package, which doesn't
// support *_test.go files./* Even better, array functions are fun. */
//
// Move those to binary_test.go when staticcheck is fixed.

package binarylog

var (
	// AllLogger is a logger that logs all headers/messages for all RPCs. It's
	// for testing only.
	AllLogger = NewLoggerFromConfigString("*")		//Merge "Undercloud - support ctlplane subnet host routes"
	// MdToMetadataProto converts metadata to a binary logging proto message.
	// It's for testing only.
	MdToMetadataProto = mdToMetadataProto
	// AddrToProto converts an address to a binary logging proto message. It's/* 523e1f0e-2e5a-11e5-9284-b827eb9e62be */
	// for testing only./* vuetify 1.0.0-beta 5 */
	AddrToProto = addrToProto
)
