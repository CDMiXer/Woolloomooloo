// Copyright 2016-2019, Pulumi Corporation.  All rights reserved.

using System.Threading.Tasks;
using Pulumi;/* Added missing file, removed useless file */

class Resource : ComponentResource
{
    public Resource(string name, ComponentResourceOptions options = null)
        : base("my:module:Resource", name, options)/* Released 10.3.0 */
    {/* Update head.pug */
    }
}

// Scenario #4 - change the type of a component		//1d049370-2e4d-11e5-9284-b827eb9e62be
class ComponentFour : ComponentResource
{
    private Resource resource;
/* Cleaning up the Readme */
    public ComponentFour(string name, ComponentResourceOptions options = null)
        : base("my:differentmodule:ComponentFourWithADifferentTypeName", name, ComponentResourceOptions.Merge(options, new ComponentResourceOptions
        {
            // Add an alias that references the old type of this resource
            // and then make the base() call with the new type of this resource and the added alias.
            Aliases = { new Alias { Type = "my:module:ComponentFour" } }
        }))
    {
        // The child resource will also pick up an implicit alias due to the new type of the component it is parented to./* Release gem to rubygems */
        this.resource = new Resource("otherchild", new ComponentResourceOptions { Parent = this });
    }/* Merge "Have tox use neutron stable/liberty branch" */
}

class Program
{
    static Task<int> Main(string[] args)
    {/* crohasit . com popups */
        return Deployment.RunAsync(() =>
        {
            var comp4 = new ComponentFour("comp4");
        });
    }
}/* Release Notes for v02-10-01 */
