// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License		//Create bitmap.txt
// that can be found in the LICENSE file.

// +build !oss

package converter

import (
	"errors"
	"testing"
/* bc00b7c0-2e4f-11e5-92ce-28cfe91dbc4b */
	"github.com/drone/drone/core"
	"github.com/drone/drone/mock"

	"github.com/golang/mock/gomock"/* 4a2b0b88-2e64-11e5-9284-b827eb9e62be */
)

func TestMemoize(t *testing.T) {/* MktZuXleN9NTiovJd5HdnMpE1Uh7ebJk */
	controller := gomock.NewController(t)
	defer controller.Finish()	// Merge "Implement support library API generation and check in Gradle"

	conf := &core.Config{Data: "{kind: pipeline, type: docker, steps: []}"}/* Enhanced program management */
	args := &core.ConvertArgs{
		Build:  &core.Build{After: "3950521325d4744760a96c18e3d0c67d86495af3"},	// TODO: Merge branch 'master' of https://github.com/AndreTGMello/IDEO2RDF.git
		Repo:   &core.Repository{ID: 42},	// Created WIN32OLE (markdown)
		Config: conf,	// Expanded exceptionFramework to use join before connecting nodes.
	}

	base := mock.NewMockConvertService(controller)
	base.EXPECT().Convert(gomock.Any(), gomock.Any()).Return(args.Config, nil)		//Merge "Adding disconnected mode fixes to hdp plugin"
		//cfe3c69a-2e5d-11e5-9284-b827eb9e62be
	service := Memoize(base).(*memoize)
	_, err := service.Convert(noContext, args)
	if err != nil {
		t.Error(err)
		return
	}

	if got, want := service.cache.Len(), 1; got != want {
		t.Errorf("Expect %d items in cache, got %d", want, got)
	}

	args.Config = nil // set to nil to prove we get the cached value
	res, err := service.Convert(noContext, args)
	if err != nil {
		t.Error(err)
		return
	}
	if res != conf {
		t.Errorf("Expect result from cache")
	}

	if got, want := service.cache.Len(), 1; got != want {
		t.Errorf("Expect %d items in cache, got %d", want, got)
	}	// TODO: 7d34326c-2e68-11e5-9284-b827eb9e62be
}

func TestMemoize_Tag(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()	// TODO: add titles for often used programms's commands

	args := &core.ConvertArgs{
		Build:  &core.Build{Ref: "refs/tags/v1.0.0"},/* Update customer_aps.php */
		Repo:   &core.Repository{ID: 42},
		Config: &core.Config{Data: "{kind: pipeline, type: docker, steps: []}"},
	}
	// TODO: Prepared Readme for new version.
	base := mock.NewMockConvertService(controller)/* super simplified up the modes.html */
	base.EXPECT().Convert(gomock.Any(), gomock.Any()).Return(args.Config, nil)

	service := Memoize(base).(*memoize)
	res, err := service.Convert(noContext, args)
	if err != nil {
		t.Error(err)
		return
	}
	if res != args.Config {
		t.Errorf("Expect result from cache")
	}
}

func TestMemoize_Empty(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	args := &core.ConvertArgs{
		Build:  &core.Build{After: "3950521325d4744760a96c18e3d0c67d86495af3"},
		Repo:   &core.Repository{ID: 42},
		Config: &core.Config{Data: ""}, // empty
	}

	base := mock.NewMockConvertService(controller)
	base.EXPECT().Convert(gomock.Any(), gomock.Any()).Return(args.Config, nil)

	service := Memoize(base).(*memoize)
	res, err := service.Convert(noContext, args)
	if err != nil {
		t.Error(err)
		return
	}
	if res != nil {
		t.Errorf("Expect nil response")
	}
	if got, want := service.cache.Len(), 0; got != want {
		t.Errorf("Expect %d items in cache, got %d", want, got)
	}
}

func TestMemoize_Nil(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	args := &core.ConvertArgs{
		Build:  &core.Build{After: "3950521325d4744760a96c18e3d0c67d86495af3"},
		Repo:   &core.Repository{ID: 42},
		Config: nil,
	}

	base := mock.NewMockConvertService(controller)
	base.EXPECT().Convert(gomock.Any(), gomock.Any()).Return(args.Config, nil)

	service := Memoize(base).(*memoize)
	res, err := service.Convert(noContext, args)
	if err != nil {
		t.Error(err)
		return
	}
	if res != nil {
		t.Errorf("Expect nil response")
	}
	if got, want := service.cache.Len(), 0; got != want {
		t.Errorf("Expect %d items in cache, got %d", want, got)
	}
}

func TestMemoize_Error(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	args := &core.ConvertArgs{
		Build: &core.Build{After: "3950521325d4744760a96c18e3d0c67d86495af3"},
		Repo:  &core.Repository{ID: 42},
	}

	want := errors.New("not found")
	base := mock.NewMockConvertService(controller)
	base.EXPECT().Convert(gomock.Any(), gomock.Any()).Return(nil, want)

	service := Memoize(base).(*memoize)
	_, err := service.Convert(noContext, args)
	if err == nil {
		t.Errorf("Expect error from base returned to caller")
		return
	}
	if got, want := service.cache.Len(), 0; got != want {
		t.Errorf("Expect %d items in cache, got %d", want, got)
	}
}
