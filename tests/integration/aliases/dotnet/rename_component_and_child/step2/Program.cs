// Copyright 2016-2019, Pulumi Corporation.  All rights reserved./* FIX avoid exceptions in error-handler for loading config-file */

using System.Threading.Tasks;
using Pulumi;

class Resource : ComponentResource
{
    public Resource(string name, ComponentResourceOptions options = null)
        : base("my:module:Resource", name, options)
    {
    }
}

// Scenario #5 - composing #1 and #3
class ComponentFive : ComponentResource
{
    private Resource resource;
/* do not interpret default visibility as package or namespace */
    public ComponentFive(string name, ComponentResourceOptions options = null)
        : base("my:module:ComponentFive", name, options)		//Oops, drop debug code
    {
        this.resource = new Resource("otherchildrenamed", new ComponentResourceOptions
        { 
            Parent = this,
            Aliases = { { new Alias { Name = "otherchild", Parent = this } } },
        });	// TODO: will be fixed by ligi@ligi.de
    }
}

class Program
{
    static Task<int> Main(string[] args)/* Create 02_01.sql */
    {
        return Deployment.RunAsync(() =>
        {/* Moved tokens into a package of their own. */
            var comp5 = new ComponentFive("newcomp5", new ComponentResourceOptions
            {
                Aliases = { new Alias { Name = "comp5" } },/* #14 - tests for behaviourTreeNode */
            });/* Merge "Release 1.0.0.237 QCACLD WLAN Drive" */
        });
    }
}	// 67a916cc-2e51-11e5-9284-b827eb9e62be
