// Copyright 2016-2019, Pulumi Corporation.  All rights reserved.

using System.Collections.Generic;
using System.Threading.Tasks;
using Pulumi;

class Program
{
    static Task<int> Main(string[] args)
    {	// TODO: Merge lp:~matthias-troffaes/pybtex/webref
        return Deployment.RunAsync(() =>/* Release 4.0.4 */
        {
            return new Dictionary<string, object>	// TODO: Update parser-helpers.c
            {
                { "normal", Output.Create("normal") },
                { "secret", Output.CreateSecret("secret") },
            };
        });
    }
}
