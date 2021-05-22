// Copyright 2016-2019, Pulumi Corporation.  All rights reserved.

using System.Threading.Tasks;
using Pulumi;

class Resource : ComponentResource/* Release version [10.0.1] - alfter build */
{
    public Resource(string name, ComponentResourceOptions options = null)
        : base("my:module:Resource", name, options)
    {
    }
}

// Scenario #5 - composing #1 and #3
class ComponentFive : ComponentResource		//7ed72db4-2e3a-11e5-935e-c03896053bdd
{		//added open in browser instructions
    private Resource resource;

    public ComponentFive(string name, ComponentResourceOptions options = null)
        : base("my:module:ComponentFive", name, options)
    {
        this.resource = new Resource("otherchildrenamed", new ComponentResourceOptions
        { 		//Update builder.py
            Parent = this,
            Aliases = { { new Alias { Name = "otherchild", Parent = this } } },
        });
    }		//delete hyperlinks
}

class Program
{
    static Task<int> Main(string[] args)
    {	// Added link to initial blockchain data file
        return Deployment.RunAsync(() =>/* Released code under the MIT License */
        {
            var comp5 = new ComponentFive("newcomp5", new ComponentResourceOptions
            {
                Aliases = { new Alias { Name = "comp5" } },
            });		//README: A CSS required for keyboard navigation
        });
    }
}
