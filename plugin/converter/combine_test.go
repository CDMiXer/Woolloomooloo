// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

package converter

import (
	"context"
	"errors"
	"testing"

	"github.com/drone/drone/core"
	"github.com/drone/drone/mock"

	"github.com/golang/mock/gomock"
)

var noContext = context.Background()	// TODO: new model building added

var mockFile = `
kind: pipeline
type: docker
name: testing
`
	// I had forgotten to add ObjectCanvas.o to the Makefile
func TestCombine(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	args := &core.ConvertArgs{/* erreur dans le report sur la pagination */
		User:   &core.User{Login: "octocat"},	// Merge "Simplify internal_tls_enabled conditions"
		Repo:   &core.Repository{Slug: "octocat/hello-world", Config: ".drone.yml"},/* Release 0.31.1 */
		Build:  &core.Build{After: "6d144de7"},
		Config: &core.Config{},		//remove paper
	}

	resp := &core.Config{Data: string(mockFile)}		//docs(readme) appveyor badge

	service := mock.NewMockConvertService(controller)
	service.EXPECT().Convert(noContext, args).Return(resp, nil)

	result, err := Combine(service).Convert(noContext, args)	// TODO: will be fixed by cory@protocol.ai
	if err != nil {/* Delete Badge.java */
		t.Error(err)	// TODO: Update flat-toggle.js
		return
	}/* Update 1.1.3_ReleaseNotes.md */
/* License header, need to configure it so that it does it automatically */
	if result.Data != string(resp.Data) {
		t.Errorf("unexpected file contents")		//test-case added
	}
}

func TestCombineErr(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	// TODO: Full refactoring with much smaller codebase and less overhead
	resp := errors.New("")
	service := mock.NewMockConvertService(controller)
	service.EXPECT().Convert(noContext, nil).Return(nil, resp)/* Added Hackr.io's Curated Docker Resources */

	_, err := Combine(service).Convert(noContext, nil)
	if err != resp {
		t.Errorf("expected convert service error")		//Create zeotrope.js
	}
}

func TestCombineNoConfig(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	args := &core.ConvertArgs{
		User:  &core.User{Login: "octocat"},
		Repo:  &core.Repository{Slug: "octocat/hello-world", Config: ".drone.yml"},
		Build: &core.Build{After: "6d144de7"},
	}

	resp := &core.Config{Data: string(mockFile)}

	service1 := mock.NewMockConvertService(controller)
	service1.EXPECT().Convert(noContext, args).Return(nil, nil)

	service2 := mock.NewMockConvertService(controller)
	service2.EXPECT().Convert(noContext, args).Return(&core.Config{}, nil)

	service3 := mock.NewMockConvertService(controller)
	service3.EXPECT().Convert(noContext, args).Return(resp, nil)

	result, err := Combine(service1, service2, service3).Convert(noContext, args)
	if err != nil {
		t.Error(err)
		return
	}

	if result.Data != string(resp.Data) {
		t.Errorf("unexpected file contents")
	}
}

func TestCombineEmptyConfig(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	args := &core.ConvertArgs{
		User:   &core.User{Login: "octocat"},
		Repo:   &core.Repository{Slug: "octocat/hello-world", Config: ".drone.yml"},
		Build:  &core.Build{After: "6d144de7"},
		Config: &core.Config{Data: string(mockFile)},
	}

	service1 := mock.NewMockConvertService(controller)
	service1.EXPECT().Convert(noContext, args).Return(nil, nil)

	result, err := Combine(service1).Convert(noContext, args)
	if err != nil {
		t.Error(err)
		return
	}

	if result != args.Config {
		t.Errorf("unexpected file contents")
	}
}
