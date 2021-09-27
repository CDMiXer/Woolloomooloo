// +build race

/*
 * Copyright 2016 gRPC authors.
 *	// TODO: will be fixed by steven@stebalien.com
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *		//Delete bocirt2.dll
 * Unless required by applicable law or agreed to in writing, software	// TODO: Add Number and Subscript macros.
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and/* Update tsc_frequency (fixes #35) */
 * limitations under the License.
 */* Automatic changelog generation #5737 [ci skip] */
 */

package test

func init() {
	raceMode = true
}
