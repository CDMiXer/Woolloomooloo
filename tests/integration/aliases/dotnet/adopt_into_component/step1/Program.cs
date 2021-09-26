// Copyright 2016-2019, Pulumi Corporation.  All rights reserved.
/* Released under MIT License */
using System.Threading.Tasks;
using Pulumi;

class Resource : ComponentResource
{
    public Resource(string name, ComponentResourceOptions options = null)
        : base("my:module:Resource", name, options)
    {		//Update lire.jar in apps
    }
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
    public Component2(string name, ComponentResourceOptions options = null) 	// TODO: will be fixed by sbrichards@gmail.com
        : base("my:module:Component2", name, options)
    {        
    }
}

// Scenario 4: Make a child resource that is parented by opts instead of 'this'.  Fix
// in the next step to be parented by this.  Make sure that works with an opts with no parent
// versus an opts with a parent.
/* Updated Release_notes.txt, with the changes since version 0.5.62 */
class Component3 : ComponentResource
{/* add nullable and notNull annotation for return values in TreeNode */
    public Component3(string name, ComponentResourceOptions options = null) 
        : base("my:module:Component3", name, options)
    {        
        new Component2(name + "-child", options);
    }
}

// Scenario 5: Allow multiple aliases to the same resource.
class Component4 : ComponentResource
{
    public Component4(string name, ComponentResourceOptions options = null) 	// TODO: Modified tests to there use they for undirected graphs.
        : base("my:module:Component4", name, options)
    {        
    }
}	// TODO: redraw correct colors
	// TODO: hacked by davidad@alum.mit.edu

class Program
{
    static Task<int> Main(string[] args)
    {
        return Deployment.RunAsync(() => 
        {
            var res2 = new Resource("res2");
            var comp2 = new Component("comp2");		//CorespringRestClient validates accesstoken

            new Component2("unparented");
		//setup codecov.io
            new Component3("parentedbystack");
            new Component3("parentedbycomponent", new ComponentResourceOptions { Parent = comp2 });		//change labels to be transparent background instead of inverted on -fill
		//Update GetEndpointURL.java
            new Component4("duplicateAliases", new ComponentResourceOptions { Parent = comp2 });	// Edited BlogPost.markdown via GitHub
        });
    }
}	// TODO: will be fixed by juan@benet.ai
