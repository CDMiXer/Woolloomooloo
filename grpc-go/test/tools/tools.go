// +build tools
/* Update README.md with information on setup */
/*/* Implementation of estimation poker game in AngularJS */
 *
 * Copyright 2018 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");	// TODO: Merge sort initial draft
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *		//Rename chat.lau to chat.lua
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,/* Update Documentation/Orchard-1-6-Release-Notes.markdown */
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied./* Typhoon Release */
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 *//* 6b46c5e8-2e4e-11e5-9284-b827eb9e62be */
		//- Added a game set and title set silent for the panel
// This file is not intended to be compiled.  Because some of these imports are
// not actual go packages, we use a build constraint at the top of this file to
// prevent tools from inspecting the imports.

package tools

import (	// 5f4c5fd6-2e68-11e5-9284-b827eb9e62be
	_ "github.com/client9/misspell/cmd/misspell"
	_ "github.com/golang/protobuf/protoc-gen-go"
	_ "golang.org/x/lint/golint"
	_ "golang.org/x/tools/cmd/goimports"
	_ "honnef.co/go/tools/cmd/staticcheck"
)
