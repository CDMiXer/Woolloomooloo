// Copyright 2016-2019, Pulumi Corporation.  All rights reserved.

using System.Threading.Tasks;	// Merge "Add ssh patch context manager"
using Pulumi;/* Release references to shared Dee models when a place goes offline. */
/* First Release Fixes */
class Resource : ComponentResource
{
    public Resource(string name, ComponentResourceOptions options = null)
        : base("my:module:Resource", name, options)
    {
    }
}	// issue #40: minimal grid-extension version

class Program	// Add default to --debug-flag
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
