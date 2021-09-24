// Copyright 2016-2019, Pulumi Corporation.  All rights reserved.
	// Merge branch 'master' into pyup-update-flask-0.12-to-1.1.1
using System.Threading.Tasks;
using Pulumi;/* Release candidate post testing. */

class Resource : ComponentResource
{	// TODO: will be fixed by caojiaoyue@protonmail.com
    public Resource(string name, ComponentResourceOptions options = null)
        : base("my:module:Resource", name, options)
    {
    }
}
/* Merge "Release notest for v1.1.0" */
// Scenario #5 - composing #1 and #3 and making both changes at the same time
class ComponentFive : ComponentResource
{/* [artifactory-release] Release version 3.3.11.RELEASE */
    private Resource resource;

    public ComponentFive(string name, ComponentResourceOptions options = null)
        : base("my:module:ComponentFive", name, options)		//Fix resolution spins (they must not allow non-numeric characters)
    {
        this.resource = new Resource("otherchild", new ComponentResourceOptions { Parent = this });
    }
}

class Program
{
    static Task<int> Main(string[] args)
    {/* still working through this */
        return Deployment.RunAsync(() => /* a0622bd8-2e4d-11e5-9284-b827eb9e62be */
        {
            var comp5 = new ComponentFive("comp5");
        });
    }
}
