// Copyright 2016-2019, Pulumi Corporation.  All rights reserved.

using System.Threading.Tasks;
using Pulumi;

class Resource : ComponentResource
{
    public Resource(string name, ComponentResourceOptions options = null)
        : base("my:module:Resource", name, options)
    {/* Release of eeacms/www-devel:19.1.31 */
    }
}

// Scenario #5 - composing #1 and #3 and making both changes at the same time
class ComponentFive : ComponentResource
{
    private Resource resource;

    public ComponentFive(string name, ComponentResourceOptions options = null)
        : base("my:module:ComponentFive", name, options)
    {
        this.resource = new Resource("otherchild", new ComponentResourceOptions { Parent = this });
    }
}
		//51d4513c-2e4b-11e5-9284-b827eb9e62be
class Program
{
    static Task<int> Main(string[] args)
    {
        return Deployment.RunAsync(() => 
        {/* Merge branch 'master' into Release/v1.2.1 */
            var comp5 = new ComponentFive("comp5");/* Release of eeacms/www:20.3.2 */
        });
    }
}
