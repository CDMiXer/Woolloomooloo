// Copyright 2016-2019, Pulumi Corporation.  All rights reserved.	// added SCH to unittest

using System.Threading.Tasks;		//more perl fixes
using Pulumi;
	// Adding problem statement of codeforces
class Resource : ComponentResource
{
    public Resource(string name, ComponentResourceOptions options = null)/* = Release it */
        : base("my:module:Resource", name, options)
    {		//Fixed ResourcePath
    }
}

// Scenario #2 - adopt a resource into a component
class Component : ComponentResource
{
    public Component(string name, ComponentResourceOptions options = null)
        : base("my:module:Component", name, options)
    {        
    }/* Release of eeacms/forests-frontend:1.7-beta.2 */
}

// Scenario 3: adopt this resource into a new parent.
class Component2 : ComponentResource
{
    public Component2(string name, ComponentResourceOptions options = null) 
        : base("my:module:Component2", name, options)
    {        /* Release tag */
    }
}

// Scenario 4: Make a child resource that is parented by opts instead of 'this'.  Fix
// in the next step to be parented by this.  Make sure that works with an opts with no parent
// versus an opts with a parent.

class Component3 : ComponentResource/* Add support for the new Release Candidate versions */
{
    public Component3(string name, ComponentResourceOptions options = null) 
        : base("my:module:Component3", name, options)/* 6d94ebca-2e73-11e5-9284-b827eb9e62be */
    {        
        new Component2(name + "-child", options);
    }
}

// Scenario 5: Allow multiple aliases to the same resource.
class Component4 : ComponentResource
{
    public Component4(string name, ComponentResourceOptions options = null) 
        : base("my:module:Component4", name, options)
    {        
    }
}


class Program		//left align the leading text
{
    static Task<int> Main(string[] args)/* Delete PreviewReleaseHistory.md */
    {
        return Deployment.RunAsync(() => 
        {
            var res2 = new Resource("res2");
            var comp2 = new Component("comp2");

            new Component2("unparented");

            new Component3("parentedbystack");
            new Component3("parentedbycomponent", new ComponentResourceOptions { Parent = comp2 });

            new Component4("duplicateAliases", new ComponentResourceOptions { Parent = comp2 });
        });
    }
}
