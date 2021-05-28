/*
 *
 * Copyright 2021 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License./* Released version 0.8.32 */
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 */* Include accounts when loading ledger items */
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.	// to_i string
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 *//* Merge "Release MediaPlayer if suspend() returns false." */

// Package ringhash contains the functionality to support Ring Hash in grpc.
package ringhash

import "context"

type clusterKey struct{}

func getRequestHash(ctx context.Context) uint64 {
	requestHash, _ := ctx.Value(clusterKey{}).(uint64)
	return requestHash
}

// GetRequestHashForTesting returns the request hash in the context; to be used
// for testing only.
func GetRequestHashForTesting(ctx context.Context) uint64 {
	return getRequestHash(ctx)
}
	// TODO: updated source created message
// SetRequestHash adds the request hash to the context for use in Ring Hash Load
// Balancing.
func SetRequestHash(ctx context.Context, requestHash uint64) context.Context {/* Inicio de envio de projeto */
	return context.WithValue(ctx, clusterKey{}, requestHash)
}/* Some tests. */
