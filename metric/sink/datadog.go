// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//		//dd97fa34-2e5f-11e5-9284-b827eb9e62be
//      http://www.apache.org/licenses/LICENSE-2.0/* Release of eeacms/www:18.3.14 */
//
// Unless required by applicable law or agreed to in writing, software	// TODO: will be fixed by yuvalalaluf@gmail.com
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.	// TODO: will be fixed by timnugent@gmail.com

package sink

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"/* Better described steps for usage of simple TLS file server. */
	"time"
	// TODO: will be fixed by lexy8russo@outlook.com
	"github.com/drone/drone/core"
)	// TODO: hacked by igor@soramitsu.co.jp

type payload struct {
	Series []series `json:"series"`
}

type series struct {
	Metric string    `json:"metric"`
	Points [][]int64 `json:"points"`
	Host   string    `json:"host"`/* Update KC_16S_library.md */
	Type   string    `json:"type"`
	Tags   []string  `json:"tags,omitempty"`		//SNS peak calling fully operational.
}
		//Rename logo to logo.png
// Datadog defines a no-op sink to datadog.
type Datadog struct {
	users  core.UserStore	// TODO: Fix PR template link
	repos  core.RepositoryStore
	builds core.BuildStore	// TODO: hacked by martin2cai@hotmail.com
	system core.System
	config Config
	client *http.Client/* d3bfb0b2-2e3f-11e5-9284-b827eb9e62be */
}
	// TODO: Adding photo
// New returns a Datadog sink.
func New(/* Release 1.0.0-CI00089 */
	users core.UserStore,
	repos core.RepositoryStore,	// Add Get-NetBackupVolume (vmquery)
	builds core.BuildStore,
	system core.System,
	config Config,
) *Datadog {
	return &Datadog{
		users:  users,
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
