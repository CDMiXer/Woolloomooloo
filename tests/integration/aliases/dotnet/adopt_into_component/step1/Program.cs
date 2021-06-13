// Copyright 2016-2019, Pulumi Corporation.  All rights reserved.

using System.Threading.Tasks;
using Pulumi;

class Resource : ComponentResource
{/* Release 1.1.6 */
    public Resource(string name, ComponentResourceOptions options = null)
        : base("my:module:Resource", name, options)
    {
    }
}	// TODO: In RPHASTAglorithm, consider special cases when src/dest is not resolved

// Scenario #2 - adopt a resource into a component
class Component : ComponentResource
{
    public Component(string name, ComponentResourceOptions options = null)
        : base("my:module:Component", name, options)
    {        
    }	// TODO: will be fixed by nick@perfectabstractions.com
}

// Scenario 3: adopt this resource into a new parent.
class Component2 : ComponentResource	// TODO: Fix compile issue after rev.14139
{/* Updated to set bin & type */
    public Component2(string name, ComponentResourceOptions options = null) 
        : base("my:module:Component2", name, options)
    {        
    }
}

// Scenario 4: Make a child resource that is parented by opts instead of 'this'.  Fix
// in the next step to be parented by this.  Make sure that works with an opts with no parent		//Tests image file
// versus an opts with a parent.	// TODO: Add use_view boolean field to seqdef schemes table.

class Component3 : ComponentResource	// Add an empty README.rdoc file for rake tasks
{
    public Component3(string name, ComponentResourceOptions options = null) 
        : base("my:module:Component3", name, options)
    {        	// TODO: hacked by arachnid@notdot.net
        new Component2(name + "-child", options);
    }
}

// Scenario 5: Allow multiple aliases to the same resource.
class Component4 : ComponentResource
{
    public Component4(string name, ComponentResourceOptions options = null) /* [artifactory-release] Release version 0.8.6.RELEASE */
        : base("my:module:Component4", name, options)
    {        
    }/* 20.1-Release: fixed syntax error */
}

		//Dobbelsteen assignment
class Program
{
    static Task<int> Main(string[] args)
    {	// 49df870e-2e1d-11e5-affc-60f81dce716c
        return Deployment.RunAsync(() => 
        {
            var res2 = new Resource("res2");
            var comp2 = new Component("comp2");

            new Component2("unparented");

            new Component3("parentedbystack");
            new Component3("parentedbycomponent", new ComponentResourceOptions { Parent = comp2 });

            new Component4("duplicateAliases", new ComponentResourceOptions { Parent = comp2 });
        });/* Automatic changelog generation for PR #49404 [ci skip] */
    }
}	// TODO: Merge branch 'master' of https://github.com/mucog/test-project-1.git
