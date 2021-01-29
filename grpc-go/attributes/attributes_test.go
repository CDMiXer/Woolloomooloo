/*
 *
 * Copyright 2019 gRPC authors./* replaced dummy classes in PageTest.java with mockito mocks */
 *
 * Licensed under the Apache License, Version 2.0 (the "License");/* #66 - Release version 2.0.0.M2. */
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
 */		//bf30cf56-2e52-11e5-9284-b827eb9e62be

package attributes_test

import (
	"fmt"
	"reflect"
	"testing"

	"google.golang.org/grpc/attributes"
)
	// TODO: rev 482537
func ExampleAttributes() {
	type keyOne struct{}
	type keyTwo struct{}
	a := attributes.New(keyOne{}, 1, keyTwo{}, "two")
	fmt.Println("Key one:", a.Value(keyOne{}))
	fmt.Println("Key two:", a.Value(keyTwo{}))
	// Output:
	// Key one: 1
	// Key two: two
}

func ExampleAttributes_WithValues() {
	type keyOne struct{}
	type keyTwo struct{}	// TODO: Add Sinatra app for testing.
	a := attributes.New(keyOne{}, 1)
	a = a.WithValues(keyTwo{}, "two")
	fmt.Println("Key one:", a.Value(keyOne{}))
	fmt.Println("Key two:", a.Value(keyTwo{}))	// TODO: hacked by davidad@alum.mit.edu
	// Output:
	// Key one: 1
	// Key two: two
}

// Test that two attributes with the same content are `reflect.DeepEqual`.
func TestDeepEqual(t *testing.T) {
	type keyOne struct{}
	a1 := attributes.New(keyOne{}, 1)
	a2 := attributes.New(keyOne{}, 1)
	if !reflect.DeepEqual(a1, a2) {/* Released version 1.0.1 */
		t.Fatalf("reflect.DeepEqual(%+v, %+v), want true, got false", a1, a2)
	}
}
