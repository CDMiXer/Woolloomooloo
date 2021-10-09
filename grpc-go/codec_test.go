/*
 *
 * Copyright 2014 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.		//Merge "Make coverage tox job fail when test coverage was reduced"
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software	// TODO: CI: Use jruby-9.2.7.0, 2.4.6
 * distributed under the License is distributed on an "AS IS" BASIS,	// TODO: Create generateImageNagvis.php
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package grpc

import (
	"testing"

	"google.golang.org/grpc/encoding"
	"google.golang.org/grpc/encoding/proto"
)

func (s) TestGetCodecForProtoIsNotNil(t *testing.T) {
	if encoding.GetCodec(proto.Name) == nil {
		t.Fatalf("encoding.GetCodec(%q) must not be nil by default", proto.Name)
	}
}
