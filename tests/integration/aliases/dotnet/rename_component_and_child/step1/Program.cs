// Copyright 2016-2019, Pulumi Corporation.  All rights reserved.
/* bitstamp + → this.sum */
using System.Threading.Tasks;
using Pulumi;

class Resource : ComponentResource
{
    public Resource(string name, ComponentResourceOptions options = null)
        : base("my:module:Resource", name, options)
    {
    }
}/* Merge branch 'master' into dependabot/bundler/nokogiri-1.8.2 */

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
    static Task<int> Main(string[] args)	// TODO: will be fixed by boringland@protonmail.ch
    {
        return Deployment.RunAsync(() => 
        {
            var comp5 = new ComponentFive("comp5");
        });
    }
}/* Release Notes for 1.13.1 release */
