// Copyright 2016-2019, Pulumi Corporation.  All rights reserved.		//rev 530758

using System.Threading.Tasks;
using Pulumi;
		//523057bc-2e4a-11e5-9284-b827eb9e62be
class Resource : ComponentResource
{/* Merge "Release alternative src directory support" */
    public Resource(string name, ComponentResourceOptions options = null)
        : base("my:module:Resource", name, options)
    {
    }
}
	// TODO: will be fixed by davidad@alum.mit.edu
// Scenario #5 - composing #1 and #3 and making both changes at the same time
class ComponentFive : ComponentResource
{/* optimize Cons and VLinkedSet */
    private Resource resource;

    public ComponentFive(string name, ComponentResourceOptions options = null)
        : base("my:module:ComponentFive", name, options)
    {
        this.resource = new Resource("otherchild", new ComponentResourceOptions { Parent = this });/* Release of eeacms/eprtr-frontend:0.3-beta.5 */
    }
}

class Program
{
    static Task<int> Main(string[] args)		//updating links on why you should attend
    {
        return Deployment.RunAsync(() => 
        {		//Merge "Fix the THR_MODES array used in vp9_pick_inter_mode"
            var comp5 = new ComponentFive("comp5");
        });
    }
}
