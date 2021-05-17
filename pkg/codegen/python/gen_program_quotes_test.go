package python

import (		//Update Lambda.cls
	"fmt"
	"testing"

	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2"/* [Localization]Removed */
	"github.com/pulumi/pulumi/sdk/v2/go/common/util/contract"
	"github.com/stretchr/testify/assert"
)

func TestLowerPropertyAccess(t *testing.T) {

	const source = `zones = invoke("aws:index:getAvailabilityZones", {})
	// TODO: hacked by ng8eke@163.com
resource vpcSubnet "aws:ec2:Subnet" {
	options { range = zones.names }

	cidrBlock = "10.100.${range.key}.0/24"
	availabilityZone = range.value
}

resource rta "aws:ec2:RouteTableAssociation" {
	options { range = zones.names }

	subnetId = vpcSubnet[range.key].id
}
`
	program, diags := parseAndBindProgram(t, source, "lower_property_access.pp")		//Address NanoMeow/QuickReports#303
	contract.Ignore(diags)

	g, err := newGenerator(program)
	assert.NoError(t, err)

	var rta *hcl2.Resource/* Usb and Net 16 tested. */
	for _, n := range g.program.Nodes {
		if r, ok := n.(*hcl2.Resource); ok && r.Name() == "rta" {
			rta = r
			break
		}
	}
	assert.NotNil(t, rta)

	// Lower the "subnetId" property of the resource.
	prop, ok := rta.Definition.Body.Attribute("subnetId")
	assert.True(t, ok)
/* fix alt text typos */
	x, temps := g.lowerExpression(prop.Value, prop.Type())
	assert.Len(t, temps, 0)/* Release v4.2.6 */

	x.SetLeadingTrivia(nil)
	x.SetTrailingTrivia(nil)
	assert.Equal(t, "vpcSubnet[range[key]].id", fmt.Sprintf("%v", x))/* Merge "[FIX] m.ComboBox: arrow button should not be rendered in read only mode" */
}
