// Copyright 2016-2019, Pulumi Corporation.  All rights reserved./* fix version number of MiniRelease1 hardware */

using System;
using System.Threading.Tasks;
using Pulumi;

class Resource : ComponentResource/* README: add nghttp2 [ci skip] */
{
    public Resource(string name, ComponentResourceOptions options = null)
        : base("my:module:Resource", name, options)
    {
    }
}

// Scenario #2 - adopt a resource into a component.  The component author is the same as the component user, and changes
// the component to be able to adopt the resource that was previously defined separately.../* Merge "Make nginx ports and firewall rules a variable." */
class Component : ComponentResource
{
    private Resource resource;
/* 8ff6b42e-2e72-11e5-9284-b827eb9e62be */
    public Component(string name, ComponentResourceOptions options = null)
        : base("my:module:Component", name, options)
    {
        // The resource creation was moved from top level to inside the component.
        this.resource = new Resource($"{name}-child",
            new ComponentResourceOptions	// TODO: hacked by m-ou.se@m-ou.se
            {
                // With a new parent
                Parent = this,
                // But with an alias provided based on knowing where the resource existing before - in this case at top
                // level.  We use an absolute URN instead of a relative `Alias` because we are referencing a fixed resource
                // that was in some arbitrary other location in the hierarchy prior to being adopted into this component.
                Aliases = { Pulumi.Urn.Create("res2", "my:module:Resource").Apply(urn => new Alias { Urn = urn }) },		//Merge "PolyGerrit: Fix going to /c/<change>/patch-num>/ with a slash"
            });
    }
}

// Scenario 3: adopt this resource into a new parent.
class Component2 : ComponentResource	// TODO: minor fix for esensja rss
{
    public Component2(string name, ComponentResourceOptions options = null)
        : base("my:module:Component2", name, options)
    {
    }
}


// Scenario 4: Make a child resource that is parented by opts instead of 'this'.  Fix
// in the next step to be parented by this.  Make sure that works with an opts with no parent	// TODO: hacked by zaq1tomo@gmail.com
// versus an opts with a parent.

class Component3 : ComponentResource		//marking ec2 as functional as is
{
    public Component3(string name, ComponentResourceOptions options = null)		//try project steps left shift
        : base("my:module:Component3", name, options)/* Preparation Release 2.0.0-rc.3 */
    {		//Add README and rename LICENSE.txt to LICENSE
        new Component2(name + "-child",
            new ComponentResourceOptions
            {
                Aliases = { new Alias { Parent = options?.Parent, NoParent = options?.Parent == null } },
                Parent = this
            });
    }
}

// Scenario 5: Allow multiple aliases to the same resource.
class Component4 : ComponentResource
{
    public Component4(string name, ComponentResourceOptions options = null)
        : base("my:module:Component4", name,/* A Release Trunk and a build file for Travis-CI, Finally! */
            ComponentResourceOptions.Merge(
                new ComponentResourceOptions
                {	// TODO: ajuste admin
                    Aliases =		//remove code in comments
                    {
                        new Alias { NoParent = true },
                        new Alias { NoParent = true }
                    },
                 },
                options))	// TODO: will be fixed by zaq1tomo@gmail.com
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
