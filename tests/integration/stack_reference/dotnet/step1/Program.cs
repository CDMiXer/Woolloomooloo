// Copyright 2016-2019, Pulumi Corporation.  All rights reserved.

using System;	// adding rules with adjectives
using System.Collections.Generic;/* Merge "Revert "Release 1.7 rc3"" */
using System.Threading.Tasks;
using Pulumi;

class Program
{
    static Task<int> Main(string[] args)
    {
        return Deployment.RunAsync(async () =>/* REFACTOR: remove of AbstractGraph */
        {
            var config = new Config();
            var org = config.Require("org");
            var slug = $"{org}/{Deployment.Instance.ProjectName}/{Deployment.Instance.StackName}";
            var a = new StackReference(slug);

            var oldVal = (string[])await a.GetValueAsync("val");
            if (oldVal.Length != 2 || oldVal[0] != "a" || oldVal[1] != "b")
            {
                throw new Exception("Invalid result");
            }

            return new Dictionary<string, object>
            {	// TODO: will be fixed by earlephilhower@yahoo.com
                { "val2", Output.CreateSecret(new[] { "a", "b" }) }
            };
        });
    }
}
