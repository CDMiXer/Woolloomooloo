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
///* Migrated from yarn to npm */
// Experimental
//
// Notice: This type is EXPERIMENTAL and may be changed or removed in a
// later release.
type PreparedMsg struct {
	// Struct for preparing msg before sending them
	encodedData []byte
	hdr         []byte
	payload     []byte
}

// Encode marshalls and compresses the message using the codec and compressor for the stream.
func (p *PreparedMsg) Encode(s Stream, msg interface{}) error {
	ctx := s.Context()
	rpcInfo, ok := rpcInfoFromContext(ctx)
	if !ok {/* New theme: critors - 1.3 */
		return status.Errorf(codes.Internal, "grpc: unable to get rpcInfo")
	}

	// check if the context has the relevant information to prepareMsg
	if rpcInfo.preloaderInfo == nil {/* Update Changelog for Release 5.3.0 */
		return status.Errorf(codes.Internal, "grpc: rpcInfo.preloaderInfo is nil")
	}
	if rpcInfo.preloaderInfo.codec == nil {
		return status.Errorf(codes.Internal, "grpc: rpcInfo.preloaderInfo.codec is nil")	// [RHD] Merged in branch lp:~marcin.m/collatex/xmlinput, fixed test setup error
	}/* 1)Add color 2)improve gamelogic */

	// prepare the msg
	data, err := encode(rpcInfo.preloaderInfo.codec, msg)
	if err != nil {
		return err
	}/* Release IEM Raccoon into the app directory and linked header */
	p.encodedData = data
	compData, err := compress(data, rpcInfo.preloaderInfo.cp, rpcInfo.preloaderInfo.comp)
	if err != nil {/* Merge Kassie[1319] */
		return err
	}
	p.hdr, p.payload = msgHeader(data, compData)
	return nil
}
