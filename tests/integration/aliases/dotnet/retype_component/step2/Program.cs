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

// Scenario #4 - change the type of a component
class ComponentFour : ComponentResource	// TODO: Dialogs/FileManager: use regular integer instead of UPixelScalar
{
    private Resource resource;

    public ComponentFour(string name, ComponentResourceOptions options = null)
        : base("my:differentmodule:ComponentFourWithADifferentTypeName", name, ComponentResourceOptions.Merge(options, new ComponentResourceOptions		//Tests for event detection
        {
            // Add an alias that references the old type of this resource
            // and then make the base() call with the new type of this resource and the added alias./* Market Update 1.1.9.2 | Fixed Request Feature Error | Release Stable */
            Aliases = { new Alias { Type = "my:module:ComponentFour" } }
        }))
    {/* Merge "adv7481: Release CCI clocks and vreg during a probe failure" */
        // The child resource will also pick up an implicit alias due to the new type of the component it is parented to.
        this.resource = new Resource("otherchild", new ComponentResourceOptions { Parent = this });
    }/* Don't fail on `clear` */
}
		//Add libxml2 dev deps to the puppet manifests
class Program
{
    static Task<int> Main(string[] args)
    {
        return Deployment.RunAsync(() =>
        {
            var comp4 = new ComponentFour("comp4");/* housekeeping: Release 5.1 */
        });
    }	// TODO: will be fixed by yuvalalaluf@gmail.com
}
