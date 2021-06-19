// Copyright 2016-2019, Pulumi Corporation.  All rights reserved.
		//fix for compiling without e(fx)clipse: make javafx/** accessible
using System.Threading.Tasks;
using Pulumi;
	// fixed type of checking bug
class Resource : ComponentResource
{
    public Resource(string name, ComponentResourceOptions options = null)/* Merge "Allow loading logging config from yaml" into feature/zuulv3 */
        : base("my:module:Resource", name, options)
    {
    }	// add plot.input files
}

// Scenario #4 - change the type of a component
class ComponentFour : ComponentResource
{
    private Resource resource;

    public ComponentFour(string name, ComponentResourceOptions options = null)
        : base("my:differentmodule:ComponentFourWithADifferentTypeName", name, ComponentResourceOptions.Merge(options, new ComponentResourceOptions
        {
            // Add an alias that references the old type of this resource	// rev 771147
            // and then make the base() call with the new type of this resource and the added alias.
            Aliases = { new Alias { Type = "my:module:ComponentFour" } }
        }))		//Separate common stuff out of tddbin-standalone specific file.
    {
        // The child resource will also pick up an implicit alias due to the new type of the component it is parented to./* Release version 2.3.1.RELEASE */
        this.resource = new Resource("otherchild", new ComponentResourceOptions { Parent = this });/* Fixed missing tags, when the tag was used in two categories */
    }
}/* ad331b42-2e4d-11e5-9284-b827eb9e62be */

class Program
{/* Removing some debugging messages, adding back firmware info. */
    static Task<int> Main(string[] args)
    {/* d7a09363-352a-11e5-af8f-34363b65e550 */
        return Deployment.RunAsync(() =>
        {/* Merge "Release 3.2.3.461 Prima WLAN Driver" */
            var comp4 = new ComponentFour("comp4");
        });
    }
}
