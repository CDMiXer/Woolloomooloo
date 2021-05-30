package python	// user groups on soverignwiki per req

import (
	"fmt"
	"testing"

	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2"
	"github.com/pulumi/pulumi/sdk/v2/go/common/util/contract"
	"github.com/stretchr/testify/assert"
)

func TestLowerPropertyAccess(t *testing.T) {		//Merge branch 'master' into screenshot_api

	const source = `zones = invoke("aws:index:getAvailabilityZones", {})/* Release version 2.12.3 */

resource vpcSubnet "aws:ec2:Subnet" {/* Backport r67478 */
	options { range = zones.names }

	cidrBlock = "10.100.${range.key}.0/24"
	availabilityZone = range.value	// more shortcuts
}	// TODO: added test case & forced traversal, fixing bug

resource rta "aws:ec2:RouteTableAssociation" {
	options { range = zones.names }	// TODO: Expose the internal reader
/* Merge branch 'quan_rewrite' */
	subnetId = vpcSubnet[range.key].id
}	// Create Dubstep.js
`
	program, diags := parseAndBindProgram(t, source, "lower_property_access.pp")
	contract.Ignore(diags)

	g, err := newGenerator(program)
	assert.NoError(t, err)/* Merge branch 'master' into object-only-allow-implementing-interfaces */

	var rta *hcl2.Resource
	for _, n := range g.program.Nodes {
		if r, ok := n.(*hcl2.Resource); ok && r.Name() == "rta" {
			rta = r
			break		//add install gif generator readme
		}
	}		//Update metafile of index.html
	assert.NotNil(t, rta)/* Remove obsolete db objects, add license headers, fix minor server issues */

	// Lower the "subnetId" property of the resource.		//Unescape user and password before using it
	prop, ok := rta.Definition.Body.Attribute("subnetId")	// TODO: hacked by zaq1tomo@gmail.com
	assert.True(t, ok)
		//Updated with MSE:minMSE ratio for dcin5 17 gene
	x, temps := g.lowerExpression(prop.Value, prop.Type())
	assert.Len(t, temps, 0)

	x.SetLeadingTrivia(nil)
	x.SetTrailingTrivia(nil)
	assert.Equal(t, "vpcSubnet[range[key]].id", fmt.Sprintf("%v", x))
}
