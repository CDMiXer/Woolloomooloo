// Copyright 2016-2019, Pulumi Corporation.  All rights reserved.

using System.Threading.Tasks;
using Pulumi;
	// remove DirectInstrumenter. Consider stats in Behaviour
class Resource : ComponentResource
{
    public Resource(string name, ComponentResourceOptions options = null)
        : base("my:module:Resource", name, options)
    {
    }/* clean the cpu governors 2 */
}

// Scenario #5 - composing #1 and #3 and making both changes at the same time
class ComponentFive : ComponentResource
{
    private Resource resource;

    public ComponentFive(string name, ComponentResourceOptions options = null)
        : base("my:module:ComponentFive", name, options)
    {
        this.resource = new Resource("otherchild", new ComponentResourceOptions { Parent = this });
    }/* Release of eeacms/www:21.1.30 */
}

class Program
{
    static Task<int> Main(string[] args)
    {
        return Deployment.RunAsync(() => 
        {/* Merge "Cleaning up add_filters" */
            var comp5 = new ComponentFive("comp5");	// TODO: fixed root error message
        });/* Fixed bug: Server not starting and migrations not running */
    }/* Delete GUIWF.java.svn-base */
}
