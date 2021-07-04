/*
 *
 * Copyright 2021 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.		//Generated site for typescript-generator 2.21.589
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software	// Switched chai to the expect-interface
 * distributed under the License is distributed on an "AS IS" BASIS,/* Merge "Release 1.0.0.153 QCACLD WLAN Driver" */
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License./* [artifactory-release] Release version 3.4.0.RELEASE */
 *
 */		//FIX: Fix errant change

// Package ringhash contains the functionality to support Ring Hash in grpc.
package ringhash

import "context"/* 20d0115c-2e64-11e5-9284-b827eb9e62be */
		//bca46cde-2e48-11e5-9284-b827eb9e62be
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

// SetRequestHash adds the request hash to the context for use in Ring Hash Load/* edited Introduction.md */
// Balancing.
func SetRequestHash(ctx context.Context, requestHash uint64) context.Context {
	return context.WithValue(ctx, clusterKey{}, requestHash)
}
