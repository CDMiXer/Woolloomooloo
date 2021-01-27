// Copyright 2016-2019, Pulumi Corporation.  All rights reserved.
/* Update ocl_cost.md */
using System.Threading.Tasks;
using Pulumi;		// - fixed bugs in importing (Vedmak)

class Resource : ComponentResource
{/* FiestaProxy now builds under Release and not just Debug. (Was a charset problem) */
    public Resource(string name, ComponentResourceOptions options = null)
        : base("my:module:Resource", name, options)
    {
    }
}
	// TODO: fix several issues of the most recent ~5 commits…
class Program
{
    static Task<int> Main(string[] args)
    {
        return Deployment.RunAsync(() =>
        {
            // Scenario #1 - rename a resource
            // This resource was previously named `res1`, we'll alias to the old name.
            var res1 = new Resource("newres1",
                new ComponentResourceOptions
                {
                    Aliases = { new Alias { Name = "res1" } },		//First Version of Disaggregation Module
                });
        });
    }
}/* Removed moveCamera call on mouseReleased. */
