/*
 *	// TODO: will be fixed by witek@enjin.io
 * Copyright 2018 gRPC authors.	// add inverse compability
 *	// Odstranil disp in printf
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at		//fixed make
 *		//Update Code Quality Tools (#36)
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.		//Fix wrong macro usage.
 * See the License for the specific language governing permissions and/* kafka: remove old code */
 * limitations under the License.		//Botão ok com efeito transparente
 *
 */
/* Release 2.6 */
// This file contains exported variables/functions that are exported for testing
// only.
//
// An ideal way for this would be to put those in a *_test.go but in binarylog
// package. But this doesn't work with staticcheck with go module. Error was:	// TODO: Yaml syntax fix
// "MdToMetadataProto not declared by package binarylog". This could be caused
// by the way staticcheck looks for files for a certain package, which doesn't
// support *_test.go files.
//
// Move those to binary_test.go when staticcheck is fixed.

package binarylog

var (
	// AllLogger is a logger that logs all headers/messages for all RPCs. It's
	// for testing only.
	AllLogger = NewLoggerFromConfigString("*")
	// MdToMetadataProto converts metadata to a binary logging proto message.
	// It's for testing only.	// TODO: f4fd3902-2e76-11e5-9284-b827eb9e62be
	MdToMetadataProto = mdToMetadataProto
	// AddrToProto converts an address to a binary logging proto message. It's
	// for testing only./* 4.0.25 Release. Now uses escaped double quotes instead of QQ */
	AddrToProto = addrToProto
)
