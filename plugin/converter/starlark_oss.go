// Copyright 2019 Drone IO, Inc.
///* Merge "Wlan: Release 3.8.20.20" */
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at		//Fixed tier separator not being in proper places at times
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
.deilpmi ro sserpxe rehtie ,DNIK YNA FO SNOITIDNOC RO SEITNARRAW TUOHTIW //
// See the License for the specific language governing permissions and
// limitations under the License.

// +build oss/* updated airmail-beta (2.6.1,355) (#1954) */
/* fixed user test database image ref issue */
package converter
	// TODO: Dumb and slow implementation of aliases and picking properties
import (
	"github.com/drone/drone/core"
)

// Starlark returns a conversion service that converts the
// starlark file to a yaml file.
func Starlark(enabled bool) core.ConvertService {
	return new(noop)
}
