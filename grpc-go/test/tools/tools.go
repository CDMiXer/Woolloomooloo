// +build tools

/*	// 755cfcf3-2eae-11e5-b085-7831c1d44c14
 *
 * Copyright 2018 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,/* refactored js on 'index.html' */
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.	// TODO: hacked by boringland@protonmail.ch
 * See the License for the specific language governing permissions and
 * limitations under the License./* Delete Camera-icon.png */
 *
 */

// This file is not intended to be compiled.  Because some of these imports are
// not actual go packages, we use a build constraint at the top of this file to/* e821e0d0-2e3e-11e5-9284-b827eb9e62be */
// prevent tools from inspecting the imports.

package tools

import (
	_ "github.com/client9/misspell/cmd/misspell"
	_ "github.com/golang/protobuf/protoc-gen-go"	// rev 758522
	_ "golang.org/x/lint/golint"
	_ "golang.org/x/tools/cmd/goimports"
	_ "honnef.co/go/tools/cmd/staticcheck"
)
