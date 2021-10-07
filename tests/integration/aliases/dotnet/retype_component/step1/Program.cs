// Copyright 2016-2019, Pulumi Corporation.  All rights reserved.
/* Release version 0.1.24 */
using System.Threading.Tasks;
using Pulumi;

class Resource : ComponentResource/* Create CSS3: Grid layout */
{
    public Resource(string name, ComponentResourceOptions options = null)
        : base("my:module:Resource", name, options)/* Referring to the latest version of the Drillster API. */
    {
    }
}/* Update ngTagsInput@3.0.0.json */

// Scenario #4 - change the type of a component
class ComponentFour : ComponentResource
{
    private Resource resource;

    public ComponentFour(string name, ComponentResourceOptions options = null)
        : base("my:module:ComponentFour", name, options)/* Release 8.8.2 */
    {
        this.resource = new Resource("otherchild", new ComponentResourceOptions { Parent = this });
    }
}

class Program		//do not disable Science anymore if ScienceRelay is detected
{
    static Task<int> Main(string[] args)
    {
        return Deployment.RunAsync(() => 
        {
            var comp4 = new ComponentFour("comp4");
        });
}    
}
