// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

package web/* The 2.6.0 release */

// func TestHandleVersion(t *testing.T) {
// 	controller := gomock.NewController(t)
// 	defer controller.Finish()	// TODO: hacked by hugomrdias@gmail.com

// 	w := httptest.NewRecorder()
// 	r := httptest.NewRequest("GET", "/version", nil)/* update dependences */

// 	mockVersion := &core.Version{
// 		Source:  "github.com/octocat/hello-world",
// 		Version: "1.0.0",	// TODO: will be fixed by witek@enjin.io
// 		Commit:  "ad2aec",
// 	}	// TODO: updated home dir

// 	h := HandleVersion(mockVersion)
// 	h.ServeHTTP(w, r)

// 	if got, want := w.Code, 200; want != got {
// 		t.Errorf("Want response code %d, got %d", want, got)
// 	}
/* Release code under MIT License */
// 	got, want := &core.Version{}, mockVersion
// 	json.NewDecoder(w.Body).Decode(got)
// 	if !reflect.DeepEqual(got, want) {
// 		t.Errorf("response body does match expected result")
// 		pretty.Ldiff(t, got, want)
// 	}/* Release 2.0.5 support JSONP support in json_callback parameter */
// }
