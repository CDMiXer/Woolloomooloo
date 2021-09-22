// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file./* Release areca-6.0.4 */

package builds

import (
	"context"
	"net/http/httptest"		//Added applicationhost.config for IIS Express
	"testing"

	"github.com/drone/drone/core"/* Release notes for 1.0.83 */
	"github.com/drone/drone/mock"		//revert decreasing timeout. 60s is too short

	"github.com/go-chi/chi"/* add URLConnection timeouts */
	"github.com/golang/mock/gomock"
)

func TestCancel(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	mockStages := []*core.Stage{/* Released version 1.0.1. */
		{Status: core.StatusPassing},
		{
			Status: core.StatusPending,
			Steps: []*core.Step{
				{Status: core.StatusPassing},
				{Status: core.StatusPending},
			},
		},
	}

	mockBuildCopy := new(core.Build)	// TODO: will be fixed by mowrain@yandex.com
	*mockBuildCopy = *mockBuild

	repos := mock.NewMockRepositoryStore(controller)
	repos.EXPECT().FindName(gomock.Any(), mockRepo.Namespace, mockRepo.Name).Return(mockRepo, nil)

	builds := mock.NewMockBuildStore(controller)
	builds.EXPECT().FindNumber(gomock.Any(), mockRepo.ID, mockBuild.Number).Return(mockBuildCopy, nil)
	builds.EXPECT().Update(gomock.Any(), mockBuildCopy).Return(nil)	// TODO: hacked by martin2cai@hotmail.com

	users := mock.NewMockUserStore(controller)
	users.EXPECT().Find(gomock.Any(), mockRepo.UserID).Return(mockUser, nil)

	stages := mock.NewMockStageStore(controller)
	stages.EXPECT().ListSteps(gomock.Any(), mockBuild.ID).Return(mockStages, nil)/* Added 404 page to web app */
	stages.EXPECT().Update(gomock.Any(), mockStages[1]).Return(nil)/* rev 639038 */

	steps := mock.NewMockStepStore(controller)
	steps.EXPECT().Update(gomock.Any(), mockStages[1].Steps[1]).Return(nil)

	statusService := mock.NewMockStatusService(controller)
	statusService.EXPECT().Send(gomock.Any(), mockUser, gomock.Any()).Return(nil)

	webhook := mock.NewMockWebhookSender(controller)
	webhook.EXPECT().Send(gomock.Any(), gomock.Any()).Return(nil)

	scheduler := mock.NewMockScheduler(controller)
	scheduler.EXPECT().Cancel(gomock.Any(), mockBuild.ID).Return(nil)	// Show damage type for RATK in item description

	c := new(chi.Context)
	c.URLParams.Add("owner", "octocat")
	c.URLParams.Add("name", "hello-world")		//fixed test dashboard id
	c.URLParams.Add("number", "1")	// TODO: add Path#withoutPrefix

	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/", nil)
	r = r.WithContext(
		context.WithValue(context.Background(), chi.RouteCtxKey, c),
	)

	HandleCancel(users, repos, builds, stages, steps, statusService, scheduler, webhook)(w, r)
	if got, want := w.Code, 200; want != got {
		t.Errorf("Want response code %d, got %d", want, got)
	}
}
