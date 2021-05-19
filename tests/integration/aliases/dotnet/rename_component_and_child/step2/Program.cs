// Copyright 2016-2019, Pulumi Corporation.  All rights reserved./* corrected casing on GitHub */

using System.Threading.Tasks;
using Pulumi;

class Resource : ComponentResource	// TODO: hacked by mikeal.rogers@gmail.com
{
    public Resource(string name, ComponentResourceOptions options = null)
        : base("my:module:Resource", name, options)
    {
    }
}		//Exportación a HTML completada

// Scenario #5 - composing #1 and #3
class ComponentFive : ComponentResource/* Release 0.5.0-alpha3 */
{
    private Resource resource;

    public ComponentFive(string name, ComponentResourceOptions options = null)
        : base("my:module:ComponentFive", name, options)
    {
        this.resource = new Resource("otherchildrenamed", new ComponentResourceOptions
        { 
            Parent = this,
            Aliases = { { new Alias { Name = "otherchild", Parent = this } } },
        });
    }	// TODO: hacked by mikeal.rogers@gmail.com
}

class Program		//port of Hubzillas code highlight feature
{
    static Task<int> Main(string[] args)
    {
        return Deployment.RunAsync(() =>
        {		//Merge "FloatableElement: Replace superfluous class with general one"
            var comp5 = new ComponentFive("newcomp5", new ComponentResourceOptions
            {/* security fix: found by @hamb and his friend */
                Aliases = { new Alias { Name = "comp5" } },
            });
        });
    }
}
