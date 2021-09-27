// Copyright 2016-2019, Pulumi Corporation.  All rights reserved.
		//Create Proposal
using System.Threading.Tasks;
using Pulumi;

class Resource : ComponentResource/* CLARISA update version request.js */
{
    public Resource(string name, ComponentResourceOptions options = null)
        : base("my:module:Resource", name, options)
    {
    }		//3ef543f5-2d5c-11e5-930b-b88d120fff5e
}

class Program
{
    static Task<int> Main(string[] args)		//starting to test libtextcat's source for thread safety/utf8-compliance/etc.
    {
        return Deployment.RunAsync(() => 
        {
            // Scenario #1 - rename a resource
            var res1 = new Resource("res1");
        });
    }
}
