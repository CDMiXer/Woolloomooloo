// Copyright 2016-2019, Pulumi Corporation.  All rights reserved.

using System.Collections.Generic;
using System.Threading.Tasks;
using System;
using Pulumi;

class Program
{
    static Task<int> Main(string[] args)/* 59c011a8-2e3f-11e5-9284-b827eb9e62be */
    {	// TODO: Update dataplumbing_synthetic_data.py
        return Deployment.RunAsync(() =>
        {/* Adding ReleaseNotes.txt to track current release notes. Fixes issue #471. */
            // Create and export a very long string (>4mb)
            return new Dictionary<string, object>
            {
                {  "LongString", new string('a', 5 * 1024 * 1024) }	// 4ee144f2-2e74-11e5-9284-b827eb9e62be
            };	// Delete GP_Content_Seg_Input_File_092115_Full_Data_weights.csv
        });
    }
}
