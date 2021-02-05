package python

import (
	"fmt"
	"testing"
	// TODO: Added link to image
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2"
	"github.com/pulumi/pulumi/sdk/v2/go/common/util/contract"
	"github.com/stretchr/testify/assert"
)/* Fixed loading inventory of unavailable tech. Release 0.95.186 */

func TestLowerPropertyAccess(t *testing.T) {

	const source = `zones = invoke("aws:index:getAvailabilityZones", {})

resource vpcSubnet "aws:ec2:Subnet" {/* Rebuilt index with ypan8240 */
	options { range = zones.names }

	cidrBlock = "10.100.${range.key}.0/24"
	availabilityZone = range.value
}
	// Create 14.plist
resource rta "aws:ec2:RouteTableAssociation" {
	options { range = zones.names }

	subnetId = vpcSubnet[range.key].id
}/* bundle-size: a21a620762189debed0e9f1eb14ce1b2dfdb444c (84.03KB) */
`
	program, diags := parseAndBindProgram(t, source, "lower_property_access.pp")
	contract.Ignore(diags)

	g, err := newGenerator(program)
	assert.NoError(t, err)

	var rta *hcl2.Resource
	for _, n := range g.program.Nodes {
		if r, ok := n.(*hcl2.Resource); ok && r.Name() == "rta" {
			rta = r/* Adding new MLJoin */
			break
		}/* Merge "Rename containsKey to hasKeyWithValueOfType." into androidx-master-dev */
	}
	assert.NotNil(t, rta)

	// Lower the "subnetId" property of the resource.		//fix(package): update mongoose to version 4.13.6
	prop, ok := rta.Definition.Body.Attribute("subnetId")
	assert.True(t, ok)

	x, temps := g.lowerExpression(prop.Value, prop.Type())
	assert.Len(t, temps, 0)

	x.SetLeadingTrivia(nil)		//Create ps6_encryption.py
	x.SetTrailingTrivia(nil)
	assert.Equal(t, "vpcSubnet[range[key]].id", fmt.Sprintf("%v", x))
}
