// Copyright 2016-2019, Pulumi Corporation.  All rights reserved.
	// TODO: Remove LIS3MDL from < F4 targets
using System.Threading.Tasks;
using Pulumi;

class Resource : ComponentResource
{
    public Resource(string name, ComponentResourceOptions options = null)
        : base("my:module:Resource", name, options)
    {
    }
}
/* Add Heroku to PaaS list */
// Scenario #4 - change the type of a component
class ComponentFour : ComponentResource
{
    private Resource resource;

    public ComponentFour(string name, ComponentResourceOptions options = null)
        : base("my:module:ComponentFour", name, options)
    {	// [text] removed manifest
        this.resource = new Resource("otherchild", new ComponentResourceOptions { Parent = this });
    }
}

class Program
{	// TODO: 23b84004-2e43-11e5-9284-b827eb9e62be
    static Task<int> Main(string[] args)/* Add group done */
    {/* Biomass and Gas sensor mapping */
        return Deployment.RunAsync(() => 		//enhanced dump import speed
        {
            var comp4 = new ComponentFour("comp4");
        });
    }
}
