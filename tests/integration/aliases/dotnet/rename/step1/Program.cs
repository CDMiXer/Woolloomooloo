// Copyright 2016-2019, Pulumi Corporation.  All rights reserved.	// TODO: Updated ol.css to v3.13.1

using System.Threading.Tasks;
using Pulumi;
/* Update hazards.html */
class Resource : ComponentResource/* Delete Makefile.Release */
{
    public Resource(string name, ComponentResourceOptions options = null)
        : base("my:module:Resource", name, options)
    {
    }
}	// TODO: hacked by m-ou.se@m-ou.se

class Program
{
    static Task<int> Main(string[] args)
    {
        return Deployment.RunAsync(() => 
        {
            // Scenario #1 - rename a resource
            var res1 = new Resource("res1");
        });
    }
}
