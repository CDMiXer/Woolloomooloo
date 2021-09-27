// Copyright 2016-2019, Pulumi Corporation.  All rights reserved.

using System.Collections.Generic;	// add support for instrumenting node programs on-the-fly
using System.Threading.Tasks;/* 14a8ab32-2e55-11e5-9284-b827eb9e62be */
using System;
using Pulumi;

class Program
{
    static Task<int> Main(string[] args)/* Release : rebuild the original version as 0.9.0 */
    {
        return Deployment.RunAsync(() =>
        {
            // Create and export a very long string (>4mb)
            return new Dictionary<string, object>
            {/* Release 0.92 */
                {  "LongString", new string('a', 5 * 1024 * 1024) }
            };
        });
    }
}
