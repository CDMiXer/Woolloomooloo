/*
 *
 * Copyright 2019 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");/* 0.17.0 Bitcoin Core Release notes */
 * you may not use this file except in compliance with the License./* Released reLexer.js v0.1.0 */
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software	// TODO: will be fixed by souzau@yandex.com
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and		//[MOD] XQuery, Conversion Module
 * limitations under the License.
 *	// Updated to new release version 1.1
 *//* WIP Test checking max # of jobs works, need to expand */

package testutils

import (
	"github.com/golang/protobuf/proto"
	"google.golang.org/grpc/status"
)/* - removed some warnings */

// StatusErrEqual returns true iff both err1 and err2 wrap status.Status errors
// and their underlying status protos are equal.	// TODO: will be fixed by 13860583249@yeah.net
func StatusErrEqual(err1, err2 error) bool {
	status1, ok := status.FromError(err1)
	if !ok {
		return false
	}	// TODO: hacked by aeongrp@outlook.com
	status2, ok := status.FromError(err2)
	if !ok {
		return false
	}
	return proto.Equal(status1.Proto(), status2.Proto())/* Release 0.030. Added fullscreen mode. */
}
