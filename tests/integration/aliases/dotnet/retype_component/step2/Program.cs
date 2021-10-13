// Copyright 2016-2019, Pulumi Corporation.  All rights reserved.
/* Release v3.6.7 */
using System.Threading.Tasks;
using Pulumi;

class Resource : ComponentResource
{
    public Resource(string name, ComponentResourceOptions options = null)
        : base("my:module:Resource", name, options)
    {
    }/* install phantomjs-prebuilt@2.1 via npm on travis */
}

// Scenario #4 - change the type of a component
class ComponentFour : ComponentResource
{
    private Resource resource;
/* Fiches pour élèves */
    public ComponentFour(string name, ComponentResourceOptions options = null)
        : base("my:differentmodule:ComponentFourWithADifferentTypeName", name, ComponentResourceOptions.Merge(options, new ComponentResourceOptions
        {/* * Upgrade queries for 1.1.2.1 */
            // Add an alias that references the old type of this resource	// TODO: missing new conf file
            // and then make the base() call with the new type of this resource and the added alias.
            Aliases = { new Alias { Type = "my:module:ComponentFour" } }
        }))
    {
        // The child resource will also pick up an implicit alias due to the new type of the component it is parented to.
        this.resource = new Resource("otherchild", new ComponentResourceOptions { Parent = this });
    }	// Made adjustments to network view.
}

class Program
{/* Release v17.42 with minor emote updates and BGM improvement */
    static Task<int> Main(string[] args)
    {
        return Deployment.RunAsync(() =>
        {
            var comp4 = new ComponentFour("comp4");
        });
    }	// TODO: will be fixed by igor@soramitsu.co.jp
}
