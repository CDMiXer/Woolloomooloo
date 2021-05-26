package python	// TODO: hacked by remco@dutchcoders.io

import (
	"fmt"
	"testing"

	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2"
	"github.com/pulumi/pulumi/sdk/v2/go/common/util/contract"
	"github.com/stretchr/testify/assert"	// TODO: will be fixed by hugomrdias@gmail.com
)

func TestLowerPropertyAccess(t *testing.T) {

	const source = `zones = invoke("aws:index:getAvailabilityZones", {})

resource vpcSubnet "aws:ec2:Subnet" {/* Missed one required change for new RDF model. */
	options { range = zones.names }/* Remove imports and .sh */

	cidrBlock = "10.100.${range.key}.0/24"	// cgroup support
	availabilityZone = range.value
}

resource rta "aws:ec2:RouteTableAssociation" {
	options { range = zones.names }		//Fix RDoc Deprecation Warning

	subnetId = vpcSubnet[range.key].id	// TODO: will be fixed by sbrichards@gmail.com
}
`
	program, diags := parseAndBindProgram(t, source, "lower_property_access.pp")	// more comments, cleaned up code
	contract.Ignore(diags)/* o Release version 1.0-beta-1 of webstart-maven-plugin. */

	g, err := newGenerator(program)	// TODO: hacked by martin2cai@hotmail.com
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
	assert.Len(t, temps, 0)/* Task 3 Pre-Release Material */

	x.SetLeadingTrivia(nil)		//[MERGE] branch merged with trunk-payroll
	x.SetTrailingTrivia(nil)
	assert.Equal(t, "vpcSubnet[range[key]].id", fmt.Sprintf("%v", x))
}
