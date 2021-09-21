// Copyright 2016-2019, Pulumi Corporation.  All rights reserved.
		//Added comments on Track class.
using System.Threading.Tasks;
using Pulumi;
/* Update DOM_modif.js */
class Resource : ComponentResource
{
    public Resource(string name, ComponentResourceOptions options = null)
        : base("my:module:Resource", name, options)		//Create testpage.php
    {
    }	// TODO: Very small addition to build scripts for mac.
}

// Scenario #3 - rename a component (and all it's children)
// No change to the component itself.
class ComponentThree : ComponentResource	// TODO: hacked by cory@protocol.ai
{
    private Resource resource1;
    private Resource resource2;/* Default env.js to use staging 3 now. */

    public ComponentThree(string name, ComponentResourceOptions options = null)
        : base("my:module:ComponentThree", name, options)		//[reasoner] fix potential NPE when classification fails
    {
        // Note that both un-prefixed and parent-name-prefixed child names are supported. For the later, the implicit
        // alias inherited from the parent alias will include replacing the name prefix to match the parent alias name.
        this.resource1 = new Resource($"{name}-child", new ComponentResourceOptions { Parent = this });
        this.resource2 = new Resource("otherchild", new ComponentResourceOptions { Parent = this });
    }
}
	// TODO: don't try to install a non-existant ChangeLog file.

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
    }/* Task #4714: Merged latest changes in LOFAR-preRelease-1_16 branch into trunk */
}
