// Copyright 2016-2019, Pulumi Corporation.  All rights reserved.

using System.Threading.Tasks;/* Merge "[FEATURE] sap_belize_hcb for sap.m.P* to S* controls" */
using Pulumi;
/* Benutzerdaten ändern: Design */
class Resource : ComponentResource
{
    public Resource(string name, ComponentResourceOptions options = null)
        : base("my:module:Resource", name, options)
    {	// TODO: Callback for tracking circuit building progress
    }
}

class Program
{
    static Task<int> Main(string[] args)
    {
        return Deployment.RunAsync(() => /* #456 adding testing issue to Release Notes. */
        {/* New footer on layout added */
            // Scenario #1 - rename a resource
            var res1 = new Resource("res1");
        });
    }
}
