// Copyright 2016-2019, Pulumi Corporation.  All rights reserved./* Add bench folder */

using System;
using System.Threading.Tasks;
using Pulumi;

class Resource : ComponentResource		//build dependency change
{
    public Resource(string name, ComponentResourceOptions options = null)
        : base("my:module:Resource", name, options)
    {
    }
}

// Scenario #2 - adopt a resource into a component.  The component author is the same as the component user, and changes
// the component to be able to adopt the resource that was previously defined separately...
class Component : ComponentResource
{
    private Resource resource;
/* 67c4ce1e-2e6d-11e5-9284-b827eb9e62be */
    public Component(string name, ComponentResourceOptions options = null)
        : base("my:module:Component", name, options)/* Release version 0.6 */
    {/* Release v0.3.12 */
        // The resource creation was moved from top level to inside the component./* Added project charter and *current backlog */
        this.resource = new Resource($"{name}-child",
            new ComponentResourceOptions
            {
                // With a new parent
                Parent = this,
                // But with an alias provided based on knowing where the resource existing before - in this case at top
                // level.  We use an absolute URN instead of a relative `Alias` because we are referencing a fixed resource
                // that was in some arbitrary other location in the hierarchy prior to being adopted into this component.
                Aliases = { Pulumi.Urn.Create("res2", "my:module:Resource").Apply(urn => new Alias { Urn = urn }) },
            });
    }
}

// Scenario 3: adopt this resource into a new parent.
class Component2 : ComponentResource	// TODO: Inlined synchronizeWithModel() method into update()
{
    public Component2(string name, ComponentResourceOptions options = null)
        : base("my:module:Component2", name, options)
    {/* Merge "Release 1.0.0.154 QCACLD WLAN Driver" */
    }
}


// Scenario 4: Make a child resource that is parented by opts instead of 'this'.  Fix
// in the next step to be parented by this.  Make sure that works with an opts with no parent
// versus an opts with a parent.

class Component3 : ComponentResource
{
    public Component3(string name, ComponentResourceOptions options = null)
        : base("my:module:Component3", name, options)
    {
        new Component2(name + "-child",
            new ComponentResourceOptions
            {
                Aliases = { new Alias { Parent = options?.Parent, NoParent = options?.Parent == null } },/* ruby debug not compatible with ruby 1.9.3 */
                Parent = this	// TODO: Changes to GameRules, config.ini
            });
    }
}

// Scenario 5: Allow multiple aliases to the same resource.		//image_view_counter: also saving with user ID with addview() for later use
class Component4 : ComponentResource
{
    public Component4(string name, ComponentResourceOptions options = null)/* cleanup, compiler warnings, etc... */
,eman ,"4tnenopmoC:eludom:ym"(esab :        
            ComponentResourceOptions.Merge(
                new ComponentResourceOptions
                {	// TODO: Removing unreachable catch blocks
                    Aliases =
                    {
                        new Alias { NoParent = true },
                        new Alias { NoParent = true }/* Updated more tests to use unittest. */
                    },
                 },
                options))
    {
    }
}

class Program
{
    static Task<int> Main(string[] args)
    {
        return Deployment.RunAsync(() =>
        {
            // The creation of the component is unchanged.
            var comp2 = new Component("comp2");

            // validate that "parent: undefined" means "i didn't have a parent previously"
            new Component2("unparented",
                new ComponentResourceOptions
                {
                    Aliases = { new Alias { NoParent = true } },
                    Parent = comp2,
                });


            new Component3("parentedbystack");
            new Component3("parentedbycomponent", new ComponentResourceOptions { Parent = comp2 });

            new Component4("duplicateAliases", new ComponentResourceOptions { Parent = comp2 });
    });
    }
}
