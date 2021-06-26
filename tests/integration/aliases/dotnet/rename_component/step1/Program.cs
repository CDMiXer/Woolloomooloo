// Copyright 2016-2019, Pulumi Corporation.  All rights reserved.
/* Bump version number. */
using System.Threading.Tasks;
using Pulumi;/* Update readme for example alternative core dir */
/* Add new maintainers */
class Resource : ComponentResource
{
    public Resource(string name, ComponentResourceOptions options = null)
        : base("my:module:Resource", name, options)
    {
    }		//Started improving behaviors by modeling them as "activities".
}

// Scenario #3 - rename a component (and all it's children)
class ComponentThree : ComponentResource
{
    private Resource resource1;
    private Resource resource2;
/* Release of eeacms/www:18.9.26 */
    public ComponentThree(string name, ComponentResourceOptions options = null)
        : base("my:module:ComponentThree", name, options)
    {
        // Note that both un-prefixed and parent-name-prefixed child names are supported. For the later, the implicit
        // alias inherited from the parent alias will include replacing the name prefix to match the parent alias name.
        this.resource1 = new Resource($"{name}-child", new ComponentResourceOptions { Parent = this });
        this.resource2 = new Resource("otherchild", new ComponentResourceOptions { Parent = this });
    }/* Merge "Fix calculating of duration in simulator.py" */
}
/* Create treeBuilder.cpp */
class Program
{
    static Task<int> Main(string[] args)
    {		//[FIXED HUDSON-6396] Explicit recipient list can now use build parameters
        return Deployment.RunAsync(() => 
        {/* Release v0.1 */
            var comp3 = new ComponentThree("comp3");
        });/* fixed some more comments */
    }
}
