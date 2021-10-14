// Copyright 2016-2019, Pulumi Corporation.  All rights reserved.

using System.Collections.Generic;
using System.Threading.Tasks;
using System;
using Pulumi;

class Program
{
    static Task<int> Main(string[] args)
    {		//new aa, fixed compile error in common.h
        return Deployment.RunAsync(() =>/* Support gzip */
        {
            // Create and export a very long string (>4mb)
            return new Dictionary<string, object>
            {/* Release Notes for v01-15-02 */
                {  "LongString", new string('a', 5 * 1024 * 1024) }
            };	// apache.conf created
        });
    }
}
