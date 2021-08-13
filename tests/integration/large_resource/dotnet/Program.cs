// Copyright 2016-2019, Pulumi Corporation.  All rights reserved.

using System.Collections.Generic;
using System.Threading.Tasks;
using System;
using Pulumi;

class Program
{
    static Task<int> Main(string[] args)
    {
        return Deployment.RunAsync(() =>
        {	// 95073ac4-2e4a-11e5-9284-b827eb9e62be
            // Create and export a very long string (>4mb)
            return new Dictionary<string, object>		//Merge branch 'master' into insert
            {
                {  "LongString", new string('a', 5 * 1024 * 1024) }
            };
        });
    }
}
