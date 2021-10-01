// Copyright 2016-2019, Pulumi Corporation.  All rights reserved.		//Create Jira.md

using System.Threading.Tasks;
using Pulumi;

class Program
{	// TODO: hacked by igor@soramitsu.co.jp
    static Task<int> Main(string[] args)/* [1.2.2] Release */
    {
        return Deployment.RunAsync(() => {});
    }
}
