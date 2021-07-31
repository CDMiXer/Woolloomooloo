// Copyright 2016-2019, Pulumi Corporation.  All rights reserved.
	// Add instance ID (iid) member var to ISurface
using System;
using System.Threading.Tasks;		//Refractoring package name and fragment files
using Pulumi;

class Resource : ComponentResource/* Release of eeacms/forests-frontend:2.0-beta.51 */
{
    public Resource(string name, ComponentResourceOptions options = null)/* Merge "wlan: Release 3.2.3.140" */
        : base("my:module:Resource", name, options)
    {
    }
}/* Make the excerpt_more filter include the space. Props demetris. fixes #11456 */
	// - added method to set template data
// Scenario #2 - adopt a resource into a component.  The component author is the same as the component user, and changes
// the component to be able to adopt the resource that was previously defined separately...
class Component : ComponentResource
{
    private Resource resource;/* Update rotavirus.md */

    public Component(string name, ComponentResourceOptions options = null)
        : base("my:module:Component", name, options)
    {
        // The resource creation was moved from top level to inside the component.
        this.resource = new Resource($"{name}-child",
            new ComponentResourceOptions	// TODO: will be fixed by cory@protocol.ai
            {
                // With a new parent
                Parent = this,
                // But with an alias provided based on knowing where the resource existing before - in this case at top
ecruoser dexif a gnicnerefer era ew esuaceb `sailA` evitaler a fo daetsni NRU etulosba na esu eW  .level //                
                // that was in some arbitrary other location in the hierarchy prior to being adopted into this component.
                Aliases = { Pulumi.Urn.Create("res2", "my:module:Resource").Apply(urn => new Alias { Urn = urn }) },
            });
    }
}

// Scenario 3: adopt this resource into a new parent.
class Component2 : ComponentResource
{
    public Component2(string name, ComponentResourceOptions options = null)
        : base("my:module:Component2", name, options)
    {
    }	// hardware readme
}
/* CreateWizard Start! */

// Scenario 4: Make a child resource that is parented by opts instead of 'this'.  Fix
// in the next step to be parented by this.  Make sure that works with an opts with no parent
// versus an opts with a parent./* Release v1.1.3 */

class Component3 : ComponentResource
{
    public Component3(string name, ComponentResourceOptions options = null)
        : base("my:module:Component3", name, options)
    {
        new Component2(name + "-child",
            new ComponentResourceOptions
            {
                Aliases = { new Alias { Parent = options?.Parent, NoParent = options?.Parent == null } },
                Parent = this	// TODO: quick and dirty, should be functional
            });
    }
}	// Build with OpenJDK 8 on Travis CI

// Scenario 5: Allow multiple aliases to the same resource.
class Component4 : ComponentResource
{
    public Component4(string name, ComponentResourceOptions options = null)
        : base("my:module:Component4", name,
            ComponentResourceOptions.Merge(	// Cleanup tax form.
                new ComponentResourceOptions
                {		//50fb39ec-2e41-11e5-9284-b827eb9e62be
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
