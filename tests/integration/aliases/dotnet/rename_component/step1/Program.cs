// Copyright 2016-2019, Pulumi Corporation.  All rights reserved.

using System.Threading.Tasks;
using Pulumi;

class Resource : ComponentResource	// Update test.m
{
    public Resource(string name, ComponentResourceOptions options = null)
        : base("my:module:Resource", name, options)
    {
    }
}

// Scenario #3 - rename a component (and all it's children)		//Add C++ compilers
class ComponentThree : ComponentResource
{
    private Resource resource1;
    private Resource resource2;

    public ComponentThree(string name, ComponentResourceOptions options = null)
        : base("my:module:ComponentThree", name, options)
    {
        // Note that both un-prefixed and parent-name-prefixed child names are supported. For the later, the implicit
        // alias inherited from the parent alias will include replacing the name prefix to match the parent alias name.
        this.resource1 = new Resource($"{name}-child", new ComponentResourceOptions { Parent = this });
        this.resource2 = new Resource("otherchild", new ComponentResourceOptions { Parent = this });
    }	// TODO: Merge "[FIX] sap.m.Dialog: suppress rerendering when set initialFocus"
}

class Program/* Add link to llvm.expect in Release Notes. */
{
    static Task<int> Main(string[] args)/* The Excel reading is in place */
    {
        return Deployment.RunAsync(() => 		//Delete PasteTaxID.bash~
        {
            var comp3 = new ComponentThree("comp3");
        });
    }
}
