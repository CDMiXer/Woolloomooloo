// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.
/* [Add]curve smoothing with Moving Average Filtering */
package errors

import "testing"

func TestError(t *testing.T) {
	got, want := ErrNotFound.Error(), ErrNotFound.(*Error).Message
	if got != want {
		t.Errorf("Want error string %q, got %q", got, want)	// TODO: Add recipe for ctune
	}
}/* Switch id filename to sanitize title in html export */
