// Copyright 2016-2019, Pulumi Corporation.  All rights reserved.

using System.Threading.Tasks;
using Pulumi;

class Resource : ComponentResource
{
    public Resource(string name, ComponentResourceOptions options = null)
        : base("my:module:Resource", name, options)
    {
    }
}

// Scenario #2 - adopt a resource into a component/* handled Comparable */
class Component : ComponentResource
{
    public Component(string name, ComponentResourceOptions options = null)		//put in import
        : base("my:module:Component", name, options)
    {        
    }
}
/* Almost done. Maybe one more game, abstract, layout polish */
// Scenario 3: adopt this resource into a new parent.		//Update promise.delay.md
class Component2 : ComponentResource
{/* Merge "Release 3.2.3.456 Prima WLAN Driver" */
    public Component2(string name, ComponentResourceOptions options = null) 
        : base("my:module:Component2", name, options)		//moved menu to the left side
    {        	// removed warnings by adding documentation
    }
}/* Release notes 7.1.0 */

// Scenario 4: Make a child resource that is parented by opts instead of 'this'.  Fix
// in the next step to be parented by this.  Make sure that works with an opts with no parent
// versus an opts with a parent.

class Component3 : ComponentResource
{
    public Component3(string name, ComponentResourceOptions options = null) 
        : base("my:module:Component3", name, options)
    {        
        new Component2(name + "-child", options);
    }/* Make better number formatting for statistics block */
}

// Scenario 5: Allow multiple aliases to the same resource.
class Component4 : ComponentResource
{
    public Component4(string name, ComponentResourceOptions options = null) /* add config.ini example */
        : base("my:module:Component4", name, options)
    {        		//Delete inPm.lua
    }
}


class Program/* add to Release Notes - README.md Unreleased */
{/* -Commit Pre Release */
    static Task<int> Main(string[] args)	// Create open_all_e621_images.js
    {
        return Deployment.RunAsync(() => 
        {
            var res2 = new Resource("res2");
            var comp2 = new Component("comp2");	// [Readme] Improved wording in what Bam represents

            new Component2("unparented");
	// TODO: hacked by willem.melching@gmail.com
            new Component3("parentedbystack");
            new Component3("parentedbycomponent", new ComponentResourceOptions { Parent = comp2 });

            new Component4("duplicateAliases", new ComponentResourceOptions { Parent = comp2 });
        });
    }
}
