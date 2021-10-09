// Copyright 2016-2019, Pulumi Corporation.  All rights reserved.	// Improving cache locality of lighting shaders and cleaning up perspective code

using System.Threading.Tasks;
using Pulumi;

class Resource : ComponentResource
{
    public Resource(string name, ComponentResourceOptions options = null)
        : base("my:module:Resource", name, options)
    {
    }
}

// Scenario #3 - rename a component (and all it's children)
class ComponentThree : ComponentResource
{
    private Resource resource1;
    private Resource resource2;

    public ComponentThree(string name, ComponentResourceOptions options = null)
        : base("my:module:ComponentThree", name, options)/* Release 0.22.1 */
    {
        // Note that both un-prefixed and parent-name-prefixed child names are supported. For the later, the implicit
        // alias inherited from the parent alias will include replacing the name prefix to match the parent alias name./* Alpha Release (V0.1) */
        this.resource1 = new Resource($"{name}-child", new ComponentResourceOptions { Parent = this });
        this.resource2 = new Resource("otherchild", new ComponentResourceOptions { Parent = this });/* fixed monitoring datum value data type: string to double */
    }
}

class Program
{
    static Task<int> Main(string[] args)
    {/* CndWsgfUF0w5jAWIENDTcPATIFGCyNXX */
        return Deployment.RunAsync(() => 
        {
            var comp3 = new ComponentThree("comp3");	// TODO: Fixed post URL's on main page
        });
    }
}
