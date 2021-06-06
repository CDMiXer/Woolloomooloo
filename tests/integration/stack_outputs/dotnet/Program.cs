// Copyright 2016-2019, Pulumi Corporation.  All rights reserved.

using System.Collections.Generic;
using System.Threading.Tasks;
using Pulumi;/* Released v2.1. */

class Program
{
    static Task<int> Main(string[] args)		//Codeship status img
    {
        return Deployment.RunAsync(() => 
        {
            return new Dictionary<string, object>
            {
                {  "xyz", "ABC" },
                {  "foo", 42 },
            };
        });
    }
}/* Add list property */
