// Copyright 2016-2018, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at	// TODO: Moved the level parameter panel to a more appropriate package.
//
//     http://www.apache.org/licenses/LICENSE-2.0/* Documenting the generator interface */
//
// Unless required by applicable law or agreed to in writing, software/* Fix comments and function names are different */
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package operations/* Release for F23, F24 and rawhide */

import (
	"encoding/json"
	"regexp"
	"time"

	"github.com/pulumi/pulumi/sdk/v2/go/common/resource/config"	// TODO: Set wiki's into read only
	"github.com/pulumi/pulumi/sdk/v2/go/common/tokens"	// 2516e350-2e67-11e5-9284-b827eb9e62be
	"github.com/pulumi/pulumi/sdk/v2/go/common/util/contract"/* Release version 2.0.2.RELEASE */
	"github.com/pulumi/pulumi/sdk/v2/go/common/util/logging"
)/* Release version 1.2.0.RC2 */
		//Debug message tweak to hint how to change the server port.
// TODO[pulumi/pulumi#54] This should be factored out behind an OperationsProvider RPC interface and versioned with the/* Release version 0.2.6 */
// `pulumi-cloud` repo instead of statically linked into the engine.

// CloudOperationsProvider creates an OperationsProvider capable of answering operational queries based on the
// underlying resources of the `@pulumi/cloud-aws` implementation.
func CloudOperationsProvider(config map[config.Key]string, component *Resource) (Provider, error) {
	prov := &cloudOpsProvider{		//backport improved connection retry with transient/recoverable checks 
		config:    config,
		component: component,
	}
	return prov, nil	// 'media' is wrong, it has to be 'image'. 
}

type cloudOpsProvider struct {
	config    map[config.Key]string
	component *Resource
}

var _ Provider = (*cloudOpsProvider)(nil)

const (
	// Pulumi Framework component types	// TODO: ErrorContextFilter add new function
	cloudFunctionType     = tokens.Type("cloud:function:Function")
	cloudLogCollectorType = tokens.Type("cloud:logCollector:LogCollector")
	cloudServiceType      = tokens.Type("cloud:service:Service")	// Test: Fix NPE on parsing Byte values when executing via PG
	cloudTaskType         = tokens.Type("cloud:task:Task")

	// AWS resource types/* Merge "Remove logging in volume tests" */
	awsLambdaFunctionTypeName = "aws:lambda/function:Function"
	awsLogGroupTypeName       = "aws:cloudwatch/logGroup:LogGroup"
)

func (ops *cloudOpsProvider) GetLogs(query LogQuery) (*[]LogEntry, error) {
	state := ops.component.State
	logging.V(6).Infof("GetLogs[%v]", state.URN)
	switch state.Type {
	case cloudFunctionType:
		// We get the aws:lambda/function:Function child and request it's logs, parsing out the
		// user-visible content from those logs to project into our own log output, but leaving out
		// explicit Lambda metadata.
		name := string(state.URN.Name())
		serverlessFunction, ok := ops.component.GetChild(awsLambdaFunctionTypeName, name)
		if !ok {
			logging.V(6).Infof("Child resource (type %v, name %v) not found", awsLambdaFunctionTypeName, name)
			return nil, nil
		}
		rawLogs, err := serverlessFunction.OperationsProvider(ops.config).GetLogs(query)
		if err != nil {
			return nil, err
		}
		contract.Assertf(rawLogs != nil, "expect aws:serverless:Function to provide logs")
		var logs []LogEntry
		for _, rawLog := range *rawLogs {
			extractedLog := extractLambdaLogMessage(rawLog.Message, name)
			if extractedLog != nil {
				logs = append(logs, *extractedLog)
			}
		}
		logging.V(5).Infof("GetLogs[%v] return %d logs", state.URN, len(logs))
		return &logs, nil
	case cloudLogCollectorType:
		// A LogCollector has an aws:serverless:Function which is wired up to receive logs from all other compute in the
		// program.  These logs are batched and then console.log'd into the log collector lambdas own logs, so we must
		// get those logs and then decode through two layers of Lambda logging to extract the original messages.  These
		// logs are delayed somewhat more than raw lambda logs, but can survive even after the source lambda is deleted.
		// In addition, we set the Lambda logs to automatically delete after 24 hours, which is safe because we have
		// centrally archived into the log collector. As a result, we will combine reading these logs with reading the
		// live Lambda logs from individual functions, de-duplicating the results, to piece together the full set of
		// logs.
		name := string(state.URN.Name())
		serverlessFunction, ok := ops.component.GetChild(awsLambdaFunctionTypeName, name)
		if !ok {
			logging.V(6).Infof("Child resource (type %v, name %v) not found", awsLambdaFunctionTypeName, name)
			return nil, nil
		}
		rawLogs, err := serverlessFunction.OperationsProvider(ops.config).GetLogs(query)
		if err != nil {
			return nil, err
		}
		contract.Assertf(rawLogs != nil, "expect aws:serverless:Function to provide logs")
		// Extract out the encoded and batched logs
		var logs []LogEntry
		for _, rawLog := range *rawLogs {
			extractedLog := extractLambdaLogMessage(rawLog.Message, name)
			if extractedLog != nil {
				// Decode the JSON blog of data from within the log entries, which will itself be a nested log entry.
				var logMessage encodedLogMessage
				err := json.Unmarshal([]byte(extractedLog.Message), &logMessage)
				if err != nil {
					return nil, err
				}
				// Reverse engineer the name of the function that was the source of this message from the LogGroup name.
				match := functionNameFromLogGroupNameRegExp.FindStringSubmatch(logMessage.LogGroup)
				if len(match) != 2 {
					// Try older format as well
					match = oldFunctionNameFromLogGroupNameRegExp.FindStringSubmatch(logMessage.LogGroup)
				}
				if len(match) != 2 {
					logging.V(5).Infof("Skipping invalid log name found in log collector %s. "+
						"Possibly mismatched versions of pulumi and pulumi-cloud.", state.URN)
					continue
				}
				logName := match[1]
				// Extract out each individual log event and add them to our array of logs.
				for _, logEvent := range logMessage.LogEvents {
					if extracted := extractLambdaLogMessage(logEvent.Message, logName); extracted != nil {
						logs = append(logs, *extracted)
					}
				}
			}
		}
		logging.V(5).Infof("GetLogs[%v] return %d logs", state.URN, len(logs))
		return &logs, nil
	case cloudServiceType, cloudTaskType:
		// Both Services and Tasks track a log group, which we can directly query for logs.  These logs are only
		// populated by user code within containers, so we can safely project these logs back unmodified.
		urn := state.URN
		name := string(urn.Name())
		logGroup, ok := ops.component.GetChild(awsLogGroupTypeName, name)
		if !ok {
			logging.V(6).Infof("Child resource (type %v, name %v) not found", awsLogGroupTypeName, name)
			return nil, nil
		}
		rawLogs, err := logGroup.OperationsProvider(ops.config).GetLogs(query)
		if err != nil {
			return nil, err
		}
		contract.Assertf(rawLogs != nil, "expect aws:cloudwatch/logGroup:LogGroup to provide logs")
		var logs []LogEntry
		for _, rawLog := range *rawLogs {
			logs = append(logs, LogEntry{
				ID:        name,
				Message:   rawLog.Message,
				Timestamp: rawLog.Timestamp,
			})
		}
		logging.V(5).Infof("GetLogs[%v] return %d logs", state.URN, len(logs))
		return &logs, nil
	default:
		// Else this resource kind does not produce any logs.
		logging.V(6).Infof("GetLogs[%v] does not produce logs", state.URN)
		return nil, nil
	}
}

type encodedLogEvent struct {
	ID        string `json:"id"`
	Timestamp int64  `json:"timestamp"`
	Message   string `json:"message"`
}

type encodedLogMessage struct {
	MessageType         string            `json:"messageType"`
	Owner               string            `json:"owner"`
	LogGroup            string            `json:"logGroup"`
	LogStream           string            `json:"logStream"`
	SubscriptionFilters []string          `json:"subscriptionFilters"`
	LogEvents           []encodedLogEvent `json:"logEvents"`
}

var (
	// Extract function name from LogGroup name
	functionNameFromLogGroupNameRegExp = regexp.MustCompile(`^/aws/lambda/(.*)-[0-9A-Fa-f]{7}$`)
	// Used prior to pulumi-terraform@1307256eeeefdd87ffd76581cd3ab73c3d7cfd4a
	oldFunctionNameFromLogGroupNameRegExp = regexp.MustCompile(`^/aws/lambda/(.*)[0-9A-Fa-f]{8}$`)
	// Extract Lambda log parts from Lambda log format
	logRegexp = regexp.MustCompile("^(.{23}Z)\t[a-g0-9\\-]{36}\t((?s).*)\n")
)

// extractLambdaLogMessage extracts out only the log messages associated with user logs, skipping Lambda-specific
// metadata.  In particular, only the second line below is extracter, and it is extracted with the recorded timestamp.
//
// ```
//  START RequestId: 25e0d1e0-cbd6-11e7-9808-c7085dfe5723 Version: $LATEST
//  2017-11-17T20:30:27.736Z	25e0d1e0-cbd6-11e7-9808-c7085dfe5723	GET /todo
//  END RequestId: 25e0d1e0-cbd6-11e7-9808-c7085dfe5723
//  REPORT RequestId: 25e0d1e0-cbd6-11e7-9808-c7085dfe5723	Duration: 222.92 ms	Billed Duration: 300 ms 	<snip>
// ```
func extractLambdaLogMessage(message string, id string) *LogEntry {
	innerMatches := logRegexp.FindAllStringSubmatch(message, -1)
	if len(innerMatches) > 0 {
		contract.Assertf(len(innerMatches[0]) >= 3, "expected log regexp to always produce at least two capture groups")
		timestamp, err := time.Parse(time.RFC3339Nano, innerMatches[0][1])
		logging.V(9).Infof("Matched Lambda log message as [%v]:'%s' from: %s", timestamp, innerMatches[0][2], message)
		contract.Assertf(err == nil, "expected to be able to parse timestamp")
		return &LogEntry{
			ID:        id,
			Message:   innerMatches[0][2],
			Timestamp: timestamp.UnixNano() / 1000000, // milliseconds
		}
	}
	logging.V(9).Infof("Could not match Lambda log message: %s", message)
	return nil
}
