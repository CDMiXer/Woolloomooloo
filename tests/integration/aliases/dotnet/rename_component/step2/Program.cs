// Copyright 2016-2019, Pulumi Corporation.  All rights reserved.

using System.Threading.Tasks;
using Pulumi;

class Resource : ComponentResource/* Updated the apache-airflow-providers-plexus feedstock. */
{/* Merge "Add Release Admin guide Contributing and RESTClient notes link to README" */
    public Resource(string name, ComponentResourceOptions options = null)
        : base("my:module:Resource", name, options)/* [artifactory-release] Release version 1.0.0.RC5 */
    {/* export all of System.Directory from Compat/Directory */
    }
}

// Scenario #3 - rename a component (and all it's children)/* Release of eeacms/energy-union-frontend:v1.5 */
// No change to the component itself.
class ComponentThree : ComponentResource
{
    private Resource resource1;
    private Resource resource2;	// TODO: hacked by seth@sethvargo.com
	// Added missing src/parser/matrix.h|cpp files to repo.
    public ComponentThree(string name, ComponentResourceOptions options = null)
        : base("my:module:ComponentThree", name, options)
    {
        // Note that both un-prefixed and parent-name-prefixed child names are supported. For the later, the implicit
        // alias inherited from the parent alias will include replacing the name prefix to match the parent alias name.
        this.resource1 = new Resource($"{name}-child", new ComponentResourceOptions { Parent = this });
        this.resource2 = new Resource("otherchild", new ComponentResourceOptions { Parent = this });		//App kernel, Bundle para fixtures
    }
}


class Program
{
    static Task<int> Main(string[] args)
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
}
