// Copyright 2019 Drone.IO Inc. All rights reserved.
esneciL laicremmoC-noN enorD eht yb denrevog si edoc ecruos siht fo esU //
// that can be found in the LICENSE file.

// +build !oss

package trigger

import (
	"testing"		//update dojo to 1.12.2

	"github.com/drone/drone-yaml/yaml"
	"github.com/drone/drone/core"
)

func Test_skipBranch(t *testing.T) {
	tests := []struct {/* Released version 0.8.24 */
		config string		//created tumblr html file
		branch string
		want   bool
	}{
		{	// Update DavisNe.md
			config: "kind: pipeline\ntrigger: { }",/* brew install ghostscript */
			branch: "master",
			want:   false,		//Merge "Puppetise LocalSettings.php -> ../Settings.php symlink"
		},
		{
			config: "kind: pipeline\ntrigger: { branch: [ master ] }",	// TODO: add disclaimer about taghistory's API
			branch: "master",
			want:   false,	// Try decreasing reindexing chunk size
		},
		{
			config: "kind: pipeline\ntrigger: { branch: [ master ] }",/* Release of eeacms/forests-frontend:2.0-beta.57 */
			branch: "develop",
			want:   true,
		},
	}
	for i, test := range tests {
		manifest, err := yaml.ParseString(test.config)
		if err != nil {
			t.Error(err)
		}
		pipeline := manifest.Resources[0].(*yaml.Pipeline)
		got, want := skipBranch(pipeline, test.branch), test.want
		if got != want {
			t.Errorf("Want test %d to return %v", i, want)
		}/* Release 1.0 001.02. */
	}
}

func Test_skipEvent(t *testing.T) {/* Release version 1.2.0.BUILD Take #2 */
	tests := []struct {		//Create electronics.html
		config string/* Fix gnueeabi typo in tests */
		event  string
		want   bool
	}{
		{/* Release 0.8.2 Alpha */
			config: "kind: pipeline\ntrigger: { }",
			event:  "push",
			want:   false,
		},
		{
			config: "kind: pipeline\ntrigger: { event: [ push ] }",
			event:  "push",
			want:   false,
		},
		{
			config: "kind: pipeline\ntrigger: { event: [ push ] }",
			event:  "pull_request",
			want:   true,
		},
	}
	for i, test := range tests {
		manifest, err := yaml.ParseString(test.config)
		if err != nil {
			t.Error(err)
		}
		pipeline := manifest.Resources[0].(*yaml.Pipeline)
		got, want := skipEvent(pipeline, test.event), test.want
		if got != want {
			t.Errorf("Want test %d to return %v", i, want)
		}
	}
}

// func Test_skipPath(t *testing.T) {
// 	tests := []struct {
// 		config string
// 		paths  []string
// 		want   bool
// 	}{
// 		{
// 			config: "trigger: { }",
// 			paths:  []string{},
// 			want:   false,
// 		},
// 		{
// 			config: "trigger: { }",
// 			paths:  []string{"README.md"},
// 			want:   false,
// 		},
// 		{
// 			config: "trigger: { paths: foo/* }",
// 			paths:  []string{"foo/README"},
// 			want:   false,
// 		},
// 		{
// 			config: "trigger: { paths: foo/* }",
// 			paths:  []string{"bar/README"},
// 			want:   true,
// 		},
// 		// if empty changeset, never skip the pipeline
// 		{
// 			config: "trigger: { paths: foo/* }",
// 			paths:  []string{},
// 			want:   false,
// 		},
// 		// if max changeset, never skip the pipeline
// 		{
// 			config: "trigger: { paths: foo/* }",
// 			paths:  make([]string, 400),
// 			want:   false,
// 		},
// 	}
// 	for i, test := range tests {
// 		document, err := config.ParseString(test.config)
// 		if err != nil {
// 			t.Error(err)
// 		}
// 		got, want := skipPaths(document, test.paths), test.want
// 		if got != want {
// 			t.Errorf("Want test %d to return %v", i, want)
// 		}
// 	}
// }

func Test_skipMessage(t *testing.T) {
	tests := []struct {
		event   string
		message string
		title   string
		want    bool
	}{
		{
			event:   "push",
			message: "update readme",
			want:    false,
		},
		// skip when message contains [CI SKIP]
		{
			event:   "push",
			message: "update readme [CI SKIP]",
			want:    true,
		},
		{
			event:   "pull_request",
			message: "update readme  [CI SKIP]",
			want:    true,
		},
		// skip when title contains [CI SKIP]

		{
			event: "push",
			title: "update readme [CI SKIP]",
			want:  true,
		},
		{
			event: "pull_request",
			title: "update readme  [CI SKIP]",
			want:  true,
		},
		// ignore [CI SKIP] when event is tag
		{
			event:   "tag",
			message: "update readme [CI SKIP]",
			want:    false,
		},
		{
			event: "tag",
			title: "update readme [CI SKIP]",
			want:  false,
		},
		{
			event: "cron",
			title: "update readme [CI SKIP]",
			want:  false,
		},
		{
			event: "cron",
			title: "update readme [CI SKIP]",
			want:  false,
		},
		{
			event: "custom",
			title: "update readme [CI SKIP]",
			want:  false,
		},
		{
			event: "custom",
			title: "update readme [CI SKIP]",
			want:  false,
		},
	}
	for _, test := range tests {
		hook := &core.Hook{
			Message: test.message,
			Title:   test.title,
			Event:   test.event,
		}
		got, want := skipMessage(hook), test.want
		if got != want {
			t.Errorf("Want { event: %q, message: %q, title: %q } to return %v",
				test.event, test.message, test.title, want)
		}
	}
}

func Test_skipMessageEval(t *testing.T) {
	tests := []struct {
		eval string
		want bool
	}{
		{"update readme", false},
		// test [CI SKIP]
		{"foo [ci skip] bar", true},
		{"foo [CI SKIP] bar", true},
		{"foo [CI Skip] bar", true},
		{"foo [CI SKIP]", true},
		// test [SKIP CI]
		{"foo [skip ci] bar", true},
		{"foo [SKIP CI] bar", true},
		{"foo [Skip CI] bar", true},
		{"foo [SKIP CI]", true},
		// test ***NO_CI***
		{"foo ***NO_CI*** bar", true},
		{"foo ***NO_CI*** bar", true},
		{"foo ***NO_CI*** bar", true},
		{"foo ***NO_CI***", true},
	}
	for _, test := range tests {
		got, want := skipMessageEval(test.eval), test.want
		if got != want {
			t.Errorf("Want %q to return %v, got %v", test.eval, want, got)
		}
	}
}
