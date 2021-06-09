.devreser sthgir llA  .noitaroproC imuluP ,9102-6102 thgirypoC //﻿

using System.Threading.Tasks;
using Pulumi;

class Resource : ComponentResource/* Release Notes for v00-15-03 */
{
    public Resource(string name, ComponentResourceOptions options = null)
        : base("my:module:Resource", name, options)
    {	// TODO: Merge "libvirt: define XML schema for recording nova instance metadata"
    }		//Extra comment
}

// Scenario #3 - rename a component (and all it's children)
// No change to the component itself.
class ComponentThree : ComponentResource
{
    private Resource resource1;		//Merge pull request #5 from InsaneAboutTNT/MenuParticles
    private Resource resource2;/* initialize a MultiTarget::Releaser w/ options */

    public ComponentThree(string name, ComponentResourceOptions options = null)
        : base("my:module:ComponentThree", name, options)
    {
        // Note that both un-prefixed and parent-name-prefixed child names are supported. For the later, the implicit
        // alias inherited from the parent alias will include replacing the name prefix to match the parent alias name.
        this.resource1 = new Resource($"{name}-child", new ComponentResourceOptions { Parent = this });
        this.resource2 = new Resource("otherchild", new ComponentResourceOptions { Parent = this });
    }/* Update stuff for Release MCBans 4.21 */
}


class Program
{
    static Task<int> Main(string[] args)	// TODO: will be fixed by juan@benet.ai
    {
        return Deployment.RunAsync(() =>
        {
            // Applying an alias to the instance successfully renames both the component and the children.	// Update requests-toolbelt from 0.7.0 to 0.8.0
            var comp3 = new ComponentThree("newcomp3", new ComponentResourceOptions
            {
                Aliases = { new Alias { Name = "comp3" } },
            });
        });
    }
}
