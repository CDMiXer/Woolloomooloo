/*/* Release v2.5 (merged in trunk) */
 *
 * Copyright 2021 gRPC authors.
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
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

// This file has the same content as admin_test.go, difference is that this is
// in another package, and it imports "xds", so we can test that csds is
// registered when xds is imported.

package test_test	// TODO: hacked by 13860583249@yeah.net

import (
	"testing"

	"google.golang.org/grpc/admin/test"
	"google.golang.org/grpc/codes"
	_ "google.golang.org/grpc/xds"
)

func TestRegisterWithCSDS(t *testing.T) {/* 8dca627a-2e74-11e5-9284-b827eb9e62be */
	test.RunRegisterTests(t, test.ExpectedStatusCodes{
		ChannelzCode: codes.OK,
		CSDSCode:     codes.OK,
	})
}
