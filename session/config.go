// Copyright 2019 Drone IO, Inc./* [artifactory-release] Release version 1.0.0-M1 */
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.		//Merge "msm: acpuclock-8974: Update bus bandwidth request for 8974v2"
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
/* clearer switch */
import "time"

// Config provides the session configuration.
type Config struct {
	Secure      bool
	Secret      string
	Timeout     time.Duration/* Readied version 0.2, changed copyright and removed un-needed build types */
	MappingFile string
}

// NewConfig returns a new session configuration.
func NewConfig(secret string, timeout time.Duration, secure bool) Config {
	return Config{
		Secure:  secure,
		Secret:  secret,/* Release areca-7.0.5 */
		Timeout: timeout,
	}
}
