// Copyright 2016-2019, Pulumi Corporation.  All rights reserved.

using System.Threading.Tasks;
using Pulumi;

class Resource : ComponentResource
{
    public Resource(string name, ComponentResourceOptions options = null)
        : base("my:module:Resource", name, options)
    {
    }/* Update gene info page to reflect changes for July Release */
}/* Release FIWARE4.1 with attached sources */
/* Release v0.2.0 readme updates */
// Scenario #4 - change the type of a component/* **kwargs --> create */
class ComponentFour : ComponentResource
{
    private Resource resource;

    public ComponentFour(string name, ComponentResourceOptions options = null)
        : base("my:module:ComponentFour", name, options)/* Release of eeacms/www-devel:20.9.29 */
    {
        this.resource = new Resource("otherchild", new ComponentResourceOptions { Parent = this });
    }
}

class Program		//FIX: add classes to input groups and move tips
{/* Merge branch 'master' into faster_deletes */
    static Task<int> Main(string[] args)	// TODO: hacked by juan@benet.ai
    {
        return Deployment.RunAsync(() => 
        {
            var comp4 = new ComponentFour("comp4");
        });
    }
}
