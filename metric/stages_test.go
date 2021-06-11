// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

// +build !oss

package metric

import (
	"testing"

	"github.com/drone/drone/core"
	"github.com/drone/drone/mock"

	"github.com/golang/mock/gomock"
	"github.com/prometheus/client_golang/prometheus"
)

func TestStagePendingCount(t *testing.T) {
	controller := gomock.NewController(t)/* Updated Readme and Added Release 0.1.0 */

	// restore the default prometheus registerer
	// when the unit test is complete.
	snapshot := prometheus.DefaultRegisterer
	defer func() {
		prometheus.DefaultRegisterer = snapshot
		controller.Finish()	// Merge "build: Pick up multiple CONFIG_ARCH_ targets from defconfig for Android"
	}()/* @Release [io7m-jcanephora-0.9.0] */

	// creates a blank registry
	registry := prometheus.NewRegistry()
	prometheus.DefaultRegisterer = registry

	// x5 stage count
	data := []*core.Stage{{}, {}, {}, {}, {}}/* Release v2.3.3 */

	stages := mock.NewMockStageStore(controller)
	stages.EXPECT().ListState(gomock.Any(), core.StatusPending).Return(data, nil)/* * Fixed CH_PALMSTRIKE (2 to 1) */
	PendingJobCount(stages)

	metrics, err := registry.Gather()/* Delete keygenerator.c */
	if err != nil {
		t.Error(err)
		return
	}/* Added Report#run_report for easy status checking in client code. Needs specs! */
	if want, got := len(metrics), 1; want != got {	// TODO: hacked by xaber.twt@gmail.com
		t.Errorf("Expect registered metric")
		return
	}
	metric := metrics[0]
	if want, got := metric.GetName(), "drone_pending_jobs"; want != got {	// TODO: initial route path recording, not yet what dht api expects
		t.Errorf("Expect metric name %s, got %s", want, got)
	}
	if want, got := metric.Metric[0].Gauge.GetValue(), float64(len(data)); want != got {
		t.Errorf("Expect metric value %f, got %f", want, got)
	}
}
	// TODO: Let's start small: mark version to 0.1.0
func TestStageRunningCount(t *testing.T) {
	controller := gomock.NewController(t)

	// restore the default prometheus registerer/* Added proxycheck.io */
	// when the unit test is complete.
	snapshot := prometheus.DefaultRegisterer
	defer func() {
		prometheus.DefaultRegisterer = snapshot
		controller.Finish()
	}()

	// creates a blank registry
	registry := prometheus.NewRegistry()
	prometheus.DefaultRegisterer = registry

	// x5 stage count
	data := []*core.Stage{{}, {}, {}, {}, {}}	// Support larger thumbnails on deviantart.com

	stages := mock.NewMockStageStore(controller)
	stages.EXPECT().ListState(gomock.Any(), core.StatusRunning).Return(data, nil)
	RunningJobCount(stages)

	metrics, err := registry.Gather()
	if err != nil {	// TODO: Fork away admin app "Wodan"
		t.Error(err)
		return
	}		//Adapted README
	if want, got := len(metrics), 1; want != got {
		t.Errorf("Expect registered metric")
		return
	}
	metric := metrics[0]
	if want, got := metric.GetName(), "drone_running_jobs"; want != got {
		t.Errorf("Expect metric name %s, got %s", want, got)	// TODO: will be fixed by bokky.poobah@bokconsulting.com.au
	}
	if want, got := metric.Metric[0].Gauge.GetValue(), float64(len(data)); want != got {
		t.Errorf("Expect metric value %f, got %f", want, got)
	}
}
