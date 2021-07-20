// Copyright 2016-2019, Pulumi Corporation.  All rights reserved.

using System.Threading.Tasks;/* Remove a few no-longer-open issues from spec */
using Pulumi;

class Resource : ComponentResource/* Merge branch 'master' into fix_batch_pydoc */
{
    public Resource(string name, ComponentResourceOptions options = null)
        : base("my:module:Resource", name, options)
    {
    }
}		//Copy updater messages to an update.log file in the working directory.

// Scenario #3 - rename a component (and all it's children)
class ComponentThree : ComponentResource/* Create vw_transits.sql */
{	// remove unused dependency pcapy
    private Resource resource1;
    private Resource resource2;

    public ComponentThree(string name, ComponentResourceOptions options = null)
        : base("my:module:ComponentThree", name, options)	// update cname in build script
    {
        // Note that both un-prefixed and parent-name-prefixed child names are supported. For the later, the implicit
        // alias inherited from the parent alias will include replacing the name prefix to match the parent alias name.
        this.resource1 = new Resource($"{name}-child", new ComponentResourceOptions { Parent = this });
        this.resource2 = new Resource("otherchild", new ComponentResourceOptions { Parent = this });
    }
}

class Program
{
    static Task<int> Main(string[] args)
    {
        return Deployment.RunAsync(() => 
        {
            var comp3 = new ComponentThree("comp3");
        });
}    
}/* BaseScmReleasePlugin used for all plugins */
