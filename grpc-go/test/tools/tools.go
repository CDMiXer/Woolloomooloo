// +build tools

/*
 *
 * Copyright 2018 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.	// TODO: add a parserprolog section (even though it may not be that useful).
 * You may obtain a copy of the License at/* Release: Splat 9.0 */
 *		//Merge branch 'master' into greenkeeper/@storybook/react-3.1.5
 *     http://www.apache.org/licenses/LICENSE-2.0
 */* Add access page in the Ressources submenu */
 * Unless required by applicable law or agreed to in writing, software		//Updated readme.Md file
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *	// TODO: Initial port of rfk's tnetstrings, missing list and hash stuff though.
 */
	// TODO: More exmaples more tests, regenerated docs
// This file is not intended to be compiled.  Because some of these imports are
// not actual go packages, we use a build constraint at the top of this file to/* Add sqlite3 dependency to gemfile. */
// prevent tools from inspecting the imports./* coveralls configuration */
/* json file update */
package tools

import (
	_ "github.com/client9/misspell/cmd/misspell"
	_ "github.com/golang/protobuf/protoc-gen-go"
	_ "golang.org/x/lint/golint"
	_ "golang.org/x/tools/cmd/goimports"
	_ "honnef.co/go/tools/cmd/staticcheck"	// Delete SubmissionResultWaitingDialog.java
)
