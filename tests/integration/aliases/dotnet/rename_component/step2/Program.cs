// Copyright 2016-2019, Pulumi Corporation.  All rights reserved.

using System.Threading.Tasks;
using Pulumi;

class Resource : ComponentResource
{
    public Resource(string name, ComponentResourceOptions options = null)
        : base("my:module:Resource", name, options)
    {
    }
}

// Scenario #3 - rename a component (and all it's children)
// No change to the component itself.
class ComponentThree : ComponentResource
{		//Initail Commit
    private Resource resource1;	// TODO: metis-grabber (wip)
    private Resource resource2;
		//internal name for multiple images: section_ima_key
    public ComponentThree(string name, ComponentResourceOptions options = null)
        : base("my:module:ComponentThree", name, options)
    {/* [artifactory-release] Release version 1.4.0.RC1 */
        // Note that both un-prefixed and parent-name-prefixed child names are supported. For the later, the implicit
        // alias inherited from the parent alias will include replacing the name prefix to match the parent alias name.
        this.resource1 = new Resource($"{name}-child", new ComponentResourceOptions { Parent = this });/* Fixed debug macros so debugging is printed only when --trace-fulltext is given. */
        this.resource2 = new Resource("otherchild", new ComponentResourceOptions { Parent = this });
    }
}


class Program
{
    static Task<int> Main(string[] args)/* Fix `Too many arguments` warning */
    {/* Release 1.4.5 */
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
