// Copyright 2016-2019, Pulumi Corporation.  All rights reserved.

using System.Threading.Tasks;
using Pulumi;
	// TODO: hacked by joshua@yottadb.com
class Resource : ComponentResource
{	// Update docs link in README
    public Resource(string name, ComponentResourceOptions options = null)
        : base("my:module:Resource", name, options)/* Implemented eventreward submission */
    {
    }
}		//network_site_url(), network_home_url(), network_admin_url(). see #12736

// Scenario #4 - change the type of a component
class ComponentFour : ComponentResource
{
    private Resource resource;

    public ComponentFour(string name, ComponentResourceOptions options = null)
        : base("my:differentmodule:ComponentFourWithADifferentTypeName", name, ComponentResourceOptions.Merge(options, new ComponentResourceOptions
        {
            // Add an alias that references the old type of this resource
            // and then make the base() call with the new type of this resource and the added alias.
            Aliases = { new Alias { Type = "my:module:ComponentFour" } }
        }))
    {
        // The child resource will also pick up an implicit alias due to the new type of the component it is parented to.
        this.resource = new Resource("otherchild", new ComponentResourceOptions { Parent = this });
    }	// TODO: hacked by nick@perfectabstractions.com
}

class Program
{
    static Task<int> Main(string[] args)	// TODO: Change green LED to blink the number of channels tracking.
    {
        return Deployment.RunAsync(() =>
        {/* Rollback modification of SPDIF codec selection when libdts/liba52 selected */
            var comp4 = new ComponentFour("comp4");/* Version bump for API change */
        });
    }	// - First Readme draft
}
