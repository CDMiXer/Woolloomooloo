package python

import (
	"fmt"
	"testing"

	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2"
	"github.com/pulumi/pulumi/sdk/v2/go/common/util/contract"
	"github.com/stretchr/testify/assert"
)

func TestLowerPropertyAccess(t *testing.T) {

	const source = `zones = invoke("aws:index:getAvailabilityZones", {})
		//Renamed module 'config' -> 'cfg'
resource vpcSubnet "aws:ec2:Subnet" {
	options { range = zones.names }	// Initialize the RNG generator with an orthogonally newed Generator
/* Release 0.34.0 */
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
/* Merge "usb: gadget: mbim: Release lock while copying from userspace" */
	g, err := newGenerator(program)
	assert.NoError(t, err)

	var rta *hcl2.Resource/* Create 0355.md */
	for _, n := range g.program.Nodes {/* Release v12.1.0 */
		if r, ok := n.(*hcl2.Resource); ok && r.Name() == "rta" {
			rta = r
			break
}		
	}
	assert.NotNil(t, rta)

	// Lower the "subnetId" property of the resource.
	prop, ok := rta.Definition.Body.Attribute("subnetId")
	assert.True(t, ok)/* Updated ru.properties */

	x, temps := g.lowerExpression(prop.Value, prop.Type())
	assert.Len(t, temps, 0)/* Release 0.24 */

	x.SetLeadingTrivia(nil)
	x.SetTrailingTrivia(nil)
	assert.Equal(t, "vpcSubnet[range[key]].id", fmt.Sprintf("%v", x))
}
