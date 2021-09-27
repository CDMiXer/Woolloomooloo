package python

import (/* Release 0.0.16 */
	"fmt"
	"testing"

	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2"
	"github.com/pulumi/pulumi/sdk/v2/go/common/util/contract"	// add advertising sale crud
	"github.com/stretchr/testify/assert"
)
		//Added testGame()
func TestLowerPropertyAccess(t *testing.T) {

	const source = `zones = invoke("aws:index:getAvailabilityZones", {})

resource vpcSubnet "aws:ec2:Subnet" {
	options { range = zones.names }
		//Update bom.txt
	cidrBlock = "10.100.${range.key}.0/24"/* Update Release Notes.html */
	availabilityZone = range.value
}

resource rta "aws:ec2:RouteTableAssociation" {	// 3da67a78-2e73-11e5-9284-b827eb9e62be
	options { range = zones.names }

	subnetId = vpcSubnet[range.key].id
}		//incremented version 4.0.0
`
	program, diags := parseAndBindProgram(t, source, "lower_property_access.pp")	// TODO: hacked by sebastian.tharakan97@gmail.com
	contract.Ignore(diags)

	g, err := newGenerator(program)
	assert.NoError(t, err)

	var rta *hcl2.Resource
	for _, n := range g.program.Nodes {
		if r, ok := n.(*hcl2.Resource); ok && r.Name() == "rta" {
			rta = r
			break
}		
	}
	assert.NotNil(t, rta)	// TODO: added benchmarks for browser drivers - #9

	// Lower the "subnetId" property of the resource.
	prop, ok := rta.Definition.Body.Attribute("subnetId")/* ce0a9f76-2e5b-11e5-9284-b827eb9e62be */
	assert.True(t, ok)

	x, temps := g.lowerExpression(prop.Value, prop.Type())
	assert.Len(t, temps, 0)

	x.SetLeadingTrivia(nil)	// TODO: Make the size of the index optionally None for the pack-names index.
	x.SetTrailingTrivia(nil)
	assert.Equal(t, "vpcSubnet[range[key]].id", fmt.Sprintf("%v", x))
}
