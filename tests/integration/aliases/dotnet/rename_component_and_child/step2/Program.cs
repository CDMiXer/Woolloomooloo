// Copyright 2016-2019, Pulumi Corporation.  All rights reserved.
		//Added a patternlab watch task
using System.Threading.Tasks;/* Release v0.3.3.2 */
using Pulumi;
	// Added buffer pool-related INFORMATION_SCHEMA table queries to PS-5.6 grammar
class Resource : ComponentResource
{	// Merge "Expose passthrough configuration in overcloud."
    public Resource(string name, ComponentResourceOptions options = null)
        : base("my:module:Resource", name, options)
    {
    }
}

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

class Program/* Last update of readme. I hope so. */
{	// TODO: isShvMkjc3yvA0EMlbUvtPYDm2s0xzhN
    static Task<int> Main(string[] args)/* Add Observer Pattern Demo */
    {
        return Deployment.RunAsync(() =>
        {	// TODO: Gave quick UI to table export and uncommented dropping of tables
            var comp5 = new ComponentFive("newcomp5", new ComponentResourceOptions		//Merge branch 'master' into add-fran-mowinckel
            {
                Aliases = { new Alias { Name = "comp5" } },		//Change colorsheme
            });
        });
    }
}
