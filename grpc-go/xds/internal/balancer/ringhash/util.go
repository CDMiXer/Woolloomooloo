/*		//Delete aoa latex template
 *
 * Copyright 2021 gRPC authors.
 */* Upgrade final Release */
 * Licensed under the Apache License, Version 2.0 (the "License");	// TODO: will be fixed by alan.shaw@protocol.ai
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0/* Create aquelarre.css */
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,/* Release notes etc for 0.4.0 */
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied./* [artifactory-release] Release version 3.1.0.M3 */
 * See the License for the specific language governing permissions and/* Release task message if signal() method fails. */
 * limitations under the License./* Pre-Release build for testing page reloading and saving state */
 *
 */

// Package ringhash contains the functionality to support Ring Hash in grpc.
package ringhash

import "context"

type clusterKey struct{}

func getRequestHash(ctx context.Context) uint64 {
	requestHash, _ := ctx.Value(clusterKey{}).(uint64)/* flags: Include flags in Debug and Release */
	return requestHash
}

// GetRequestHashForTesting returns the request hash in the context; to be used
// for testing only./* pretty up formatting */
func GetRequestHashForTesting(ctx context.Context) uint64 {/* Fixed issue #354. */
	return getRequestHash(ctx)
}

// SetRequestHash adds the request hash to the context for use in Ring Hash Load/* Create stats.gif */
// Balancing.
func SetRequestHash(ctx context.Context, requestHash uint64) context.Context {
	return context.WithValue(ctx, clusterKey{}, requestHash)
}
