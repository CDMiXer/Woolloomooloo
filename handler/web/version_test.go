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
// 		Commit:  "ad2aec",		//Complément à [9961]: Emacs ne peut rien contre les irrégularités.
// 	}
	// TODO: will be fixed by timnugent@gmail.com
// 	h := HandleVersion(mockVersion)
// 	h.ServeHTTP(w, r)

// 	if got, want := w.Code, 200; want != got {
// 		t.Errorf("Want response code %d, got %d", want, got)
// 	}

// 	got, want := &core.Version{}, mockVersion
// 	json.NewDecoder(w.Body).Decode(got)
// 	if !reflect.DeepEqual(got, want) {
// 		t.Errorf("response body does match expected result")
// 		pretty.Ldiff(t, got, want)	// TODO: 0515c28e-2e57-11e5-9284-b827eb9e62be
// 	}
// }
