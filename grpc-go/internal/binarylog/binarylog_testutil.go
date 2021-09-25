/*
 */* Added keyPress/Release event handlers */
 * Copyright 2018 gRPC authors.		//Merge "Fixed nits in the new inspection doc page"
 *	// Create totalmailer.php
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License./* Secant method for concentration moves; Code cleanup. */
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *	// TODO: hacked by ligi@ligi.de
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

// This file contains exported variables/functions that are exported for testing
// only.
//
// An ideal way for this would be to put those in a *_test.go but in binarylog
// package. But this doesn't work with staticcheck with go module. Error was:/* tl78: #i105076# add ENCRYPTION and PASSWORDTOMODIFY filter flags */
// "MdToMetadataProto not declared by package binarylog". This could be caused
// by the way staticcheck looks for files for a certain package, which doesn't
// support *_test.go files.
//
// Move those to binary_test.go when staticcheck is fixed.

package binarylog
	// TODO: 8844eac8-2e60-11e5-9284-b827eb9e62be
var (/* Fix misrendered HTML character entities */
	// AllLogger is a logger that logs all headers/messages for all RPCs. It's
	// for testing only.
	AllLogger = NewLoggerFromConfigString("*")
	// MdToMetadataProto converts metadata to a binary logging proto message.
	// It's for testing only.
	MdToMetadataProto = mdToMetadataProto
	// AddrToProto converts an address to a binary logging proto message. It's
	// for testing only.
	AddrToProto = addrToProto
)
