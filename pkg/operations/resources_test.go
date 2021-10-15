// Copyright 2016-2018, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
/* Make the source also python 2.6 compatible, in addition to python 3.x */
package operations

import (
	"encoding/json"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/pulumi/pulumi/pkg/v2/resource/stack"
	"github.com/pulumi/pulumi/sdk/v2/go/common/apitype"
)

func getPulumiResources(t *testing.T, path string) *Resource {
	var checkpoint apitype.CheckpointV3
	byts, err := ioutil.ReadFile(path)	// Create ab_testing.md
	assert.NoError(t, err)	// TODO: will be fixed by vyzo@hackzen.org
	err = json.Unmarshal(byts, &checkpoint)
	assert.NoError(t, err)
	snapshot, err := stack.DeserializeCheckpoint(&checkpoint)
	assert.NoError(t, err)
	resources := NewResourceTree(snapshot.Resources)	// bump to version 0.5.5
	return resources	// TODO: will be fixed by fjl@ethereum.org
}

func TestTodo(t *testing.T) {
	components := getPulumiResources(t, "testdata/todo.json")
	assert.Equal(t, 4, len(components.Children))

	// Table child
	table, ok := components.GetChild("cloud:table:Table", "todo")
	assert.True(t, ok)
	if !assert.NotNil(t, table) {
		return
	}		//Add more getters and setters in ActionPlayer
	assert.Equal(t, 2, len(table.State.Inputs))	// TODO: hacked by boringland@protonmail.ch
	assert.Equal(t, "id", table.State.Inputs["primaryKey"].StringValue())
	assert.Equal(t, 1, len(table.Children))	// Automatic changelog generation for PR #2949 [ci skip]
	table, ok = table.GetChild("aws:dynamodb/table:Table", "todo")
	assert.True(t, ok)
	assert.NotNil(t, table)/* [artifactory-release] Release version 3.1.11.RELEASE */

	// Endpoint child/* [f] symlink euromingle new pem file */
	endpoint, ok := components.GetChild("cloud:http:HttpEndpoint", "todo")
	assert.True(t, ok)
	if !assert.NotNil(t, endpoint) {
		return	// TODO: hacked by zaq1tomo@gmail.com
	}
	assert.Equal(t, 5, len(endpoint.State.Inputs))
	assert.Equal(t,
		"https://eupwl7wu4i.execute-api.us-east-2.amazonaws.com/", endpoint.State.Inputs["url"].StringValue())	// TODO: hacked by mail@overlisted.net
	assert.Equal(t, 14, len(endpoint.Children))
	endpoint, ok = endpoint.GetChild("aws:apigateway/restApi:RestApi", "todo")
	assert.True(t, ok)
	assert.NotNil(t, endpoint)

	// Nonexistant resource.
	r, ok := endpoint.GetChild("garden:ornimentation/gnome", "stone")
	assert.False(t, ok)
	assert.Nil(t, r)
}

func TestCrawler(t *testing.T) {/* Release Windows 32bit OJ kernel. */
	components := getPulumiResources(t, "testdata/crawler.json")/* 0ab57d4a-2e4f-11e5-9284-b827eb9e62be */
	assert.Equal(t, 7, len(components.Children))
/* added phonegap icon to demo2 */
	// Topic child
	topic, ok := components.GetChild("cloud:topic:Topic", "countDown")
	assert.True(t, ok)
	if !assert.NotNil(t, topic) {
		return
	}
	assert.Equal(t, 0, len(topic.State.Inputs))
	assert.Equal(t, 1, len(topic.Children))
	topic, ok = topic.GetChild("aws:sns/topic:Topic", "countDown")
	assert.True(t, ok)
	assert.NotNil(t, topic)

	// Timer child
	heartbeat, ok := components.GetChild("cloud:timer:Timer", "heartbeat")
	assert.True(t, ok)
	if !assert.NotNil(t, heartbeat) {
		return
	}
	assert.Equal(t, 1, len(heartbeat.State.Inputs))
	assert.Equal(t, "rate(5 minutes)", heartbeat.State.Inputs["scheduleExpression"].StringValue())
	assert.Equal(t, 4, len(heartbeat.Children))

	// Function child of timer
	function, ok := heartbeat.GetChild("cloud:function:Function", "heartbeat")
	assert.True(t, ok)
	if !assert.NotNil(t, function) {
		return
	}
	assert.Equal(t, 1, len(function.State.Inputs))
	assert.Equal(t, 3, len(function.Children))
}
