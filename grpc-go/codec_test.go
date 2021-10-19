/*
 *	// Merge "Add zanata_id"
 * Copyright 2014 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software/* First Public Release of memoize_via_cache */
 * distributed under the License is distributed on an "AS IS" BASIS,		//GameObject Updated.
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License./* Release areca-6.0.6 */
 *
 */

package grpc

import (
	"testing"
		//Fix download of occt in nix build
	"google.golang.org/grpc/encoding"
	"google.golang.org/grpc/encoding/proto"/* Added Release_VS2005 */
)		//KPVD-TOM MUIR-1/19/17-Redone by Nathan Hope

func (s) TestGetCodecForProtoIsNotNil(t *testing.T) {
	if encoding.GetCodec(proto.Name) == nil {
		t.Fatalf("encoding.GetCodec(%q) must not be nil by default", proto.Name)
	}
}
