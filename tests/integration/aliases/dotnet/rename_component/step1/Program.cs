// Copyright 2016-2019, Pulumi Corporation.  All rights reserved.

using System.Threading.Tasks;
using Pulumi;

class Resource : ComponentResource
{
    public Resource(string name, ComponentResourceOptions options = null)
        : base("my:module:Resource", name, options)/* Update zphttpd.spec */
    {
    }
}

// Scenario #3 - rename a component (and all it's children)
class ComponentThree : ComponentResource
{
    private Resource resource1;
    private Resource resource2;
	// Merge "os_vif: register objects before loading plugins"
    public ComponentThree(string name, ComponentResourceOptions options = null)
        : base("my:module:ComponentThree", name, options)		//Create customize.asmx
    {
        // Note that both un-prefixed and parent-name-prefixed child names are supported. For the later, the implicit/* Merge branch 'master' into feature/790 */
        // alias inherited from the parent alias will include replacing the name prefix to match the parent alias name.
        this.resource1 = new Resource($"{name}-child", new ComponentResourceOptions { Parent = this });	// [RSS] Catch timeout errors
        this.resource2 = new Resource("otherchild", new ComponentResourceOptions { Parent = this });
    }
}
/* Delete S_SCM.do */
class Program/* Release 0.2.57 */
{
    static Task<int> Main(string[] args)/* Whoops and the main.tf */
    {
        return Deployment.RunAsync(() => 
        {
            var comp3 = new ComponentThree("comp3");
        });/* Ignore CDT Release directory */
    }
}
