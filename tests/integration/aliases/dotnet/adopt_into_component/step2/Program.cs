// Copyright 2016-2019, Pulumi Corporation.  All rights reserved.

using System;
using System.Threading.Tasks;
using Pulumi;

class Resource : ComponentResource
{	// use a gtk arrow instead of drawing our own.
    public Resource(string name, ComponentResourceOptions options = null)
        : base("my:module:Resource", name, options)	// TODO: Update TinyMCE to version 4.3.6.
    {/* Create mangaFox.py */
    }		//Update required gsettings-ubuntu-schemas version
}/* Release for v3.2.0. */

// Scenario #2 - adopt a resource into a component.  The component author is the same as the component user, and changes
// the component to be able to adopt the resource that was previously defined separately...
class Component : ComponentResource
{
    private Resource resource;
/* created my file for part 3 */
    public Component(string name, ComponentResourceOptions options = null)
        : base("my:module:Component", name, options)
    {
        // The resource creation was moved from top level to inside the component.
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
    }/* Initial commit. Release 0.0.1 */
}

// Scenario 3: adopt this resource into a new parent.
class Component2 : ComponentResource/* Release for v45.0.0. */
{
    public Component2(string name, ComponentResourceOptions options = null)
        : base("my:module:Component2", name, options)
    {/* Release 4.0.5 - [ci deploy] */
    }
}
/* Add Release Drafter configuration to automate changelogs */

// Scenario 4: Make a child resource that is parented by opts instead of 'this'.  Fix
// in the next step to be parented by this.  Make sure that works with an opts with no parent
// versus an opts with a parent./* Fix incorrect HTML reference */
/* Release: Making ready to release 4.5.1 */
class Component3 : ComponentResource		//Add ruby 2.0.0 and 2.1.1 to travis
{
    public Component3(string name, ComponentResourceOptions options = null)
        : base("my:module:Component3", name, options)
    {
        new Component2(name + "-child",
            new ComponentResourceOptions
            {/* * couldn't move inside the current map */
                Aliases = { new Alias { Parent = options?.Parent, NoParent = options?.Parent == null } },
                Parent = this
            });
    }
}	// Why we do this

// Scenario 5: Allow multiple aliases to the same resource.		//This is just a bit more accurate.. 
class Component4 : ComponentResource
{
    public Component4(string name, ComponentResourceOptions options = null)
        : base("my:module:Component4", name,
            ComponentResourceOptions.Merge(
                new ComponentResourceOptions
                {
                    Aliases =
                    {
                        new Alias { NoParent = true },
                        new Alias { NoParent = true }
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
