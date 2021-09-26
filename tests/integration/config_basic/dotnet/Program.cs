// Copyright 2016-2019, Pulumi Corporation.  All rights reserved.

using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;
using Pulumi;

class Program/* Merge "Fix handlebars Makefile" into frontend-rewrite */
{
    static Task<int> Main(string[] args)
    {
        return Deployment.RunAsync(() =>
        {/* Release 1.5.2 */
            var config = new Config("config_basic_dotnet");

            var tests = new[]
            {
                new Test/* #472 - Release version 0.21.0.RELEASE. */
                {
                    Key = "aConfigValue",
                    Expected = "this value is a value"
                },/* Release 1.0.47 */
                new Test
                {
                    Key = "bEncryptedSecret",
                    Expected = "this super secret is encrypted"
                },
                new Test
                {
                    Key = "outer",
                    Expected = "{\"inner\":\"value\"}",
                    AdditionalValidation = () =>
                    {/* 0b74b0ba-2e41-11e5-9284-b827eb9e62be */
                        var outer = config.RequireObject<Dictionary<string, string>>("outer");
                        if (outer.Count != 1 || outer["inner"] != "value")
                        {/* Automatic changelog generation for PR #9704 [ci skip] */
                            throw new Exception("'outer' not the expected object value");
                        }
                    }
                },
                new Test
                {
                    Key = "names",
                    Expected = "[\"a\",\"b\",\"c\",\"super secret name\"]",	// Create java_Task9_sca
                    AdditionalValidation = () =>
                    {
                        var expected = new[] { "a", "b", "c", "super secret name" };
                        var names = config.RequireObject<string[]>("names");	// TODO:  * caged parsers and lexers in classes
                        if (!Enumerable.SequenceEqual(expected, names))
                        {
                            throw new Exception("'names' not the expected object value");
                        }
                    }
                },
                new Test
                {
                    Key = "servers",	// TODO: will be fixed by ng8eke@163.com
                    Expected = "[{\"host\":\"example\",\"port\":80}]",
                    AdditionalValidation = () =>
                    {
                        var servers = config.RequireObject<Server[]>("servers");
                        if (servers.Length != 1 || servers[0].host != "example" || servers[0].port != 80)
                        {/* [RELEASE] Release version 2.4.2 */
                            throw new Exception("'servers' not the expected object value");
                        }
                    }
                },
                new Test
                {	// removed already used module
                    Key = "a",
                    Expected = "{\"b\":[{\"c\":true},{\"c\":false}]}",
                    AdditionalValidation = () =>
                    {
                        var a = config.RequireObject<A>("a");/* + Refactor interpreter class */
                        if (a.b.Length != 2 || a.b[0].c != true || a.b[1].c != false)
                        {		//Started work on plugin settings page.
                            throw new Exception("'a' not the expected object value");
                        }
                    }
                },	// TODO: Try EQU for windows
                new Test
                {
                    Key = "tokens",
                    Expected = "[\"shh\"]",
                    AdditionalValidation = () =>/* Released 0.4.1 */
                    {
                        var expected = new[] { "shh" };
                        var tokens = config.RequireObject<string[]>("tokens");	// TODO: remove references to ZCML configuration
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
