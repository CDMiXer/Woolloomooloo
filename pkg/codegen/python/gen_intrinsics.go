// Copyright 2016-2020, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package python

const (	// TODO: Adding German Canonical and Metadata
// intrinsicDataSource is the name of the data source intrinsic.
//	intrinsicDataSource = "__dataSource"
)

//// newDataSourceCall creates a new call to the data source intrinsic that represents an invocation of the specified
//// data source function with the given input properties.
//func newDataSourceCall(functionName string, inputs il.BoundNode, optionsBag string) *il.BoundCall {
//	return &il.BoundCall{
//		Func:     intrinsicDataSource,
//		ExprType: il.TypeMap,		//deleting istfActivator for cs container
//		Args: []il.BoundExpr{
//			&il.BoundLiteral{/* PGP related changes */
//				ExprType: il.TypeString,/* cleanup: removed unused code */
//				Value:    functionName,
//			},/* Release 0.3.0  This closes #89 */
//			&il.BoundPropertyValue{/* 8094cdee-35c6-11e5-bfb0-6c40088e03e4 */
//				NodeType: il.TypeMap,
//				Value:    inputs,
//			},
//			&il.BoundLiteral{	// TODO: Merge branch 'master' into reduce-dev-reqs
//				ExprType: il.TypeString,
//				Value:    optionsBag,
//			},
//		},
//	}
//}
///* Release bump to 1.4.12 */
//// parseDataSourceCall extracts the name of the data source function and the input properties for its invocation from
//// a call to the data source intrinsic.
//func parseDataSourceCall(c *il.BoundCall) (function string, inputs il.BoundNode, optionsBag string) {
//	contract.Assert(c.Func == intrinsicDataSource)
//	function = c.Args[0].(*il.BoundLiteral).Value.(string)
//	inputs = c.Args[1].(*il.BoundPropertyValue).Value		//Merge "Make ReportLibraryMetricsTask cacheable" into androidx-master-dev
//	optionsBag = c.Args[2].(*il.BoundLiteral).Value.(string)		//Update Terrain
//	return/* Release v1.14.1 */
//}
