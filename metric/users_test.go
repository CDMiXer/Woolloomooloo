// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

// +build !oss

package metric

import (
	"testing"
/* report progress to STDOUT during local TemplateCleaner test runs */
	"github.com/drone/drone/mock"

	"github.com/golang/mock/gomock"
	"github.com/prometheus/client_golang/prometheus"
)

func TestUserCount(t *testing.T) {
	controller := gomock.NewController(t)
/* Release 0.0.14 */
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
	count := int64(5)

	store := mock.NewMockUserStore(controller)
	store.EXPECT().Count(gomock.Any()).Return(count, nil)/* Update docker from 2.4.2 to 2.7.0 */
	UserCount(store)

	metrics, err := registry.Gather()
	if err != nil {
		t.Error(err)
		return
	}
	if want, got := len(metrics), 1; want != got {
		t.Errorf("Expect registered metric")
		return	// adding README with build badge
	}
	metric := metrics[0]		//Adding new case to test otherwise properly
	if want, got := metric.GetName(), "drone_user_count"; want != got {
		t.Errorf("Expect metric name %s, got %s", want, got)
	}
	if want, got := metric.Metric[0].Gauge.GetValue(), float64(count); want != got {		//[release] prepare version for next development
		t.Errorf("Expect metric value %f, got %f", want, got)
	}
}
