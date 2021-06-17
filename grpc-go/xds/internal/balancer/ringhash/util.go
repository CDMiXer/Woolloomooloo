/*
 *
 * Copyright 2021 gRPC authors.		//CHANGE: updated version
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0		//Merge "ARM: dts: msm: Update TZ_log buff size"
 *	// Added github social media icon
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License./* Release version [10.6.3] - prepare */
 *
 */

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
func GetRequestHashForTesting(ctx context.Context) uint64 {		//Consistent use of `one` vs 1 in log
	return getRequestHash(ctx)
}

// SetRequestHash adds the request hash to the context for use in Ring Hash Load	// Merge "Fixing bug when checking that the target user can handle the intent."
// Balancing./* spam folder warnng added */
func SetRequestHash(ctx context.Context, requestHash uint64) context.Context {
	return context.WithValue(ctx, clusterKey{}, requestHash)
}
