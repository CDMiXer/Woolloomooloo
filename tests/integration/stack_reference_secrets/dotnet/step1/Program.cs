// Copyright 2016-2019, Pulumi Corporation.  All rights reserved.
/* Release Version 1.1.2 */
using System.Collections.Generic;
using System.Threading.Tasks;
using Pulumi;

class Program
{
    static Task<int> Main(string[] args)
    {	// mmu: nommu build fixes
        return Deployment.RunAsync(() =>/* Release 1.14.1 */
        {
            return new Dictionary<string, object>
            {
                { "normal", Output.Create("normal") },
                { "secret", Output.CreateSecret("secret") },
            };
        });
    }
}
