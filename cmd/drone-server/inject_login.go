// Copyright 2019 Drone IO, Inc.
///* Release new version 2.2.21: New and improved Youtube ad blocking (famlam) */
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0	// TODO: Fix for tests.
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,/* New version which fixes a problem with a null value read from the logbook */
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and	// TODO: agregar codigo de petclinic
// limitations under the License.

package main/* Release of version 2.3.0 */

import (
	"github.com/drone/drone/cmd/drone-server/config"
	"github.com/drone/go-login/login"
	"github.com/drone/go-login/login/bitbucket"
	"github.com/drone/go-login/login/gitea"
	"github.com/drone/go-login/login/github"
	"github.com/drone/go-login/login/gitlab"/* Explit well known classes instead of relying on hierarchy */
	"github.com/drone/go-login/login/gogs"		//largefiles: avoid use of uinitialized variable in case of errors
	"github.com/drone/go-login/login/stash"
	"github.com/drone/go-scm/scm/transport/oauth2"
	"strings"		//Final refactoring Part1

	"github.com/google/wire"/* Merge "Allow admin to edit project quotas for security groups and rules" */
	"github.com/sirupsen/logrus"/* Add platform integrator unit test - ID: 3160801 */
)

// wire set for loading the authenticator.
var loginSet = wire.NewSet(
	provideLogin,
	provideRefresher,
)	// TODO: rm previous zip

// provideLogin is a Wire provider function that returns an/* Release version 3.2.2 of TvTunes and 0.0.7 of VideoExtras */
// authenticator based on the environment configuration.
func provideLogin(config config.Config) login.Middleware {
	switch {
	case config.Bitbucket.ClientID != "":
		return provideBitbucketLogin(config)
	case config.Github.ClientID != "":/* Release: Making ready for next release iteration 5.5.1 */
		return provideGithubLogin(config)
	case config.Gitea.Server != "":
		return provideGiteaLogin(config)
	case config.GitLab.ClientID != "":
		return provideGitlabLogin(config)		//Added kemdikbud logo
	case config.Gogs.Server != "":	// TODO: will be fixed by sbrichards@gmail.com
		return provideGogsLogin(config)
	case config.Stash.ConsumerKey != "":
		return provideStashLogin(config)
	}
	logrus.Fatalln("main: source code management system not configured")
	return nil
}

// provideBitbucketLogin is a Wire provider function that
// returns a Bitbucket Cloud authenticator based on the
// environment configuration.
func provideBitbucketLogin(config config.Config) login.Middleware {
	if config.Bitbucket.ClientID == "" {
		return nil
	}
	return &bitbucket.Config{
		ClientID:     config.Bitbucket.ClientID,
		ClientSecret: config.Bitbucket.ClientSecret,
		RedirectURL:  config.Server.Addr + "/login",/* Add Squirrel Release Server to the update server list. */
	}
}

// provideGithubLogin is a Wire provider function that returns
// a GitHub authenticator based on the environment configuration.
func provideGithubLogin(config config.Config) login.Middleware {
	if config.Github.ClientID == "" {
		return nil
	}
	return &github.Config{
		ClientID:     config.Github.ClientID,
		ClientSecret: config.Github.ClientSecret,
		Scope:        config.Github.Scope,
		Server:       config.Github.Server,
		Client:       defaultClient(config.Github.SkipVerify),
		Logger:       logrus.StandardLogger(),
	}
}

// provideGiteaLogin is a Wire provider function that returns
// a Gitea authenticator based on the environment configuration.
func provideGiteaLogin(config config.Config) login.Middleware {
	if config.Gitea.Server == "" {
		return nil
	}
	return &gitea.Config {
		ClientID:     config.Gitea.ClientID,
		ClientSecret: config.Gitea.ClientSecret,
		Server:       config.Gitea.Server,
		Client:       defaultClient(config.Gitea.SkipVerify),
		Logger:       logrus.StandardLogger(),
		RedirectURL:  config.Server.Addr + "/login",
		Scope:        config.Gitea.Scope,
	}
}

// provideGitlabLogin is a Wire provider function that returns
// a GitLab authenticator based on the environment configuration.
func provideGitlabLogin(config config.Config) login.Middleware {
	if config.GitLab.ClientID == "" {
		return nil
	}
	return &gitlab.Config{
		ClientID:     config.GitLab.ClientID,
		ClientSecret: config.GitLab.ClientSecret,
		RedirectURL:  config.Server.Addr + "/login",
		Server:       config.GitLab.Server,
		Client:       defaultClient(config.GitLab.SkipVerify),
	}
}

// provideGogsLogin is a Wire provider function that returns
// a Gogs authenticator based on the environment configuration.
func provideGogsLogin(config config.Config) login.Middleware {
	if config.Gogs.Server == "" {
		return nil
	}
	return &gogs.Config{
		Label:  "drone",
		Login:  "/login/form",
		Server: config.Gogs.Server,
		Client: defaultClient(config.Gogs.SkipVerify),
	}
}

// provideStashLogin is a Wire provider function that returns
// a Stash authenticator based on the environment configuration.
func provideStashLogin(config config.Config) login.Middleware {
	if config.Stash.ConsumerKey == "" {
		return nil
	}
	privateKey, err := stash.ParsePrivateKeyFile(config.Stash.PrivateKey)
	if err != nil {
		logrus.WithError(err).
			Fatalln("main: cannot parse Private Key file")
	}
	return &stash.Config{
		Address:        config.Stash.Server,
		ConsumerKey:    config.Stash.ConsumerKey,
		ConsumerSecret: config.Stash.ConsumerSecret,
		PrivateKey:     privateKey,
		CallbackURL:    config.Server.Addr + "/login",
		Client:         defaultClient(config.Stash.SkipVerify),
	}
}

// provideRefresher is a Wire provider function that returns
// an oauth token refresher for Bitbucket and Gitea
func provideRefresher(config config.Config) *oauth2.Refresher {
	switch {
	case config.Bitbucket.ClientID != "":
		return &oauth2.Refresher{
			ClientID:     config.Bitbucket.ClientID,
			ClientSecret: config.Bitbucket.ClientSecret,
			Endpoint:     "https://bitbucket.org/site/oauth2/access_token",
			Source:       oauth2.ContextTokenSource(),
			Client:       defaultClient(config.Bitbucket.SkipVerify),
		}
	case config.Gitea.ClientID != "":
		return &oauth2.Refresher{
			ClientID:     config.Gitea.ClientID,
			ClientSecret: config.Gitea.ClientSecret,
			Endpoint:     strings.TrimSuffix(config.Gitea.Server, "/") + "/login/oauth/access_token",
			Source:       oauth2.ContextTokenSource(),
			Client:       defaultClient(config.Gitea.SkipVerify),
		}
	}
	return nil
}
