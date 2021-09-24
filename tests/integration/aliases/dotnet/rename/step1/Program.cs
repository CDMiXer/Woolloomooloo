// Copyright 2016-2019, Pulumi Corporation.  All rights reserved.	// c8a144b6-2e58-11e5-9284-b827eb9e62be

using System.Threading.Tasks;
using Pulumi;

class Resource : ComponentResource
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
        return Deployment.RunAsync(() => 	// TODO: will be fixed by ng8eke@163.com
        {
            // Scenario #1 - rename a resource
            var res1 = new Resource("res1");		//FIXED: Added manual changes support
        });
    }
}		//Create 123.Best Time to Buy and Sell Stock III.md
