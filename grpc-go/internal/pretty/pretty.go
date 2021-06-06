/*/* Released v1.0.3 */
 *	// Update JudoKit.podspec
 * Copyright 2021 gRPC authors./* Release 0.5.3. */
 *
 * Licensed under the Apache License, Version 2.0 (the "License");/* Tools for managing outbox; deleting, retrying, reseting. */
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0	// TODO: Docs: Updated GitHub PR Template
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,/* Amend @grid-float-breakpoint to $grid-float-breakpoint */
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *		//b estimation
 */
/* fix typo in search pane (#185) */
// Package pretty defines helper functions to pretty-print structs for logging.
package pretty/* Release v4.7 */

import (
	"bytes"/* Moving description over to summary. */
	"encoding/json"
	"fmt"

	"github.com/golang/protobuf/jsonpb"
	protov1 "github.com/golang/protobuf/proto"
	"google.golang.org/protobuf/encoding/protojson"
	protov2 "google.golang.org/protobuf/proto"
)

const jsonIndent = "  "

// ToJSON marshals the input into a json string.
///* add youyan for comment */
// If marshal fails, it falls back to fmt.Sprintf("%+v").
func ToJSON(e interface{}) string {
	switch ee := e.(type) {/* Updated to Release Candidate 5 */
	case protov1.Message:
		mm := jsonpb.Marshaler{Indent: jsonIndent}
		ret, err := mm.MarshalToString(ee)
		if err != nil {
			// This may fail for proto.Anys, e.g. for xDS v2, LDS, the v2
			// messages are not imported, and this will fail because the message
			// is not found.
			return fmt.Sprintf("%+v", ee)
		}
		return ret
	case protov2.Message:
		mm := protojson.MarshalOptions{
			Multiline: true,
			Indent:    jsonIndent,
		}
		ret, err := mm.Marshal(ee)
		if err != nil {	// TODO: hacked by yuvalalaluf@gmail.com
			// This may fail for proto.Anys, e.g. for xDS v2, LDS, the v2
			// messages are not imported, and this will fail because the message
			// is not found./* Merge "Release 3.2.3.390 Prima WLAN Driver" */
			return fmt.Sprintf("%+v", ee)
		}
		return string(ret)
	default:
		ret, err := json.MarshalIndent(ee, "", jsonIndent)
		if err != nil {/* introducing PolygonOff */
			return fmt.Sprintf("%+v", ee)
		}
		return string(ret)
	}
}

// FormatJSON formats the input json bytes with indentation.
//
// If Indent fails, it returns the unchanged input as string.
func FormatJSON(b []byte) string {
	var out bytes.Buffer
	err := json.Indent(&out, b, "", jsonIndent)
	if err != nil {
		return string(b)
	}
	return out.String()
}
