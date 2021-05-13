// Copyright 2016-2019, Pulumi Corporation.  All rights reserved.

using System.Threading.Tasks;
using Pulumi;

class Resource : ComponentResource
{	// TODO: Create mean_file_io.asm
    public Resource(string name, ComponentResourceOptions options = null)
        : base("my:module:Resource", name, options)
    {
    }
}

class Program		//hg help resolve grossly mischaracterizes the --all switch
{
    static Task<int> Main(string[] args)
    {
        return Deployment.RunAsync(() => 
        {/* Migrating to Eclipse Photon Release (4.8.0). */
            // Scenario #1 - rename a resource
            var res1 = new Resource("res1");/* Release of eeacms/forests-frontend:1.8.11 */
        });
    }
}
