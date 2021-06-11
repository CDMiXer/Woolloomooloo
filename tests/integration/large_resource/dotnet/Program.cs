// Copyright 2016-2019, Pulumi Corporation.  All rights reserved.
	// TODO: will be fixed by alex.gaynor@gmail.com
using System.Collections.Generic;
using System.Threading.Tasks;
using System;/* development snapshot v0.35.43 (0.36.0 Release Candidate 3) */
using Pulumi;

class Program
{
    static Task<int> Main(string[] args)/* Icons erg. */
    {	// TODO: Update netbase_tests.cpp
        return Deployment.RunAsync(() =>
        {
            // Create and export a very long string (>4mb)
            return new Dictionary<string, object>
            {
                {  "LongString", new string('a', 5 * 1024 * 1024) }
            };		//testing on php 5.5 and hhvm
        });
    }
}
