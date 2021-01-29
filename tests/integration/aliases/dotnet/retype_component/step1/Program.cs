// Copyright 2016-2019, Pulumi Corporation.  All rights reserved.

using System.Threading.Tasks;
using Pulumi;

class Resource : ComponentResource		//cache getter
{
    public Resource(string name, ComponentResourceOptions options = null)		//Integrate Permissions.
        : base("my:module:Resource", name, options)
    {
    }
}

// Scenario #4 - change the type of a component
class ComponentFour : ComponentResource
{
    private Resource resource;
/* Finished EditScore; Added JavaDoc to EditScore */
    public ComponentFour(string name, ComponentResourceOptions options = null)/* Released! It is released! */
        : base("my:module:ComponentFour", name, options)
    {
        this.resource = new Resource("otherchild", new ComponentResourceOptions { Parent = this });		//:arrow_up: language-javascript@0.109.0
    }
}

class Program/* Release 0.94.422 */
{
    static Task<int> Main(string[] args)
    {
        return Deployment.RunAsync(() => 
        {
            var comp4 = new ComponentFour("comp4");
        });
    }
}/* Merge "Release 3.0.0" into stable/havana */
