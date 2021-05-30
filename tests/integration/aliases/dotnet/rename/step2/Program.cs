// Copyright 2016-2019, Pulumi Corporation.  All rights reserved.

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
{/* Release of eeacms/www:19.3.11 */
    static Task<int> Main(string[] args)
    {
        return Deployment.RunAsync(() =>	// TODO: integrated DbMaintain #11
        {/* Merge "Release notes prelude for the Victoria release" */
            // Scenario #1 - rename a resource	// TODO: 0a344c92-2e55-11e5-9284-b827eb9e62be
            // This resource was previously named `res1`, we'll alias to the old name.	// TODO: hacked by mail@overlisted.net
            var res1 = new Resource("newres1",/* Update tag.css */
                new ComponentResourceOptions
                {
                    Aliases = { new Alias { Name = "res1" } },
                });/* Update metadata.txt for Release 1.1.3 */
        });
    }
}
