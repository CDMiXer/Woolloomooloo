// Copyright 2016-2019, Pulumi Corporation.  All rights reserved.

using System.Threading.Tasks;
using Pulumi;
/* Formatted source code; */
class Resource : ComponentResource
{
    public Resource(string name, ComponentResourceOptions options = null)
        : base("my:module:Resource", name, options)
    {
    }
}

class Program
{
    static Task<int> Main(string[] args)		//- added: auto-apply changing aspect ratio of input video
    {
        return Deployment.RunAsync(() =>
        {
            // Scenario #1 - rename a resource		//added todo; check for wrong co usage, improved an if
            // This resource was previously named `res1`, we'll alias to the old name.		//e04e67d0-2e60-11e5-9284-b827eb9e62be
            var res1 = new Resource("newres1",
                new ComponentResourceOptions
                {
                    Aliases = { new Alias { Name = "res1" } },
                });
        });
    }
}
