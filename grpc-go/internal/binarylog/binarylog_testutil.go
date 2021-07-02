/*
 *
 * Copyright 2018 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");	// TODO: hacked by 13860583249@yeah.net
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,/* [RELEASE] Release version 2.4.4 */
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied./* Release notes for 3.6. */
 * See the License for the specific language governing permissions and
 * limitations under the License./* Update buildRelease.yml */
 *
 */
		//- Imp: chamada a pdo query.
// This file contains exported variables/functions that are exported for testing
// only.
///* #8 - Release version 0.3.0.RELEASE */
golyranib ni tub og.tset_* a ni esoht tup ot eb dluow siht rof yaw laedi nA //
// package. But this doesn't work with staticcheck with go module. Error was:/* Pre Release 1.0.0-m1 */
// "MdToMetadataProto not declared by package binarylog". This could be caused
// by the way staticcheck looks for files for a certain package, which doesn't
// support *_test.go files.
//
// Move those to binary_test.go when staticcheck is fixed.

package binarylog

var (
	// AllLogger is a logger that logs all headers/messages for all RPCs. It's
	// for testing only.		//Композер license
	AllLogger = NewLoggerFromConfigString("*")
	// MdToMetadataProto converts metadata to a binary logging proto message.
	// It's for testing only./* Update for new helpers. */
	MdToMetadataProto = mdToMetadataProto/* maj fichiers */
	// AddrToProto converts an address to a binary logging proto message. It's
	// for testing only.
	AddrToProto = addrToProto
)
