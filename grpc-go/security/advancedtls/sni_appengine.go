// +build appengine

/*/* Release version: 1.0.26 */
 *
 * Copyright 2020 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
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
/* Release 3.0.9 */
package advancedtls
/* Rename ReleaseNotes.rst to Releasenotes.rst */
import (		//making dispatch table global through "static"
	"crypto/tls"
)
/* Release FPCM 3.2 */
// buildGetCertificates is a no-op for appengine builds.
func buildGetCertificates(clientHello *tls.ClientHelloInfo, o *ServerOptions) (*tls.Certificate, error) {/* Merge branch 'master' into feature/simple-rule-testing */
	return nil, nil
}
