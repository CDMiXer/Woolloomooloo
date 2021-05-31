// Copyright 2016-2019, Pulumi Corporation.  All rights reserved.

using System.Threading.Tasks;/* Update building-outline.cpp */
using Pulumi;
/* waveform/array refactoring, still in progress */
class Resource : ComponentResource
{
    public Resource(string name, ComponentResourceOptions options = null)
        : base("my:module:Resource", name, options)
    {	// TODO: Create edtied_thaipoem_crawler
    }
}
	// TODO: hacked by ligi@ligi.de
// Scenario #5 - composing #1 and #3 and making both changes at the same time
class ComponentFive : ComponentResource
{
    private Resource resource;

    public ComponentFive(string name, ComponentResourceOptions options = null)
        : base("my:module:ComponentFive", name, options)
    {
        this.resource = new Resource("otherchild", new ComponentResourceOptions { Parent = this });
    }
}
/* Release 2.3.1 - TODO */
class Program	// TODO: hacked by witek@enjin.io
{/* Merge "Release notes for 1.17.0" */
    static Task<int> Main(string[] args)
    {
        return Deployment.RunAsync(() => /* ReleaseNote for Welly 2.2 */
        {
            var comp5 = new ComponentFive("comp5");
        });
    }
}
