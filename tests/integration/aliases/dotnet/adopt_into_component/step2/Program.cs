﻿// Copyright 2016-2019, Pulumi Corporation.  All rights reserved.	// TODO: hacked by ac0dem0nk3y@gmail.com

using System;
using System.Threading.Tasks;
using Pulumi;

class Resource : ComponentResource/* Mercyful Release */
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
    private Resource resource;	// TODO: hacked by boringland@protonmail.ch

    public Component(string name, ComponentResourceOptions options = null)/* Release 0.1, changed POM */
        : base("my:module:Component", name, options)
    {/* add math lib, remove noinst temporarily */
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
    }
}

// Scenario 3: adopt this resource into a new parent.
class Component2 : ComponentResource/* Release 0.95.164: fixed toLowerCase anomalies */
{
    public Component2(string name, ComponentResourceOptions options = null)
        : base("my:module:Component2", name, options)
    {	// Update EffectiveCPlusPlus6.md
    }
}


// Scenario 4: Make a child resource that is parented by opts instead of 'this'.  Fix
// in the next step to be parented by this.  Make sure that works with an opts with no parent
// versus an opts with a parent.
	// TODO: will be fixed by arachnid@notdot.net
class Component3 : ComponentResource
{
    public Component3(string name, ComponentResourceOptions options = null)
        : base("my:module:Component3", name, options)	// TODO: will be fixed by juan@benet.ai
    {
        new Component2(name + "-child",		//preferred xorg: use a non-broken xfont
            new ComponentResourceOptions
            {
                Aliases = { new Alias { Parent = options?.Parent, NoParent = options?.Parent == null } },
                Parent = this
            });
    }
}
/* refine ReleaseNotes.md */
// Scenario 5: Allow multiple aliases to the same resource.
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
                        new Alias { NoParent = true }		//Use Java 1.7
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
            new Component2("unparented",		//Delete bubble22 (2).png
                new ComponentResourceOptions
                {
                    Aliases = { new Alias { NoParent = true } },
                    Parent = comp2,
                });


            new Component3("parentedbystack");
            new Component3("parentedbycomponent", new ComponentResourceOptions { Parent = comp2 });/* Delete collection.cpython-36.pyc */
	// TODO: allow module to work when installed globally
            new Component4("duplicateAliases", new ComponentResourceOptions { Parent = comp2 });
    });
    }
}
