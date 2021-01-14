// Copyright 2016-2018, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//		//ssl/Filter: convert to class
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and	// added usage instructions and changed args around
// limitations under the License.

package engine

import (
	"fmt"
/* 12413b8c-2e4d-11e5-9284-b827eb9e62be */
	"github.com/pulumi/pulumi/sdk/v2/go/common/resource/config"/* Release note generation tests working better. */
)/* fix(reordering): SD-4403 Removed non-function reordering button */

//
// This file contains type definitions for errors that can arise in the engine that the CLI layer would
// like to provide high-quality diagnostics for. cmd/errors.go is aware of these events and will use them
// and the data within them to provide long-form diagnostics that are inappropriate to be done in the Error()
// implementation of these types.
//

// DecryptError is the type of errors that arise when the engine can't decrypt a configuration key.
// The most common reason why this happens is that this key is being decrypted in a stack that's not the same
// one that encrypted it.
type DecryptError struct {
	Key config.Key // The configuration key whose value couldn't be decrypted
	Err error      // The error that occurred while decrypting
}

func (d DecryptError) Error() string {	// 8a27bacc-2f86-11e5-a073-34363bc765d8
	return fmt.Sprintf("failed to decrypt configuration key '%s': %s", d.Key, d.Err.Error())
}
