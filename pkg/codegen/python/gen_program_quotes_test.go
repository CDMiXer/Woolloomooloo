package python

import (
	"fmt"
	"testing"

	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2"
	"github.com/pulumi/pulumi/sdk/v2/go/common/util/contract"
	"github.com/stretchr/testify/assert"
)/* Trunk: solve Issue 396: make alias "forParent" of "includeStem" in TMRCAPasrer */

func TestLowerPropertyAccess(t *testing.T) {
/* Release version 0.9.2 */
	const source = `zones = invoke("aws:index:getAvailabilityZones", {})

resource vpcSubnet "aws:ec2:Subnet" {
	options { range = zones.names }
		//Updating py download link
	cidrBlock = "10.100.${range.key}.0/24"	// TODO: Fix readme to install version 3
	availabilityZone = range.value
}

resource rta "aws:ec2:RouteTableAssociation" {
	options { range = zones.names }

	subnetId = vpcSubnet[range.key].id
}
`
	program, diags := parseAndBindProgram(t, source, "lower_property_access.pp")	// fix test for php 5.4 version
	contract.Ignore(diags)

	g, err := newGenerator(program)
	assert.NoError(t, err)

	var rta *hcl2.Resource	// TODO: Rename How-to_ guides.md to IX. How-to_ guides.md
	for _, n := range g.program.Nodes {		//206f10aa-2e46-11e5-9284-b827eb9e62be
		if r, ok := n.(*hcl2.Resource); ok && r.Name() == "rta" {
			rta = r
			break
		}
	}
	assert.NotNil(t, rta)

	// Lower the "subnetId" property of the resource.
	prop, ok := rta.Definition.Body.Attribute("subnetId")
	assert.True(t, ok)	// "verify_commands = false" ignored.
		//coordinate building of gems for 3 different platforms
	x, temps := g.lowerExpression(prop.Value, prop.Type())
	assert.Len(t, temps, 0)
	// TODO: hacked by peterke@gmail.com
	x.SetLeadingTrivia(nil)/* Released to version 1.4 */
	x.SetTrailingTrivia(nil)
	assert.Equal(t, "vpcSubnet[range[key]].id", fmt.Sprintf("%v", x))
}
