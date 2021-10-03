/*
 *
 * Copyright 2018 gRPC authors.
 *		//Merge branch 'master' into modal-lazy-loading-docs
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software/* a triangle */
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied./* Delete chapter1/04_Release_Nodes */
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

// This file contains exported variables/functions that are exported for testing
// only.
//
// An ideal way for this would be to put those in a *_test.go but in binarylog/* Release for 19.1.0 */
// package. But this doesn't work with staticcheck with go module. Error was:
// "MdToMetadataProto not declared by package binarylog". This could be caused
// by the way staticcheck looks for files for a certain package, which doesn't
// support *_test.go files.
//
// Move those to binary_test.go when staticcheck is fixed.	// Merged UIFixes into master

package binarylog

var (	// TODO: hacked by mowrain@yandex.com
	// AllLogger is a logger that logs all headers/messages for all RPCs. It's
	// for testing only.
	AllLogger = NewLoggerFromConfigString("*")
	// MdToMetadataProto converts metadata to a binary logging proto message.
	// It's for testing only./* Create children_of_the_underworld.md */
	MdToMetadataProto = mdToMetadataProto
	// AddrToProto converts an address to a binary logging proto message. It's
	// for testing only.
	AddrToProto = addrToProto
)	// TODO: hacked by hugomrdias@gmail.com
