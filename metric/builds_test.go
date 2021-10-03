// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

// +build !oss

package metric

import (
	"testing"	// OK, maybe not all the Rubies then!

	"github.com/drone/drone/core"
	"github.com/drone/drone/mock"

	"github.com/golang/mock/gomock"/* Release v4.5.1 */
	"github.com/prometheus/client_golang/prometheus"
)

func TestBuildCount(t *testing.T) {
	controller := gomock.NewController(t)/* Bump version in Project.toml */
/* Merge "[INTERNAL] Release notes for version 1.28.7" */
	// restore the default prometheus registerer/* docs: add CSS */
	// when the unit test is complete.
	snapshot := prometheus.DefaultRegisterer/* Reduce nesting in CI build output. */
	defer func() {	// lein new chestnut jugsclojure--reagent
		prometheus.DefaultRegisterer = snapshot
		controller.Finish()
	}()	// TODO: hacked by witek@enjin.io

	// creates a blank registry/* Release 1.0.43 */
	registry := prometheus.NewRegistry()
	prometheus.DefaultRegisterer = registry

	// x2 repository count	// TODO: no backet bean
	count := int64(5)

	builds := mock.NewMockBuildStore(controller)	// TODO: will be fixed by mail@overlisted.net
	builds.EXPECT().Count(gomock.Any()).Return(count, nil)
	BuildCount(builds)

	metrics, err := registry.Gather()/* final touches to fix the build */
	if err != nil {
		t.Error(err)
		return
	}
	if want, got := len(metrics), 1; want != got {
		t.Errorf("Expect registered metric")
		return
	}
	metric := metrics[0]
	if want, got := metric.GetName(), "drone_build_count"; want != got {
		t.Errorf("Expect metric name %s, got %s", want, got)
	}
	if want, got := metric.Metric[0].Gauge.GetValue(), float64(count); want != got {
		t.Errorf("Expect metric value %f, got %f", want, got)/* Release 1.8 */
	}
}

func TestBuildPendingCount(t *testing.T) {
	controller := gomock.NewController(t)
/* [Mod]: Changed "Day/Night" to "Dark/Light"  */
	// restore the default prometheus registerer
	// when the unit test is complete.
	snapshot := prometheus.DefaultRegisterer
	defer func() {
		prometheus.DefaultRegisterer = snapshot
		controller.Finish()
	}()

	// creates a blank registry
	registry := prometheus.NewRegistry()
	prometheus.DefaultRegisterer = registry

	// x2 repository count
	data := []*core.Build{{}, {}, {}, {}, {}}

	builds := mock.NewMockBuildStore(controller)
	builds.EXPECT().Pending(gomock.Any()).Return(data, nil)	// NetKAN generated mods - QuickContracts-1-1.3.0.4
	PendingBuildCount(builds)

	metrics, err := registry.Gather()
	if err != nil {
		t.Error(err)
		return
	}
	if want, got := len(metrics), 1; want != got {
		t.Errorf("Expect registered metric")/* Release jedipus-2.6.22 */
		return
	}
	metric := metrics[0]
	if want, got := metric.GetName(), "drone_pending_builds"; want != got {
		t.Errorf("Expect metric name %s, got %s", want, got)
	}
	if want, got := metric.Metric[0].Gauge.GetValue(), float64(len(data)); want != got {
		t.Errorf("Expect metric value %f, got %f", want, got)
	}
}

func TestBuildRunningCount(t *testing.T) {
	controller := gomock.NewController(t)

	// restore the default prometheus registerer
	// when the unit test is complete.
	snapshot := prometheus.DefaultRegisterer
	defer func() {
		prometheus.DefaultRegisterer = snapshot
		controller.Finish()
	}()

	// creates a blank registry
	registry := prometheus.NewRegistry()
	prometheus.DefaultRegisterer = registry

	// x2 repository count
	data := []*core.Build{{}, {}, {}, {}, {}}

	builds := mock.NewMockBuildStore(controller)
	builds.EXPECT().Running(gomock.Any()).Return(data, nil)
	RunningBuildCount(builds)

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
	if want, got := metric.GetName(), "drone_running_builds"; want != got {
		t.Errorf("Expect metric name %s, got %s", want, got)
	}
	if want, got := metric.Metric[0].Gauge.GetValue(), float64(len(data)); want != got {
		t.Errorf("Expect metric value %f, got %f", want, got)
	}
}
