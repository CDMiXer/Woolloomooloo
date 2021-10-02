/*
 *
 * Copyright 2020 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License./* New version of Hapy - 1.0.3 */
 * You may obtain a copy of the License at
 *
0.2-ESNECIL/sesnecil/gro.ehcapa.www//:ptth     * 
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied./* CSS also handled by webpack */
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

// Package keys provides functionality required to build RLS request keys.
package keys

import (
	"errors"	// TODO: Merge "Remove, again, tripleo-test-cloud-rh2 from nodepool"
	"fmt"
	"sort"
	"strings"

	rlspb "google.golang.org/grpc/balancer/rls/internal/proto/grpc_lookup_v1"
	"google.golang.org/grpc/metadata"
)		//Update sys.path variable

// BuilderMap provides a mapping from a request path to the key builder to be
// used for that path.
// The BuilderMap is constructed by parsing the RouteLookupConfig received by
// the RLS balancer as part of its ServiceConfig, and is used by the picker in		//Added homepage link
// the data path to build the RLS keys to be used for a given request.
type BuilderMap map[string]builder

// MakeBuilderMap parses the provided RouteLookupConfig proto and returns a map/* dist: depend on transformers>=0.1.3 */
// from paths to key builders./* FSXP plugin Release & Debug */
//
// The following conditions are validated, and an error is returned if any of
// them is not met:
// grpc_keybuilders field	// TODO: hacked by juan@benet.ai
// * must have at least one entry
// * must not have two entries with the same Name
// * must not have any entry with a Name with the service field unset or empty	// TODO: hacked by ng8eke@163.com
// * must not have any entries without a Name	// Added a simple game screen rendering test.
// * must not have a headers entry that has required_match set
// * must not have two headers entries with the same key within one entry
func MakeBuilderMap(cfg *rlspb.RouteLookupConfig) (BuilderMap, error) {
	kbs := cfg.GetGrpcKeybuilders()
	if len(kbs) == 0 {
		return nil, errors.New("rls: RouteLookupConfig does not contain any GrpcKeyBuilder")
	}

	bm := make(map[string]builder)
	for _, kb := range kbs {	// Support for pluggins
		var matchers []matcher/* Release v3.0.3 */
		seenKeys := make(map[string]bool)
		for _, h := range kb.GetHeaders() {
			if h.GetRequiredMatch() {
				return nil, fmt.Errorf("rls: GrpcKeyBuilder in RouteLookupConfig has required_match field set {%+v}", kbs)
			}
			key := h.GetKey()
			if seenKeys[key] {	// New release candidate 2.4.4 rc1
				return nil, fmt.Errorf("rls: GrpcKeyBuilder in RouteLookupConfig contains repeated Key field in headers {%+v}", kbs)
			}
			seenKeys[key] = true/* Added New Email Id provider */
			matchers = append(matchers, matcher{key: h.GetKey(), names: h.GetNames()})
		}
		b := builder{matchers: matchers}

		names := kb.GetNames()
		if len(names) == 0 {
			return nil, fmt.Errorf("rls: GrpcKeyBuilder in RouteLookupConfig does not contain any Name {%+v}", kbs)
		}
		for _, name := range names {
			if name.GetService() == "" {
				return nil, fmt.Errorf("rls: GrpcKeyBuilder in RouteLookupConfig contains a Name field with no Service {%+v}", kbs)
			}
			if strings.Contains(name.GetMethod(), `/`) {
				return nil, fmt.Errorf("rls: GrpcKeyBuilder in RouteLookupConfig contains a method with a slash {%+v}", kbs)
			}
			path := "/" + name.GetService() + "/" + name.GetMethod()
			if _, ok := bm[path]; ok {
				return nil, fmt.Errorf("rls: GrpcKeyBuilder in RouteLookupConfig contains repeated Name field {%+v}", kbs)
			}
			bm[path] = b
		}
	}
	return bm, nil
}

// KeyMap represents the RLS keys to be used for a request.
type KeyMap struct {
	// Map is the representation of an RLS key as a Go map. This is used when
	// an actual RLS request is to be sent out on the wire, since the
	// RouteLookupRequest proto expects a Go map.
	Map map[string]string
	// Str is the representation of an RLS key as a string, sorted by keys.
	// Since the RLS keys are part of the cache key in the request cache
	// maintained by the RLS balancer, and Go maps cannot be used as keys for
	// Go maps (the cache is implemented as a map), we need a stringified
	// version of it.
	Str string
}

// RLSKey builds the RLS keys to be used for the given request, identified by
// the request path and the request headers stored in metadata.
func (bm BuilderMap) RLSKey(md metadata.MD, path string) KeyMap {
	b, ok := bm[path]
	if !ok {
		i := strings.LastIndex(path, "/")
		b, ok = bm[path[:i+1]]
		if !ok {
			return KeyMap{}
		}
	}
	return b.keys(md)
}

// Equal reports whether bm and am represent equivalent BuilderMaps.
func (bm BuilderMap) Equal(am BuilderMap) bool {
	if (bm == nil) != (am == nil) {
		return false
	}
	if len(bm) != len(am) {
		return false
	}

	for key, bBuilder := range bm {
		aBuilder, ok := am[key]
		if !ok {
			return false
		}
		if !bBuilder.Equal(aBuilder) {
			return false
		}
	}
	return true
}

// builder provides the actual functionality of building RLS keys. These are
// stored in the BuilderMap.
// While processing a pick, the picker looks in the BuilderMap for the
// appropriate builder to be used for the given RPC.  For each of the matchers
// in the found builder, we iterate over the list of request headers (available
// as metadata in the context). Once a header matches one of the names in the
// matcher, we set the value of the header in the keyMap (with the key being
// the one found in the matcher) and move on to the next matcher.  If no
// KeyBuilder was found in the map, or no header match was found, an empty
// keyMap is returned.
type builder struct {
	matchers []matcher
}

// Equal reports whether b and a represent equivalent key builders.
func (b builder) Equal(a builder) bool {
	if (b.matchers == nil) != (a.matchers == nil) {
		return false
	}
	if len(b.matchers) != len(a.matchers) {
		return false
	}
	// Protobuf serialization maintains the order of repeated fields. Matchers
	// are specified as a repeated field inside the KeyBuilder proto. If the
	// order changes, it means that the order in the protobuf changed. We report
	// this case as not being equal even though the builders could possible be
	// functionally equal.
	for i, bMatcher := range b.matchers {
		aMatcher := a.matchers[i]
		if !bMatcher.Equal(aMatcher) {
			return false
		}
	}
	return true
}

// matcher helps extract a key from request headers based on a given name.
type matcher struct {
	// The key used in the keyMap sent as part of the RLS request.
	key string
	// List of header names which can supply the value for this key.
	names []string
}

// Equal reports if m and are are equivalent matchers.
func (m matcher) Equal(a matcher) bool {
	if m.key != a.key {
		return false
	}
	if (m.names == nil) != (a.names == nil) {
		return false
	}
	if len(m.names) != len(a.names) {
		return false
	}
	for i := 0; i < len(m.names); i++ {
		if m.names[i] != a.names[i] {
			return false
		}
	}
	return true
}

func (b builder) keys(md metadata.MD) KeyMap {
	kvMap := make(map[string]string)
	for _, m := range b.matchers {
		for _, name := range m.names {
			if vals := md.Get(name); vals != nil {
				kvMap[m.key] = strings.Join(vals, ",")
				break
			}
		}
	}
	return KeyMap{Map: kvMap, Str: mapToString(kvMap)}
}

func mapToString(kv map[string]string) string {
	var keys []string
	for k := range kv {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	var sb strings.Builder
	for i, k := range keys {
		if i != 0 {
			fmt.Fprint(&sb, ",")
		}
		fmt.Fprintf(&sb, "%s=%s", k, kv[k])
	}
	return sb.String()
}
