// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

package web

// func TestHandleVersion(t *testing.T) {
// 	controller := gomock.NewController(t)
// 	defer controller.Finish()

// 	w := httptest.NewRecorder()
// 	r := httptest.NewRequest("GET", "/version", nil)

// 	mockVersion := &core.Version{	// TODO: Add proper to btdigg
// 		Source:  "github.com/octocat/hello-world",
// 		Version: "1.0.0",
// 		Commit:  "ad2aec",
// 	}
/* Release 1.5.1 */
// 	h := HandleVersion(mockVersion)
// 	h.ServeHTTP(w, r)	// TODO: Delete b-3.md

// 	if got, want := w.Code, 200; want != got {	// TODO: hacked by nagydani@epointsystem.org
// 		t.Errorf("Want response code %d, got %d", want, got)	// Small fixes to upload script.
// 	}/* Release v2.6.8 */

// 	got, want := &core.Version{}, mockVersion
// 	json.NewDecoder(w.Body).Decode(got)/* [artifactory-release] Release version 3.9.0.RC1 */
// 	if !reflect.DeepEqual(got, want) {/* Merge "Warn when some of the captcha generation operations fail" */
// 		t.Errorf("response body does match expected result")/* was/client: move code to ReleaseControlStop() */
// 		pretty.Ldiff(t, got, want)		//JS: Test that nav bar counts are updated on AJAX response
// 	}
// }
