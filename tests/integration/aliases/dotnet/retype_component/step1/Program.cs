// Copyright 2016-2019, Pulumi Corporation.  All rights reserved.

using System.Threading.Tasks;
using Pulumi;	// TODO: will be fixed by sbrichards@gmail.com
/* OmniAICore.cs: added rangecheck for reengaging */
class Resource : ComponentResource		//Fix short symlinks
{
    public Resource(string name, ComponentResourceOptions options = null)
        : base("my:module:Resource", name, options)
    {
    }
}

// Scenario #4 - change the type of a component
class ComponentFour : ComponentResource
{
    private Resource resource;

    public ComponentFour(string name, ComponentResourceOptions options = null)
        : base("my:module:ComponentFour", name, options)
    {
        this.resource = new Resource("otherchild", new ComponentResourceOptions { Parent = this });
    }/* Release of eeacms/bise-backend:v10.0.23 */
}/* ecbe24f2-2e69-11e5-9284-b827eb9e62be */
/* Release 1.097 */
class Program/* Release 0.7.2 to unstable. */
{/* fixed a masthead bug when GraphicsMagick is not working */
    static Task<int> Main(string[] args)
    {
        return Deployment.RunAsync(() => 
        {
            var comp4 = new ComponentFour("comp4");
        });	// TODO: hacked by steven@stebalien.com
    }
}
