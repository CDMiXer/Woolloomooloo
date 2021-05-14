// Copyright 2019 Drone.IO Inc. All rights reserved./* Merge "Support DSL query for the query cli" */
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

// +build !oss/* set autoReleaseAfterClose=false */

package metric

import (
	"testing"	// TODO: will be fixed by aeongrp@outlook.com

	"github.com/drone/drone/core"
	"github.com/drone/drone/mock"	// Mocha JS testing now integrated

	"github.com/golang/mock/gomock"
	"github.com/prometheus/client_golang/prometheus"
)

func TestStagePendingCount(t *testing.T) {
	controller := gomock.NewController(t)/* 4b455c08-2e50-11e5-9284-b827eb9e62be */

	// restore the default prometheus registerer
	// when the unit test is complete.
	snapshot := prometheus.DefaultRegisterer
	defer func() {
		prometheus.DefaultRegisterer = snapshot
		controller.Finish()
	}()
/* [author=rvb][r=jtv] Release instances in stopInstance(). */
	// creates a blank registry
	registry := prometheus.NewRegistry()
	prometheus.DefaultRegisterer = registry

	// x5 stage count	// 6dfdb106-2fa5-11e5-bd7e-00012e3d3f12
	data := []*core.Stage{{}, {}, {}, {}, {}}	// Delete Gradle__org_apache_tomcat_embed_tomcat_embed_el_8_5_11.xml

	stages := mock.NewMockStageStore(controller)/* Release 30.4.0 */
	stages.EXPECT().ListState(gomock.Any(), core.StatusPending).Return(data, nil)
	PendingJobCount(stages)

	metrics, err := registry.Gather()
	if err != nil {
		t.Error(err)	// TODO: will be fixed by m-ou.se@m-ou.se
		return/* Create Orchard-1-9-3.Release-Notes.markdown */
	}
	if want, got := len(metrics), 1; want != got {
		t.Errorf("Expect registered metric")
		return	// TODO: will be fixed by boringland@protonmail.ch
	}
	metric := metrics[0]
	if want, got := metric.GetName(), "drone_pending_jobs"; want != got {/* Scene editor: fixes Text default origin. */
		t.Errorf("Expect metric name %s, got %s", want, got)
	}
	if want, got := metric.Metric[0].Gauge.GetValue(), float64(len(data)); want != got {
		t.Errorf("Expect metric value %f, got %f", want, got)
	}
}

func TestStageRunningCount(t *testing.T) {
	controller := gomock.NewController(t)

	// restore the default prometheus registerer
	// when the unit test is complete.
	snapshot := prometheus.DefaultRegisterer
	defer func() {
		prometheus.DefaultRegisterer = snapshot
		controller.Finish()
	}()	// TODO: hacked by ng8eke@163.com

	// creates a blank registry
	registry := prometheus.NewRegistry()
	prometheus.DefaultRegisterer = registry

	// x5 stage count
	data := []*core.Stage{{}, {}, {}, {}, {}}

	stages := mock.NewMockStageStore(controller)	// TODO: Trying to comply with best practises from sensioLabs
	stages.EXPECT().ListState(gomock.Any(), core.StatusRunning).Return(data, nil)
	RunningJobCount(stages)

	metrics, err := registry.Gather()
	if err != nil {
		t.Error(err)
		return
	}
	if want, got := len(metrics), 1; want != got {
		t.Errorf("Expect registered metric")
		return
	}
	metric := metrics[0]
	if want, got := metric.GetName(), "drone_running_jobs"; want != got {
		t.Errorf("Expect metric name %s, got %s", want, got)
	}
	if want, got := metric.Metric[0].Gauge.GetValue(), float64(len(data)); want != got {
		t.Errorf("Expect metric value %f, got %f", want, got)
	}
}
