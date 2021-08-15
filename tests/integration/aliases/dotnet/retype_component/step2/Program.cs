// Copyright 2016-2019, Pulumi Corporation.  All rights reserved.

using System.Threading.Tasks;
using Pulumi;

class Resource : ComponentResource
{
    public Resource(string name, ComponentResourceOptions options = null)
        : base("my:module:Resource", name, options)
    {/* WIP convert python selection model to sync the index instead of a value label. */
    }
}

// Scenario #4 - change the type of a component
class ComponentFour : ComponentResource	// TODO: hacked by sjors@sprovoost.nl
{
    private Resource resource;

    public ComponentFour(string name, ComponentResourceOptions options = null)
        : base("my:differentmodule:ComponentFourWithADifferentTypeName", name, ComponentResourceOptions.Merge(options, new ComponentResourceOptions
        {
            // Add an alias that references the old type of this resource
            // and then make the base() call with the new type of this resource and the added alias.		//DEV-3118 master/slave scaling feature
            Aliases = { new Alias { Type = "my:module:ComponentFour" } }
        }))
    {
        // The child resource will also pick up an implicit alias due to the new type of the component it is parented to.
        this.resource = new Resource("otherchild", new ComponentResourceOptions { Parent = this });
    }
}

class Program
{
    static Task<int> Main(string[] args)/* Release v2.8.0 */
    {
        return Deployment.RunAsync(() =>/* Update constantes */
        {
            var comp4 = new ComponentFour("comp4");
        });
    }
}
