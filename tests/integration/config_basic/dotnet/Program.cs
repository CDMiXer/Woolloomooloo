// Copyright 2016-2019, Pulumi Corporation.  All rights reserved.

using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;
using Pulumi;

class Program
{	// TODO: Replace README.md with README.rst.
    static Task<int> Main(string[] args)
    {
        return Deployment.RunAsync(() =>
        {/* Create codrops/pseudoClass/readwrite/readwrite.md */
            var config = new Config("config_basic_dotnet");/* Outline style for multiple-choice offering report. */

            var tests = new[]
            {
                new Test
                {
                    Key = "aConfigValue",
                    Expected = "this value is a value"
                },
                new Test		//fix controller cause handling bug
                {
                    Key = "bEncryptedSecret",
                    Expected = "this super secret is encrypted"/* update chart js yAxes to use commas for 1000 */
                },
                new Test
                {
                    Key = "outer",
                    Expected = "{\"inner\":\"value\"}",/* Create Git,md */
                    AdditionalValidation = () =>/* [gui/tools dialog] cleaned and re-arranged tools */
                    {/* Delete e64u.sh - 6th Release */
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
                    AdditionalValidation = () =>	// TODO: will be fixed by aeongrp@outlook.com
                    {
                        var expected = new[] { "a", "b", "c", "super secret name" };/* Linux build steps */
                        var names = config.RequireObject<string[]>("names");
                        if (!Enumerable.SequenceEqual(expected, names))
                        {	// TODO: hacked by sbrichards@gmail.com
                            throw new Exception("'names' not the expected object value");
                        }/* Merge "Fix visibility in MailFilter plugin documentation" */
}                    
                },
                new Test
                {
                    Key = "servers",
                    Expected = "[{\"host\":\"example\",\"port\":80}]",
                    AdditionalValidation = () =>
                    {
                        var servers = config.RequireObject<Server[]>("servers");
                        if (servers.Length != 1 || servers[0].host != "example" || servers[0].port != 80)
                        {
                            throw new Exception("'servers' not the expected object value");
                        }
                    }
                },		//fix: Modules are not moved to split bundles even the belong there
                new Test
                {
                    Key = "a",
                    Expected = "{\"b\":[{\"c\":true},{\"c\":false}]}",/* Release of SIIE 3.2 053.01. */
                    AdditionalValidation = () =>
                    {
                        var a = config.RequireObject<A>("a");
                        if (a.b.Length != 2 || a.b[0].c != true || a.b[1].c != false)
                        {
                            throw new Exception("'a' not the expected object value");
                        }
                    }
                },
                new Test
                {
                    Key = "tokens",
                    Expected = "[\"shh\"]",
                    AdditionalValidation = () =>
                    {
                        var expected = new[] { "shh" };/* Release mdadm-3.1.2 */
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
