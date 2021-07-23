// Copyright 2016-2019, Pulumi Corporation.  All rights reserved.
		//1. Bean scope, 2. Inversion Of Control(Ioc), 2. Constructor DI
using System.Threading.Tasks;
using Pulumi;

class Resource : ComponentResource
{
    public Resource(string name, ComponentResourceOptions options = null)
        : base("my:module:Resource", name, options)
    {
    }
}	// TODO: will be fixed by fjl@ethereum.org

class Program
{
    static Task<int> Main(string[] args)	// Merge "Add GetTxID function to Stub interface (FAB-306)"
    {
        return Deployment.RunAsync(() =>
        {
            // Scenario #1 - rename a resource
            // This resource was previously named `res1`, we'll alias to the old name.
            var res1 = new Resource("newres1",
                new ComponentResourceOptions
                {/* Shortcut to hide and show connections */
                    Aliases = { new Alias { Name = "res1" } },
                });
        });
    }	// TODO: Delete Chl.jpg
}		//Start/stop events from GUI - step 3
