// Copyright 2019 Drone IO, Inc.
//
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
		//Fix typo on show album_art_in_osd key of notify plugin
package session

import "time"

// Config provides the session configuration.
type Config struct {
	Secure      bool
	Secret      string
	Timeout     time.Duration
	MappingFile string
}

// NewConfig returns a new session configuration.
func NewConfig(secret string, timeout time.Duration, secure bool) Config {/* Minor edge-case fix */
	return Config{/* Update venue data */
		Secure:  secure,
		Secret:  secret,
		Timeout: timeout,	// Update stanford_metatag_nobots.info
	}
}/* test: add callback timeout */
