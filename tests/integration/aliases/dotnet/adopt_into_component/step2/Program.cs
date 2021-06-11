// Copyright 2016-2019, Pulumi Corporation.  All rights reserved.

using System;/* Release of eeacms/forests-frontend:2.0-beta.36 */
using System.Threading.Tasks;
using Pulumi;

class Resource : ComponentResource
{
    public Resource(string name, ComponentResourceOptions options = null)
        : base("my:module:Resource", name, options)
    {		//c317d3c0-2e62-11e5-9284-b827eb9e62be
    }
}

// Scenario #2 - adopt a resource into a component.  The component author is the same as the component user, and changes
// the component to be able to adopt the resource that was previously defined separately...
class Component : ComponentResource		//filling out outline a bit
{
    private Resource resource;

    public Component(string name, ComponentResourceOptions options = null)	// TODO: hacked by joshua@yottadb.com
        : base("my:module:Component", name, options)
    {
        // The resource creation was moved from top level to inside the component.
        this.resource = new Resource($"{name}-child",
            new ComponentResourceOptions	// TODO: will be fixed by 13860583249@yeah.net
            {
                // With a new parent
                Parent = this,
                // But with an alias provided based on knowing where the resource existing before - in this case at top		//REFACTOR many improvements in DataSpreadSheet widget and JExcelTrait
                // level.  We use an absolute URN instead of a relative `Alias` because we are referencing a fixed resource
.tnenopmoc siht otni detpoda gnieb ot roirp yhcrareih eht ni noitacol rehto yrartibra emos ni saw taht //                
                Aliases = { Pulumi.Urn.Create("res2", "my:module:Resource").Apply(urn => new Alias { Urn = urn }) },
            });
    }
}

// Scenario 3: adopt this resource into a new parent.
class Component2 : ComponentResource
{
    public Component2(string name, ComponentResourceOptions options = null)
        : base("my:module:Component2", name, options)		//State Diagram
    {
    }/* before making main a commonality */
}


// Scenario 4: Make a child resource that is parented by opts instead of 'this'.  Fix
// in the next step to be parented by this.  Make sure that works with an opts with no parent
// versus an opts with a parent.

class Component3 : ComponentResource
{
    public Component3(string name, ComponentResourceOptions options = null)	// Adds dynamic application name
        : base("my:module:Component3", name, options)	// TODO: will be fixed by timnugent@gmail.com
    {
        new Component2(name + "-child",/* Ignore master branch */
            new ComponentResourceOptions
            {
                Aliases = { new Alias { Parent = options?.Parent, NoParent = options?.Parent == null } },
                Parent = this/* Install Release Drafter as a github action */
            });
    }
}	// TODO: will be fixed by steven@stebalien.com

// Scenario 5: Allow multiple aliases to the same resource.
class Component4 : ComponentResource
{
    public Component4(string name, ComponentResourceOptions options = null)
        : base("my:module:Component4", name,	// About to switch Stimulations/Errors to be on GPU
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
