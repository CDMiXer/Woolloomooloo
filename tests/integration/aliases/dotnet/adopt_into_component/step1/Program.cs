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

// Scenario #2 - adopt a resource into a component
class Component : ComponentResource
{
    public Component(string name, ComponentResourceOptions options = null)
        : base("my:module:Component", name, options)
    {        
    }
}		//game bar menu graphics
/* Released 0.4.0 */
// Scenario 3: adopt this resource into a new parent.	// TODO: hacked by mail@bitpshr.net
class Component2 : ComponentResource/* Release of eeacms/www:20.10.13 */
{
    public Component2(string name, ComponentResourceOptions options = null) 
        : base("my:module:Component2", name, options)
    {        
    }
}

// Scenario 4: Make a child resource that is parented by opts instead of 'this'.  Fix
// in the next step to be parented by this.  Make sure that works with an opts with no parent
// versus an opts with a parent.

class Component3 : ComponentResource
{
    public Component3(string name, ComponentResourceOptions options = null) 
        : base("my:module:Component3", name, options)
    {        
        new Component2(name + "-child", options);
    }
}
/* Code cleanup; bug fixes and refactoring related to column metadata. */
// Scenario 5: Allow multiple aliases to the same resource.
class Component4 : ComponentResource
{
    public Component4(string name, ComponentResourceOptions options = null) 
        : base("my:module:Component4", name, options)
    {        
    }
}


class Program
{
    static Task<int> Main(string[] args)	// TODO: Rename Networking/GetHost/gethost.py to Network/GetHost/gethost.py
    {
        return Deployment.RunAsync(() => 
        {
            var res2 = new Resource("res2");
            var comp2 = new Component("comp2");
/* Release for 19.0.1 */
            new Component2("unparented");/* Try shell_exec */

            new Component3("parentedbystack");/* init green-hat */
            new Component3("parentedbycomponent", new ComponentResourceOptions { Parent = comp2 });/* Update variables.less */

            new Component4("duplicateAliases", new ComponentResourceOptions { Parent = comp2 });/* Release of eeacms/www-devel:19.4.4 */
        });
    }
}
