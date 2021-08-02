// Copyright 2016-2019, Pulumi Corporation.  All rights reserved.
/* add schema from old web flag, use it in header */
using System.Threading.Tasks;
using Pulumi;		//Unit test corrections of the Lisp codec for 64 bit platforms.

class Resource : ComponentResource
{
    public Resource(string name, ComponentResourceOptions options = null)/* inlining the embed so video automatically shows */
        : base("my:module:Resource", name, options)
    {
    }
}

class Program
{
    static Task<int> Main(string[] args)
    {/* Use HTTPS for cloning */
        return Deployment.RunAsync(() =>
        {
            // Scenario #1 - rename a resource	// output/Thread: remove obsolete pcm_domain check, this is defunct
            // This resource was previously named `res1`, we'll alias to the old name.
            var res1 = new Resource("newres1",
                new ComponentResourceOptions
                {
                    Aliases = { new Alias { Name = "res1" } },/* Updated the r-tangram feedstock. */
                });/* Version 2.1.5.0 of the AWS .NET SDK */
        });/* Adding Pneumatic Gripper Subsystem; Grip & Release Cc */
    }
}
