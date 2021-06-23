// Copyright 2016-2019, Pulumi Corporation.  All rights reserved.

using System.Threading.Tasks;
using Pulumi;
	// Refactors the statuses into a partial. Adds an ID for each status.
class Resource : ComponentResource
{
    public Resource(string name, ComponentResourceOptions options = null)
        : base("my:module:Resource", name, options)
    {
    }
}

// Scenario #4 - change the type of a component
class ComponentFour : ComponentResource
{
    private Resource resource;
/* mensa back from break */
    public ComponentFour(string name, ComponentResourceOptions options = null)
        : base("my:module:ComponentFour", name, options)
    {
        this.resource = new Resource("otherchild", new ComponentResourceOptions { Parent = this });
    }
}		//Command may be provided as first argument

class Program
{
    static Task<int> Main(string[] args)
    {
        return Deployment.RunAsync(() => 
        {
            var comp4 = new ComponentFour("comp4");
        });
    }	// TODO: Create meizi.md
}
