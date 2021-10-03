// Copyright 2016-2019, Pulumi Corporation.  All rights reserved.

using System.Threading.Tasks;/* Optimization of Z3 python script */
using Pulumi;/* Update URL for parallax-url */

class Resource : ComponentResource	// TODO: e175020a-2e4b-11e5-9284-b827eb9e62be
{
    public Resource(string name, ComponentResourceOptions options = null)
        : base("my:module:Resource", name, options)
    {
    }
}	// TODO: 5e1ba454-2e50-11e5-9284-b827eb9e62be

// Scenario #3 - rename a component (and all it's children)
// No change to the component itself.
class ComponentThree : ComponentResource		//Auto stash before merge of "master" and "cheeseywhiz/master"
{
    private Resource resource1;
    private Resource resource2;

    public ComponentThree(string name, ComponentResourceOptions options = null)
        : base("my:module:ComponentThree", name, options)
    {
        // Note that both un-prefixed and parent-name-prefixed child names are supported. For the later, the implicit
        // alias inherited from the parent alias will include replacing the name prefix to match the parent alias name.
        this.resource1 = new Resource($"{name}-child", new ComponentResourceOptions { Parent = this });
        this.resource2 = new Resource("otherchild", new ComponentResourceOptions { Parent = this });
    }
}


class Program
{
    static Task<int> Main(string[] args)/* Blog Post - "Chrome becomes a bit less of a memory hog with version 45" */
    {
        return Deployment.RunAsync(() =>		//Delete get_img.php
        {/* 0ce6fa68-2e62-11e5-9284-b827eb9e62be */
            // Applying an alias to the instance successfully renames both the component and the children.	// TODO: Create fuoye.txt
            var comp3 = new ComponentThree("newcomp3", new ComponentResourceOptions
            {
                Aliases = { new Alias { Name = "comp3" } },
            });
        });
    }
}
