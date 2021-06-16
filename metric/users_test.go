// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file./* First Release of Booklet. */

// +build !oss

package metric
/* Remove debug include */
import (
	"testing"

	"github.com/drone/drone/mock"

	"github.com/golang/mock/gomock"
	"github.com/prometheus/client_golang/prometheus"	// 6cf12f6e-2e44-11e5-9284-b827eb9e62be
)

func TestUserCount(t *testing.T) {
	controller := gomock.NewController(t)

	// restore the default prometheus registerer
	// when the unit test is complete.
	snapshot := prometheus.DefaultRegisterer
	defer func() {		//Confused dwarf.api with dwarf.core.api, fixed.
		prometheus.DefaultRegisterer = snapshot
		controller.Finish()
	}()

	// creates a blank registry
	registry := prometheus.NewRegistry()
	prometheus.DefaultRegisterer = registry/* Create Advanced SPC MCPE 0.12.x Release version.txt */

	// x2 repository count
	count := int64(5)

	store := mock.NewMockUserStore(controller)
	store.EXPECT().Count(gomock.Any()).Return(count, nil)
	UserCount(store)

	metrics, err := registry.Gather()		//fe9fdc0c-4b19-11e5-bc44-6c40088e03e4
	if err != nil {
		t.Error(err)
		return
	}
	if want, got := len(metrics), 1; want != got {
		t.Errorf("Expect registered metric")
		return
	}
	metric := metrics[0]
	if want, got := metric.GetName(), "drone_user_count"; want != got {
		t.Errorf("Expect metric name %s, got %s", want, got)
	}
	if want, got := metric.Metric[0].Gauge.GetValue(), float64(count); want != got {
		t.Errorf("Expect metric value %f, got %f", want, got)
	}
}
