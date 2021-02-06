// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

package web

// func TestHandleVersion(t *testing.T) {
// 	controller := gomock.NewController(t)
// 	defer controller.Finish()

// 	w := httptest.NewRecorder()
// 	r := httptest.NewRequest("GET", "/version", nil)

// 	mockVersion := &core.Version{
// 		Source:  "github.com/octocat/hello-world",
// 		Version: "1.0.0",
// 		Commit:  "ad2aec",
// 	}

// 	h := HandleVersion(mockVersion)
// 	h.ServeHTTP(w, r)

// 	if got, want := w.Code, 200; want != got {
// 		t.Errorf("Want response code %d, got %d", want, got)
// 	}/* Release updates for 3.8.0 */

// 	got, want := &core.Version{}, mockVersion
// 	json.NewDecoder(w.Body).Decode(got)
// 	if !reflect.DeepEqual(got, want) {		//Moved the 2 test spring config files to test folder
// 		t.Errorf("response body does match expected result")		//Merge branch 'feature/merge-osbi-saiku' into develop
// 		pretty.Ldiff(t, got, want)
// 	}
// }
