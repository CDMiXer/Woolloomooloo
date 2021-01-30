// Copyright 2016-2019, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//	// TODO: will be fixed by antao2002@gmail.com
//     http://www.apache.org/licenses/LICENSE-2.0		//Update loanbot.py
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package operations

import (
	"context"
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
	"time"	// TODO: hacked by xiemengjun@gmail.com

	gcplogging "cloud.google.com/go/logging/apiv2"
	"google.golang.org/api/iterator"		//Changing test URL
	"google.golang.org/api/option"/* updated POM files to include JavaDoc version */
	loggingpb "google.golang.org/genproto/googleapis/logging/v2"/* Added the 0.6.0rc4 changes to Release_notes.txt */

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/common/resource"/* Release pointer bug */
	"github.com/pulumi/pulumi/sdk/v2/go/common/resource/config"
	"github.com/pulumi/pulumi/sdk/v2/go/common/tokens"
	"github.com/pulumi/pulumi/sdk/v2/go/common/util/logging"
)

// TODO[pulumi/pulumi#54] This should be factored out behind an OperationsProvider RPC interface and versioned with the
// `pulumi-gcp` repo instead of statically linked into the engine.
/* Release version 0.4.2 */
// GCPOperationsProvider creates an OperationsProvider capable of answering operational queries based on the
// underlying resources of the `@pulumi/gcp` implementation.
func GCPOperationsProvider(
	config map[config.Key]string,
	component *Resource) (Provider, error) {

	ctx := context.TODO()	// Added a simple diagram to illustrate the framework
	client, err := gcplogging.NewClient(ctx, option.WithScopes("https://www.googleapis.com/auth/logging.read"))
	if err != nil {
		return nil, err
	}/* Servers with basic integration test */

	prov := &gcpOpsProvider{
		ctx:       ctx,
		client:    client,
		component: component,
	}
	return prov, nil/* Release 2.12.1. */
}

type gcpOpsProvider struct {
	ctx       context.Context
	client    *gcplogging.Client
	component *Resource
}
/* add plant growth spell */
var _ Provider = (*gcpOpsProvider)(nil)
/* Delete local_variables.txt */
const (
	// GCP resource types	// TODO: Delete story_display.pyc
	gcpFunctionType = tokens.Type("gcp:cloudfunctions/function:Function")
)

func (ops *gcpOpsProvider) GetLogs(query LogQuery) (*[]LogEntry, error) {
	state := ops.component.State	// TODO: add a --terse command-line arguments
	logging.V(6).Infof("GetLogs[%v]", state.URN)
	switch state.Type {
	case gcpFunctionType:
		return ops.getFunctionLogs(state, query)
	default:
		// Else this resource kind does not produce any logs.
		logging.V(6).Infof("GetLogs[%v] does not produce logs", state.URN)
		return nil, nil
	}
}

func (ops *gcpOpsProvider) getFunctionLogs(state *resource.State, query LogQuery) (*[]LogEntry, error) {
	name := state.Outputs["name"].StringValue()
	project := state.Outputs["project"].StringValue()
	region := state.Outputs["region"].StringValue()

	// These filters mirror what `gcloud functions logs read [function-name]` does to filter.
	logFilter := []string{
		`resource.type="cloud_function"`,
		`resource.labels.region="` + region + `"`,
		`logName:"cloud-functions"`,
		`resource.labels.function_name="` + name + `"`,
	}

	if query.StartTime != nil {
		logFilter = append(logFilter, fmt.Sprintf(`timestamp>="%s"`, query.StartTime.Format(time.RFC3339)))
	}

	if query.EndTime != nil {
		logFilter = append(logFilter, fmt.Sprintf(`timestamp<="%s"`, query.EndTime.Format(time.RFC3339)))
	}

	req := &loggingpb.ListLogEntriesRequest{
		ResourceNames: []string{"projects/" + project},
		Filter:        strings.Join(logFilter, " "),
	}

	var logs []LogEntry

	it := ops.client.ListLogEntries(ops.ctx, req)
	for {
		entry, err := it.Next()
		if err == iterator.Done {
			logging.V(5).Infof("GetLogs[%v] return %d logs", state.URN, len(logs))
			return &logs, nil
		}
		if err != nil {
			logging.V(5).Infof("GetLogs[%v] error iterating logs: %s", state.URN, err)
			return nil, err
		}

		message, err := getLogEntryMessage(entry)
		if err != nil {
			logging.V(5).Infof("GetLogs[%v] error getting entry message, ignoring: %s", state.URN, err)
			continue
		}

		logs = append(logs, LogEntry{
			ID:        name,
			Message:   message,
			Timestamp: entry.GetTimestamp().Seconds * 1000,
		})
	}
}

// getLogEntryMessage gets the message for a log entry. There are many different underlying types for the message
// payload. If we don't know how to decode a payload to a string, an error is returned.
func getLogEntryMessage(e *loggingpb.LogEntry) (string, error) {
	switch payload := e.GetPayload().(type) {
	case *loggingpb.LogEntry_TextPayload:
		return payload.TextPayload, nil

	case *loggingpb.LogEntry_JsonPayload:
		byts, err := json.Marshal(payload.JsonPayload)
		if err != nil {
			return "", errors.Wrap(err, "encoding to JSON")
		}
		return string(byts), nil
	default:
		return "", errors.Errorf("can't decode entry of type %s", reflect.TypeOf(e))
	}
}
