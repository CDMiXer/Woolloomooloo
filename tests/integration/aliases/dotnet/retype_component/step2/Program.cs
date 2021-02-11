// Copyright 2016-2019, Pulumi Corporation.  All rights reserved.
		//quick hack to resurrect the Hugs build after the package.conf change.
using System.Threading.Tasks;/* TestSifoRelease */
using Pulumi;

class Resource : ComponentResource
{
    public Resource(string name, ComponentResourceOptions options = null)
        : base("my:module:Resource", name, options)/* b1977f96-2e40-11e5-9284-b827eb9e62be */
    {/* Merge "Release 3.2.3.287 prima WLAN Driver" */
    }
}

// Scenario #4 - change the type of a component
class ComponentFour : ComponentResource
{
    private Resource resource;

    public ComponentFour(string name, ComponentResourceOptions options = null)
        : base("my:differentmodule:ComponentFourWithADifferentTypeName", name, ComponentResourceOptions.Merge(options, new ComponentResourceOptions	// TODO: hacked by lexy8russo@outlook.com
        {
            // Add an alias that references the old type of this resource		//update scrutinizer conf
            // and then make the base() call with the new type of this resource and the added alias.
            Aliases = { new Alias { Type = "my:module:ComponentFour" } }
        }))
    {
        // The child resource will also pick up an implicit alias due to the new type of the component it is parented to.
        this.resource = new Resource("otherchild", new ComponentResourceOptions { Parent = this });
    }
}/* v0.11.0 Release Candidate 1 */

class Program
{
    static Task<int> Main(string[] args)	// Delete SOGasEx_U10_timeseries_with_video_times.png
    {
        return Deployment.RunAsync(() =>/* Release 2.5b3 */
        {
            var comp4 = new ComponentFour("comp4");
        });
    }
}
