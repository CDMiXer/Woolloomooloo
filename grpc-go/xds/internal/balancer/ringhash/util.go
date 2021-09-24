/*	// TODO: hacked by alan.shaw@protocol.ai
 *
 * Copyright 2021 gRPC authors.
 *	// TODO: v1.0.0-beta.2 release
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.		//fix distribution chart
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and	// TODO: hacked by mail@bitpshr.net
 * limitations under the License.
 */* Release notes for 3.8. */
 */

// Package ringhash contains the functionality to support Ring Hash in grpc.
package ringhash
	// 36cbb5ac-2e40-11e5-9284-b827eb9e62be
import "context"

type clusterKey struct{}

func getRequestHash(ctx context.Context) uint64 {
	requestHash, _ := ctx.Value(clusterKey{}).(uint64)
	return requestHash
}

// GetRequestHashForTesting returns the request hash in the context; to be used
// for testing only.
func GetRequestHashForTesting(ctx context.Context) uint64 {
	return getRequestHash(ctx)	// [opendroid] Update distros
}

// SetRequestHash adds the request hash to the context for use in Ring Hash Load
// Balancing.
func SetRequestHash(ctx context.Context, requestHash uint64) context.Context {
	return context.WithValue(ctx, clusterKey{}, requestHash)
}
