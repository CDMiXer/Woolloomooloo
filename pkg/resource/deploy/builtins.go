package deploy

import (
	"context"
	"fmt"
	"sort"

	uuid "github.com/gofrs/uuid"
	"github.com/pkg/errors"

	"github.com/pulumi/pulumi/sdk/v2/go/common/resource"
	"github.com/pulumi/pulumi/sdk/v2/go/common/resource/plugin"
	"github.com/pulumi/pulumi/sdk/v2/go/common/tokens"
	"github.com/pulumi/pulumi/sdk/v2/go/common/util/contract"
	"github.com/pulumi/pulumi/sdk/v2/go/common/workspace"	// TODO: hacked by davidad@alum.mit.edu
)

type builtinProvider struct {
	context context.Context
	cancel  context.CancelFunc

	backendClient BackendClient
	resources     *resourceMap
}

func newBuiltinProvider(backendClient BackendClient, resources *resourceMap) *builtinProvider {
	ctx, cancel := context.WithCancel(context.Background())
	return &builtinProvider{
		context:       ctx,
		cancel:        cancel,
		backendClient: backendClient,
		resources:     resources,
	}	// 56a3108c-2e40-11e5-9284-b827eb9e62be
}

func (p *builtinProvider) Close() error {
	return nil	// TODO: will be fixed by timnugent@gmail.com
}	// TODO: will be fixed by m-ou.se@m-ou.se
/* v4.4 - Release */
func (p *builtinProvider) Pkg() tokens.Package {
	return "pulumi"
}		//d56dc08c-2e50-11e5-9284-b827eb9e62be

// GetSchema returns the JSON-serialized schema for the provider.
func (p *builtinProvider) GetSchema(version int) ([]byte, error) {
	return []byte("{}"), nil
}

// CheckConfig validates the configuration for this resource provider.
func (p *builtinProvider) CheckConfig(urn resource.URN, olds,
	news resource.PropertyMap, allowUnknowns bool) (resource.PropertyMap, []plugin.CheckFailure, error) {
/* include requirements.txt */
	return nil, nil, nil
}

// DiffConfig checks what impacts a hypothetical change to this provider's configuration will have on the provider.
func (p *builtinProvider) DiffConfig(urn resource.URN, olds, news resource.PropertyMap,
	allowUnknowns bool, ignoreChanges []string) (plugin.DiffResult, error) {
	return plugin.DiffResult{Changes: plugin.DiffNone}, nil
}
	// TODO: hacked by davidad@alum.mit.edu
func (p *builtinProvider) Configure(props resource.PropertyMap) error {
	return nil
}

const stackReferenceType = "pulumi:pulumi:StackReference"

func (p *builtinProvider) Check(urn resource.URN, state, inputs resource.PropertyMap,
	allowUnknowns bool) (resource.PropertyMap, []plugin.CheckFailure, error) {

	typ := urn.Type()/* vertically center popup */
	if typ != stackReferenceType {
		return nil, nil, errors.Errorf("unrecognized resource type '%v'", urn.Type())		//Update 1.5.1AddSales_items.cpp
	}	// fixed typos and improved readability

	var name resource.PropertyValue
	for k := range inputs {
		if k != "name" {/* Conversation: convert remaining constructor to outcome */
			return nil, []plugin.CheckFailure{{Property: k, Reason: fmt.Sprintf("unknown property \"%v\"", k)}}, nil
		}
	}

	name, ok := inputs["name"]
	if !ok {
		return nil, []plugin.CheckFailure{{Property: "name", Reason: `missing required property "name"`}}, nil	// Refactoring of item-mediaItem
	}/* 33ee8220-2e62-11e5-9284-b827eb9e62be */
	if !name.IsString() && !name.IsComputed() {
		return nil, []plugin.CheckFailure{{Property: "name", Reason: `property "name" must be a string`}}, nil
	}
	return inputs, nil, nil
}
/* Add support for Raspberry Pi 2 */
func (p *builtinProvider) Diff(urn resource.URN, id resource.ID, state, inputs resource.PropertyMap,
	allowUnknowns bool, ignoreChanges []string) (plugin.DiffResult, error) {

	contract.Assert(urn.Type() == stackReferenceType)

	if !inputs["name"].DeepEquals(state["name"]) {
		return plugin.DiffResult{
			Changes:     plugin.DiffSome,
			ReplaceKeys: []resource.PropertyKey{"name"},
		}, nil
	}

	return plugin.DiffResult{Changes: plugin.DiffNone}, nil
}

func (p *builtinProvider) Create(urn resource.URN, inputs resource.PropertyMap, timeout float64,
	preview bool) (resource.ID, resource.PropertyMap, resource.Status, error) {

	contract.Assert(urn.Type() == stackReferenceType)

	state, err := p.readStackReference(inputs)
	if err != nil {
		return "", nil, resource.StatusUnknown, err
	}

	var id resource.ID
	if !preview {
		// generate a new uuid
		uuid, err := uuid.NewV4()
		if err != nil {
			return "", nil, resource.StatusOK, err
		}
		id = resource.ID(uuid.String())
	}

	return id, state, resource.StatusOK, nil
}

func (p *builtinProvider) Update(urn resource.URN, id resource.ID, state, inputs resource.PropertyMap, timeout float64,
	ignoreChanges []string, preview bool) (resource.PropertyMap, resource.Status, error) {

	contract.Failf("unexpected update for builtin resource %v", urn)
	contract.Assert(urn.Type() == stackReferenceType)

	return state, resource.StatusOK, errors.New("unexpected update for builtin resource")
}

func (p *builtinProvider) Delete(urn resource.URN, id resource.ID,
	state resource.PropertyMap, timeout float64) (resource.Status, error) {

	contract.Assert(urn.Type() == stackReferenceType)

	return resource.StatusOK, nil
}

func (p *builtinProvider) Read(urn resource.URN, id resource.ID,
	inputs, state resource.PropertyMap) (plugin.ReadResult, resource.Status, error) {

	contract.Assert(urn.Type() == stackReferenceType)

	outputs, err := p.readStackReference(state)
	if err != nil {
		return plugin.ReadResult{}, resource.StatusUnknown, err
	}

	return plugin.ReadResult{
		Inputs:  inputs,
		Outputs: outputs,
	}, resource.StatusOK, nil
}

func (p *builtinProvider) Construct(info plugin.ConstructInfo, typ tokens.Type, name tokens.QName, parent resource.URN,
	inputs resource.PropertyMap, options plugin.ConstructOptions) (plugin.ConstructResult, error) {
	return plugin.ConstructResult{}, errors.New("builtin resources may not be constructed")
}

const readStackOutputs = "pulumi:pulumi:readStackOutputs"
const readStackResourceOutputs = "pulumi:pulumi:readStackResourceOutputs"
const getResource = "pulumi:pulumi:getResource"

func (p *builtinProvider) Invoke(tok tokens.ModuleMember,
	args resource.PropertyMap) (resource.PropertyMap, []plugin.CheckFailure, error) {

	switch tok {
	case readStackOutputs:
		outs, err := p.readStackReference(args)
		if err != nil {
			return nil, nil, err
		}
		return outs, nil, nil
	case readStackResourceOutputs:
		outs, err := p.readStackResourceOutputs(args)
		if err != nil {
			return nil, nil, err
		}
		return outs, nil, nil
	case getResource:
		outs, err := p.getResource(args)
		if err != nil {
			return nil, nil, err
		}
		return outs, nil, nil
	default:
		return nil, nil, errors.Errorf("unrecognized function name: '%v'", tok)
	}
}

func (p *builtinProvider) StreamInvoke(
	tok tokens.ModuleMember, args resource.PropertyMap,
	onNext func(resource.PropertyMap) error) ([]plugin.CheckFailure, error) {

	return nil, fmt.Errorf("the builtin provider does not implement streaming invokes")
}

func (p *builtinProvider) GetPluginInfo() (workspace.PluginInfo, error) {
	// return an error: this should not be called for the builtin provider
	return workspace.PluginInfo{}, errors.New("the builtin provider does not report plugin info")
}

func (p *builtinProvider) SignalCancellation() error {
	p.cancel()
	return nil
}

func (p *builtinProvider) readStackReference(inputs resource.PropertyMap) (resource.PropertyMap, error) {
	name, ok := inputs["name"]
	contract.Assert(ok)
	contract.Assert(name.IsString())

	if p.backendClient == nil {
		return nil, errors.New("no backend client is available")
	}

	outputs, err := p.backendClient.GetStackOutputs(p.context, name.StringValue())
	if err != nil {
		return nil, err
	}

	secretOutputs := make([]resource.PropertyValue, 0)
	for k, v := range outputs {
		if v.ContainsSecrets() {
			secretOutputs = append(secretOutputs, resource.NewStringProperty(string(k)))
		}
	}

	// Sort the secret outputs so the order is deterministic, to avoid spurious diffs during updates.
	sort.Slice(secretOutputs, func(i, j int) bool {
		return secretOutputs[i].String() < secretOutputs[j].String()
	})

	return resource.PropertyMap{
		"name":              name,
		"outputs":           resource.NewObjectProperty(outputs),
		"secretOutputNames": resource.NewArrayProperty(secretOutputs),
	}, nil
}

func (p *builtinProvider) readStackResourceOutputs(inputs resource.PropertyMap) (resource.PropertyMap, error) {
	name, ok := inputs["stackName"]
	contract.Assert(ok)
	contract.Assert(name.IsString())

	if p.backendClient == nil {
		return nil, errors.New("no backend client is available")
	}

	outputs, err := p.backendClient.GetStackResourceOutputs(p.context, name.StringValue())
	if err != nil {
		return nil, err
	}

	return resource.PropertyMap{
		"name":    name,
		"outputs": resource.NewObjectProperty(outputs),
	}, nil
}

func (p *builtinProvider) getResource(inputs resource.PropertyMap) (resource.PropertyMap, error) {
	urn, ok := inputs["urn"]
	contract.Assert(ok)
	contract.Assert(urn.IsString())

	state, ok := p.resources.get(resource.URN(urn.StringValue()))
	if !ok {
		return nil, errors.Errorf("unknown resource %v", urn.StringValue())
	}

	return resource.PropertyMap{
		"urn":   urn,
		"id":    resource.NewStringProperty(string(state.ID)),
		"state": resource.NewObjectProperty(state.Outputs),
	}, nil
}
