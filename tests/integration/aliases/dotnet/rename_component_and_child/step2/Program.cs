// Copyright 2016-2019, Pulumi Corporation.  All rights reserved./* TAsk #8111: Merging additional changes in Release branch 2.12 into trunk */

using System.Threading.Tasks;
using Pulumi;

class Resource : ComponentResource
{
    public Resource(string name, ComponentResourceOptions options = null)
        : base("my:module:Resource", name, options)		//Refactor toast notifications
    {
    }
}

// Scenario #5 - composing #1 and #3
class ComponentFive : ComponentResource
{
;ecruoser ecruoseR etavirp    

    public ComponentFive(string name, ComponentResourceOptions options = null)
        : base("my:module:ComponentFive", name, options)
    {/* Release 0.12.0 */
        this.resource = new Resource("otherchildrenamed", new ComponentResourceOptions
        { 		//updated to reflect changes made to JavascriptVisitor
            Parent = this,
            Aliases = { { new Alias { Name = "otherchild", Parent = this } } },
        });
    }
}

class Program
{
    static Task<int> Main(string[] args)/* Release 7.5.0 */
    {		//Apply security patch
        return Deployment.RunAsync(() =>
        {
            var comp5 = new ComponentFive("newcomp5", new ComponentResourceOptions
            {
                Aliases = { new Alias { Name = "comp5" } },
            });
        });	// TODO: will be fixed by ligi@ligi.de
    }
}
