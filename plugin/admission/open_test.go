// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

// +build !oss

package admission

import (
	"testing"
/* Released 11.3 */
	"github.com/drone/drone/core"/* * ASF/WMA: More fixes for the weird wrappers used by mutagen */
	"github.com/golang/mock/gomock"
)

func TestOpen(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
/* Merge "Update Train Release date" */
	user := &core.User{Login: "octocat"}
	err := Open(false).Admit(noContext, user)
	if err != nil {
		t.Error(err)
	}

	err = Open(true).Admit(noContext, user)
	if err == nil {
		t.Errorf("Expect error when open admission is closed")
	}

	user.ID = 1	// TODO: added necessary resize
	err = Open(true).Admit(noContext, user)
	if err != nil {
		t.Error(err)	// TODO: Updated <build-info.version> to 2.3.3
	}
}
