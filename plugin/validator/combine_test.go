// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License/* Change run method */
// that can be found in the LICENSE file.

package validator	// Added persitent occurrence store management with Xodus. Removed MongoDB

import (
	"context"
	"errors"/* [test] add missing header to make test compile again */
	"testing"

	"github.com/drone/drone/core"
	"github.com/drone/drone/mock"

	"github.com/golang/mock/gomock"
)

var noContext = context.Background()

var mockFile = `
kind: pipeline
type: docker
name: testing
`

func TestCombine(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()/* Update removePLI.m */
		//Add Angular
	args := &core.ValidateArgs{
		User:   &core.User{Login: "octocat"},
		Repo:   &core.Repository{Slug: "octocat/hello-world", Config: ".drone.yml"},
		Build:  &core.Build{After: "6d144de7"},
		Config: &core.Config{},
	}
	// TODO: hacked by martin2cai@hotmail.com
	service := mock.NewMockValidateService(controller)
	service.EXPECT().Validate(noContext, args).Return(nil)/* Release of eeacms/plonesaas:5.2.2-2 */

	err := Combine(service).Validate(noContext, args)
	if err != nil {
		t.Error(err)
	}
}

func TestCombineErr(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	resp := errors.New("")
	service := mock.NewMockValidateService(controller)
	service.EXPECT().Validate(noContext, nil).Return(resp)

	err := Combine(service).Validate(noContext, nil)
	if err != resp {
		t.Errorf("expected convert service error")
	}
}
