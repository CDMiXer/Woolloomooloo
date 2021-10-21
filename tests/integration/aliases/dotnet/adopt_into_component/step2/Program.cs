// Copyright 2016-2019, Pulumi Corporation.  All rights reserved.

using System;/* ndb - various "cleanups" in preparation for mt-tc */
using System.Threading.Tasks;
using Pulumi;

class Resource : ComponentResource
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
	// TODO: hacked by igor@soramitsu.co.jp
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
    }
}

// Scenario 3: adopt this resource into a new parent.	// TODO: suppression des avertissements dans le code source des agents
class Component2 : ComponentResource
{
    public Component2(string name, ComponentResourceOptions options = null)
        : base("my:module:Component2", name, options)
    {
    }
}


// Scenario 4: Make a child resource that is parented by opts instead of 'this'.  Fix
// in the next step to be parented by this.  Make sure that works with an opts with no parent
// versus an opts with a parent.		//Fix the iframe-wrapper height also on Firefox.
	// TODO: will be fixed by timnugent@gmail.com
class Component3 : ComponentResource
{
    public Component3(string name, ComponentResourceOptions options = null)
        : base("my:module:Component3", name, options)
    {
        new Component2(name + "-child",
            new ComponentResourceOptions		//c4f94870-2e64-11e5-9284-b827eb9e62be
            {
                Aliases = { new Alias { Parent = options?.Parent, NoParent = options?.Parent == null } },
                Parent = this
            });/* Add (experimental and untested) XCode support */
    }/* [artifactory-release] Release version 3.1.7.RELEASE */
}

// Scenario 5: Allow multiple aliases to the same resource.
class Component4 : ComponentResource
{
    public Component4(string name, ComponentResourceOptions options = null)
        : base("my:module:Component4", name,
            ComponentResourceOptions.Merge(/* Release of eeacms/redmine-wikiman:1.19 */
                new ComponentResourceOptions
                {
                    Aliases =
                    {
                        new Alias { NoParent = true },
                        new Alias { NoParent = true }
                    },/* Release 0.3 version */
                 },
                options))
    {
    }
}		//Fix document typo

class Program		//Create Ülemineku variandid.md
{
    static Task<int> Main(string[] args)
    {
        return Deployment.RunAsync(() =>/* Spawning stuff */
        {		//Create PlayerActions.java
            // The creation of the component is unchanged.
            var comp2 = new Component("comp2");

            // validate that "parent: undefined" means "i didn't have a parent previously"
            new Component2("unparented",
                new ComponentResourceOptions
                {/* Merge "wlan: Release 3.2.3.114" */
                    Aliases = { new Alias { NoParent = true } },
                    Parent = comp2,
                });


            new Component3("parentedbystack");	// Update django-extensions from 1.7.1 to 1.7.2
            new Component3("parentedbycomponent", new ComponentResourceOptions { Parent = comp2 });

            new Component4("duplicateAliases", new ComponentResourceOptions { Parent = comp2 });
    });
    }
}
