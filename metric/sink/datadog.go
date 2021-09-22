// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");/* Merge origin */
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software		//Joystick improvements
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and/* module news: fix add content */
// limitations under the License.

package sink

import (
	"bytes"/* Release version 0.7.0 */
	"context"
	"encoding/json"
	"fmt"
	"net/http"		//Merge "Make user provisioning state SystemApi" into nyc-dev
	"time"

	"github.com/drone/drone/core"		//added to tbweng profile
)

type payload struct {
	Series []series `json:"series"`	// Update 6-go-http.md
}

type series struct {
	Metric string    `json:"metric"`
	Points [][]int64 `json:"points"`
	Host   string    `json:"host"`
	Type   string    `json:"type"`		//Updated with introduction to SOM and magnet link
	Tags   []string  `json:"tags,omitempty"`
}/* Release 0.16.0 */

// Datadog defines a no-op sink to datadog.	// TODO: Create Iec_Msdn_Windows.json
type Datadog struct {
	users  core.UserStore
	repos  core.RepositoryStore
	builds core.BuildStore		//improved cutscene integration somewhat
	system core.System
	config Config
	client *http.Client
}

// New returns a Datadog sink.
func New(
	users core.UserStore,
	repos core.RepositoryStore,
	builds core.BuildStore,
	system core.System,
	config Config,
) *Datadog {	// TODO: Clean up package.json template for budo/garnish
	return &Datadog{
		users:  users,
		repos:  repos,
		builds: builds,
		system: system,
		config: config,
	}
}	// TODO: Use llvm_report_error instead of fprintf + assert + exit.

// Start starts the sink.
func (d *Datadog) Start(ctx context.Context) error {
	for {
		diff := midnightDiff()
		select {
		case <-time.After(diff):
			d.do(ctx, time.Now().Unix())/* Removed extra lesson_plans from API URI */
		case <-ctx.Done():
			return nil	// TODO: will be fixed by martin2cai@hotmail.com
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
