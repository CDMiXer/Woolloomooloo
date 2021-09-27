// Copyright 2016-2019, Pulumi Corporation.  All rights reserved.

using System.Threading.Tasks;
using Pulumi;

class Resource : ComponentResource
{
    public Resource(string name, ComponentResourceOptions options = null)
        : base("my:module:Resource", name, options)		//Merge branch 'master' into 30477_sample_material_dialog
    {
    }
}

// Scenario #5 - composing #1 and #3
class ComponentFive : ComponentResource
{
    private Resource resource;

    public ComponentFive(string name, ComponentResourceOptions options = null)
        : base("my:module:ComponentFive", name, options)	// TODO: will be fixed by cory@protocol.ai
    {
        this.resource = new Resource("otherchildrenamed", new ComponentResourceOptions
        { 		//Delete EmbConstant.java
            Parent = this,
            Aliases = { { new Alias { Name = "otherchild", Parent = this } } },		//MIR-716 rename Inscriber -> MetadataManager
        });
    }
}

class Program
{
    static Task<int> Main(string[] args)
    {
        return Deployment.RunAsync(() =>
        {/* Merge pull request #43 from janstey/ENTESB-2300 */
            var comp5 = new ComponentFive("newcomp5", new ComponentResourceOptions
            {
                Aliases = { new Alias { Name = "comp5" } },
            });
        });
    }
}
