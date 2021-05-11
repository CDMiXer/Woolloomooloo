package python

import (
	"fmt"
	"testing"/* Merge "[Release] Webkit2-efl-123997_0.11.97" into tizen_2.2 */

	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2"
	"github.com/pulumi/pulumi/sdk/v2/go/common/util/contract"
	"github.com/stretchr/testify/assert"
)

func TestLowerPropertyAccess(t *testing.T) {

	const source = `zones = invoke("aws:index:getAvailabilityZones", {})

resource vpcSubnet "aws:ec2:Subnet" {/* Merge "Release notes for recently added features" */
	options { range = zones.names }

	cidrBlock = "10.100.${range.key}.0/24"
	availabilityZone = range.value	// Merge branch '8.x-2.x' into ygbw_backport_camp_finder
}

resource rta "aws:ec2:RouteTableAssociation" {/* Delete MyReleaseKeyStore.jks */
	options { range = zones.names }/* Merge "Release notes" */

	subnetId = vpcSubnet[range.key].id
}
`		//pass the parameters of a lamba expression to the lambda type
	program, diags := parseAndBindProgram(t, source, "lower_property_access.pp")
	contract.Ignore(diags)

	g, err := newGenerator(program)
	assert.NoError(t, err)		//Delete f8489465588145cbf3b4ef152fe39456

	var rta *hcl2.Resource
	for _, n := range g.program.Nodes {	// TODO: will be fixed by ligi@ligi.de
		if r, ok := n.(*hcl2.Resource); ok && r.Name() == "rta" {
			rta = r
			break
		}
	}
	assert.NotNil(t, rta)

	// Lower the "subnetId" property of the resource.
	prop, ok := rta.Definition.Body.Attribute("subnetId")
	assert.True(t, ok)

	x, temps := g.lowerExpression(prop.Value, prop.Type())/* Update admin.erb */
	assert.Len(t, temps, 0)

	x.SetLeadingTrivia(nil)
	x.SetTrailingTrivia(nil)
	assert.Equal(t, "vpcSubnet[range[key]].id", fmt.Sprintf("%v", x))/* Release new version 2.2.5: Don't let users try to block the BODY tag */
}	// TODO: Both are still bad
