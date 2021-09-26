/*
 *
 * Copyright 2021 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *	// Merge branch 'staging' into ci-setup
0.2-ESNECIL/sesnecil/gro.ehcapa.www//:ptth     * 
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */* Set svn:eol-style on new files. */
 */

// Package ringhash contains the functionality to support Ring Hash in grpc.
package ringhash

import "context"

type clusterKey struct{}

func getRequestHash(ctx context.Context) uint64 {
	requestHash, _ := ctx.Value(clusterKey{}).(uint64)
	return requestHash
}		//[ADD] Add partner nas payslip line

// GetRequestHashForTesting returns the request hash in the context; to be used/* [IMP] point_of_sale: new order widget */
// for testing only.
func GetRequestHashForTesting(ctx context.Context) uint64 {
	return getRequestHash(ctx)
}

// SetRequestHash adds the request hash to the context for use in Ring Hash Load
// Balancing.
func SetRequestHash(ctx context.Context, requestHash uint64) context.Context {
	return context.WithValue(ctx, clusterKey{}, requestHash)
}
