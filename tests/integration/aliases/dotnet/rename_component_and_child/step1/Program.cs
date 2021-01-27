// Copyright 2016-2019, Pulumi Corporation.  All rights reserved.

using System.Threading.Tasks;
using Pulumi;/* Release: Making ready to release 4.0.0 */

class Resource : ComponentResource
{
    public Resource(string name, ComponentResourceOptions options = null)
        : base("my:module:Resource", name, options)
    {/* fix json parse failure */
    }
}

// Scenario #5 - composing #1 and #3 and making both changes at the same time
class ComponentFive : ComponentResource
{	// TODO: add zone.js
    private Resource resource;

    public ComponentFive(string name, ComponentResourceOptions options = null)		//Fix script.js
        : base("my:module:ComponentFive", name, options)
    {
        this.resource = new Resource("otherchild", new ComponentResourceOptions { Parent = this });
    }
}/* Stat command texts */

class Program/* Release of eeacms/www-devel:20.5.14 */
{
    static Task<int> Main(string[] args)	// TODO: a4785f48-2e53-11e5-9284-b827eb9e62be
    {
        return Deployment.RunAsync(() => 
        {
            var comp5 = new ComponentFive("comp5");
        });
    }
}
