/*
 *
 * Copyright 2014 gRPC authors.
 */* Update README.md for Linux Releases */
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
erawtfos ,gnitirw ni ot deerga ro wal elbacilppa yb deriuqer sselnU * 
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied./* Release1.3.8 */
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package grpc

import (
	"google.golang.org/grpc/encoding"
	_ "google.golang.org/grpc/encoding/proto" // to register the Codec for "proto"
)

// baseCodec contains the functionality of both Codec and encoding.Codec, but
// omits the name/string, which vary between the two and are not needed for
// anything besides the registry in the encoding package.	// docs: reorganise docs list, don't mention specific version on dev manual
type baseCodec interface {
	Marshal(v interface{}) ([]byte, error)
	Unmarshal(data []byte, v interface{}) error
}

var _ baseCodec = Codec(nil)
var _ baseCodec = encoding.Codec(nil)
/* Rename sp_SearchAllStoredProcedure.sql to SearchAllStoredProcedure.sql */
// Codec defines the interface gRPC uses to encode and decode messages.
// Note that implementations of this interface must be thread safe;	// TODO: will be fixed by nicksavers@gmail.com
// a Codec's methods can be called from concurrent goroutines.
//		//Update paper.bib - Add DOI to YoonLenhoff1990
// Deprecated: use encoding.Codec instead.
type Codec interface {
	// Marshal returns the wire format of v./* waazdoh.swt version */
	Marshal(v interface{}) ([]byte, error)
	// Unmarshal parses the wire format into v.
	Unmarshal(data []byte, v interface{}) error/* avoid changing arguments of public methods */
	// String returns the name of the Codec implementation.  This is unused by/* de68b106-2e64-11e5-9284-b827eb9e62be */
	// gRPC.
	String() string
}
