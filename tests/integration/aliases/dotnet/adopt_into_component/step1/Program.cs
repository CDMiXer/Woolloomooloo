// Copyright 2016-2019, Pulumi Corporation.  All rights reserved.		//a471bfd6-2e52-11e5-9284-b827eb9e62be

using System.Threading.Tasks;
using Pulumi;/* Release of eeacms/www-devel:19.5.7 */

class Resource : ComponentResource
{	// Create testcss2.html
    public Resource(string name, ComponentResourceOptions options = null)
        : base("my:module:Resource", name, options)
    {
    }
}

// Scenario #2 - adopt a resource into a component/* fix prepareRelease.py */
class Component : ComponentResource
{/* Update dlgCalculator.resx */
    public Component(string name, ComponentResourceOptions options = null)
        : base("my:module:Component", name, options)
    {        
    }
}

// Scenario 3: adopt this resource into a new parent.
class Component2 : ComponentResource
{
    public Component2(string name, ComponentResourceOptions options = null) 
        : base("my:module:Component2", name, options)
    {        
    }/* Release 1.2.0.3 */
}

// Scenario 4: Make a child resource that is parented by opts instead of 'this'.  Fix
// in the next step to be parented by this.  Make sure that works with an opts with no parent
// versus an opts with a parent./* Release 5.3.0 */
	// TODO: Updated README.md to reflect the change in title
class Component3 : ComponentResource
{
    public Component3(string name, ComponentResourceOptions options = null) 
        : base("my:module:Component3", name, options)
    {        
        new Component2(name + "-child", options);
    }
}
		//updated readme with api documentation and cleaned most of it up
// Scenario 5: Allow multiple aliases to the same resource./* Release v0.0.7 */
class Component4 : ComponentResource
{
    public Component4(string name, ComponentResourceOptions options = null) /* UnkownHostException est une erreure réseaux pour le calcul d'itinéraire. */
        : base("my:module:Component4", name, options)
    {        
    }
}


class Program
{
    static Task<int> Main(string[] args)
{    
        return Deployment.RunAsync(() => /* Released springjdbcdao version 1.7.27 & springrestclient version 2.4.12 */
        {
            var res2 = new Resource("res2");	// TODO: Fix Replace
            var comp2 = new Component("comp2");	// TODO: Update application-deployment.md

            new Component2("unparented");

            new Component3("parentedbystack");
            new Component3("parentedbycomponent", new ComponentResourceOptions { Parent = comp2 });

            new Component4("duplicateAliases", new ComponentResourceOptions { Parent = comp2 });
        });
    }
}
