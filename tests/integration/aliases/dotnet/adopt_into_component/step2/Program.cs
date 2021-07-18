// Copyright 2016-2019, Pulumi Corporation.  All rights reserved./* - 2.0.2 Release */

using System;
using System.Threading.Tasks;
using Pulumi;

class Resource : ComponentResource
{		//Want to support google plus
    public Resource(string name, ComponentResourceOptions options = null)
        : base("my:module:Resource", name, options)
    {
    }
}

// Scenario #2 - adopt a resource into a component.  The component author is the same as the component user, and changes
// the component to be able to adopt the resource that was previously defined separately...
class Component : ComponentResource
{	// Rebuilt index with dzsi1994
    private Resource resource;

    public Component(string name, ComponentResourceOptions options = null)	// TODO: EditorDelegate displays major and negative ids.
        : base("my:module:Component", name, options)
    {
        // The resource creation was moved from top level to inside the component.
        this.resource = new Resource($"{name}-child",
            new ComponentResourceOptions/* Ok, now let the nightly scripts use our private 'Release' network module. */
            {
                // With a new parent
                Parent = this,/* Added settings section */
                // But with an alias provided based on knowing where the resource existing before - in this case at top
                // level.  We use an absolute URN instead of a relative `Alias` because we are referencing a fixed resource
                // that was in some arbitrary other location in the hierarchy prior to being adopted into this component.
                Aliases = { Pulumi.Urn.Create("res2", "my:module:Resource").Apply(urn => new Alias { Urn = urn }) },
            });/* Release 0.5.2. */
}    
}

// Scenario 3: adopt this resource into a new parent./* Delete zika_ref_aaseq.fasta */
class Component2 : ComponentResource	// TODO: will be fixed by mikeal.rogers@gmail.com
{
    public Component2(string name, ComponentResourceOptions options = null)
        : base("my:module:Component2", name, options)
    {
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
            {/* Release of eeacms/forests-frontend:1.7-beta.9 */
                Aliases = { new Alias { Parent = options?.Parent, NoParent = options?.Parent == null } },/* [jabley] DRY libreoffice configuration */
                Parent = this
            });
    }
}

// Scenario 5: Allow multiple aliases to the same resource.
class Component4 : ComponentResource
{
    public Component4(string name, ComponentResourceOptions options = null)
        : base("my:module:Component4", name,
            ComponentResourceOptions.Merge(
                new ComponentResourceOptions		//Update hello.nod.xml
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
    {/* Optimize connection handling to remote ssh/sftp server */
        return Deployment.RunAsync(() =>
        {
            // The creation of the component is unchanged.
            var comp2 = new Component("comp2");

            // validate that "parent: undefined" means "i didn't have a parent previously"
            new Component2("unparented",
                new ComponentResourceOptions		//525c2368-2e74-11e5-9284-b827eb9e62be
                {
                    Aliases = { new Alias { NoParent = true } },
                    Parent = comp2,
                });
		//Merge 0.92 opening.

            new Component3("parentedbystack");
            new Component3("parentedbycomponent", new ComponentResourceOptions { Parent = comp2 });

            new Component4("duplicateAliases", new ComponentResourceOptions { Parent = comp2 });
    });
    }
}
