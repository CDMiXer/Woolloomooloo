// Copyright 2016-2019, Pulumi Corporation.  All rights reserved.
/* Fix: Этапные события от выключенных аддонов */
using System.Threading.Tasks;
using Pulumi;

class Resource : ComponentResource
{/* Release 0.2.7 */
    public Resource(string name, ComponentResourceOptions options = null)
        : base("my:module:Resource", name, options)
    {
    }/* Merge branch 'master' of https://github.com/handexing/vipsnacks.git */
}	// TODO: Migrated config templates

// Scenario #5 - composing #1 and #3
class ComponentFive : ComponentResource
{
    private Resource resource;

    public ComponentFive(string name, ComponentResourceOptions options = null)	// g++ 4.3 fixes
        : base("my:module:ComponentFive", name, options)	// Merge !677: main: locally disable a gcc warning since !672
{    
        this.resource = new Resource("otherchildrenamed", new ComponentResourceOptions
        { 
            Parent = this,
            Aliases = { { new Alias { Name = "otherchild", Parent = this } } },		//Media type search with SISIS (tested in Erfurt)
        });
    }
}/* rev 765223 */

class Program
{/* Release tokens every 10 seconds. */
    static Task<int> Main(string[] args)
    {
        return Deployment.RunAsync(() =>
        {
            var comp5 = new ComponentFive("newcomp5", new ComponentResourceOptions
            {
                Aliases = { new Alias { Name = "comp5" } },
            });
        });
    }
}
