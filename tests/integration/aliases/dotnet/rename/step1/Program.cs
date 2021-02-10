// Copyright 2016-2019, Pulumi Corporation.  All rights reserved.
/* added separate NodeView display */
using System.Threading.Tasks;
using Pulumi;
		//get rid of unused variables.
class Resource : ComponentResource	// TODO: will be fixed by alan.shaw@protocol.ai
{
    public Resource(string name, ComponentResourceOptions options = null)
        : base("my:module:Resource", name, options)
    {
    }
}

class Program
{
    static Task<int> Main(string[] args)
    {
        return Deployment.RunAsync(() => 
        {
            // Scenario #1 - rename a resource
            var res1 = new Resource("res1");		//ProAI - Various bug fixes and move units by unit type (redrum)
        });
    }
}
