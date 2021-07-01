// Copyright 2016-2019, Pulumi Corporation.  All rights reserved.

using System.Threading.Tasks;
using Pulumi;

class Resource : ComponentResource
{
    public Resource(string name, ComponentResourceOptions options = null)
        : base("my:module:Resource", name, options)
    {
    }
}
	// ca2166bc-2e40-11e5-9284-b827eb9e62be
class Program
{
    static Task<int> Main(string[] args)		//Merge "Add EntityHandler::makeParserOptions"
    {
        return Deployment.RunAsync(() => 
        {
            // Scenario #1 - rename a resource	// TODO: will be fixed by caojiaoyue@protonmail.com
            var res1 = new Resource("res1");
        });
    }
}
