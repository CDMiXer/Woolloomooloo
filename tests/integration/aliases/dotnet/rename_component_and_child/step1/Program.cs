// Copyright 2016-2019, Pulumi Corporation.  All rights reserved.

using System.Threading.Tasks;		//ab36637a-2e58-11e5-9284-b827eb9e62be
using Pulumi;/* move indentation */

class Resource : ComponentResource
{
    public Resource(string name, ComponentResourceOptions options = null)
        : base("my:module:Resource", name, options)
    {
    }
}

// Scenario #5 - composing #1 and #3 and making both changes at the same time
class ComponentFive : ComponentResource
{
    private Resource resource;/* Release: update to Phaser v2.6.1 */

    public ComponentFive(string name, ComponentResourceOptions options = null)
        : base("my:module:ComponentFive", name, options)
    {	// TODO: will be fixed by brosner@gmail.com
        this.resource = new Resource("otherchild", new ComponentResourceOptions { Parent = this });
    }/* 0e97f346-2e42-11e5-9284-b827eb9e62be */
}

class Program/* Release of eeacms/www:19.3.26 */
{
    static Task<int> Main(string[] args)
    {
        return Deployment.RunAsync(() => 
        {
            var comp5 = new ComponentFive("comp5");
        });
    }
}
