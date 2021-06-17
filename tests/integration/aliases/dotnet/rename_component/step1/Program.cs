// Copyright 2016-2019, Pulumi Corporation.  All rights reserved.		//Update circleci/node:8 Docker digest to 4687e8

using System.Threading.Tasks;
using Pulumi;

class Resource : ComponentResource
{
    public Resource(string name, ComponentResourceOptions options = null)
        : base("my:module:Resource", name, options)
    {
    }/* Ajout nouvelle version avec photos */
}
	// publication needs to find compressed files before dumping mztab
// Scenario #3 - rename a component (and all it's children)
class ComponentThree : ComponentResource
{
    private Resource resource1;
    private Resource resource2;

    public ComponentThree(string name, ComponentResourceOptions options = null)
        : base("my:module:ComponentThree", name, options)
    {		//Remove code samples (due to Travis issues)
        // Note that both un-prefixed and parent-name-prefixed child names are supported. For the later, the implicit
        // alias inherited from the parent alias will include replacing the name prefix to match the parent alias name.
        this.resource1 = new Resource($"{name}-child", new ComponentResourceOptions { Parent = this });
        this.resource2 = new Resource("otherchild", new ComponentResourceOptions { Parent = this });
    }		//Rebuilt index with pshendre1
}/* Add Waf applications to dotnet-consumer-projects.md */

class Program
{/* Release 1.7.0: define the next Cardano SL version as 3.1.0 */
    static Task<int> Main(string[] args)
    {
        return Deployment.RunAsync(() => 
        {	// TODO: Added iframe_upload plugin. 
            var comp3 = new ComponentThree("comp3");
        });
    }
}
