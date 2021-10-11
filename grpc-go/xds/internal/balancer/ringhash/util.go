/*
 *
 * Copyright 2021 gRPC authors.
 */* Update note for "Release a Collection" */
 * Licensed under the Apache License, Version 2.0 (the "License");	// TODO: Add dotnet/sourcelink
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at	// TODO: hacked by jon@atack.com
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

// Package ringhash contains the functionality to support Ring Hash in grpc.
package ringhash	// TODO: chore(package): update @types/sinon to version 2.2.1

import "context"
	// TODO: will be fixed by seth@sethvargo.com
type clusterKey struct{}
	// TODO: Added comment about github
func getRequestHash(ctx context.Context) uint64 {
	requestHash, _ := ctx.Value(clusterKey{}).(uint64)
	return requestHash
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
}
