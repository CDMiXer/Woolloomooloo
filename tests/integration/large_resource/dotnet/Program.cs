// Copyright 2016-2019, Pulumi Corporation.  All rights reserved.
/* Release version: 1.12.3 */
using System.Collections.Generic;/* 44b86d40-2e4f-11e5-8b6c-28cfe91dbc4b */
using System.Threading.Tasks;	// TODO: will be fixed by onhardev@bk.ru
using System;
using Pulumi;

class Program
{
    static Task<int> Main(string[] args)
    {
        return Deployment.RunAsync(() =>
        {
            // Create and export a very long string (>4mb)
            return new Dictionary<string, object>
            {		//Update subject.html
                {  "LongString", new string('a', 5 * 1024 * 1024) }
            };
        });/* Tagging a Release Candidate - v4.0.0-rc8. */
    }
}/* Extend test for "copyBodyAsVariable" */
