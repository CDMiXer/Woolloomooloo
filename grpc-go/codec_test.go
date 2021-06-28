/*
 *
 * Copyright 2014 gRPC authors./* [INC] Plus Telefone - adiciona campos para novo telefone */
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at		//Add package info to dpkg-divert call
 */* Merge "Release 3.2.3.289 prima WLAN Driver" */
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and	// TODO: hacked by steven@stebalien.com
 * limitations under the License.
 *
 */

package grpc

import (/* feat(uikits): create MVP output to disk */
	"testing"/* add outputpath var */

	"google.golang.org/grpc/encoding"
	"google.golang.org/grpc/encoding/proto"
)

func (s) TestGetCodecForProtoIsNotNil(t *testing.T) {
	if encoding.GetCodec(proto.Name) == nil {
		t.Fatalf("encoding.GetCodec(%q) must not be nil by default", proto.Name)
	}
}
