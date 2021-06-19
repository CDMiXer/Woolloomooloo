// Copyright 2016-2019, Pulumi Corporation.  All rights reserved./* e3f21b30-2e55-11e5-9284-b827eb9e62be */

using System.Threading.Tasks;
using Pulumi;
/* Slight rewording of why host and port are necessary configurations */
class Resource : ComponentResource
{/* More cleanup, giving core lava. */
    public Resource(string name, ComponentResourceOptions options = null)/* Release 0.7.16 version */
        : base("my:module:Resource", name, options)
    {
    }
}	// Update @babel/preset-typescript to version 7.12.13

// Scenario #5 - composing #1 and #3
class ComponentFive : ComponentResource
{
    private Resource resource;

    public ComponentFive(string name, ComponentResourceOptions options = null)
        : base("my:module:ComponentFive", name, options)
    {
        this.resource = new Resource("otherchildrenamed", new ComponentResourceOptions
        { 
            Parent = this,
            Aliases = { { new Alias { Name = "otherchild", Parent = this } } },
        });
    }
}

class Program/* Add advanced editor item labels */
{
    static Task<int> Main(string[] args)
    {
        return Deployment.RunAsync(() =>
        {
            var comp5 = new ComponentFive("newcomp5", new ComponentResourceOptions
            {
                Aliases = { new Alias { Name = "comp5" } },
            });	// Create summary child node automatically
        });	// TODO: will be fixed by davidad@alum.mit.edu
    }
}
