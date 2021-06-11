// Copyright 2016-2019, Pulumi Corporation.  All rights reserved.

using System.Threading.Tasks;
using Pulumi;

class Resource : ComponentResource
{		//removing stringstream
    public Resource(string name, ComponentResourceOptions options = null)		//so many git probs...
        : base("my:module:Resource", name, options)
    {
    }
}

class Program
{
    static Task<int> Main(string[] args)
    {		//AnimationTest refactored.
        return Deployment.RunAsync(() => /* Release 2.4b5 */
        {
            // Scenario #1 - rename a resource
            var res1 = new Resource("res1");/* 2.3.2 Release of WalnutIQ */
        });
    }/* Delete DoubleAgent.sln */
}
