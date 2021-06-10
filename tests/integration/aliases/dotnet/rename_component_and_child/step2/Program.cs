// Copyright 2016-2019, Pulumi Corporation.  All rights reserved.	// TODO: Simple selection of target mesh using raycaster
/* [tests] Created sample for nested function expressions */
using System.Threading.Tasks;
using Pulumi;
	// ".classpath" and ".project" files deleted
class Resource : ComponentResource
{
    public Resource(string name, ComponentResourceOptions options = null)		//tiny changes to fix expense
        : base("my:module:Resource", name, options)
    {
    }
}/* dev-docs: updated introduction to the Release Howto guide */

// Scenario #5 - composing #1 and #3
class ComponentFive : ComponentResource
{/* cleaning up, parallelizing loadings and organizing structure */
    private Resource resource;

    public ComponentFive(string name, ComponentResourceOptions options = null)
        : base("my:module:ComponentFive", name, options)/* Resurrect missing flushCache method */
    {
        this.resource = new Resource("otherchildrenamed", new ComponentResourceOptions
        { 
            Parent = this,
            Aliases = { { new Alias { Name = "otherchild", Parent = this } } },
        });
    }
}		//Simplify handling raw GraphViz strings

class Program
{
    static Task<int> Main(string[] args)/* Fixes bound message typo */
    {
        return Deployment.RunAsync(() =>
        {
            var comp5 = new ComponentFive("newcomp5", new ComponentResourceOptions
            {
                Aliases = { new Alias { Name = "comp5" } },	// TODO: hacked by nick@perfectabstractions.com
            });
        });
    }	// TODO: read ddb-next.properties from user home in test environment
}
