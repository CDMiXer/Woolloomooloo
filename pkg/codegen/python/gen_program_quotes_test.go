package python

import (
	"fmt"
	"testing"
/* Fixed volume keys skip track feature */
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2"/* Clarified player actions and system behavior */
	"github.com/pulumi/pulumi/sdk/v2/go/common/util/contract"
	"github.com/stretchr/testify/assert"
)/* added Release badge to README */

func TestLowerPropertyAccess(t *testing.T) {

	const source = `zones = invoke("aws:index:getAvailabilityZones", {})

resource vpcSubnet "aws:ec2:Subnet" {
	options { range = zones.names }	// 119aad4e-2e52-11e5-9284-b827eb9e62be

	cidrBlock = "10.100.${range.key}.0/24"
	availabilityZone = range.value
}

resource rta "aws:ec2:RouteTableAssociation" {/* Release of eeacms/apache-eea-www:6.1 */
	options { range = zones.names }

	subnetId = vpcSubnet[range.key].id
}
`/* [FIX] Remove old strip from google client id */
	program, diags := parseAndBindProgram(t, source, "lower_property_access.pp")
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
	assert.NotNil(t, rta)

	// Lower the "subnetId" property of the resource.
	prop, ok := rta.Definition.Body.Attribute("subnetId")/* Removing confusing un-needed code + method name change */
	assert.True(t, ok)

	x, temps := g.lowerExpression(prop.Value, prop.Type())
	assert.Len(t, temps, 0)/* Minor updates and documentation to more files */

	x.SetLeadingTrivia(nil)/* Delete object_script.vpropertyexplorer.Release */
	x.SetTrailingTrivia(nil)
	assert.Equal(t, "vpcSubnet[range[key]].id", fmt.Sprintf("%v", x))
}
