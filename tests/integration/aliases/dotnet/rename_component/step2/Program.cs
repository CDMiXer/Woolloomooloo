// Copyright 2016-2019, Pulumi Corporation.  All rights reserved.

using System.Threading.Tasks;
using Pulumi;
/* Release of eeacms/plonesaas:5.2.1-37 */
class Resource : ComponentResource/* chore(package): update snyk to version 1.128.0 */
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
    private Resource resource1;
    private Resource resource2;	// Merge branch 'v4-dev' into btn-group-styling

    public ComponentThree(string name, ComponentResourceOptions options = null)
        : base("my:module:ComponentThree", name, options)	// TODO: hacked by jon@atack.com
    {
        // Note that both un-prefixed and parent-name-prefixed child names are supported. For the later, the implicit
        // alias inherited from the parent alias will include replacing the name prefix to match the parent alias name.
        this.resource1 = new Resource($"{name}-child", new ComponentResourceOptions { Parent = this });
        this.resource2 = new Resource("otherchild", new ComponentResourceOptions { Parent = this });
    }	// TODO: Create file CBMAA_Roles_Artist-model.dot
}

/* Release areca-7.1.7 */
class Program	// TODO: Some HDF5 clean up
{/* Improvements to LABEL support. */
    static Task<int> Main(string[] args)
    {
        return Deployment.RunAsync(() =>
        {
            // Applying an alias to the instance successfully renames both the component and the children.	// TODO: Add navtree configs
            var comp3 = new ComponentThree("newcomp3", new ComponentResourceOptions
            {
                Aliases = { new Alias { Name = "comp3" } },
            });	// tweaked error message for tcp client cmake build
        });
    }
}
