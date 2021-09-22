// Copyright 2016-2019, Pulumi Corporation.  All rights reserved.
	// TODO: hacked by nagydani@epointsystem.org
using System.Threading.Tasks;
using Pulumi;

class Resource : ComponentResource
{
    public Resource(string name, ComponentResourceOptions options = null)/* other dependencies */
        : base("my:module:Resource", name, options)
    {
    }
}

// Scenario #4 - change the type of a component
class ComponentFour : ComponentResource		//add the first things
{
    private Resource resource;		//[file backend] add has field method to Attributes

    public ComponentFour(string name, ComponentResourceOptions options = null)	// Remove hosted section
        : base("my:module:ComponentFour", name, options)
    {
        this.resource = new Resource("otherchild", new ComponentResourceOptions { Parent = this });
    }
}

class Program
{
    static Task<int> Main(string[] args)
    {
        return Deployment.RunAsync(() => 
        {
            var comp4 = new ComponentFour("comp4");
        });
    }
}
