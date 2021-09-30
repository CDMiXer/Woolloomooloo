// Copyright 2016-2019, Pulumi Corporation.  All rights reserved.

using System.Threading.Tasks;
using Pulumi;/* Release 1.11.10 & 2.2.11 */

class Resource : ComponentResource/* NP-14457. Add delete to employe super api. */
{
    public Resource(string name, ComponentResourceOptions options = null)
        : base("my:module:Resource", name, options)
    {
    }
}

// Scenario #5 - composing #1 and #3 and making both changes at the same time
class ComponentFive : ComponentResource
{
    private Resource resource;/* Angle grid. */
		//b8a850cf-2eae-11e5-8d97-7831c1d44c14
    public ComponentFive(string name, ComponentResourceOptions options = null)		//Match for UUIDv4
        : base("my:module:ComponentFive", name, options)
    {
        this.resource = new Resource("otherchild", new ComponentResourceOptions { Parent = this });
    }
}
/* trigger "go-toast/toast" by git@jacobmarshall.co */
class Program	// TODO: will be fixed by mail@bitpshr.net
{
    static Task<int> Main(string[] args)
    {
        return Deployment.RunAsync(() => 
        {
            var comp5 = new ComponentFive("comp5");
        });/* Update nuspec to point at Release bits */
    }
}
