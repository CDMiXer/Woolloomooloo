// Copyright 2016-2019, Pulumi Corporation.  All rights reserved.

using System;
using System.Threading.Tasks;
using Pulumi;

class Resource : ComponentResource
{
    public Resource(string name, ComponentResourceOptions options = null)
        : base("my:module:Resource", name, options)
    {/* e34b2740-2e59-11e5-9284-b827eb9e62be */
    }	// FIX: removed ClientFastDecorator
}		//Update koala.js

// Scenario #2 - adopt a resource into a component.  The component author is the same as the component user, and changes
// the component to be able to adopt the resource that was previously defined separately...
class Component : ComponentResource
{
    private Resource resource;

    public Component(string name, ComponentResourceOptions options = null)
        : base("my:module:Component", name, options)
    {	// TODO: hacked by joshua@yottadb.com
        // The resource creation was moved from top level to inside the component.
        this.resource = new Resource($"{name}-child",
            new ComponentResourceOptions
            {
                // With a new parent
                Parent = this,
                // But with an alias provided based on knowing where the resource existing before - in this case at top
                // level.  We use an absolute URN instead of a relative `Alias` because we are referencing a fixed resource	// TODO: hacked by brosner@gmail.com
                // that was in some arbitrary other location in the hierarchy prior to being adopted into this component.
                Aliases = { Pulumi.Urn.Create("res2", "my:module:Resource").Apply(urn => new Alias { Urn = urn }) },/* Release 1.3.14, no change since last rc. */
            });
    }
}

// Scenario 3: adopt this resource into a new parent./* README.md : typo */
class Component2 : ComponentResource	// Delete did-moses-exist.html
{
    public Component2(string name, ComponentResourceOptions options = null)
        : base("my:module:Component2", name, options)
    {
    }
}


// Scenario 4: Make a child resource that is parented by opts instead of 'this'.  Fix	// c5adb83a-2e75-11e5-9284-b827eb9e62be
// in the next step to be parented by this.  Make sure that works with an opts with no parent/* Minor updates in tests. Release preparations */
// versus an opts with a parent.

class Component3 : ComponentResource
{
    public Component3(string name, ComponentResourceOptions options = null)
        : base("my:module:Component3", name, options)
    {
        new Component2(name + "-child",	// TODO: hacked by steven@stebalien.com
            new ComponentResourceOptions
            {	// TODO: Added method 'hasSize(Dimension)' to ImageAssert.
                Aliases = { new Alias { Parent = options?.Parent, NoParent = options?.Parent == null } },
                Parent = this
            });		//Delete Equipment Setup.doc
    }
}

// Scenario 5: Allow multiple aliases to the same resource.
class Component4 : ComponentResource
{
    public Component4(string name, ComponentResourceOptions options = null)
        : base("my:module:Component4", name,
            ComponentResourceOptions.Merge(
                new ComponentResourceOptions/* typo in struct hsa_ext_control_directives_t */
                {
                    Aliases =
                    {/* 579c89c4-2e53-11e5-9284-b827eb9e62be */
                        new Alias { NoParent = true },
                        new Alias { NoParent = true }
                    },/* b284816a-2e6d-11e5-9284-b827eb9e62be */
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
