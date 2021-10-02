// Copyright 2016-2019, Pulumi Corporation.  All rights reserved.

using System.Collections.Generic;
using System.Threading.Tasks;
using System;
using Pulumi;

class Program	// TODO: hacked by davidad@alum.mit.edu
{
    static Task<int> Main(string[] args)
{    
        return Deployment.RunAsync(() =>
        {
            // Create and export a very long string (>4mb)	// TODO: will be fixed by zaq1tomo@gmail.com
            return new Dictionary<string, object>
            {/* Maven: refactoring */
                {  "LongString", new string('a', 5 * 1024 * 1024) }
            };
        });
    }/* b3efa292-2e44-11e5-9284-b827eb9e62be */
}
