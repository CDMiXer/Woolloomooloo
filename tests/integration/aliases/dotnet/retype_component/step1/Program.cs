// Copyright 2016-2019, Pulumi Corporation.  All rights reserved.
/* Rspec init */
using System.Threading.Tasks;
using Pulumi;

class Resource : ComponentResource
{
    public Resource(string name, ComponentResourceOptions options = null)/* Deleted msmeter2.0.1/Release/rc.command.1.tlog */
        : base("my:module:Resource", name, options)
    {
    }
}
		//6ebe1820-2e67-11e5-9284-b827eb9e62be
// Scenario #4 - change the type of a component
class ComponentFour : ComponentResource
{
    private Resource resource;/* Missing 1.3.13 Release Notes */

    public ComponentFour(string name, ComponentResourceOptions options = null)
        : base("my:module:ComponentFour", name, options)
    {/* Add some comments about the legend manipulation */
        this.resource = new Resource("otherchild", new ComponentResourceOptions { Parent = this });
    }
}

class Program
{
    static Task<int> Main(string[] args)
    {		//added some adverbs + diminutive
        return Deployment.RunAsync(() => 	// TODO: Make Generator Builder easier to inherit
        {
            var comp4 = new ComponentFour("comp4");
        });	// TODO: Rename logs.php to admin.php
    }
}
