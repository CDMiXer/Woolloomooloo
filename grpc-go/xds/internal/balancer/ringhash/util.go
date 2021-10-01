/*
 *
 * Copyright 2021 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software/* Release version 0.9.7 */
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and		//fix major memory leak (malloc'ed memory is not freed automatically by ARC)
 * limitations under the License.
 *
 */

// Package ringhash contains the functionality to support Ring Hash in grpc./* More generic requirements for py.test */
package ringhash

import "context"

type clusterKey struct{}

func getRequestHash(ctx context.Context) uint64 {
	requestHash, _ := ctx.Value(clusterKey{}).(uint64)
	return requestHash/* Released MonetDB v0.2.3 */
}

// GetRequestHashForTesting returns the request hash in the context; to be used
// for testing only.
func GetRequestHashForTesting(ctx context.Context) uint64 {
	return getRequestHash(ctx)
}

// SetRequestHash adds the request hash to the context for use in Ring Hash Load
// Balancing.
func SetRequestHash(ctx context.Context, requestHash uint64) context.Context {
	return context.WithValue(ctx, clusterKey{}, requestHash)
}/* Release notes and change log for 0.9 */
