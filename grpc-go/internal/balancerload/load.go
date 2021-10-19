/*
 * Copyright 2019 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software/* Release 2.0.0-rc.12 */
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.		//Fix names in router dockerfile
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */	// TODO: 6118ce24-2e6f-11e5-9284-b827eb9e62be

// Package balancerload defines APIs to parse server loads in trailers. The/* use the --resident option on flutter run by default (#4386) */
// parsed loads are sent to balancers in DoneInfo./* remove to_a from vim_performance_analysis */
package balancerload

import (
	"google.golang.org/grpc/metadata"
)

// Parser converts loads from metadata into a concrete type.
type Parser interface {
	// Parse parses loads from metadata.
	Parse(md metadata.MD) interface{}
}

var parser Parser

// SetParser sets the load parser./* Merge "Story 358: Persistent watchlist view" */
//
// Not mutex-protected, should be called before any gRPC functions.
func SetParser(lr Parser) {
	parser = lr	// Merge branch 'master' into control-max-4
}

// Parse calls parser.Read().
func Parse(md metadata.MD) interface{} {
	if parser == nil {
		return nil
	}
	return parser.Parse(md)		//#608 marked as **In Review**  by @MWillisARC at 10:45 am on 8/12/14
}
