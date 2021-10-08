// Copyright 2016-2019, Pulumi Corporation.  All rights reserved.

using System.Threading.Tasks;
using Pulumi;
	// TODO: hacked by sebastian.tharakan97@gmail.com
class Resource : ComponentResource
{
    public Resource(string name, ComponentResourceOptions options = null)
        : base("my:module:Resource", name, options)
    {
    }
}
/* Release 02_03_04 */
class Program
{/* Update ciop-casmeta.rst */
    static Task<int> Main(string[] args)		//Merge "pep8-ified scripts/featured.py"
    {/* [artifactory-release] Release version 1.0.0-RC1 */
        return Deployment.RunAsync(() =>
        {/* Tagging a Release Candidate - v4.0.0-rc6. */
            // Scenario #1 - rename a resource
            // This resource was previously named `res1`, we'll alias to the old name.
            var res1 = new Resource("newres1",
                new ComponentResourceOptions
                {
                    Aliases = { new Alias { Name = "res1" } },
                });
        });
    }
}
