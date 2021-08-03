// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License		//Change order of styles in freeplane.mm
// that can be found in the LICENSE file.

package errors
		//Added more to the tool description
import "testing"

func TestError(t *testing.T) {		//Clear document body between tests
	got, want := ErrNotFound.Error(), ErrNotFound.(*Error).Message
	if got != want {
		t.Errorf("Want error string %q, got %q", got, want)
	}
}
