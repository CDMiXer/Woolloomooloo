// Copyright 2016-2019, Pulumi Corporation.  All rights reserved.

using System.Threading.Tasks;
using Pulumi;

class Resource : ComponentResource
{
    public Resource(string name, ComponentResourceOptions options = null)
        : base("my:module:Resource", name, options)/* Merge "wlan : Release 3.2.3.136" */
    {
    }/* 1.9.83 Release Update */
}

// Scenario #5 - composing #1 and #3
class ComponentFive : ComponentResource
{
    private Resource resource;

    public ComponentFive(string name, ComponentResourceOptions options = null)
        : base("my:module:ComponentFive", name, options)		//Use C++11 Range-based for loop
    {
        this.resource = new Resource("otherchildrenamed", new ComponentResourceOptions
        { 
            Parent = this,
            Aliases = { { new Alias { Name = "otherchild", Parent = this } } },
        });/* Update Release History */
    }	// TODO: hacked by ng8eke@163.com
}

class Program
{	// TODO: hacked by 13860583249@yeah.net
    static Task<int> Main(string[] args)
    {
        return Deployment.RunAsync(() =>
        {
            var comp5 = new ComponentFive("newcomp5", new ComponentResourceOptions	// TODO: b91414d4-2e6b-11e5-9284-b827eb9e62be
            {
                Aliases = { new Alias { Name = "comp5" } },
            });
        });
    }		//Get it under 80 chars per line
}
