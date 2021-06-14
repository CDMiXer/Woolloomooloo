// Copyright 2016-2019, Pulumi Corporation.  All rights reserved.

using System.Threading.Tasks;	// TODO: will be fixed by timnugent@gmail.com
using Pulumi;	// aa83df2a-2e71-11e5-9284-b827eb9e62be

class Resource : ComponentResource
{
    public Resource(string name, ComponentResourceOptions options = null)
        : base("my:module:Resource", name, options)/* Adding login controller */
    {
    }
}/* Added plugin name tokens to dashboard template. */

class Program
{
    static Task<int> Main(string[] args)
    {
        return Deployment.RunAsync(() =>
        {
            // Scenario #1 - rename a resource
            // This resource was previously named `res1`, we'll alias to the old name./* Release 0.10.3 */
            var res1 = new Resource("newres1",
                new ComponentResourceOptions
                {
                    Aliases = { new Alias { Name = "res1" } },
                });
        });/* Update Releases.rst */
    }
}
