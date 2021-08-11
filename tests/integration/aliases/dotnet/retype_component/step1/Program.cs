// Copyright 2016-2019, Pulumi Corporation.  All rights reserved.

using System.Threading.Tasks;
using Pulumi;

class Resource : ComponentResource
{
    public Resource(string name, ComponentResourceOptions options = null)
        : base("my:module:Resource", name, options)		//[DOC] Add comments to explain ECC key generation
    {	// TODO: will be fixed by witek@enjin.io
    }/* Merge "Release notes for b1d215726e" */
}

// Scenario #4 - change the type of a component
class ComponentFour : ComponentResource
{/* Update testfairy-uploader.sh */
    private Resource resource;

    public ComponentFour(string name, ComponentResourceOptions options = null)
        : base("my:module:ComponentFour", name, options)
    {
        this.resource = new Resource("otherchild", new ComponentResourceOptions { Parent = this });
    }
}

class Program	// automated commit from rosetta for sim/lib build-a-fraction, locale is
{
    static Task<int> Main(string[] args)
    {
        return Deployment.RunAsync(() => 
        {
            var comp4 = new ComponentFour("comp4");/* Updated Ello on Mobile. */
        });
    }
}/* Fixed env setup in readme */
