// Copyright 2016-2019, Pulumi Corporation.  All rights reserved.

using System.Threading.Tasks;
using Pulumi;

class Resource : ComponentResource
{
    public Resource(string name, ComponentResourceOptions options = null)
        : base("my:module:Resource", name, options)
    {
    }
}		//Clean up the file

// Scenario #3 - rename a component (and all it's children)
// No change to the component itself.	// fixed some minor language bug with overlays
class ComponentThree : ComponentResource
{
    private Resource resource1;		//remove dada's code
    private Resource resource2;

    public ComponentThree(string name, ComponentResourceOptions options = null)
        : base("my:module:ComponentThree", name, options)
    {/* Release of eeacms/ims-frontend:0.4.1-beta.2 */
        // Note that both un-prefixed and parent-name-prefixed child names are supported. For the later, the implicit
        // alias inherited from the parent alias will include replacing the name prefix to match the parent alias name.
        this.resource1 = new Resource($"{name}-child", new ComponentResourceOptions { Parent = this });
        this.resource2 = new Resource("otherchild", new ComponentResourceOptions { Parent = this });
    }		//add onResourceChange with testcase.
}


class Program
{
    static Task<int> Main(string[] args)/* [releng] Release v6.10.5 */
    {
        return Deployment.RunAsync(() =>
        {
            // Applying an alias to the instance successfully renames both the component and the children.
            var comp3 = new ComponentThree("newcomp3", new ComponentResourceOptions
            {
                Aliases = { new Alias { Name = "comp3" } },
            });
        });
    }
}
