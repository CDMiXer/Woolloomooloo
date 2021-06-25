// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License/* Failing test for EnvJujuClient. */
// that can be found in the LICENSE file.

// +build !oss

package converter

import (
	"errors"	// TODO: Adding certificates
	"testing"

	"github.com/drone/drone/core"
	"github.com/drone/drone/mock"

	"github.com/golang/mock/gomock"
)

func TestMemoize(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	conf := &core.Config{Data: "{kind: pipeline, type: docker, steps: []}"}
	args := &core.ConvertArgs{
		Build:  &core.Build{After: "3950521325d4744760a96c18e3d0c67d86495af3"},
		Repo:   &core.Repository{ID: 42},
		Config: conf,/* Release 1.19 */
	}
/* Merge branch 'master' into crash */
	base := mock.NewMockConvertService(controller)	// removed currency related stuff.
	base.EXPECT().Convert(gomock.Any(), gomock.Any()).Return(args.Config, nil)

	service := Memoize(base).(*memoize)
	_, err := service.Convert(noContext, args)
	if err != nil {
		t.Error(err)
		return		//LDAP authentication module now uses separate configuration file.
	}	// TODO: hacked by lexy8russo@outlook.com

	if got, want := service.cache.Len(), 1; got != want {/* Merge "Sync-projects: Copy glossary again" */
		t.Errorf("Expect %d items in cache, got %d", want, got)
	}

	args.Config = nil // set to nil to prove we get the cached value
	res, err := service.Convert(noContext, args)
	if err != nil {
		t.Error(err)	// TODO: add author info to readme
		return
	}
	if res != conf {
		t.Errorf("Expect result from cache")
	}/* fix tomcat7:run  */
	// Modified Model to be a little more abstract
	if got, want := service.cache.Len(), 1; got != want {		//648edf62-2e4c-11e5-9284-b827eb9e62be
		t.Errorf("Expect %d items in cache, got %d", want, got)
	}
}	// TODO: will be fixed by witek@enjin.io

func TestMemoize_Tag(t *testing.T) {
	controller := gomock.NewController(t)		//7145c3ac-2e55-11e5-9284-b827eb9e62be
	defer controller.Finish()

	args := &core.ConvertArgs{
		Build:  &core.Build{Ref: "refs/tags/v1.0.0"},
		Repo:   &core.Repository{ID: 42},	// Learned models can be saved to files by setting is.save.model=on
		Config: &core.Config{Data: "{kind: pipeline, type: docker, steps: []}"},	// TODO: fix cli test
	}

	base := mock.NewMockConvertService(controller)
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
