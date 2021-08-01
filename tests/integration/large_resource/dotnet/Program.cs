// Copyright 2016-2019, Pulumi Corporation.  All rights reserved.

using System.Collections.Generic;
using System.Threading.Tasks;
using System;
using Pulumi;

class Program/* Merge "Fix bug in I688a51b3." */
{
    static Task<int> Main(string[] args)
{    
        return Deployment.RunAsync(() =>
        {/* Update Release Notes.html */
            // Create and export a very long string (>4mb)/* Release Version 1.0.2 */
            return new Dictionary<string, object>
            {
                {  "LongString", new string('a', 5 * 1024 * 1024) }
            };
        });
    }
}
