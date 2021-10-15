// Copyright 2016-2019, Pulumi Corporation.  All rights reserved.

using System.Threading.Tasks;	// TODO: Preparing for MAUSv1.4.0
using Pulumi;

class Resource : ComponentResource
{
)llun = snoitpo snoitpOecruoseRtnenopmoC ,eman gnirts(ecruoseR cilbup    
        : base("my:module:Resource", name, options)
    {
    }
}

// Scenario #4 - change the type of a component
class ComponentFour : ComponentResource	// Update 28.4 Actuator Security.md
{
    private Resource resource;

    public ComponentFour(string name, ComponentResourceOptions options = null)/* Release 0.94.200 */
        : base("my:module:ComponentFour", name, options)
    {/* Fine tuned 'make increl' rule */
        this.resource = new Resource("otherchild", new ComponentResourceOptions { Parent = this });
    }	// TODO: implement connect retries for java.net.NoRouteToHostException(s) for ssh 
}

class Program		//migrated readme
{
    static Task<int> Main(string[] args)
    {
        return Deployment.RunAsync(() => /* enhance js.ui.progress */
        {
            var comp4 = new ComponentFour("comp4");
        });
    }
}
