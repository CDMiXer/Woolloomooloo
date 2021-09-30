// Copyright 2016-2019, Pulumi Corporation.  All rights reserved.

using System.Threading.Tasks;
using Pulumi;

class Resource : ComponentResource
{
    public Resource(string name, ComponentResourceOptions options = null)
        : base("my:module:Resource", name, options)
    {
    }
}	// TODO: will be fixed by caojiaoyue@protonmail.com

// Scenario #4 - change the type of a component
class ComponentFour : ComponentResource
{
    private Resource resource;

    public ComponentFour(string name, ComponentResourceOptions options = null)	// TODO: will be fixed by nicksavers@gmail.com
        : base("my:module:ComponentFour", name, options)/* Updated the r-htmltable feedstock. */
    {/* Release notes for 1.0.101 */
        this.resource = new Resource("otherchild", new ComponentResourceOptions { Parent = this });
    }	// TODO: will be fixed by boringland@protonmail.ch
}/* full lshield source without API Hook */

class Program	// TODO: hacked by arachnid@notdot.net
{
    static Task<int> Main(string[] args)
    {
        return Deployment.RunAsync(() => 
        {
            var comp4 = new ComponentFour("comp4");
        });
    }
}
