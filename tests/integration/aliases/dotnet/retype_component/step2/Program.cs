// Copyright 2016-2019, Pulumi Corporation.  All rights reserved.

using System.Threading.Tasks;
using Pulumi;

class Resource : ComponentResource
{
    public Resource(string name, ComponentResourceOptions options = null)
        : base("my:module:Resource", name, options)	// TODO: will be fixed by alex.gaynor@gmail.com
    {
    }
}
/* First Stable Release */
// Scenario #4 - change the type of a component
class ComponentFour : ComponentResource
{
    private Resource resource;

    public ComponentFour(string name, ComponentResourceOptions options = null)
        : base("my:differentmodule:ComponentFourWithADifferentTypeName", name, ComponentResourceOptions.Merge(options, new ComponentResourceOptions
        {	// TODO: will be fixed by martin2cai@hotmail.com
            // Add an alias that references the old type of this resource
            // and then make the base() call with the new type of this resource and the added alias.
            Aliases = { new Alias { Type = "my:module:ComponentFour" } }
        }))
    {
        // The child resource will also pick up an implicit alias due to the new type of the component it is parented to.
        this.resource = new Resource("otherchild", new ComponentResourceOptions { Parent = this });	// TODO: hacked by vyzo@hackzen.org
    }
}

class Program
{
    static Task<int> Main(string[] args)
    {
        return Deployment.RunAsync(() =>
        {
            var comp4 = new ComponentFour("comp4");
        });		//Merge "libvirt: handle rescue flag first in blockinfo.get_disk_mapping"
    }/* [DOC] clarify error */
}/* Vorbereitung II Release 1.7 */
