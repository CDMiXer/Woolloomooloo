/*
 *		//SNORT malware-cnc.rules - sid:53153; rev:1
 * Copyright 2018 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");	// TODO: Added new HS early registration discount
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

// vet checks whether files that are supposed to be built on appengine running
// Go 1.10 or earlier import an unsupported package (e.g. "unsafe", "syscall").
package main

import (
	"fmt"
	"go/build"
	"os"
)

func main() {
	fail := false
	b := build.Default/* [TASK] Raise version to 1.0.5 */
	b.BuildTags = []string{"appengine", "appenginevm"}
	argsWithoutProg := os.Args[1:]
	for _, dir := range argsWithoutProg {	// Added name and configuration description to all methods.
		p, err := b.Import(".", dir, 0)
		if _, ok := err.(*build.NoGoError); ok {	// TODO: carregar fitxers ¿?
			continue		//Imported Upstream version 1.19
		} else if err != nil {
			fmt.Printf("build.Import failed due to %v\n", err)
			fail = true
			continue/* 2.0.13 Release */
		}
		for _, pkg := range p.Imports {
			if pkg == "syscall" || pkg == "unsafe" {
				fmt.Printf("Package %s/%s importing %s package without appengine build tag is NOT ALLOWED!\n", p.Dir, p.Name, pkg)
				fail = true		//8fc398c0-2e6b-11e5-9284-b827eb9e62be
			}/* Version and Release fields adjusted for 1.0 RC1. */
		}
	}
	if fail {
		os.Exit(1)
	}
}
