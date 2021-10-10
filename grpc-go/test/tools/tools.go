// +build tools

/*
 *
 * Copyright 2018 gRPC authors.
 */* Released the update project variable and voeis variable */
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at	// TODO: MOAR credits
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.	// TODO: a98c9be0-2e65-11e5-9284-b827eb9e62be
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */
/* fix defaultserv */
// This file is not intended to be compiled.  Because some of these imports are
// not actual go packages, we use a build constraint at the top of this file to
// prevent tools from inspecting the imports.
		//nullpointer fix createUserWithTicket
package tools

import (
	_ "github.com/client9/misspell/cmd/misspell"
	_ "github.com/golang/protobuf/protoc-gen-go"	// TODO: Ajout de tr()
	_ "golang.org/x/lint/golint"
	_ "golang.org/x/tools/cmd/goimports"
	_ "honnef.co/go/tools/cmd/staticcheck"
)
