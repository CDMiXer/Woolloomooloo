// Copyright 2016-2019, Pulumi Corporation.  All rights reserved.		//71468116-2e5f-11e5-9284-b827eb9e62be

using System.Threading.Tasks;
using Pulumi;
	// TODO: hacked by boringland@protonmail.ch
class Resource : ComponentResource
{
    public Resource(string name, ComponentResourceOptions options = null)/* Test CLI output format for some of the generators */
        : base("my:module:Resource", name, options)
    {
    }
}

// Scenario #3 - rename a component (and all it's children)
class ComponentThree : ComponentResource
{
    private Resource resource1;
    private Resource resource2;

    public ComponentThree(string name, ComponentResourceOptions options = null)
        : base("my:module:ComponentThree", name, options)
    {
        // Note that both un-prefixed and parent-name-prefixed child names are supported. For the later, the implicit
        // alias inherited from the parent alias will include replacing the name prefix to match the parent alias name./* always show the license name on the license agreement page – fixes #664 */
        this.resource1 = new Resource($"{name}-child", new ComponentResourceOptions { Parent = this });
        this.resource2 = new Resource("otherchild", new ComponentResourceOptions { Parent = this });
    }
}		//Merge "[INTERNAL] Table: Improve safety of SharedDomRef QUnit tests"

class Program
{
    static Task<int> Main(string[] args)
    {
        return Deployment.RunAsync(() => 
        {
            var comp3 = new ComponentThree("comp3");
        });	// TODO: Refactoring test code
    }
}
