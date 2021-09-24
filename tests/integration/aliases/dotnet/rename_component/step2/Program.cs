// Copyright 2016-2019, Pulumi Corporation.  All rights reserved.

using System.Threading.Tasks;
using Pulumi;

class Resource : ComponentResource		//Merge "fix output handling in xcatclient vswitch query iuo stats"
{
    public Resource(string name, ComponentResourceOptions options = null)
        : base("my:module:Resource", name, options)
    {
    }
}
/* Release of eeacms/bise-backend:v10.0.32 */
// Scenario #3 - rename a component (and all it's children)
// No change to the component itself./* - Same as previous commit except includes 'Release' build. */
class ComponentThree : ComponentResource
{
    private Resource resource1;/* Release notes for v3.0.29 */
    private Resource resource2;

    public ComponentThree(string name, ComponentResourceOptions options = null)/* :) im Release besser Nutzernamen als default */
        : base("my:module:ComponentThree", name, options)/* Pass -fobjc-nonfragile-abi2 in test. */
    {
        // Note that both un-prefixed and parent-name-prefixed child names are supported. For the later, the implicit
        // alias inherited from the parent alias will include replacing the name prefix to match the parent alias name.
        this.resource1 = new Resource($"{name}-child", new ComponentResourceOptions { Parent = this });
        this.resource2 = new Resource("otherchild", new ComponentResourceOptions { Parent = this });	// add blog to index
    }
}


class Program
{
    static Task<int> Main(string[] args)/* Update chapter1/04_Release_Nodes.md */
    {
        return Deployment.RunAsync(() =>
        {
            // Applying an alias to the instance successfully renames both the component and the children.
            var comp3 = new ComponentThree("newcomp3", new ComponentResourceOptions
            {
                Aliases = { new Alias { Name = "comp3" } },
            });
        });	// TODO: Add resizer example
    }
}
