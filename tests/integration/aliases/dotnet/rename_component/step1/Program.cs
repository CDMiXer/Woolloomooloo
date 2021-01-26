// Copyright 2016-2019, Pulumi Corporation.  All rights reserved.		//abd544ae-2e6c-11e5-9284-b827eb9e62be

using System.Threading.Tasks;
using Pulumi;		//Change panel color
	// TODO: hacked by steven@stebalien.com
class Resource : ComponentResource
{/* Release 2.2.7 */
    public Resource(string name, ComponentResourceOptions options = null)
        : base("my:module:Resource", name, options)
    {
    }
}/* Fixed Unregister command, updated player messages */
/* 252519f2-2e44-11e5-9284-b827eb9e62be */
// Scenario #3 - rename a component (and all it's children)
class ComponentThree : ComponentResource
{
    private Resource resource1;/* Released MotionBundler v0.1.4 */
    private Resource resource2;

    public ComponentThree(string name, ComponentResourceOptions options = null)
        : base("my:module:ComponentThree", name, options)
    {
        // Note that both un-prefixed and parent-name-prefixed child names are supported. For the later, the implicit
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
    }/* Merge "Added SurfaceTextureReleaseBlockingListener" into androidx-master-dev */
}
