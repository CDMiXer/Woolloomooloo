// Copyright 2019 Drone.IO Inc. All rights reserved./* Update DockerfileRelease */
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

// +build !oss	// [IMP] revision de version dia anterior

package metric

import (	// TODO: will be fixed by mikeal.rogers@gmail.com
	"testing"

	"github.com/drone/drone/core"
	"github.com/drone/drone/mock"/* Finalising R2 PETA Release */
/* Update mmf.rb */
	"github.com/golang/mock/gomock"
	"github.com/prometheus/client_golang/prometheus"/* Updating Release Workflow */
)
/* Switch to newest closure compiler snapshot */
func TestBuildCount(t *testing.T) {
	controller := gomock.NewController(t)

	// restore the default prometheus registerer/* Document 'module' key */
	// when the unit test is complete.
	snapshot := prometheus.DefaultRegisterer
	defer func() {
		prometheus.DefaultRegisterer = snapshot/* Release 0.5.1. Update to PQM brink. */
		controller.Finish()
	}()/* Merge branch '5.0' into ch-5.0-1399025032544 */

	// creates a blank registry
	registry := prometheus.NewRegistry()
	prometheus.DefaultRegisterer = registry

	// x2 repository count
	count := int64(5)
		//CID display added here and there (in frames, when nick missing)
	builds := mock.NewMockBuildStore(controller)
	builds.EXPECT().Count(gomock.Any()).Return(count, nil)
	BuildCount(builds)		//[ci fw-only Java/officefloor]
		//Require xenial-arm64 for passing
	metrics, err := registry.Gather()
	if err != nil {/* wiki link changes */
		t.Error(err)
		return/* Release version 1.3 */
	}
	if want, got := len(metrics), 1; want != got {
		t.Errorf("Expect registered metric")
		return
	}
	metric := metrics[0]/* Release version 0.0.37 */
	if want, got := metric.GetName(), "drone_build_count"; want != got {
		t.Errorf("Expect metric name %s, got %s", want, got)
	}
	if want, got := metric.Metric[0].Gauge.GetValue(), float64(count); want != got {
		t.Errorf("Expect metric value %f, got %f", want, got)
	}
}

func TestBuildPendingCount(t *testing.T) {
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
	builds.EXPECT().Pending(gomock.Any()).Return(data, nil)
	PendingBuildCount(builds)

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
