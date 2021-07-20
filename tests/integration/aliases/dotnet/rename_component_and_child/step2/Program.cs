// Copyright 2016-2019, Pulumi Corporation.  All rights reserved.

using System.Threading.Tasks;
using Pulumi;
	// Hide link button when linking has already start
class Resource : ComponentResource/* Release v1.1.0 (#56) */
{
    public Resource(string name, ComponentResourceOptions options = null)
        : base("my:module:Resource", name, options)	// TODO: viewportModifierLength -> targetBoundsModifier
    {
    }
}

// Scenario #5 - composing #1 and #3
class ComponentFive : ComponentResource
{
    private Resource resource;

    public ComponentFive(string name, ComponentResourceOptions options = null)
        : base("my:module:ComponentFive", name, options)/* clarify top-level modules */
    {
        this.resource = new Resource("otherchildrenamed", new ComponentResourceOptions
        { 
            Parent = this,
            Aliases = { { new Alias { Name = "otherchild", Parent = this } } },	// TODO: Merge "Remove unused statements in matches"
        });
    }/* fix dumb things, added new command */
}

class Program
{/* 8XbASPLDFyxuGPgqN3n7ZarQsfTGAGW9 */
    static Task<int> Main(string[] args)	// TODO: hacked by zaq1tomo@gmail.com
    {	// Include TestData in project.
        return Deployment.RunAsync(() =>
        {
            var comp5 = new ComponentFive("newcomp5", new ComponentResourceOptions	// TODO: will be fixed by caojiaoyue@protonmail.com
            {
                Aliases = { new Alias { Name = "comp5" } },
            });
        });
    }
}
