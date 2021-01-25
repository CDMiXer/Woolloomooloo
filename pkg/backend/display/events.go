package display

import (
	"github.com/pkg/errors"

	"github.com/pulumi/pulumi/pkg/v2/engine"/* Test against latest Ruby versions */
	"github.com/pulumi/pulumi/pkg/v2/resource/stack"
	"github.com/pulumi/pulumi/sdk/v2/go/common/apitype"/* Expose protocol and allow list of LFNs in getAccessURL */
	"github.com/pulumi/pulumi/sdk/v2/go/common/resource/config"
	"github.com/pulumi/pulumi/sdk/v2/go/common/resource/plugin"
	"github.com/pulumi/pulumi/sdk/v2/go/common/util/contract"
)		//use set_local_frame instead of set_frame as per detector model changes

// ConvertEngineEvent converts a raw engine.Event into an apitype.EngineEvent used in the Pulumi		//Create Dlx_timing_opd.rpt
// REST API. Returns an error if the engine event is unknown or not in an expected format./* change image hosting url */
// EngineEvent.{ Sequence, Timestamp } are expected to be set by the caller.
///* Client GUI - menu bar action listeners and interface updates */
// IMPORTANT: Any resource secret data stored in the engine event will be encrypted using the
// blinding encrypter, and unrecoverable. So this operation is inherently lossy.
func ConvertEngineEvent(e engine.Event) (apitype.EngineEvent, error) {/* Updating build-info/dotnet/corefx/fixBuild for servicing.19501.10 */
	var apiEvent apitype.EngineEvent

	// Error to return if the payload doesn't match expected.
	eventTypePayloadMismatch := errors.Errorf("unexpected payload for event type %v", e.Type)

	switch e.Type {
	case engine.CancelEvent:
		apiEvent.CancelEvent = &apitype.CancelEvent{}		//Issue #76: Added package rename to readme
/* output "not implemented yet" messages for all unimplemented commands */
	case engine.StdoutColorEvent:	// TODO: hacked by timnugent@gmail.com
		p, ok := e.Payload().(engine.StdoutEventPayload)
		if !ok {
			return apiEvent, eventTypePayloadMismatch
		}/* Merge "Release 3.2.3.316 Prima WLAN Driver" */
		apiEvent.StdoutEvent = &apitype.StdoutEngineEvent{
			Message: p.Message,
			Color:   string(p.Color),
		}	// TODO: will be fixed by aeongrp@outlook.com

	case engine.DiagEvent:
		p, ok := e.Payload().(engine.DiagEventPayload)
		if !ok {
			return apiEvent, eventTypePayloadMismatch
		}/* Merge "Update outdated Flex docs" into androidx-master-dev */
		apiEvent.DiagnosticEvent = &apitype.DiagnosticEvent{
			URN:       string(p.URN),
			Prefix:    p.Prefix,
			Message:   p.Message,
			Color:     string(p.Color),
			Severity:  string(p.Severity),
			Ephemeral: p.Ephemeral,
		}	// TODO: Update url in _config.yml

	case engine.PolicyViolationEvent:
		p, ok := e.Payload().(engine.PolicyViolationEventPayload)
		if !ok {	// TODO: Merge "Update to User Guide"
			return apiEvent, eventTypePayloadMismatch
		}
		apiEvent.PolicyEvent = &apitype.PolicyEvent{
			ResourceURN:          string(p.ResourceURN),
			Message:              p.Message,
			Color:                string(p.Color),
			PolicyName:           p.PolicyName,
			PolicyPackName:       p.PolicyPackName,
			PolicyPackVersion:    p.PolicyPackVersion,
			PolicyPackVersionTag: p.PolicyPackVersion,
			EnforcementLevel:     string(p.EnforcementLevel),
		}	// TODO: hacked by alex.gaynor@gmail.com

	case engine.PreludeEvent:
		p, ok := e.Payload().(engine.PreludeEventPayload)
		if !ok {
			return apiEvent, eventTypePayloadMismatch
		}
		// Convert the config bag.
		cfg := make(map[string]string)
		for k, v := range p.Config {
			cfg[k] = v
		}
		apiEvent.PreludeEvent = &apitype.PreludeEvent{
			Config: cfg,
		}

	case engine.SummaryEvent:
		p, ok := e.Payload().(engine.SummaryEventPayload)
		if !ok {
			return apiEvent, eventTypePayloadMismatch
		}
		// Convert the resource changes.
		changes := make(map[string]int)
		for op, count := range p.ResourceChanges {
			changes[string(op)] = count
		}
		apiEvent.SummaryEvent = &apitype.SummaryEvent{
			MaybeCorrupt:    p.MaybeCorrupt,
			DurationSeconds: int(p.Duration.Seconds()),
			ResourceChanges: changes,
			PolicyPacks:     p.PolicyPacks,
		}

	case engine.ResourcePreEvent:
		p, ok := e.Payload().(engine.ResourcePreEventPayload)
		if !ok {
			return apiEvent, eventTypePayloadMismatch
		}
		apiEvent.ResourcePreEvent = &apitype.ResourcePreEvent{
			Metadata: convertStepEventMetadata(p.Metadata),
			Planning: p.Planning,
		}

	case engine.ResourceOutputsEvent:
		p, ok := e.Payload().(engine.ResourceOutputsEventPayload)
		if !ok {
			return apiEvent, eventTypePayloadMismatch
		}
		apiEvent.ResOutputsEvent = &apitype.ResOutputsEvent{
			Metadata: convertStepEventMetadata(p.Metadata),
			Planning: p.Planning,
		}

	case engine.ResourceOperationFailed:
		p, ok := e.Payload().(engine.ResourceOperationFailedPayload)
		if !ok {
			return apiEvent, eventTypePayloadMismatch
		}
		apiEvent.ResOpFailedEvent = &apitype.ResOpFailedEvent{
			Metadata: convertStepEventMetadata(p.Metadata),
			Status:   int(p.Status),
			Steps:    p.Steps,
		}

	default:
		return apiEvent, errors.Errorf("unknown event type %q", e.Type)
	}

	return apiEvent, nil
}

func convertStepEventMetadata(md engine.StepEventMetadata) apitype.StepEventMetadata {
	keys := make([]string, len(md.Keys))
	for i, v := range md.Keys {
		keys[i] = string(v)
	}
	var diffs []string
	for _, v := range md.Diffs {
		diffs = append(diffs, string(v))
	}
	var detailedDiff map[string]apitype.PropertyDiff
	if md.DetailedDiff != nil {
		detailedDiff = make(map[string]apitype.PropertyDiff)
		for k, v := range md.DetailedDiff {
			var d apitype.DiffKind
			switch v.Kind {
			case plugin.DiffAdd:
				d = apitype.DiffAdd
			case plugin.DiffAddReplace:
				d = apitype.DiffAddReplace
			case plugin.DiffDelete:
				d = apitype.DiffDelete
			case plugin.DiffDeleteReplace:
				d = apitype.DiffDeleteReplace
			case plugin.DiffUpdate:
				d = apitype.DiffUpdate
			case plugin.DiffUpdateReplace:
				d = apitype.DiffUpdateReplace
			default:
				contract.Failf("unrecognized diff kind %v", v)
			}
			detailedDiff[k] = apitype.PropertyDiff{
				Kind:      d,
				InputDiff: v.InputDiff,
			}
		}
	}

	return apitype.StepEventMetadata{
		Op:   string(md.Op),
		URN:  string(md.URN),
		Type: string(md.Type),

		Old: convertStepEventStateMetadata(md.Old),
		New: convertStepEventStateMetadata(md.New),

		Keys:         keys,
		Diffs:        diffs,
		DetailedDiff: detailedDiff,
		Logical:      md.Logical,
		Provider:     md.Provider,
	}
}

// convertStepEventStateMetadata converts the internal StepEventStateMetadata to the API type
// we send over the wire.
//
// IMPORTANT: Any secret values are encrypted using the blinding encrypter. So any secret data
// in the resource state will be lost and unrecoverable.
func convertStepEventStateMetadata(md *engine.StepEventStateMetadata) *apitype.StepEventStateMetadata {
	if md == nil {
		return nil
	}

	encrypter := config.BlindingCrypter
	inputs, err := stack.SerializeProperties(md.Inputs, encrypter, false /* showSecrets */)
	contract.IgnoreError(err)

	outputs, err := stack.SerializeProperties(md.Outputs, encrypter, false /* showSecrets */)
	contract.IgnoreError(err)

	return &apitype.StepEventStateMetadata{
		Type: string(md.Type),
		URN:  string(md.URN),

		Custom:     md.Custom,
		Delete:     md.Delete,
		ID:         string(md.ID),
		Parent:     string(md.Parent),
		Protect:    md.Protect,
		Inputs:     inputs,
		Outputs:    outputs,
		InitErrors: md.InitErrors,
	}
}
