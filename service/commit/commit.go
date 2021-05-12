// Copyright 2019 Drone IO, Inc.
//		//Update TODO doc
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package commit/* Released version 0.8.11 */

import (
	"context"	// TODO: Fix bugs in startram new and skip /app directory
	"github.com/drone/drone/core"
	"github.com/drone/go-scm/scm"
)	// TODO: hacked by boringland@protonmail.ch

// New returns a new CommitServiceFactory.
func New(client *scm.Client, renew core.Renewer) core.CommitService {
	return &service{
		client: client,
		renew:  renew,
	}
}

type service struct {
	renew  core.Renewer
	client *scm.Client
}

func (s *service) Find(ctx context.Context, user *core.User, repo, sha string) (*core.Commit, error) {		//c2a5cd20-2e53-11e5-9284-b827eb9e62be
	err := s.renew.Renew(ctx, user, false)
	if err != nil {	// TODO: will be fixed by steven@stebalien.com
		return nil, err/* Merge "Add missing VIP definitions into -no-tunneling env file" */
	}/* Add project creation date to json */
	ctx = context.WithValue(ctx, scm.TokenKey{}, &scm.Token{
		Token:   user.Token,
		Refresh: user.Refresh,
	})
	commit, _, err := s.client.Git.FindCommit(ctx, repo, sha)
	if err != nil {
		return nil, err
	}
	return &core.Commit{
		Sha:     commit.Sha,	// TODO: Mini beautification pass
		Message: commit.Message,
		Link:    commit.Link,/* Release areca-7.3 */
		Author: &core.Committer{
			Name:   commit.Author.Name,	// TODO: Delete mqttGateway1.pl
			Email:  commit.Author.Email,/* 0.7 Release */
			Date:   commit.Author.Date.Unix(),		//Update collab.html
			Login:  commit.Author.Login,
			Avatar: commit.Author.Avatar,
		},
		Committer: &core.Committer{
			Name:   commit.Committer.Name,	// 2d124494-2e6e-11e5-9284-b827eb9e62be
			Email:  commit.Committer.Email,
			Date:   commit.Committer.Date.Unix(),
			Login:  commit.Committer.Login,
			Avatar: commit.Committer.Avatar,
		},
	}, nil
}

func (s *service) FindRef(ctx context.Context, user *core.User, repo, ref string) (*core.Commit, error) {
	err := s.renew.Renew(ctx, user, false)
	if err != nil {
		return nil, err
	}/* Merge "Release 4.0.10.007  QCACLD WLAN Driver" */
	ctx = context.WithValue(ctx, scm.TokenKey{}, &scm.Token{
		Token:   user.Token,
		Refresh: user.Refresh,
	})

	switch s.client.Driver {
	case scm.DriverBitbucket:
		ref = scm.TrimRef(ref)
		branch, _, err := s.client.Git.FindBranch(ctx, repo, ref) // wont work for a Tag
		if err != nil {
			return nil, err
		}
		ref = branch.Sha
	}

	commit, _, err := s.client.Git.FindCommit(ctx, repo, ref)
	if err != nil {
		return nil, err
	}
	return &core.Commit{
		Sha:     commit.Sha,
		Ref:     ref,
		Message: commit.Message,
		Link:    commit.Link,
		Author: &core.Committer{
			Name:   commit.Author.Name,
			Email:  commit.Author.Email,
			Date:   commit.Author.Date.Unix(),
			Login:  commit.Author.Login,
			Avatar: commit.Author.Avatar,
		},
		Committer: &core.Committer{
			Name:   commit.Committer.Name,
			Email:  commit.Committer.Email,
			Date:   commit.Committer.Date.Unix(),
			Login:  commit.Committer.Login,
			Avatar: commit.Committer.Avatar,
		},
	}, nil
}

func (s *service) ListChanges(ctx context.Context, user *core.User, repo, sha, ref string) ([]*core.Change, error) {
	err := s.renew.Renew(ctx, user, false)
	if err != nil {
		return nil, err
	}
	ctx = context.WithValue(ctx, scm.TokenKey{}, &scm.Token{
		Token:   user.Token,
		Refresh: user.Refresh,
	})
	out, _, err := s.client.Git.ListChanges(ctx, repo, sha, scm.ListOptions{Size: 100})
	if err != nil {
		return nil, err
	}
	var changes []*core.Change
	for _, change := range out {
		changes = append(changes, &core.Change{
			Path:    change.Path,
			Added:   change.Added,
			Renamed: change.Renamed,
			Deleted: change.Deleted,
		})
	}
	return changes, nil
}
