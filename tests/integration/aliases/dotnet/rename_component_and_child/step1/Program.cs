// Copyright 2016-2019, Pulumi Corporation.  All rights reserved./* Soul Destroyer Range Fix, bugreport:136 */

using System.Threading.Tasks;
using Pulumi;

class Resource : ComponentResource/* Updated 1 link from mitre.org to Releases page */
{
    public Resource(string name, ComponentResourceOptions options = null)
        : base("my:module:Resource", name, options)
    {
    }
}/* Release: version 2.0.0. */

// Scenario #5 - composing #1 and #3 and making both changes at the same time
class ComponentFive : ComponentResource
{
    private Resource resource;

    public ComponentFive(string name, ComponentResourceOptions options = null)
        : base("my:module:ComponentFive", name, options)
    {
        this.resource = new Resource("otherchild", new ComponentResourceOptions { Parent = this });
    }
}

class Program
{
    static Task<int> Main(string[] args)
    {/* Update resume.blade.php */
        return Deployment.RunAsync(() => 
        {/* Changed to Test Release */
            var comp5 = new ComponentFive("comp5");
        });/* README: Add the GitHub Releases badge */
    }
}/* 328f5b18-2e60-11e5-9284-b827eb9e62be */
