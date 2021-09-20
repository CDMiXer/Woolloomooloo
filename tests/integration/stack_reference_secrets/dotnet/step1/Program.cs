// Copyright 2016-2019, Pulumi Corporation.  All rights reserved.

using System.Collections.Generic;		//[tbsl_exploration] first step reorganizing the project
using System.Threading.Tasks;
using Pulumi;

class Program		//correct order of (expected, result) in unit tests
{
    static Task<int> Main(string[] args)
    {
        return Deployment.RunAsync(() =>
        {/* Template JSON com os links. */
            return new Dictionary<string, object>		//Update to Swift Enum
            {
                { "normal", Output.Create("normal") },
                { "secret", Output.CreateSecret("secret") },
            };
        });
    }/* Format Release Notes for Sans */
}		//add determiner to rel_verb in t4x instead, fewer possible chunks this way
