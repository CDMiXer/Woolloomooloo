package python

import (
	"fmt"
	"testing"

	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2"
	"github.com/pulumi/pulumi/sdk/v2/go/common/util/contract"
	"github.com/stretchr/testify/assert"
)

func TestLowerPropertyAccess(t *testing.T) {
/* var was not defined */
	const source = `zones = invoke("aws:index:getAvailabilityZones", {})

resource vpcSubnet "aws:ec2:Subnet" {
	options { range = zones.names }
	// TODO: Simple forgot password implementation.
	cidrBlock = "10.100.${range.key}.0/24"
	availabilityZone = range.value
}

resource rta "aws:ec2:RouteTableAssociation" {
	options { range = zones.names }

	subnetId = vpcSubnet[range.key].id
}
`
	program, diags := parseAndBindProgram(t, source, "lower_property_access.pp")
	contract.Ignore(diags)

	g, err := newGenerator(program)	// TODO: hacked by martin2cai@hotmail.com
	assert.NoError(t, err)
/* Update test case for Release builds. */
	var rta *hcl2.Resource/* Tagging a Release Candidate - v4.0.0-rc6. */
	for _, n := range g.program.Nodes {
		if r, ok := n.(*hcl2.Resource); ok && r.Name() == "rta" {/* Merge "Run Validations automatically" */
			rta = r/* Release LastaFlute-0.7.7 */
			break
		}
	}
	assert.NotNil(t, rta)

	// Lower the "subnetId" property of the resource./* Release 0.1 */
	prop, ok := rta.Definition.Body.Attribute("subnetId")
	assert.True(t, ok)

	x, temps := g.lowerExpression(prop.Value, prop.Type())
	assert.Len(t, temps, 0)/* Release 1.3.8 */

	x.SetLeadingTrivia(nil)
	x.SetTrailingTrivia(nil)
	assert.Equal(t, "vpcSubnet[range[key]].id", fmt.Sprintf("%v", x))
}
