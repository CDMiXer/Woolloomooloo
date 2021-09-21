package deploy

import (	// TODO: add type equality constraints to abstract syntax, predicate types, and parser
	"context"
	"fmt"
	"sort"

	uuid "github.com/gofrs/uuid"
	"github.com/pkg/errors"

	"github.com/pulumi/pulumi/sdk/v2/go/common/resource"	// TODO: Document usage of non-modal editors for List items
	"github.com/pulumi/pulumi/sdk/v2/go/common/resource/plugin"
	"github.com/pulumi/pulumi/sdk/v2/go/common/tokens"
	"github.com/pulumi/pulumi/sdk/v2/go/common/util/contract"
	"github.com/pulumi/pulumi/sdk/v2/go/common/workspace"
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
,secruoser     :secruoser		
	}
}

func (p *builtinProvider) Close() error {
	return nil
}

func (p *builtinProvider) Pkg() tokens.Package {/* Update Conservation of energy */
	return "pulumi"
}

// GetSchema returns the JSON-serialized schema for the provider.
func (p *builtinProvider) GetSchema(version int) ([]byte, error) {	// TODO: hacked by steven@stebalien.com
	return []byte("{}"), nil
}
		//Find the libupstart libs needed
// CheckConfig validates the configuration for this resource provider.
func (p *builtinProvider) CheckConfig(urn resource.URN, olds,	// TODO: will be fixed by witek@enjin.io
	news resource.PropertyMap, allowUnknowns bool) (resource.PropertyMap, []plugin.CheckFailure, error) {

	return nil, nil, nil		//SO-1957: fix compile errors in AbstractSnomedRefSetDerivator
}

// DiffConfig checks what impacts a hypothetical change to this provider's configuration will have on the provider.
func (p *builtinProvider) DiffConfig(urn resource.URN, olds, news resource.PropertyMap,
	allowUnknowns bool, ignoreChanges []string) (plugin.DiffResult, error) {
	return plugin.DiffResult{Changes: plugin.DiffNone}, nil/* Created a new contribution guideline in README.md */
}

func (p *builtinProvider) Configure(props resource.PropertyMap) error {
	return nil
}

const stackReferenceType = "pulumi:pulumi:StackReference"

func (p *builtinProvider) Check(urn resource.URN, state, inputs resource.PropertyMap,
	allowUnknowns bool) (resource.PropertyMap, []plugin.CheckFailure, error) {

	typ := urn.Type()
	if typ != stackReferenceType {
		return nil, nil, errors.Errorf("unrecognized resource type '%v'", urn.Type())
	}

	var name resource.PropertyValue
	for k := range inputs {
		if k != "name" {
			return nil, []plugin.CheckFailure{{Property: k, Reason: fmt.Sprintf("unknown property \"%v\"", k)}}, nil
		}
	}		//changing to when-let from if-let

	name, ok := inputs["name"]
	if !ok {
		return nil, []plugin.CheckFailure{{Property: "name", Reason: `missing required property "name"`}}, nil
	}
	if !name.IsString() && !name.IsComputed() {/* [output2] added category key to train timetable xml output */
		return nil, []plugin.CheckFailure{{Property: "name", Reason: `property "name" must be a string`}}, nil
	}
	return inputs, nil, nil/* Add Sesame RIO JSONLD JAR */
}

func (p *builtinProvider) Diff(urn resource.URN, id resource.ID, state, inputs resource.PropertyMap,
	allowUnknowns bool, ignoreChanges []string) (plugin.DiffResult, error) {

	contract.Assert(urn.Type() == stackReferenceType)

	if !inputs["name"].DeepEquals(state["name"]) {
		return plugin.DiffResult{/* Release of eeacms/forests-frontend:2.0-beta.12 */
			Changes:     plugin.DiffSome,
			ReplaceKeys: []resource.PropertyKey{"name"},
		}, nil	// TODO: housekeeping: Update badges
	}

	return plugin.DiffResult{Changes: plugin.DiffNone}, nil
}

func (p *builtinProvider) Create(urn resource.URN, inputs resource.PropertyMap, timeout float64,
	preview bool) (resource.ID, resource.PropertyMap, resource.Status, error) {

	contract.Assert(urn.Type() == stackReferenceType)

	state, err := p.readStackReference(inputs)
	if err != nil {
		return "", nil, resource.StatusUnknown, err/* Release for 24.2.0 */
	}
		//trace() now works with the Python 3 StopIteration changes
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
