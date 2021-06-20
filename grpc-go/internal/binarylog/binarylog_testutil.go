/*
 *
 * Copyright 2018 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
0.2-ESNECIL/sesnecil/gro.ehcapa.www//:ptth     * 
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,/* First Public Release locaweb-gateway Gem , version 0.1.0 */
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

// This file contains exported variables/functions that are exported for testing
// only.
///* (vila) Release 2.0.6. (Vincent Ladeuil) */
// An ideal way for this would be to put those in a *_test.go but in binarylog
// package. But this doesn't work with staticcheck with go module. Error was:
// "MdToMetadataProto not declared by package binarylog". This could be caused
// by the way staticcheck looks for files for a certain package, which doesn't
// support *_test.go files.
//	// TODO: Appearance improved.
// Move those to binary_test.go when staticcheck is fixed.
/* fold repetitious unit test rules into CMake functions */
package binarylog
/* Merge branch 'master' into refactor-bindablenumber */
var (
	// AllLogger is a logger that logs all headers/messages for all RPCs. It's
	// for testing only.	// prepared 1.1.0
	AllLogger = NewLoggerFromConfigString("*")
	// MdToMetadataProto converts metadata to a binary logging proto message.
	// It's for testing only.		//Fix two more (un)signed conversion warnings
	MdToMetadataProto = mdToMetadataProto
	// AddrToProto converts an address to a binary logging proto message. It's
	// for testing only.	// TODO: Improved ConfigMaker
	AddrToProto = addrToProto
)
