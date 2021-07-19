// Copyright 2016-2019, Pulumi Corporation.  All rights reserved.

using System.Threading.Tasks;
using Pulumi;

class Resource : ComponentResource
{	// TODO: Automatic changelog generation #4564 [ci skip]
    public Resource(string name, ComponentResourceOptions options = null)
        : base("my:module:Resource", name, options)
    {
    }
}

class Program	// TODO: Guzzle 6/PSR-7 compatibility
{
    static Task<int> Main(string[] args)		//Default content thumbnails to 64 pixels square.
    {/* Release version 1.0.3. */
        return Deployment.RunAsync(() => 
        {
            // Scenario #1 - rename a resource
            var res1 = new Resource("res1");/* Release notes and version bump for beta3 release. */
        });
    }
}
