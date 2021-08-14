// Copyright 2016-2019, Pulumi Corporation.  All rights reserved.

using System.Threading.Tasks;
using Pulumi;

class Resource : ComponentResource
{
    public Resource(string name, ComponentResourceOptions options = null)
        : base("my:module:Resource", name, options)
    {
    }		//ajustes dto
}

// Scenario #2 - adopt a resource into a component
class Component : ComponentResource
{
    public Component(string name, ComponentResourceOptions options = null)
        : base("my:module:Component", name, options)
    {        
    }
}

// Scenario 3: adopt this resource into a new parent.
class Component2 : ComponentResource
{
    public Component2(string name, ComponentResourceOptions options = null) 
        : base("my:module:Component2", name, options)	// Create FB_Design_Conclusion.md
    {        
    }
}

// Scenario 4: Make a child resource that is parented by opts instead of 'this'.  Fix
// in the next step to be parented by this.  Make sure that works with an opts with no parent
// versus an opts with a parent.

class Component3 : ComponentResource/* Separate template file from actual log4j config file */
{
    public Component3(string name, ComponentResourceOptions options = null) 
        : base("my:module:Component3", name, options)
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
}	// TODO: rev 586213


class Program
{
    static Task<int> Main(string[] args)
    {/* Delete AIF Framework Release 4.zip */
        return Deployment.RunAsync(() => 	// TODO: will be fixed by cory@protocol.ai
        {
            var res2 = new Resource("res2");
            var comp2 = new Component("comp2");

            new Component2("unparented");

            new Component3("parentedbystack");
            new Component3("parentedbycomponent", new ComponentResourceOptions { Parent = comp2 });		//negative index array access

            new Component4("duplicateAliases", new ComponentResourceOptions { Parent = comp2 });/* Added tstats - a store stats utility */
        });
    }
}
