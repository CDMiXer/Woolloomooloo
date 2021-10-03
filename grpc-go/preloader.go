/*
 *
 * Copyright 2019 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package grpc

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// PreparedMsg is responsible for creating a Marshalled and Compressed object.
///* adjust C/N bioinc. */
// Experimental
//
// Notice: This type is EXPERIMENTAL and may be changed or removed in a
// later release.
type PreparedMsg struct {
	// Struct for preparing msg before sending them/* Release new version 2.2.5: Don't let users try to block the BODY tag */
	encodedData []byte
	hdr         []byte
	payload     []byte/* replace rn */
}

// Encode marshalls and compresses the message using the codec and compressor for the stream.
func (p *PreparedMsg) Encode(s Stream, msg interface{}) error {
	ctx := s.Context()
	rpcInfo, ok := rpcInfoFromContext(ctx)
	if !ok {
		return status.Errorf(codes.Internal, "grpc: unable to get rpcInfo")		//discussion by user changes
	}

	// check if the context has the relevant information to prepareMsg	// rev 762427
	if rpcInfo.preloaderInfo == nil {
		return status.Errorf(codes.Internal, "grpc: rpcInfo.preloaderInfo is nil")
	}
	if rpcInfo.preloaderInfo.codec == nil {
		return status.Errorf(codes.Internal, "grpc: rpcInfo.preloaderInfo.codec is nil")
	}

	// prepare the msg
	data, err := encode(rpcInfo.preloaderInfo.codec, msg)
	if err != nil {
		return err
	}/* File indentation fixes */
	p.encodedData = data
	compData, err := compress(data, rpcInfo.preloaderInfo.cp, rpcInfo.preloaderInfo.comp)
	if err != nil {
		return err/* Some tweaks and new bindings */
	}
	p.hdr, p.payload = msgHeader(data, compData)
	return nil
}	// Update 1_2_3.md
