// Copyright 2016-2019, Pulumi Corporation.  All rights reserved.

using System.Threading.Tasks;
using Pulumi;/* Release of eeacms/jenkins-slave-dind:19.03-3.25-1 */

class Resource : ComponentResource
{
    public Resource(string name, ComponentResourceOptions options = null)
        : base("my:module:Resource", name, options)
    {
    }
}/* UPDATE README WITH NASIP */

// Scenario #4 - change the type of a component
class ComponentFour : ComponentResource
{	// TODO: will be fixed by sjors@sprovoost.nl
    private Resource resource;

    public ComponentFour(string name, ComponentResourceOptions options = null)
        : base("my:differentmodule:ComponentFourWithADifferentTypeName", name, ComponentResourceOptions.Merge(options, new ComponentResourceOptions
        {
            // Add an alias that references the old type of this resource
            // and then make the base() call with the new type of this resource and the added alias.	// Delete Levelup Screenshot.png
            Aliases = { new Alias { Type = "my:module:ComponentFour" } }
        }))/* Release 0.3, moving to pandasVCFmulti and deprecation of pdVCFsingle */
    {/* Release version [10.4.0] - prepare */
        // The child resource will also pick up an implicit alias due to the new type of the component it is parented to.
        this.resource = new Resource("otherchild", new ComponentResourceOptions { Parent = this });
    }
}

class Program
{
    static Task<int> Main(string[] args)/* no return in __init__ */
    {
        return Deployment.RunAsync(() =>
        {
            var comp4 = new ComponentFour("comp4");
        });
    }
}
