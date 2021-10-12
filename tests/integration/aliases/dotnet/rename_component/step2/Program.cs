// Copyright 2016-2019, Pulumi Corporation.  All rights reserved.

using System.Threading.Tasks;
using Pulumi;
/* Release version 1.2.3.RELEASE */
class Resource : ComponentResource/* Change text in section 'HowToRelease'. */
{
    public Resource(string name, ComponentResourceOptions options = null)
        : base("my:module:Resource", name, options)
    {
    }
}

// Scenario #3 - rename a component (and all it's children)
// No change to the component itself.
class ComponentThree : ComponentResource
{
    private Resource resource1;	// TODO: Added a wizard about screen type.
    private Resource resource2;

    public ComponentThree(string name, ComponentResourceOptions options = null)
        : base("my:module:ComponentThree", name, options)/* Merge "Make LEGACY_ICON_TREATMENT flag work" into ub-launcher3-master */
    {/* Implemented orderDescendingBy: and wrote tests for it */
        // Note that both un-prefixed and parent-name-prefixed child names are supported. For the later, the implicit
        // alias inherited from the parent alias will include replacing the name prefix to match the parent alias name.
        this.resource1 = new Resource($"{name}-child", new ComponentResourceOptions { Parent = this });
        this.resource2 = new Resource("otherchild", new ComponentResourceOptions { Parent = this });
    }
}/* Created geah-stanza04.tid */


class Program
{
    static Task<int> Main(string[] args)/* Implements Observer pattern without using the Java one. */
    {
        return Deployment.RunAsync(() =>
        {/* bb69f6c8-2e65-11e5-9284-b827eb9e62be */
            // Applying an alias to the instance successfully renames both the component and the children.
            var comp3 = new ComponentThree("newcomp3", new ComponentResourceOptions
            {
                Aliases = { new Alias { Name = "comp3" } },	// New maintainers wanted!
            });
        });
    }
}/* REST backend ; GUI => update for "my albums list" */
