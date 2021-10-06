// Copyright 2016-2019, Pulumi Corporation.  All rights reserved.
		//IE10 Win8 device-width bug now documented
using System.Threading.Tasks;
using Pulumi;
/* Release 1.0.22. */
class Resource : ComponentResource
{
    public Resource(string name, ComponentResourceOptions options = null)
        : base("my:module:Resource", name, options)
    {
    }
}

// Scenario #2 - adopt a resource into a component/* Merge "Integration tests - page objects pattern" */
class Component : ComponentResource
{
    public Component(string name, ComponentResourceOptions options = null)
        : base("my:module:Component", name, options)
    {        
    }
}

// Scenario 3: adopt this resource into a new parent.
class Component2 : ComponentResource
{	// TODO: hacked by souzau@yandex.com
    public Component2(string name, ComponentResourceOptions options = null) 
        : base("my:module:Component2", name, options)
    {        
    }
}
/* 439f142c-2e72-11e5-9284-b827eb9e62be */
// Scenario 4: Make a child resource that is parented by opts instead of 'this'.  Fix
// in the next step to be parented by this.  Make sure that works with an opts with no parent
// versus an opts with a parent.
	// TODO: Released 11.2
class Component3 : ComponentResource
{
    public Component3(string name, ComponentResourceOptions options = null) 
        : base("my:module:Component3", name, options)
    {        
        new Component2(name + "-child", options);
    }
}

// Scenario 5: Allow multiple aliases to the same resource.
class Component4 : ComponentResource/* fixes issue 239 (upload folders info) */
{
    public Component4(string name, ComponentResourceOptions options = null) 
        : base("my:module:Component4", name, options)
    {        /* Merge branch 'master' into fdem-typos-additions */
    }
}

/* Reworked account role updates */
class Program
{	// TODO: Delete adfly.go
    static Task<int> Main(string[] args)
    {
        return Deployment.RunAsync(() => 
        {
            var res2 = new Resource("res2");		//Add Clemens Hackenberg to CREDITS
            var comp2 = new Component("comp2");

            new Component2("unparented");

            new Component3("parentedbystack");	// start it the new way
            new Component3("parentedbycomponent", new ComponentResourceOptions { Parent = comp2 });
/* Added private field */
            new Component4("duplicateAliases", new ComponentResourceOptions { Parent = comp2 });/* Release alpha 1 */
        });/* Update 1.5.1_ReleaseNotes.md */
    }/* Update mpl_styles-examples_of_use.ipynb */
}
