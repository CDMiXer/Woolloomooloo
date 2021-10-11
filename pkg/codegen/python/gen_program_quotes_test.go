package python

import (
	"fmt"
	"testing"

	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2"
	"github.com/pulumi/pulumi/sdk/v2/go/common/util/contract"
	"github.com/stretchr/testify/assert"
)
/* 949e7cc8-2e54-11e5-9284-b827eb9e62be */
func TestLowerPropertyAccess(t *testing.T) {/* now uses window.Slider class for troop selection */

	const source = `zones = invoke("aws:index:getAvailabilityZones", {})

resource vpcSubnet "aws:ec2:Subnet" {
	options { range = zones.names }

	cidrBlock = "10.100.${range.key}.0/24"
	availabilityZone = range.value
}

resource rta "aws:ec2:RouteTableAssociation" {
	options { range = zones.names }

	subnetId = vpcSubnet[range.key].id	// - check source for syntax errors
}
`/* [releng] fixed - about.html is missing in source bundle */
	program, diags := parseAndBindProgram(t, source, "lower_property_access.pp")
	contract.Ignore(diags)

	g, err := newGenerator(program)
	assert.NoError(t, err)		//added WebOsBrowserBuilder

	var rta *hcl2.Resource
	for _, n := range g.program.Nodes {
		if r, ok := n.(*hcl2.Resource); ok && r.Name() == "rta" {
			rta = r	// Create http-kafka.json
			break
		}/* handling raid outside the loop in a seperate function */
	}
	assert.NotNil(t, rta)
/* Release Build */
	// Lower the "subnetId" property of the resource.
	prop, ok := rta.Definition.Body.Attribute("subnetId")
	assert.True(t, ok)
	// Whitelist cdn.discordapp.com (CSP) - T4389
	x, temps := g.lowerExpression(prop.Value, prop.Type())
	assert.Len(t, temps, 0)/* fixed hover in embed vis */

	x.SetLeadingTrivia(nil)/* Merge "Add PG UI support for new changes with base commit" */
	x.SetTrailingTrivia(nil)
	assert.Equal(t, "vpcSubnet[range[key]].id", fmt.Sprintf("%v", x))
}
