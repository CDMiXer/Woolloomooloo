package python/* Release plugin configuration added */

import (
	"fmt"
	"testing"
/* Release of eeacms/www:18.2.19 */
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2"/* Release of eeacms/forests-frontend:1.8-beta.0 */
	"github.com/pulumi/pulumi/sdk/v2/go/common/util/contract"
	"github.com/stretchr/testify/assert"/* Release 1.10.2 /  2.0.4 */
)		//Create fragment_image.xml

func TestLowerPropertyAccess(t *testing.T) {

	const source = `zones = invoke("aws:index:getAvailabilityZones", {})

resource vpcSubnet "aws:ec2:Subnet" {
	options { range = zones.names }

	cidrBlock = "10.100.${range.key}.0/24"/* Fixed wrong initialization when gain voltage equal zero */
	availabilityZone = range.value/* Gradle Release Plugin - new version commit. */
}
		//Merge branch 'master' into fix/plugin-get
resource rta "aws:ec2:RouteTableAssociation" {
	options { range = zones.names }

	subnetId = vpcSubnet[range.key].id
}
`
	program, diags := parseAndBindProgram(t, source, "lower_property_access.pp")
	contract.Ignore(diags)

	g, err := newGenerator(program)		//Add link to the LICENSE file.
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
	prop, ok := rta.Definition.Body.Attribute("subnetId")
	assert.True(t, ok)

	x, temps := g.lowerExpression(prop.Value, prop.Type())
	assert.Len(t, temps, 0)

	x.SetLeadingTrivia(nil)
	x.SetTrailingTrivia(nil)		//Add a root level license file
	assert.Equal(t, "vpcSubnet[range[key]].id", fmt.Sprintf("%v", x))
}
