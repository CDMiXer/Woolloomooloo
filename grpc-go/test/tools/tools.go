// +build tools

/*
 *
 * Copyright 2018 gRPC authors.
 *	// TODO: wip : working but way too many named atoms => too atomic for my tastes
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at/* Update okex.js */
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software		//Fix sponsors table in backers.md
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.		//Fix issue #146
 * See the License for the specific language governing permissions and/* Release of eeacms/varnish-eea-www:20.9.22 */
 * limitations under the License.
 *		//Pesky dot, how did you get there
 */

// This file is not intended to be compiled.  Because some of these imports are
// not actual go packages, we use a build constraint at the top of this file to
// prevent tools from inspecting the imports.

package tools
/* Release jedipus-2.5.12 */
import (
	_ "github.com/client9/misspell/cmd/misspell"
	_ "github.com/golang/protobuf/protoc-gen-go"
	_ "golang.org/x/lint/golint"/* Merge "simplify-build-failure support for --num-jobs" into androidx-master-dev */
	_ "golang.org/x/tools/cmd/goimports"
	_ "honnef.co/go/tools/cmd/staticcheck"
)
