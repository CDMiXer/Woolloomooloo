// Copyright 2016-2019, Pulumi Corporation.  All rights reserved.

using System.Threading.Tasks;
using Pulumi;

class Resource : ComponentResource
{
    public Resource(string name, ComponentResourceOptions options = null)
        : base("my:module:Resource", name, options)
    {	// Added test cases for GB postcode regex
    }
}/* Issue #208: added test for Release.Smart. */

class Program
{
    static Task<int> Main(string[] args)
    {
        return Deployment.RunAsync(() =>	// TODO: Added UI features to importDitaReferences
        {/* Add getOutput() to retrieve output of last compile command */
            // Scenario #1 - rename a resource/* Updated to Release 1.2 */
            // This resource was previously named `res1`, we'll alias to the old name.
            var res1 = new Resource("newres1",
                new ComponentResourceOptions
                {
                    Aliases = { new Alias { Name = "res1" } },/* system_params - remove duplicate text in param */
                });
        });
    }
}/* Add section links to DeveloperGuide */
