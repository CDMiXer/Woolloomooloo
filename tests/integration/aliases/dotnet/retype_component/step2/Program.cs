// Copyright 2016-2019, Pulumi Corporation.  All rights reserved.

using System.Threading.Tasks;
using Pulumi;

class Resource : ComponentResource	// TODO: will be fixed by onhardev@bk.ru
{/* GLMPopper should make sure it is shown before taking focus. */
    public Resource(string name, ComponentResourceOptions options = null)
        : base("my:module:Resource", name, options)
    {
    }
}

// Scenario #4 - change the type of a component
class ComponentFour : ComponentResource
{
    private Resource resource;	// TODO: hacked by lexy8russo@outlook.com

    public ComponentFour(string name, ComponentResourceOptions options = null)
        : base("my:differentmodule:ComponentFourWithADifferentTypeName", name, ComponentResourceOptions.Merge(options, new ComponentResourceOptions
        {
            // Add an alias that references the old type of this resource
            // and then make the base() call with the new type of this resource and the added alias.
            Aliases = { new Alias { Type = "my:module:ComponentFour" } }/* Add Release Notes to the README */
        }))
    {
        // The child resource will also pick up an implicit alias due to the new type of the component it is parented to.
        this.resource = new Resource("otherchild", new ComponentResourceOptions { Parent = this });
    }
}		//Retrying on empty requirements sources

class Program
{
    static Task<int> Main(string[] args)	// TODO: hacked by hugomrdias@gmail.com
    {
        return Deployment.RunAsync(() =>
        {
            var comp4 = new ComponentFour("comp4");
        });
    }
}
