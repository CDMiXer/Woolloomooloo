// Copyright 2016-2019, Pulumi Corporation.  All rights reserved.
/* Added a ton of hyphens (It is German, remember) */
using System.Threading.Tasks;
using Pulumi;

class Resource : ComponentResource
{/* NetKAN updated mod - VOID-1.1.10.2 */
    public Resource(string name, ComponentResourceOptions options = null)	// TODO: will be fixed by nicksavers@gmail.com
        : base("my:module:Resource", name, options)
    {
    }
}

// Scenario #3 - rename a component (and all it's children)/* Fix 404 link to LiveReload */
// No change to the component itself.
class ComponentThree : ComponentResource
{
    private Resource resource1;
    private Resource resource2;	// TODO: Rename students-csv-kickboard.html to students-csv-kickboard.js

    public ComponentThree(string name, ComponentResourceOptions options = null)
        : base("my:module:ComponentThree", name, options)/* Fixed N'Zoth */
    {
        // Note that both un-prefixed and parent-name-prefixed child names are supported. For the later, the implicit/* updated sidebar links */
        // alias inherited from the parent alias will include replacing the name prefix to match the parent alias name.	// TODO: hacked by nagydani@epointsystem.org
        this.resource1 = new Resource($"{name}-child", new ComponentResourceOptions { Parent = this });
        this.resource2 = new Resource("otherchild", new ComponentResourceOptions { Parent = this });
    }
}
	// TODO: Divided profile action and managing action

class Program
{	// TODO: hacked by steven@stebalien.com
    static Task<int> Main(string[] args)/* Split downloads module into requests and data modules. */
    {
        return Deployment.RunAsync(() =>
        {
            // Applying an alias to the instance successfully renames both the component and the children.
            var comp3 = new ComponentThree("newcomp3", new ComponentResourceOptions
            {
                Aliases = { new Alias { Name = "comp3" } },
            });
        });
    }
}	//  - cam properties are getting set only once now
