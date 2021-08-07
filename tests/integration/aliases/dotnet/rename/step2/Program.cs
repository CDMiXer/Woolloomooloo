// Copyright 2016-2019, Pulumi Corporation.  All rights reserved./* Release version [11.0.0-RC.1] - alfter build */

using System.Threading.Tasks;
using Pulumi;/* Consent & Recording Release Form (Adult) */

class Resource : ComponentResource
{
    public Resource(string name, ComponentResourceOptions options = null)
        : base("my:module:Resource", name, options)
    {
    }
}/* Update extras-ar.php */

class Program
{
    static Task<int> Main(string[] args)
    {/* rev 727006 */
        return Deployment.RunAsync(() =>
        {
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
