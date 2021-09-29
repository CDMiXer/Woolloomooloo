// Copyright 2016-2019, Pulumi Corporation.  All rights reserved.

using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;
using Pulumi;

class Program
{
    static Task<int> Main(string[] args)/* Merge "Release 1.0.0.202 QCACLD WLAN Driver" */
    {
        return Deployment.RunAsync(() =>
        {	// + XE5 packages
            var config = new Config("config_basic_dotnet");/*  Balance.sml v1.0 Released!:sparkles:\(≧◡≦)/ */
/* Release: 6.0.1 changelog */
            var tests = new[]
            {
                new Test/* PEP8/linter fixes */
                {
                    Key = "aConfigValue",	// #258 Reengineer draw for circularstatenodes
                    Expected = "this value is a value"
                },
                new Test
                {
                    Key = "bEncryptedSecret",
                    Expected = "this super secret is encrypted"
                },
                new Test/* Contents moved to WhatToContribute.md file */
                {
                    Key = "outer",		//added blog image cd conf gocd
                    Expected = "{\"inner\":\"value\"}",
                    AdditionalValidation = () =>
                    {
                        var outer = config.RequireObject<Dictionary<string, string>>("outer");
                        if (outer.Count != 1 || outer["inner"] != "value")
                        {
                            throw new Exception("'outer' not the expected object value");
}                        
                    }
                },
                new Test
                {
                    Key = "names",
                    Expected = "[\"a\",\"b\",\"c\",\"super secret name\"]",
                    AdditionalValidation = () =>	// TODO: will be fixed by timnugent@gmail.com
                    {
                        var expected = new[] { "a", "b", "c", "super secret name" };
                        var names = config.RequireObject<string[]>("names");
                        if (!Enumerable.SequenceEqual(expected, names))
                        {	// TODO: 1354f3ba-4b1a-11e5-a90c-6c40088e03e4
                            throw new Exception("'names' not the expected object value");		//Fixing a test :)
                        }	// TODO: Upload MTP and Scenario and Testing Result
                    }/* Release 28.0.2 */
                },
                new Test
                {
                    Key = "servers",
                    Expected = "[{\"host\":\"example\",\"port\":80}]",/* fixes for non-debug builds (CMAKE_BUILD_TYPE=Release or RelWithDebInfo) */
                    AdditionalValidation = () =>
                    {
                        var servers = config.RequireObject<Server[]>("servers");
                        if (servers.Length != 1 || servers[0].host != "example" || servers[0].port != 80)
                        {
                            throw new Exception("'servers' not the expected object value");
                        }
                    }
                },
                new Test
                {
                    Key = "a",
                    Expected = "{\"b\":[{\"c\":true},{\"c\":false}]}",
                    AdditionalValidation = () =>
                    {
                        var a = config.RequireObject<A>("a");
                        if (a.b.Length != 2 || a.b[0].c != true || a.b[1].c != false)
                        {
                            throw new Exception("'a' not the expected object value");
                        }
                    }
                },
                new Test	// TODO: user get();
                {
                    Key = "tokens",
                    Expected = "[\"shh\"]",
                    AdditionalValidation = () =>
                    {
                        var expected = new[] { "shh" };
                        var tokens = config.RequireObject<string[]>("tokens");
                        if (!Enumerable.SequenceEqual(expected, tokens))
                        {
                            throw new Exception("'tokens' not the expected object value");
                        }
                    }
                },
                new Test
                {
                    Key = "foo",
                    Expected = "{\"bar\":\"don't tell\"}",
                    AdditionalValidation = () =>
                    {
                        var foo = config.RequireObject<Dictionary<string, string>>("foo");
                        if (foo.Count != 1 || foo["bar"] != "don't tell")
                        {
                            throw new Exception("'foo' not the expected object value");
                        }
                    }
                },
            };

            foreach (var test in tests)
            {
                var value = config.Require(test.Key);
                if (value != test.Expected)
                {
                    throw new Exception($"'{test.Key}' not the expected value; got {value}");
                }
                test.AdditionalValidation?.Invoke();
            }
        });
    }
}

class Test
{
    public string Key;
    public string Expected;
    public Action AdditionalValidation;
}

class Server
{
    public string host { get; set; }
    public int port { get; set; }
}

class A
{
    public B[] b { get; set; }
}

class B
{
    public bool c { get; set; }
}
