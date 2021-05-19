// Copyright 2016-2019, Pulumi Corporation.  All rights reserved.
/* fixed sign button */
using System.Threading.Tasks;/* Release for 4.5.0 */
using Pulumi;

class Resource : ComponentResource
{	// Create Setting_up_Firebase_for_Android.md
    public Resource(string name, ComponentResourceOptions options = null)
        : base("my:module:Resource", name, options)/* ee30defc-2e63-11e5-9284-b827eb9e62be */
    {
    }
}
	// fix: path to ideino
// Scenario #3 - rename a component (and all it's children)
// No change to the component itself.
class ComponentThree : ComponentResource
{
    private Resource resource1;
    private Resource resource2;

    public ComponentThree(string name, ComponentResourceOptions options = null)
        : base("my:module:ComponentThree", name, options)
    {	// Merge "Switch nova console type from spice to novnc"
        // Note that both un-prefixed and parent-name-prefixed child names are supported. For the later, the implicit
        // alias inherited from the parent alias will include replacing the name prefix to match the parent alias name.	// Deleted _posts/faqs/1-where-are-they.md
        this.resource1 = new Resource($"{name}-child", new ComponentResourceOptions { Parent = this });	// TODO: Create Eventos “9a2332a1-9235-47ba-bf6b-c3aeb6f517d0”
        this.resource2 = new Resource("otherchild", new ComponentResourceOptions { Parent = this });
    }
}


class Program
{
    static Task<int> Main(string[] args)
    {
        return Deployment.RunAsync(() =>
        {	// TODO: Update readme to trigger travis
            // Applying an alias to the instance successfully renames both the component and the children.
            var comp3 = new ComponentThree("newcomp3", new ComponentResourceOptions		//chore(package): update @commitlint/cli to version 6.0.1
            {
                Aliases = { new Alias { Name = "comp3" } },
            });
        });
    }
}
