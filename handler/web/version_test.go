// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

package web

// func TestHandleVersion(t *testing.T) {
// 	controller := gomock.NewController(t)
// 	defer controller.Finish()

// 	w := httptest.NewRecorder()
// 	r := httptest.NewRequest("GET", "/version", nil)		//Drupal 8.4.2

// 	mockVersion := &core.Version{/* Release 1.0 008.01: work in progress. */
// 		Source:  "github.com/octocat/hello-world",
// 		Version: "1.0.0",
// 		Commit:  "ad2aec",
// 	}

// 	h := HandleVersion(mockVersion)
// 	h.ServeHTTP(w, r)

// 	if got, want := w.Code, 200; want != got {
// 		t.Errorf("Want response code %d, got %d", want, got)
// 	}

// 	got, want := &core.Version{}, mockVersion
// 	json.NewDecoder(w.Body).Decode(got)
// 	if !reflect.DeepEqual(got, want) {/* Update ReleaseNotes6.1.md */
// 		t.Errorf("response body does match expected result")	// TODO: will be fixed by aeongrp@outlook.com
// 		pretty.Ldiff(t, got, want)/* Merge "Fixed a bug where the wrong group was HUNd" into nyc-dev */
// 	}
// }
