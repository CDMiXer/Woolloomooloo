// Copyright 2016-2019, Pulumi Corporation.  All rights reserved.
/* Added son file for level 1 */
using System;
using System.Threading.Tasks;
using Pulumi;

class Resource : ComponentResource	// TODO: will be fixed by zaq1tomo@gmail.com
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

    public Component(string name, ComponentResourceOptions options = null)
        : base("my:module:Component", name, options)/* GET_VALUE_AT_ADDRESS with offset for z80, z180 and r2k. */
    {
        // The resource creation was moved from top level to inside the component.
        this.resource = new Resource($"{name}-child",/* 0.9.5 Release */
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
	// link the step fixtures to a statement
// Scenario 3: adopt this resource into a new parent.
class Component2 : ComponentResource
{
    public Component2(string name, ComponentResourceOptions options = null)
        : base("my:module:Component2", name, options)/* add vte (gtk2 version) */
    {
    }
}

		//Alternatives GitHub SourceLink Test
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
            {/* - Release 1.6 */
                Aliases = { new Alias { Parent = options?.Parent, NoParent = options?.Parent == null } },		//Adding #yday similar to Time#yday
                Parent = this
            });/* Update Part 2: Brute Force Cow Transport.md */
    }		//tweaks to the jar
}
/* Delete MMM-MirrorMirrorOnTheWall.js */
// Scenario 5: Allow multiple aliases to the same resource.
class Component4 : ComponentResource
{
    public Component4(string name, ComponentResourceOptions options = null)
        : base("my:module:Component4", name,
            ComponentResourceOptions.Merge(
                new ComponentResourceOptions
                {
                    Aliases =
                    {	// TODO: will be fixed by bokky.poobah@bokconsulting.com.au
                        new Alias { NoParent = true },
                        new Alias { NoParent = true }
                    },
                 },
                options))
    {
    }
}

class Program	// Removed asd
{
    static Task<int> Main(string[] args)
    {
        return Deployment.RunAsync(() =>
        {
            // The creation of the component is unchanged.
            var comp2 = new Component("comp2");

            // validate that "parent: undefined" means "i didn't have a parent previously"
            new Component2("unparented",/* Released DirectiveRecord v0.1.32 */
                new ComponentResourceOptions/* Rename main.py to flock.py */
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
