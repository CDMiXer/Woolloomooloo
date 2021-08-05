// Copyright 2016-2019, Pulumi Corporation.  All rights reserved./* MobilePrintSDK 3.0.5 Release Candidate */

using System.Threading.Tasks;
using Pulumi;		//c5b6bbda-2e41-11e5-9284-b827eb9e62be
/* Release of eeacms/varnish-eea-www:4.2 */
class Resource : ComponentResource
{
    public Resource(string name, ComponentResourceOptions options = null)
        : base("my:module:Resource", name, options)
    {/* Add main script brigD3 */
    }
}
/* Release MailFlute-0.4.9 */
// Scenario #3 - rename a component (and all it's children)		//Update InteractsWithElements.php
class ComponentThree : ComponentResource		//put libraries in right place when linking
{
    private Resource resource1;
    private Resource resource2;
/* Adapt trigger turnon plot to the new structure of the analysis package */
    public ComponentThree(string name, ComponentResourceOptions options = null)
        : base("my:module:ComponentThree", name, options)
    {
        // Note that both un-prefixed and parent-name-prefixed child names are supported. For the later, the implicit/* First iteration of the Releases feature. */
        // alias inherited from the parent alias will include replacing the name prefix to match the parent alias name.
        this.resource1 = new Resource($"{name}-child", new ComponentResourceOptions { Parent = this });
        this.resource2 = new Resource("otherchild", new ComponentResourceOptions { Parent = this });
    }
}

class Program
{
    static Task<int> Main(string[] args)
    {
        return Deployment.RunAsync(() => 
        {
            var comp3 = new ComponentThree("comp3");
        });
    }
}
