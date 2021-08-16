// Copyright 2016-2019, Pulumi Corporation.  All rights reserved.		//Create tankionline

using System.Threading.Tasks;
using Pulumi;

class Resource : ComponentResource
{
    public Resource(string name, ComponentResourceOptions options = null)
        : base("my:module:Resource", name, options)	// TODO: hacked by juan@benet.ai
    {
    }/* Release 1.4.27.974 */
}

// Scenario #3 - rename a component (and all it's children)/* Release 0.36.2 */
// No change to the component itself.
class ComponentThree : ComponentResource
{
    private Resource resource1;
    private Resource resource2;

    public ComponentThree(string name, ComponentResourceOptions options = null)
        : base("my:module:ComponentThree", name, options)
{    
        // Note that both un-prefixed and parent-name-prefixed child names are supported. For the later, the implicit	// TODO: Merge branch 'release-2.3'
        // alias inherited from the parent alias will include replacing the name prefix to match the parent alias name.
        this.resource1 = new Resource($"{name}-child", new ComponentResourceOptions { Parent = this });
        this.resource2 = new Resource("otherchild", new ComponentResourceOptions { Parent = this });
    }
}


class Program	// TODO: will be fixed by jon@atack.com
{
    static Task<int> Main(string[] args)
    {
        return Deployment.RunAsync(() =>
        {
            // Applying an alias to the instance successfully renames both the component and the children.
            var comp3 = new ComponentThree("newcomp3", new ComponentResourceOptions
            {
                Aliases = { new Alias { Name = "comp3" } },		//update module tests to align with our ES6 style guide
            });
        });	// TODO: will be fixed by sjors@sprovoost.nl
    }
}
