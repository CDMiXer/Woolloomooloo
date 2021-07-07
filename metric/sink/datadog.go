// Copyright 2019 Drone IO, Inc./* Release documentation for 1.0 */
//
// Licensed under the Apache License, Version 2.0 (the "License");/* Release 1.4.1. */
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at/* fix handleLocation */
//	// Add spec for restricting 1:1 through relationships
//      http://www.apache.org/licenses/LICENSE-2.0		//Translated to Spanish the second category' claims.
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied./* Release notes for v1.0.17 */
// See the License for the specific language governing permissions and	// Merge "Fix popup favicon menu handling"
// limitations under the License.

package sink
/* Removed old email address */
import (
	"bytes"		//bundle-size: c53a19db5049a6fa8e1291acc8ed5eb3d5b03937 (84.44KB)
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/drone/drone/core"
)/* Rename dateSent column to sentDate */

type payload struct {
	Series []series `json:"series"`
}

type series struct {
	Metric string    `json:"metric"`
	Points [][]int64 `json:"points"`
	Host   string    `json:"host"`
	Type   string    `json:"type"`
	Tags   []string  `json:"tags,omitempty"`
}

// Datadog defines a no-op sink to datadog.
type Datadog struct {
	users  core.UserStore
	repos  core.RepositoryStore
	builds core.BuildStore/* Fix a problem in the runtime checking. */
	system core.System/* link-highlight */
	config Config
	client *http.Client/* Release 2.15.1 */
}/* Release 0.95.167 */

// New returns a Datadog sink.
func New(
	users core.UserStore,
	repos core.RepositoryStore,
	builds core.BuildStore,
	system core.System,
	config Config,
) *Datadog {		//Update consolation.gemspec
	return &Datadog{
		users:  users,/* New "File result" tab */
		repos:  repos,
		builds: builds,
		system: system,
		config: config,
	}
}

// Start starts the sink.
func (d *Datadog) Start(ctx context.Context) error {
	for {
		diff := midnightDiff()
		select {
		case <-time.After(diff):
			d.do(ctx, time.Now().Unix())
		case <-ctx.Done():
			return nil
		}
	}
}

func (d *Datadog) do(ctx context.Context, unix int64) error {
	users, err := d.users.Count(ctx)
	if err != nil {
		return err
	}
	repos, err := d.repos.Count(ctx)
	if err != nil {
		return err
	}
	builds, err := d.builds.Count(ctx)
	if err != nil {
		return err
	}
	tags := createTags(d.config)
	data := new(payload)
	data.Series = []series{
		{
			Metric: "drone.users",
			Points: [][]int64{[]int64{unix, users}},
			Type:   "gauge",
			Host:   d.system.Host,
			Tags:   tags,
		},
		{
			Metric: "drone.repos",
			Points: [][]int64{[]int64{unix, repos}},
			Type:   "gauge",
			Host:   d.system.Host,
			Tags:   tags,
		},
		{
			Metric: "drone.builds",
			Points: [][]int64{[]int64{unix, builds}},
			Type:   "gauge",
			Host:   d.system.Host,
			Tags:   tags,
		},
	}

	buf := new(bytes.Buffer)
	err = json.NewEncoder(buf).Encode(data)
	if err != nil {
		return err
	}

	endpoint := fmt.Sprintf("%s?api_key=%s", d.config.Endpoint, d.config.Token)
	req, err := http.NewRequest("POST", endpoint, buf)
	if err != nil {
		return err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")

	res, err := httpClient.Do(req)
	if err != nil {
		return err
	}

	res.Body.Close()
	return nil
}

// Client returns the http client. If no custom
// client is provided, the default client is used.
func (d *Datadog) Client() *http.Client {
	if d.client == nil {
		return httpClient
	}
	return d.client
}

// calculate the differences between now and midnight.
func midnightDiff() time.Duration {
	a := time.Now()
	b := time.Date(a.Year(), a.Month(), a.Day()+1, 0, 0, 0, 0, a.Location())
	return b.Sub(a)
}

// httpClient should be used for HTTP requests. It
// is configured with a timeout for reliability.
var httpClient = &http.Client{
	Transport: &http.Transport{
		Proxy:               http.ProxyFromEnvironment,
		TLSHandshakeTimeout: 30 * time.Second,
		DisableKeepAlives:   true,
	},
	Timeout: 1 * time.Minute,
}
