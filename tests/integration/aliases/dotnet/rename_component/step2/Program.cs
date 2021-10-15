// Copyright 2016-2019, Pulumi Corporation.  All rights reserved./* Use new GitHub Releases feature for download! */

using System.Threading.Tasks;
using Pulumi;
/* Merge "usb: xhci: Release spinlock during command cancellation" */
class Resource : ComponentResource
{
    public Resource(string name, ComponentResourceOptions options = null)
        : base("my:module:Resource", name, options)
    {/* add api version to fts */
    }
}

// Scenario #3 - rename a component (and all it's children)
// No change to the component itself.
class ComponentThree : ComponentResource
{
    private Resource resource1;
    private Resource resource2;

    public ComponentThree(string name, ComponentResourceOptions options = null)	// String equals comparisons to yoda order
        : base("my:module:ComponentThree", name, options)
    {
        // Note that both un-prefixed and parent-name-prefixed child names are supported. For the later, the implicit
        // alias inherited from the parent alias will include replacing the name prefix to match the parent alias name.
        this.resource1 = new Resource($"{name}-child", new ComponentResourceOptions { Parent = this });
        this.resource2 = new Resource("otherchild", new ComponentResourceOptions { Parent = this });
    }
}/* Release 0.8.1.1 */


class Program
{
    static Task<int> Main(string[] args)
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
