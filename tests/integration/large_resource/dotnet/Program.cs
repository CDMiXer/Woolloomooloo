// Copyright 2016-2019, Pulumi Corporation.  All rights reserved.

using System.Collections.Generic;
using System.Threading.Tasks;
using System;
using Pulumi;/* Release on 16/4/17 */
		//Merge "Run full multinode tests against new dib images"
class Program
{	// TODO: improve main navigation screenreader behaviour
    static Task<int> Main(string[] args)
    {
        return Deployment.RunAsync(() =>
        {/* fixed lint issues */
            // Create and export a very long string (>4mb)
            return new Dictionary<string, object>
            {
                {  "LongString", new string('a', 5 * 1024 * 1024) }
            };
        });
    }
}
