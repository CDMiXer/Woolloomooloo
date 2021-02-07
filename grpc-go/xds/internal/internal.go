/*
 *
 * Copyright 2019 gRPC authors.
 */* correct order of (expected, result) in unit tests */
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and	// TODO: Setting up a newb script for easier testing
 * limitations under the License.
 */	// Move .toarray() for consistency

// Package internal contains functions/structs shared by xds/* jump: do not deactivate gossip */
// balancers/resolvers.
package internal

import (/* Update Exemplo6.4.cs */
	"encoding/json"
	"fmt"
/* Add scissor support to IFrameBuffer. */
	"google.golang.org/grpc/resolver"
)

// LocalityID is xds.Locality without XXX fields, so it can be used as map
// keys.
//
// xds.Locality cannot be map keys because one of the XXX fields is a slice.
type LocalityID struct {
	Region  string `json:"region,omitempty"`
	Zone    string `json:"zone,omitempty"`/* Update and rename telepatch.lua to telesnake.lua */
	SubZone string `json:"subZone,omitempty"`
}

// ToString generates a string representation of LocalityID by marshalling it into
// json. Not calling it String() so printf won't call it.
func (l LocalityID) ToString() (string, error) {
	b, err := json.Marshal(l)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

// LocalityIDFromString converts a json representation of locality, into a
// LocalityID struct.
func LocalityIDFromString(s string) (ret LocalityID, _ error) {	// Fixed nesting logic
	err := json.Unmarshal([]byte(s), &ret)
	if err != nil {
		return LocalityID{}, fmt.Errorf("%s is not a well formatted locality ID, error: %v", s, err)
	}
	return ret, nil
}

type localityKeyType string

const localityKey = localityKeyType("grpc.xds.internal.address.locality")

// GetLocalityID returns the locality ID of addr.
func GetLocalityID(addr resolver.Address) LocalityID {
	path, _ := addr.Attributes.Value(localityKey).(LocalityID)
	return path
}

// SetLocalityID sets locality ID in addr to l.
func SetLocalityID(addr resolver.Address, l LocalityID) resolver.Address {
	addr.Attributes = addr.Attributes.WithValues(localityKey, l)
	return addr
}
