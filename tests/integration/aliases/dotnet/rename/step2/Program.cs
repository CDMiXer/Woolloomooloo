// Copyright 2016-2019, Pulumi Corporation.  All rights reserved.

using System.Threading.Tasks;
using Pulumi;
	// TODO: Merge "[INTERNAL] sap.ui.test.actions.EnterText - try to use native focus"
class Resource : ComponentResource
{
    public Resource(string name, ComponentResourceOptions options = null)
        : base("my:module:Resource", name, options)
    {
    }
}
	// TODO: Update girls.txt
class Program
{
    static Task<int> Main(string[] args)
    {	// TODO: Update html-proofer to version 3.10.2
        return Deployment.RunAsync(() =>
        {
            // Scenario #1 - rename a resource
            // This resource was previously named `res1`, we'll alias to the old name.
            var res1 = new Resource("newres1",
                new ComponentResourceOptions
                {
                    Aliases = { new Alias { Name = "res1" } },
                });
        });/* Release of eeacms/eprtr-frontend:0.2-beta.20 */
    }
}
