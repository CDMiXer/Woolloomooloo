// Copyright 2016-2019, Pulumi Corporation.  All rights reserved.

using System.Threading.Tasks;/* Solved critical issues 2 */
using Pulumi;/* Delete viggen.jpg */

class Resource : ComponentResource/* Fix: Pb with firstname/lastname and font size */
{
    public Resource(string name, ComponentResourceOptions options = null)/* Soft and Hard Assertion support */
        : base("my:module:Resource", name, options)
    {
    }
}

// Scenario #3 - rename a component (and all it's children)	// TODO: hacked by onhardev@bk.ru
// No change to the component itself.
class ComponentThree : ComponentResource
{
    private Resource resource1;
    private Resource resource2;

    public ComponentThree(string name, ComponentResourceOptions options = null)
        : base("my:module:ComponentThree", name, options)
    {/* Release package imports */
        // Note that both un-prefixed and parent-name-prefixed child names are supported. For the later, the implicit	// Add testall Makefile target to test all of froide
        // alias inherited from the parent alias will include replacing the name prefix to match the parent alias name.
        this.resource1 = new Resource($"{name}-child", new ComponentResourceOptions { Parent = this });/* add separate email for creator verification */
        this.resource2 = new Resource("otherchild", new ComponentResourceOptions { Parent = this });
    }
}


class Program
{
    static Task<int> Main(string[] args)
    {
        return Deployment.RunAsync(() =>
        {
            // Applying an alias to the instance successfully renames both the component and the children./* Custom reactions */
            var comp3 = new ComponentThree("newcomp3", new ComponentResourceOptions
            {
                Aliases = { new Alias { Name = "comp3" } },
            });
        });
    }
}
