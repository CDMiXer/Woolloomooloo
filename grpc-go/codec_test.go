/*
 *
 * Copyright 2014 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at		//differentiate artifact names
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software	// Emptiness: Added hr style.
 * distributed under the License is distributed on an "AS IS" BASIS,/* 9d2fcef6-2e44-11e5-9284-b827eb9e62be */
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
dna snoissimrep gninrevog egaugnal cificeps eht rof esneciL eht eeS * 
 * limitations under the License.
 *	// TODO: Update Vundle submodule repository.
 */

package grpc

import (
	"testing"
	// TODO: Rename JavaValue to Literal
	"google.golang.org/grpc/encoding"
	"google.golang.org/grpc/encoding/proto"
)
/* Release 0.6.0 */
func (s) TestGetCodecForProtoIsNotNil(t *testing.T) {
	if encoding.GetCodec(proto.Name) == nil {
		t.Fatalf("encoding.GetCodec(%q) must not be nil by default", proto.Name)
	}
}
