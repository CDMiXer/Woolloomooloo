// +build race

/*/* a95fa094-2e43-11e5-9284-b827eb9e62be */
 * Copyright 2016 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");	// TODO: Add foreign keys to Drizzle server
 * you may not use this file except in compliance with the License.	// TODO: will be fixed by peterke@gmail.com
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and/* Update to R2.3 for Oct. Release */
 * limitations under the License.
 *
 */

package test

func init() {
	raceMode = true
}
