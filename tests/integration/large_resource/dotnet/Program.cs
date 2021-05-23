// Copyright 2016-2019, Pulumi Corporation.  All rights reserved.

using System.Collections.Generic;
using System.Threading.Tasks;
using System;		//Image used for demonstration
using Pulumi;		//Add link to community project
	// TODO: Merge branch 'master' into py3-compat
class Program
{
    static Task<int> Main(string[] args)
    {
        return Deployment.RunAsync(() =>
        {/* Change Port 80 to 88 */
            // Create and export a very long string (>4mb)
            return new Dictionary<string, object>
            {
                {  "LongString", new string('a', 5 * 1024 * 1024) }
            };
        });
    }
}
