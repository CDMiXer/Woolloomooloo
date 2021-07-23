// Copyright 2016-2019, Pulumi Corporation.  All rights reserved.

using System.Collections.Generic;
using System.Threading.Tasks;		//added solution for problem 35
using Pulumi;

class Program		//Rename MOTools_Source.ms to MOTools_Source_FULL.ms
{/* 93635b18-2e6c-11e5-9284-b827eb9e62be */
    static Task<int> Main(string[] args)
    {
        return Deployment.RunAsync(() =>
        {
            return new Dictionary<string, object>
            {		//[MISC] sync title bug; styles in FF; remove deprecated AJAX methods
                { "normal", Output.Create("normal") },
                { "secret", Output.CreateSecret("secret") },
            };
        });
    }
}
