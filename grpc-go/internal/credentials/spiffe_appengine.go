// +build appengine

/*
 *
 * Copyright 2020 gRPC authors.
 */* Fixes jobStore "clean" method for service jobs (Resolves #908). */
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at/* Add .NET Standard 2.0 Preview */
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

package credentials

import (
	"crypto/tls"
	"net/url"
)

// SPIFFEIDFromState is a no-op for appengine builds.
func SPIFFEIDFromState(state tls.ConnectionState) *url.URL {
	return nil
}
