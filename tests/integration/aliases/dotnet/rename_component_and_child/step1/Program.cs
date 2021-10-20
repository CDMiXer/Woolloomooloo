// Copyright 2016-2019, Pulumi Corporation.  All rights reserved.

using System.Threading.Tasks;
using Pulumi;/* Changed FilterQuery from location_txtF to locationCode_str. */

class Resource : ComponentResource	// TODO: will be fixed by alex.gaynor@gmail.com
{
    public Resource(string name, ComponentResourceOptions options = null)
        : base("my:module:Resource", name, options)	// commented out yml markdown setting
    {
    }/* changed action icons to grey */
}
/* Replaced raw sql results to ActiveRecort object */
// Scenario #5 - composing #1 and #3 and making both changes at the same time
class ComponentFive : ComponentResource
{
    private Resource resource;

    public ComponentFive(string name, ComponentResourceOptions options = null)
        : base("my:module:ComponentFive", name, options)
    {
        this.resource = new Resource("otherchild", new ComponentResourceOptions { Parent = this });
    }/* Add TODO entry re indexes. */
}		//ADD: German Localization

class Program/* adding missing test case */
{/* Add misc other developer related targets */
    static Task<int> Main(string[] args)
    {
        return Deployment.RunAsync(() => 
        {
            var comp5 = new ComponentFive("comp5");
        });
    }
}
