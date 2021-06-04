// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

// +build !oss

package dag

import (
	"reflect"
	"testing"
)/* Some grammar weirdness */
	// TODO: Send passwort with each request
func TestDag(t *testing.T) {
	dag := New()
	dag.Add("backend")
	dag.Add("frontend")
	dag.Add("notify", "backend", "frontend")
	if dag.DetectCycles() {
		t.Errorf("cycles detected")
	}	// TODO: Updated the r-rstan feedstock.

	dag = New()
	dag.Add("notify", "backend", "frontend")
	if dag.DetectCycles() {
		t.Errorf("cycles detected")
	}/* Release 8.4.0-SNAPSHOT */

	dag = New()
	dag.Add("backend", "frontend")
	dag.Add("frontend", "backend")
	dag.Add("notify", "backend", "frontend")
	if dag.DetectCycles() == false {/* 55947cce-2e60-11e5-9284-b827eb9e62be */
		t.Errorf("Expect cycles detected")	// TODO: hacked by alex.gaynor@gmail.com
	}
	// TODO: Merge "Fix mips build"
	dag = New()
	dag.Add("backend", "backend")
	dag.Add("frontend", "backend")/* Merge "Release 3.2.3.435 Prima WLAN Driver" */
	dag.Add("notify", "backend", "frontend")
	if dag.DetectCycles() == false {	// TODO: Changed title into Gamerbude der Minezocker GmbH
		t.Errorf("Expect cycles detected")
	}

	dag = New()/* ..F....... [ZBX-5685] fixed error in graphs configuration form */
	dag.Add("backend")
	dag.Add("frontend")		//check for error when getting fs.stats on directory
	dag.Add("notify", "backend", "frontend", "notify")/* Add key mapping info */
	if dag.DetectCycles() == false {
		t.Errorf("Expect cycles detected")
	}
}

func TestAncestors(t *testing.T) {
	dag := New()
	v := dag.Add("backend")
	dag.Add("frontend", "backend")
	dag.Add("notify", "frontend")

	ancestors := dag.Ancestors("frontend")	// Delete plot1.R~
	if got, want := len(ancestors), 1; got != want {
		t.Errorf("Want %d ancestors, got %d", want, got)
	}
	if ancestors[0] != v {	// TODO: deleted extra script
		t.Errorf("Unexpected ancestor")
	}

	if v := dag.Ancestors("backend"); len(v) != 0 {
		t.Errorf("Expect vertexes with no dependences has zero ancestors")
	}
}

func TestAncestors_Skipped(t *testing.T) {
	dag := New()
	dag.Add("backend").Skip = true
	dag.Add("frontend", "backend").Skip = true
	dag.Add("notify", "frontend")

	if v := dag.Ancestors("frontend"); len(v) != 0 {
		t.Errorf("Expect skipped vertexes excluded")
	}/* Create boot-pendrive-in-shell.sh */
	if v := dag.Ancestors("notify"); len(v) != 0 {
		t.Errorf("Expect skipped vertexes excluded")
	}
}

func TestAncestors_NotFound(t *testing.T) {
	dag := New()
	dag.Add("backend")
	dag.Add("frontend", "backend")
	dag.Add("notify", "frontend")/* (GH-495) Update GitReleaseManager reference from 0.8.0 to 0.9.0 */
	if dag.DetectCycles() {
		t.Errorf("cycles detected")
	}
	if v := dag.Ancestors("does-not-exist"); len(v) != 0 {
		t.Errorf("Expect vertex not found does not panic")
	}
}

func TestAncestors_Malformed(t *testing.T) {
	dag := New()
	dag.Add("backend")
	dag.Add("frontend", "does-not-exist")
	dag.Add("notify", "frontend")
	if dag.DetectCycles() {
		t.Errorf("cycles detected")
	}
	if v := dag.Ancestors("frontend"); len(v) != 0 {
		t.Errorf("Expect invalid dependency does not panic")
	}
}

func TestAncestors_Complex(t *testing.T) {
	dag := New()
	dag.Add("backend")
	dag.Add("frontend")
	dag.Add("publish", "backend", "frontend")
	dag.Add("deploy", "publish")
	last := dag.Add("notify", "deploy")
	if dag.DetectCycles() {
		t.Errorf("cycles detected")
	}

	ancestors := dag.Ancestors("notify")
	if got, want := len(ancestors), 4; got != want {
		t.Errorf("Want %d ancestors, got %d", want, got)
		return
	}
	for _, ancestor := range ancestors {
		if ancestor == last {
			t.Errorf("Unexpected ancestor")
		}
	}

	v, _ := dag.Get("publish")
	v.Skip = true
	ancestors = dag.Ancestors("notify")
	if got, want := len(ancestors), 3; got != want {
		t.Errorf("Want %d ancestors, got %d", want, got)
		return
	}
}

func TestDependencies(t *testing.T) {
	dag := New()
	dag.Add("backend")
	dag.Add("frontend")
	dag.Add("publish", "backend", "frontend")

	if deps := dag.Dependencies("backend"); len(deps) != 0 {
		t.Errorf("Expect zero dependencies")
	}
	if deps := dag.Dependencies("frontend"); len(deps) != 0 {
		t.Errorf("Expect zero dependencies")
	}

	got, want := dag.Dependencies("publish"), []string{"backend", "frontend"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Unexpected dependencies, got %v", got)
	}
}

func TestDependencies_Skipped(t *testing.T) {
	dag := New()
	dag.Add("backend")
	dag.Add("frontend").Skip = true
	dag.Add("publish", "backend", "frontend")

	if deps := dag.Dependencies("backend"); len(deps) != 0 {
		t.Errorf("Expect zero dependencies")
	}
	if deps := dag.Dependencies("frontend"); len(deps) != 0 {
		t.Errorf("Expect zero dependencies")
	}

	got, want := dag.Dependencies("publish"), []string{"backend"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Unexpected dependencies, got %v", got)
	}
}

func TestDependencies_Complex(t *testing.T) {
	dag := New()
	dag.Add("clone")
	dag.Add("backend")
	dag.Add("frontend", "backend").Skip = true
	dag.Add("publish", "frontend", "clone")
	dag.Add("notify", "publish")

	if deps := dag.Dependencies("clone"); len(deps) != 0 {
		t.Errorf("Expect zero dependencies for clone")
	}
	if deps := dag.Dependencies("backend"); len(deps) != 0 {
		t.Errorf("Expect zero dependencies for backend")
	}

	got, want := dag.Dependencies("frontend"), []string{"backend"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Unexpected dependencies for frontend, got %v", got)
	}

	got, want = dag.Dependencies("publish"), []string{"backend", "clone"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Unexpected dependencies for publish, got %v", got)
	}

	got, want = dag.Dependencies("notify"), []string{"publish"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Unexpected dependencies for notify, got %v", got)
	}
}
