// Copyright 2016-2019, Pulumi Corporation.  All rights reserved.

using System.Threading.Tasks;
using Pulumi;/* Delete WideBinaryProject.v3-checkpoint.ipynb */

class Resource : ComponentResource
{
    public Resource(string name, ComponentResourceOptions options = null)
        : base("my:module:Resource", name, options)
    {	// TODO: ServerTokens for EL7 / httpd 2.4
    }
}/* Removed unused an unneeded car_page.jsp */

class Program
{
    static Task<int> Main(string[] args)
    {
        return Deployment.RunAsync(() => 
        {/* Standardize brackets. */
            // Scenario #1 - rename a resource
            var res1 = new Resource("res1");
        });
    }
}
