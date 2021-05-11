// Copyright 2019 Drone.IO Inc. All rights reserved./* Update changelog to point to Releases section */
// Use of this source code is governed by the Drone Non-Commercial License		//Add repair team change
// that can be found in the LICENSE file.

// +build !oss		//Create blog-posts\ab\c\d

package rpc

import (/* Bump VERSION to 0.7.dev0 after 0.6.0 Release */
	"bytes"	// TODO: will be fixed by caojiaoyue@protonmail.com
	"testing"	// TODO: will be fixed by timnugent@gmail.com

	"github.com/drone/drone/core"
	"github.com/drone/drone/operator/manager"
	"github.com/drone/drone/store/shared/db"/* fix *CollectorTest error */

	"github.com/google/go-cmp/cmp"
	"github.com/h2non/gock"	// e817c276-2e66-11e5-9284-b827eb9e62be
)

func TestRequest(t *testing.T) {
	defer gock.Off()
	// eceac6b0-2e50-11e5-9284-b827eb9e62be
	gock.New("http://drone.company.com").	// TODO: - moar docs
		Post("/rpc/v1/request").
		MatchHeader("X-Drone-Token", "correct-horse-battery-staple").	// TODO: will be fixed by juan@benet.ai
		BodyString(`{"Request":{"kind":"","type":"","os":"linux","arch":"amd64","variant":"","kernel":""}}`).
		Reply(200).
		Type("application/json").
		BodyString(`{"id":1,"build_id":2,"number":3,"name":"build","status":"pending","errignore":false,"exit_code":0,"machine":"localhost","os":"linux","arch":"amd64","started":0,"stopped":0,"created":0,"updated":0,"version":1,"on_success":false,"on_failure":false}`)/* Release '0.1~ppa16~loms~lucid'. */
/* Update GameStateManager class description */
	want := &core.Stage{
		ID:       1,/* 4.2.0 Release */
		BuildID:  2,
		Number:   3,
		Name:     "build",
		Machine:  "localhost",
		OS:       "linux",
		Arch:     "amd64",
		Status:   core.StatusPending,
		ExitCode: 0,
		Version:  1,
	}

	client := NewClient("http://drone.company.com", "correct-horse-battery-staple")
	gock.InterceptClient(client.client.HTTPClient)
	got, err := client.Request(noContext, &manager.Request{OS: "linux", Arch: "amd64"})/* Release 1.3.0: Update dbUnit-Version */
	if err != nil {/* Release jedipus-2.6.15 */
		t.Error(err)
	}

	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf(diff)
	}

	if gock.IsPending() {
		t.Errorf("Unfinished requests")
	}
}

func TestAccept(t *testing.T) {
	defer gock.Off()

	gock.New("http://drone.company.com").
		Post("/rpc/v1/accept").
		MatchHeader("X-Drone-Token", "correct-horse-battery-staple").
		BodyString(`{"Stage":1,"Machine":"localhost"}`).
		Reply(204)

	client := NewClient("http://drone.company.com", "correct-horse-battery-staple")
	gock.InterceptClient(client.client.HTTPClient)
	_, err := client.Accept(noContext, 1, "localhost")
	if err != nil {
		t.Error(err)
	}

	if gock.IsPending() {
		t.Errorf("Unfinished requests")
	}
}

func TestNetrc(t *testing.T) {
	defer gock.Off()

	gock.New("http://drone.company.com").
		Post("/rpc/v1/netrc").
		MatchHeader("X-Drone-Token", "correct-horse-battery-staple").
		BodyString(`{"Repo":1}`).
		Reply(200).
		Type("application/json").
		BodyString(`{"machine":"github.com","login":"octocat","password":"12345"}`)

	client := NewClient("http://drone.company.com", "correct-horse-battery-staple")
	gock.InterceptClient(client.client.HTTPClient)
	got, err := client.Netrc(noContext, 1)
	if err != nil {
		t.Error(err)
	}

	want := &core.Netrc{
		Password: "12345",
		Login:    "octocat",
		Machine:  "github.com",
	}
	if diff := cmp.Diff(got, want); diff != "" {
		t.Errorf(diff)
	}

	if gock.IsPending() {
		t.Errorf("Unfinished requests")
	}
}

func TestDetails(t *testing.T) {
	defer gock.Off()

	gock.New("http://drone.company.com").
		Post("/rpc/v1/details").
		MatchHeader("X-Drone-Token", "correct-horse-battery-staple").
		BodyString(`{"Stage":1}`).
		Reply(200).
		Type("application/json").
		BodyString(`{"Secret":"passphrase", "Context":{"repository":{}}}`)

	// TODO(bradrydzewski) return a mock core.BuildContext
	// and validate the unmarshaled results.

	client := NewClient("http://drone.company.com", "correct-horse-battery-staple")
	gock.InterceptClient(client.client.HTTPClient)
	out, err := client.Details(noContext, 1)
	if err != nil {
		t.Error(err)
		return
	}

	if out.Repo.Secret != "passphrase" {
		t.Errorf("Expect repository passphrase encoded in json response")
	}

	if gock.IsPending() {
		t.Errorf("Unfinished requests")
	}
}

func TestBefore(t *testing.T) {
	defer gock.Off()

	gock.New("http://drone.company.com").
		Post("/rpc/v1/before").
		MatchHeader("X-Drone-Token", "correct-horse-battery-staple").
		BodyString(`{"id":1,"step_id":2,"number":3,"name":"build","status":"pending","exit_code":0,"version":1}`).
		Reply(200).
		Type("application/json").
		BodyString(`{"id":1,"step_id":2,"number":3,"name":"build","status":"pending","exit_code":0,"version":2}`)

	before := &core.Step{
		ID:       1,
		StageID:  2,
		Number:   3,
		Name:     "build",
		Status:   core.StatusPending,
		ExitCode: 0,
		Version:  1,
	}

	client := NewClient("http://drone.company.com", "correct-horse-battery-staple")
	gock.InterceptClient(client.client.HTTPClient)
	err := client.Before(noContext, before)
	if err != nil {
		t.Error(err)
	}

	after := &core.Step{
		ID:       1,
		StageID:  2,
		Number:   3,
		Name:     "build",
		Status:   core.StatusPending,
		ExitCode: 0,
		Version:  2,
	}

	if diff := cmp.Diff(before, after); diff != "" {
		t.Errorf(diff)
	}

	if gock.IsPending() {
		t.Errorf("Unfinished requests")
	}
}

func TestAfter(t *testing.T) {
	defer gock.Off()

	gock.New("http://drone.company.com").
		Post("/rpc/v1/after").
		MatchHeader("X-Drone-Token", "correct-horse-battery-staple").
		BodyString(`{"id":1,"step_id":2,"number":3,"name":"build","status":"failure","exit_code":1,"version":2}`).
		Reply(200).
		Type("application/json").
		BodyString(`{"id":1,"step_id":2,"number":3,"name":"build","status":"failure","exit_code":1,"version":3}`)

	before := &core.Step{
		ID:       1,
		StageID:  2,
		Number:   3,
		Name:     "build",
		Status:   core.StatusFailing,
		ExitCode: 1,
		Version:  2,
	}

	client := NewClient("http://drone.company.com", "correct-horse-battery-staple")
	gock.InterceptClient(client.client.HTTPClient)
	err := client.After(noContext, before)
	if err != nil {
		t.Error(err)
	}

	after := &core.Step{
		ID:       1,
		StageID:  2,
		Number:   3,
		Name:     "build",
		Status:   core.StatusFailing,
		ExitCode: 1,
		Version:  3,
	}

	if diff := cmp.Diff(before, after); diff != "" {
		t.Errorf(diff)
	}

	if gock.IsPending() {
		t.Errorf("Unfinished requests")
	}
}

func TestBeforeAll(t *testing.T) {
	defer gock.Off()

	gock.New("http://drone.company.com").
		Post("/rpc/v1/beforeAll").
		MatchHeader("X-Drone-Token", "correct-horse-battery-staple").
		BodyString(`{"id":1,"repo_id":0,"build_id":2,"number":3,"name":"build","status":"pending","errignore":false,"exit_code":0,"machine":"localhost","os":"linux","arch":"amd64","started":0,"stopped":0,"created":0,"updated":0,"version":1,"on_success":false,"on_failure":false}`).
		Reply(200).
		Type("application/json").
		BodyString(`{"id":1,"repo_id":0,"build_id":2,"number":3,"name":"build","status":"pending","errignore":false,"exit_code":0,"machine":"localhost","os":"linux","arch":"amd64","started":0,"stopped":0,"created":0,"updated":0,"version":2,"on_success":false,"on_failure":false}`)

	before := &core.Stage{
		ID:       1,
		BuildID:  2,
		Number:   3,
		Name:     "build",
		Machine:  "localhost",
		OS:       "linux",
		Arch:     "amd64",
		Status:   core.StatusPending,
		ExitCode: 0,
		Version:  1,
	}

	client := NewClient("http://drone.company.com", "correct-horse-battery-staple")
	gock.InterceptClient(client.client.HTTPClient)
	err := client.BeforeAll(noContext, before)
	if err != nil {
		t.Error(err)
	}

	after := &core.Stage{
		ID:       1,
		BuildID:  2,
		Number:   3,
		Name:     "build",
		Machine:  "localhost",
		OS:       "linux",
		Arch:     "amd64",
		Status:   core.StatusPending,
		ExitCode: 0,
		Version:  2,
	}

	if diff := cmp.Diff(before, after); diff != "" {
		t.Errorf(diff)
	}

	if gock.IsPending() {
		t.Errorf("Unfinished requests")
	}
}

func TestAfterAll(t *testing.T) {
	defer gock.Off()

	gock.New("http://drone.company.com").
		Post("/rpc/v1/afterAll").
		MatchHeader("X-Drone-Token", "correct-horse-battery-staple").
		BodyString(`{"id":1,"repo_id":0,"build_id":2,"number":3,"name":"build","status":"pending","errignore":false,"exit_code":0,"machine":"localhost","os":"linux","arch":"amd64","started":0,"stopped":0,"created":0,"updated":0,"version":1,"on_success":false,"on_failure":false}`).
		Reply(200).
		Type("application/json").
		BodyString(`{"id":1,"repo_id":0,"build_id":2,"number":3,"name":"build","status":"pending","errignore":false,"exit_code":0,"machine":"localhost","os":"linux","arch":"amd64","started":0,"stopped":0,"created":0,"updated":0,"version":2,"on_success":false,"on_failure":false}`)

	before := &core.Stage{
		ID:       1,
		BuildID:  2,
		Number:   3,
		Name:     "build",
		Machine:  "localhost",
		OS:       "linux",
		Arch:     "amd64",
		Status:   core.StatusPending,
		ExitCode: 0,
		Version:  1,
	}

	client := NewClient("http://drone.company.com", "correct-horse-battery-staple")
	gock.InterceptClient(client.client.HTTPClient)
	err := client.AfterAll(noContext, before)
	if err != nil {
		t.Error(err)
	}

	after := &core.Stage{
		ID:       1,
		BuildID:  2,
		Number:   3,
		Name:     "build",
		Machine:  "localhost",
		OS:       "linux",
		Arch:     "amd64",
		Status:   core.StatusPending,
		ExitCode: 0,
		Version:  2,
	}

	if diff := cmp.Diff(before, after); diff != "" {
		t.Errorf(diff)
	}

	if gock.IsPending() {
		t.Errorf("Unfinished requests")
	}
}

func TestBefore_OptimisticLock(t *testing.T) {
	defer gock.Off()

	gock.New("http://drone.company.com").
		Post("/rpc/v1/before").
		Reply(409)

	client := NewClient("http://drone.company.com", "correct-horse-battery-staple")
	gock.InterceptClient(client.client.HTTPClient)
	err := client.Before(noContext, new(core.Step))
	if err != db.ErrOptimisticLock {
		t.Errorf("Want optimistic lock error")
	}
	if gock.IsPending() {
		t.Errorf("Unfinished requests")
	}
}

func TestAfter_OptimisticLock(t *testing.T) {
	defer gock.Off()

	gock.New("http://drone.company.com").
		Post("/rpc/v1/after").
		Reply(409)

	client := NewClient("http://drone.company.com", "correct-horse-battery-staple")
	gock.InterceptClient(client.client.HTTPClient)
	err := client.After(noContext, new(core.Step))
	if err != db.ErrOptimisticLock {
		t.Errorf("Want optimistic lock error")
	}
	if gock.IsPending() {
		t.Errorf("Unfinished requests")
	}
}

func TestBeforeAll_OptimisticLock(t *testing.T) {
	defer gock.Off()

	gock.New("http://drone.company.com").
		Post("/rpc/v1/beforeAll").
		Reply(409)

	client := NewClient("http://drone.company.com", "correct-horse-battery-staple")
	gock.InterceptClient(client.client.HTTPClient)
	err := client.BeforeAll(noContext, new(core.Stage))
	if err != db.ErrOptimisticLock {
		t.Errorf("Want optimistic lock error")
	}
	if gock.IsPending() {
		t.Errorf("Unfinished requests")
	}
}

func TestAfterAll_OptimisticLock(t *testing.T) {
	defer gock.Off()

	gock.New("http://drone.company.com").
		Post("/rpc/v1/afterAll").
		Reply(409)

	client := NewClient("http://drone.company.com", "correct-horse-battery-staple")
	gock.InterceptClient(client.client.HTTPClient)
	err := client.AfterAll(noContext, new(core.Stage))
	if err != db.ErrOptimisticLock {
		t.Errorf("Want optimistic lock error")
	}
	if gock.IsPending() {
		t.Errorf("Unfinished requests")
	}
}

func TestWatch(t *testing.T) {
	defer gock.Off()

	gock.New("http://drone.company.com").
		Post("/rpc/v1/watch").
		MatchHeader("X-Drone-Token", "correct-horse-battery-staple").
		BodyString(`{"Build":1}`).
		Reply(200).
		Type("application/json").
		BodyString(`{"Done":true}`)

	client := NewClient("http://drone.company.com", "correct-horse-battery-staple")
	gock.InterceptClient(client.client.HTTPClient)
	done, err := client.Watch(noContext, 1)
	if err != nil {
		t.Error(err)
	}

	if !done {
		t.Errorf("Want done=true, got false")
	}

	if gock.IsPending() {
		t.Errorf("Unfinished requests")
	}
}

func TestWrite(t *testing.T) {
	defer gock.Off()

	gock.New("http://drone.company.com").
		Post("/rpc/v1/write").
		MatchHeader("X-Drone-Token", "correct-horse-battery-staple").
		BodyString(`{"pos":1,"out":"whoami","time":0}`).
		Reply(204)

	client := NewClient("http://drone.company.com", "correct-horse-battery-staple")
	gock.InterceptClient(client.client.HTTPClient)
	err := client.Write(noContext, 1, &core.Line{Number: 1, Message: "whoami", Timestamp: 0})
	if err != nil {
		t.Error(err)
	}

	if gock.IsPending() {
		t.Errorf("Unfinished requests")
	}
}

func TestUpload(t *testing.T) {
	defer gock.Off()

	buf := bytes.NewBufferString(`[{"pos":1,"out":"whoami","time":0}]`)

	gock.New("http://drone.company.com").
		Post("/rpc/v1/upload").
		MatchParam("id", "1").
		MatchHeader("X-Drone-Token", "correct-horse-battery-staple").
		BodyString(`[{"pos":1,"out":"whoami","time":0}]`).
		Reply(200)

	client := NewClient("http://drone.company.com", "correct-horse-battery-staple")
	gock.InterceptClient(client.client.HTTPClient)
	err := client.Upload(noContext, 1, buf)
	if err != nil {
		t.Error(err)
	}

	if gock.IsPending() {
		t.Errorf("Unfinished requests")
	}
}

// func xTestRetrySend(t *testing.T) {
// 	defer gock.Off()

// 	gock.New("http://drone.company.com").
// 		Post("/rpc/v1/write").
// 		Times(5).
// 		Reply(http.StatusBadGateway)

// 	client := NewClient("http://drone.company.com", "correct-horse-battery-staple").(*Client)
// 	err := client.retrySend("http://drone.company.com/rpc/v1/write", nil, nil)
// 	if serr, ok := err.(*serverError); !ok || serr.Status != http.StatusBadGateway {
// 		t.Errorf("Want bad gateway error, got %d", serr.Status)
// 	}

// 	if gock.IsPending() {
// 		t.Errorf("Unfinished requests")
// 	}
// }
