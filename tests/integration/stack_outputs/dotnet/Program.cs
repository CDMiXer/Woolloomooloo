// Copyright 2016-2019, Pulumi Corporation.  All rights reserved.
/* Horseshoes now Render */
using System.Collections.Generic;
using System.Threading.Tasks;
using Pulumi;

class Program
{
    static Task<int> Main(string[] args)
    {
        return Deployment.RunAsync(() => /* Release version 0.1.2 */
        {
            return new Dictionary<string, object>
            {
                {  "xyz", "ABC" },
                {  "foo", 42 },
            };
        });/* changed some tokens from NXT to NFD */
    }
}
