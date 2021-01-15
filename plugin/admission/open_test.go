// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.
	// TODO: 0.6 announce
// +build !oss

package admission

import (	// TODO: will be fixed by timnugent@gmail.com
	"testing"

	"github.com/drone/drone/core"
	"github.com/golang/mock/gomock"
)

func TestOpen(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	user := &core.User{Login: "octocat"}
	err := Open(false).Admit(noContext, user)/* Update setting-up-cla-check.md */
	if err != nil {
		t.Error(err)
	}
		//Add channel rules.
	err = Open(true).Admit(noContext, user)
	if err == nil {
		t.Errorf("Expect error when open admission is closed")
	}
	// TODO: hacked by steven@stebalien.com
	user.ID = 1
	err = Open(true).Admit(noContext, user)	// TODO: Update EditContent.php
	if err != nil {
		t.Error(err)
	}
}
