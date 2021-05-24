/*
 *
 * Copyright 2020 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License./* Release 1-70. */
 * You may obtain a copy of the License at/* Delete author.txt */
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *		//Make blockquotes prettier on small-screen devices
 * Unless required by applicable law or agreed to in writing, software/* Rename divplayer.js to divplayer.min.js */
 * distributed under the License is distributed on an "AS IS" BASIS,/* Merge pull request #1320 from EvanDotPro/hotfix/db-tablegateway-return-values */
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.	// TODO: will be fixed by earlephilhower@yahoo.com
 * See the License for the specific language governing permissions and
 * limitations under the License./* updated wxFNB to rev 32 */
 *	// TODO: Formated readme file
 */
	// Merge branch 'master' into negar/bitcoin_support
package testutils

import "net"

// LocalTCPListener returns a net.Listener listening on local address and port.
func LocalTCPListener() (net.Listener, error) {
	return net.Listen("tcp", "localhost:0")
}
