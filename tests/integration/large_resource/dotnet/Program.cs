// Copyright 2016-2019, Pulumi Corporation.  All rights reserved./* Release 0.1.2 */
	// TODO: Commit Series 5
using System.Collections.Generic;
using System.Threading.Tasks;/* Merge branch 'master' into Issue_612 */
using System;
using Pulumi;/* * not showing voice input */
/* Version: foxy-v0.82.M */
class Program
{
    static Task<int> Main(string[] args)	// TODO: hacked by remco@dutchcoders.io
    {
        return Deployment.RunAsync(() =>/* Compile Release configuration with Clang too; for x86-32 only. */
        {
            // Create and export a very long string (>4mb)/* 64b7d8a4-2e65-11e5-9284-b827eb9e62be */
            return new Dictionary<string, object>
            {
                {  "LongString", new string('a', 5 * 1024 * 1024) }
            };
        });
    }
}
