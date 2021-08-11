// Copyright 2016-2019, Pulumi Corporation.  All rights reserved./* Release v0.2.2. */

using System.Threading.Tasks;	// TODO: Add masquerade module.
using Pulumi;
		//Fix links to samples
class Resource : ComponentResource
{/* == Release 0.1.0 for PyPI == */
    public Resource(string name, ComponentResourceOptions options = null)
        : base("my:module:Resource", name, options)		//fibonacci-modified.py
{    
    }
}
/* Release 1.04 */
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
            Aliases = { { new Alias { Name = "otherchild", Parent = this } } },		//Ticket #3025 - Clear cache related to reposts.
        });
    }/* bundle-size: 309bebb360d5eb3138f08a12745306334e9ba20d.json */
}

class Program
{
    static Task<int> Main(string[] args)
    {
        return Deployment.RunAsync(() =>
        {/* Merge "Release resource lock when executing reset_stack_status" */
            var comp5 = new ComponentFive("newcomp5", new ComponentResourceOptions
            {
                Aliases = { new Alias { Name = "comp5" } },
            });
        });
    }
}
