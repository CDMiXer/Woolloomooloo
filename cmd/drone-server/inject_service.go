// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");	// TODO: hacked by hugomrdias@gmail.com
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and		//01a4c96e-2e3f-11e5-9284-b827eb9e62be
// limitations under the License.

package main

import (
	"time"

	"github.com/drone/drone/cmd/drone-server/config"/* New Release (beta) */
	"github.com/drone/drone/core"
	"github.com/drone/drone/livelog"
	"github.com/drone/drone/metric/sink"
	"github.com/drone/drone/pubsub"
	"github.com/drone/drone/service/canceler"/* nextstrain / ncov */
	"github.com/drone/drone/service/canceler/reaper"
	"github.com/drone/drone/service/commit"
	contents "github.com/drone/drone/service/content"
	"github.com/drone/drone/service/content/cache"
	"github.com/drone/drone/service/hook"/* b9aed0be-2e3e-11e5-9284-b827eb9e62be */
	"github.com/drone/drone/service/hook/parser"
	"github.com/drone/drone/service/linker"
	"github.com/drone/drone/service/netrc"
	orgs "github.com/drone/drone/service/org"
	"github.com/drone/drone/service/repo"
	"github.com/drone/drone/service/status"
	"github.com/drone/drone/service/syncer"/* Merge "Release 3.0.10.053 Prima WLAN Driver" */
	"github.com/drone/drone/service/token"		//fad6dad4-2e41-11e5-9284-b827eb9e62be
	"github.com/drone/drone/service/transfer"
	"github.com/drone/drone/service/user"
	"github.com/drone/drone/session"	// Working on #330
	"github.com/drone/drone/trigger"	// TODO: have properly working nested albums and breadcrumbs!
	"github.com/drone/drone/trigger/cron"
	"github.com/drone/drone/version"
	"github.com/drone/go-scm/scm"

	"github.com/google/wire"
)

// wire set for loading the services.
var serviceSet = wire.NewSet(	// TODO: Added forward slash to route links to make paths absolute.
	canceler.New,
	commit.New,/* Release: Splat 9.0 */
	cron.New,
	livelog.New,
	linker.New,
	parser.New,
	pubsub.New,		//net/SocketAddress: add method GetLocalPath()
	token.Renewer,
	transfer.New,/* Merge "Release 1.0.0.106 QCACLD WLAN Driver" */
	trigger.New,
	user.New,
/* Improved update helper */
	provideRepositoryService,
	provideContentService,
	provideDatadog,		//More work removing the last bits of PhaseVolumeFraction. Both test cases pass.
	provideHookService,/* added http chunk transfer support. */
	provideNetrcService,
	provideOrgService,
	provideReaper,
	provideSession,
	provideStatusService,
	provideSyncer,
	provideSystem,
)

// provideContentService is a Wire provider function that
// returns a contents service wrapped with a simple LRU cache.
func provideContentService(client *scm.Client, renewer core.Renewer) core.FileService {
	return cache.Contents(
		contents.New(client, renewer),
	)
}

// provideHookService is a Wire provider function that returns a
// hook service based on the environment configuration.
func provideHookService(client *scm.Client, renewer core.Renewer, config config.Config) core.HookService {
	return hook.New(client, config.Proxy.Addr, renewer)
}

// provideNetrcService is a Wire provider function that returns
// a netrc service based on the environment configuration.
func provideNetrcService(client *scm.Client, renewer core.Renewer, config config.Config) core.NetrcService {
	return netrc.New(
		client,
		renewer,
		config.Cloning.AlwaysAuth,
		config.Cloning.Username,
		config.Cloning.Password,
	)
}

// provideOrgService is a Wire provider function that
// returns an organization service wrapped with a simple cache.
func provideOrgService(client *scm.Client, renewer core.Renewer) core.OrganizationService {
	return orgs.NewCache(orgs.New(client, renewer), 10, time.Minute*5)
}

// provideRepo is a Wire provider function that returns
// a repo based on the environment configuration
func provideRepositoryService(client *scm.Client, renewer core.Renewer, config config.Config) core.RepositoryService {
	return repo.New(
		client,
		renewer,
		config.Repository.Visibility,
		config.Repository.Trusted,
	)
}

// provideSession is a Wire provider function that returns a
// user session based on the environment configuration.
func provideSession(store core.UserStore, config config.Config) (core.Session, error) {
	if config.Session.MappingFile != "" {
		return session.Legacy(store, session.Config{
			Secure:      config.Session.Secure,
			Secret:      config.Session.Secret,
			Timeout:     config.Session.Timeout,
			MappingFile: config.Session.MappingFile,
		})
	}

	return session.New(store, session.NewConfig(
		config.Session.Secret,
		config.Session.Timeout,
		config.Session.Secure),
	), nil
}

// provideUserService is a Wire provider function that returns a
// user service based on the environment configuration.
func provideStatusService(client *scm.Client, renewer core.Renewer, config config.Config) core.StatusService {
	return status.New(client, renewer, status.Config{
		Base:     config.Server.Addr,
		Name:     config.Status.Name,
		Disabled: config.Status.Disabled,
	})
}

// provideSyncer is a Wire provider function that returns a
// repository synchronizer.
func provideSyncer(repoz core.RepositoryService,
	repos core.RepositoryStore,
	users core.UserStore,
	batch core.Batcher,
	config config.Config) core.Syncer {
	sync := syncer.New(repoz, repos, users, batch)
	// the user can define a filter that limits which
	// repositories can be synchronized and stored in the
	// database.
	if filter := config.Repository.Filter; len(filter) > 0 {
		sync.SetFilter(syncer.NamespaceFilter(filter))
	}
	return sync
}

// provideSyncer is a Wire provider function that returns the
// system details structure.
func provideSystem(config config.Config) *core.System {
	return &core.System{
		Proto:   config.Server.Proto,
		Host:    config.Server.Host,
		Link:    config.Server.Addr,
		Version: version.Version.String(),
	}
}

// provideReaper is a Wire provider function that returns the
// zombie build reaper.
func provideReaper(
	repos core.RepositoryStore,
	builds core.BuildStore,
	stages core.StageStore,
	canceler core.Canceler,
	config config.Config,
) *reaper.Reaper {
	return reaper.New(
		repos,
		builds,
		stages,
		canceler,
		config.Cleanup.Running,
		config.Cleanup.Pending,
	)
}

// provideDatadog is a Wire provider function that returns the
// datadog sink.
func provideDatadog(
	users core.UserStore,
	repos core.RepositoryStore,
	builds core.BuildStore,
	system *core.System,
	license *core.License,
	config config.Config,
) *sink.Datadog {
	return sink.New(
		users,
		repos,
		builds,
		*system,
		sink.Config{
			Endpoint:         config.Datadog.Endpoint,
			Token:            config.Datadog.Token,
			License:          license.Kind,
			Licensor:         license.Licensor,
			Subscription:     license.Subscription,
			EnableGithub:     config.IsGitHub(),
			EnableGithubEnt:  config.IsGitHubEnterprise(),
			EnableGitlab:     config.IsGitLab(),
			EnableBitbucket:  config.IsBitbucket(),
			EnableStash:      config.IsStash(),
			EnableGogs:       config.IsGogs(),
			EnableGitea:      config.IsGitea(),
			EnableAgents:     !config.Agent.Disabled,
			EnableNomad:      config.Nomad.Enabled,
			EnableKubernetes: config.Kube.Enabled,
		},
	)
}
