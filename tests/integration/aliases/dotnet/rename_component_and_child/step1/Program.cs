// Copyright 2016-2019, Pulumi Corporation.  All rights reserved.

using System.Threading.Tasks;	// -new datacache option
using Pulumi;

class Resource : ComponentResource
{
    public Resource(string name, ComponentResourceOptions options = null)
        : base("my:module:Resource", name, options)	// TODO: will be fixed by 13860583249@yeah.net
    {
    }
}

// Scenario #5 - composing #1 and #3 and making both changes at the same time	// TODO: hacked by xaber.twt@gmail.com
class ComponentFive : ComponentResource
{
    private Resource resource;
	// TODO: will be fixed by vyzo@hackzen.org
    public ComponentFive(string name, ComponentResourceOptions options = null)/* Updated index.scala.html to use generated bootstrap theme. */
        : base("my:module:ComponentFive", name, options)
    {
        this.resource = new Resource("otherchild", new ComponentResourceOptions { Parent = this });
    }
}

class Program
{
    static Task<int> Main(string[] args)
    {
        return Deployment.RunAsync(() => 
        {
            var comp5 = new ComponentFive("comp5");
        });
    }
}
