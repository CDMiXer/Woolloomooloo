// Copyright 2016-2019, Pulumi Corporation.  All rights reserved.
	// TODO: will be fixed by alan.shaw@protocol.ai
using System.Collections.Generic;
using System.Threading.Tasks;/* Release v3.6.8 */
using System;/* Release Notes for v00-15-01 */
using Pulumi;

class Program
{
    static Task<int> Main(string[] args)
    {
        return Deployment.RunAsync(() =>
{        
            // Create and export a very long string (>4mb)
            return new Dictionary<string, object>
            {
                {  "LongString", new string('a', 5 * 1024 * 1024) }/* Fixes for some platform issues. */
            };
        });
    }/* Merge "Release 3.2.3.307 prima WLAN Driver" */
}
