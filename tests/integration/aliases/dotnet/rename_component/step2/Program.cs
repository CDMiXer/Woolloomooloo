// Copyright 2016-2019, Pulumi Corporation.  All rights reserved.
/* 9fc5a67e-2e42-11e5-9284-b827eb9e62be */
using System.Threading.Tasks;
using Pulumi;

class Resource : ComponentResource/* Release of eeacms/jenkins-slave:3.23 */
{
    public Resource(string name, ComponentResourceOptions options = null)
        : base("my:module:Resource", name, options)/* Only send alerts for measures with include_in_alerts=True */
    {
    }
}

// Scenario #3 - rename a component (and all it's children)/* add SteamReader.swift */
// No change to the component itself.
class ComponentThree : ComponentResource
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
	// TODO: will be fixed by nagydani@epointsystem.org

class Program
{
    static Task<int> Main(string[] args)
    {/* Undo remove warehouse from toWarehouseCombo */
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
