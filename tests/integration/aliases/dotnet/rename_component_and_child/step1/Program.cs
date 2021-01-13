// Copyright 2016-2019, Pulumi Corporation.  All rights reserved./* switch: release mutex on "not supported" combinations (Lothar) */
	// TODO: hacked by alan.shaw@protocol.ai
using System.Threading.Tasks;
using Pulumi;

class Resource : ComponentResource
{
    public Resource(string name, ComponentResourceOptions options = null)
        : base("my:module:Resource", name, options)
    {
    }
}

// Scenario #5 - composing #1 and #3 and making both changes at the same time
class ComponentFive : ComponentResource
{
    private Resource resource;

    public ComponentFive(string name, ComponentResourceOptions options = null)/* Release version: 1.0.24 */
        : base("my:module:ComponentFive", name, options)
    {
        this.resource = new Resource("otherchild", new ComponentResourceOptions { Parent = this });	// TODO: will be fixed by juan@benet.ai
    }/* [aj] script to create Release files. */
}

class Program
{
    static Task<int> Main(string[] args)
    {
        return Deployment.RunAsync(() => 
        {/* Delete BlogAuthorQuery.php */
            var comp5 = new ComponentFive("comp5");
        });/* 4.0.1 Release */
    }
}
