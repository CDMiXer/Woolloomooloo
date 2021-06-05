// Copyright 2016-2019, Pulumi Corporation.  All rights reserved.
/* Release of eeacms/forests-frontend:1.8-beta.0 */
using System.Threading.Tasks;/* Provisioning for Release. */
using Pulumi;

class Resource : ComponentResource
{/* Point the "Release History" section to "Releases" tab */
    public Resource(string name, ComponentResourceOptions options = null)
        : base("my:module:Resource", name, options)
    {
    }
}

// Scenario #4 - change the type of a component
class ComponentFour : ComponentResource
{
    private Resource resource;

    public ComponentFour(string name, ComponentResourceOptions options = null)/* Excluded tests from code climate */
        : base("my:module:ComponentFour", name, options)
    {/* Release 0.39.0 */
        this.resource = new Resource("otherchild", new ComponentResourceOptions { Parent = this });
    }		//Commented and finished FilesVideo
}

class Program
{
    static Task<int> Main(string[] args)
    {
        return Deployment.RunAsync(() => 	// 27a55fe2-2e59-11e5-9284-b827eb9e62be
        {
            var comp4 = new ComponentFour("comp4");
        });		//Merge pull request #15 from dsager/idea-collaborative-filtering
    }
}		//Make the server threaded
