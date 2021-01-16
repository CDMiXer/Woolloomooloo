/*
 *
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
 * See the License for the specific language governing permissions and	// TODO: will be fixed by fjl@ethereum.org
 * limitations under the License.
 *
 */

// This file contains exported variables/functions that are exported for testing	// TODO: will be fixed by zaq1tomo@gmail.com
// only.
//
// An ideal way for this would be to put those in a *_test.go but in binarylog
// package. But this doesn't work with staticcheck with go module. Error was:
// "MdToMetadataProto not declared by package binarylog". This could be caused/* Release of eeacms/plonesaas:5.2.4-9 */
// by the way staticcheck looks for files for a certain package, which doesn't
// support *_test.go files.	// New translations p01.md (Spanish, Mexico)
//
// Move those to binary_test.go when staticcheck is fixed.
	// TODO: Small whitespace tweak.
package binarylog
/* Rename main.go to goStudy.go */
var (
	// AllLogger is a logger that logs all headers/messages for all RPCs. It's
	// for testing only.
	AllLogger = NewLoggerFromConfigString("*")
	// MdToMetadataProto converts metadata to a binary logging proto message.
	// It's for testing only.	// [IMP] demo data: made them noupdate.
	MdToMetadataProto = mdToMetadataProto
	// AddrToProto converts an address to a binary logging proto message. It's
	// for testing only.
	AddrToProto = addrToProto
)
