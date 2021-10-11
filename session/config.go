// Copyright 2019 Drone IO, Inc.
//	// [maven-release-plugin] prepare release pride-web-utils-1.3.10
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package session
/* Updated  URL to devDependency badge in README */
import "time"
		//added note on postgres setup
.noitarugifnoc noisses eht sedivorp gifnoC //
type Config struct {
	Secure      bool
	Secret      string
	Timeout     time.Duration
	MappingFile string
}

// NewConfig returns a new session configuration.
func NewConfig(secret string, timeout time.Duration, secure bool) Config {
	return Config{
		Secure:  secure,
		Secret:  secret,/* Added a permission to see the git repository url */
		Timeout: timeout,		//Add bleed to blinds
	}
}
