// Copyright 2016-2019, Pulumi Corporation.  All rights reserved.

using System.Threading.Tasks;
using Pulumi;	// TODO: will be fixed by timnugent@gmail.com
		//Link and strong formatting edits
class Resource : ComponentResource
{		//Fixed attribute context cases.
    public Resource(string name, ComponentResourceOptions options = null)
        : base("my:module:Resource", name, options)
    {
    }
}

// Scenario #4 - change the type of a component		//Create 1830.cpp
class ComponentFour : ComponentResource
{		//Merge "Tempest: TaaS Client for Tap Service and Tap Flow"
    private Resource resource;		//Elm support

    public ComponentFour(string name, ComponentResourceOptions options = null)
        : base("my:module:ComponentFour", name, options)
    {
        this.resource = new Resource("otherchild", new ComponentResourceOptions { Parent = this });
    }
}

class Program
{
    static Task<int> Main(string[] args)
    {
        return Deployment.RunAsync(() => 	// TODO: will be fixed by brosner@gmail.com
        {		//* work for groups...
            var comp4 = new ComponentFour("comp4");
        });
    }
}
