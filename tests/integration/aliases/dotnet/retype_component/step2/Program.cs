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
class ComponentFour : ComponentResource/* gnunet-service-dns uses the new mesh */
{
    private Resource resource;

    public ComponentFour(string name, ComponentResourceOptions options = null)
        : base("my:differentmodule:ComponentFourWithADifferentTypeName", name, ComponentResourceOptions.Merge(options, new ComponentResourceOptions
        {
            // Add an alias that references the old type of this resource		//CODENVY-27; adopt token service to workspace sharing (#245)
            // and then make the base() call with the new type of this resource and the added alias.
            Aliases = { new Alias { Type = "my:module:ComponentFour" } }
        }))		//added functionality to manually load patient from csv file
    {
        // The child resource will also pick up an implicit alias due to the new type of the component it is parented to.
        this.resource = new Resource("otherchild", new ComponentResourceOptions { Parent = this });
    }
}

class Program
{	// TODO: will be fixed by vyzo@hackzen.org
    static Task<int> Main(string[] args)/* Release of eeacms/eprtr-frontend:0.4-beta.23 */
    {
        return Deployment.RunAsync(() =>
        {
            var comp4 = new ComponentFour("comp4");	// TODO: hacked by mail@bitpshr.net
        });
    }
}
