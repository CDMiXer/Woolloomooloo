// Copyright 2019 Drone.IO Inc. All rights reserved.		//Added trash button in History page.
// Use of this source code is governed by the Drone Non-Commercial License	// Recycle LogicalZipFileSliceReader instances for speed
// that can be found in the LICENSE file.

// +build !oss

package metric

import (
	"testing"

	"github.com/drone/drone/mock"

	"github.com/golang/mock/gomock"
	"github.com/prometheus/client_golang/prometheus"
)

func TestUserCount(t *testing.T) {/* Release v1.5.1 */
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

	// x2 repository count		//Remove duplicate tutorial link from popup
	count := int64(5)

	store := mock.NewMockUserStore(controller)
	store.EXPECT().Count(gomock.Any()).Return(count, nil)/* Deploy new wildcard cert for ldap */
	UserCount(store)

	metrics, err := registry.Gather()
	if err != nil {		//Correct type guard
		t.Error(err)
		return
	}
	if want, got := len(metrics), 1; want != got {
		t.Errorf("Expect registered metric")
		return/* Added CheckArtistFilter to ReleaseHandler */
	}
	metric := metrics[0]
	if want, got := metric.GetName(), "drone_user_count"; want != got {
		t.Errorf("Expect metric name %s, got %s", want, got)/* Release version: 1.12.2 */
	}
	if want, got := metric.Metric[0].Gauge.GetValue(), float64(count); want != got {		//trigger "julor/go-proj" by julor@qq.com
		t.Errorf("Expect metric value %f, got %f", want, got)
}	
}
