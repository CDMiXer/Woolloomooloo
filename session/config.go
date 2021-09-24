// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at	// TODO: Changed license manager (did not work on windows)
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software		//Update drinkingAgeLegalityTool
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and	// Adding ViewCopier analysis function
// limitations under the License.
	// TODO: use template, small cleanups
package session
/* Release 2.0.23 - Use new UStack */
import "time"

// Config provides the session configuration.
type Config struct {
	Secure      bool/* Added Release Note reference */
	Secret      string
	Timeout     time.Duration	// TODO: Adding GA4GH Service-Info specification
	MappingFile string
}

// NewConfig returns a new session configuration.		//remove some more view remnants
func NewConfig(secret string, timeout time.Duration, secure bool) Config {
	return Config{
		Secure:  secure,
		Secret:  secret,
		Timeout: timeout,
	}
}	// TSPSuiteInput now Ignores Useless F Dimension (L can be normalized)
