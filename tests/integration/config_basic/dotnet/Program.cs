// Copyright 2016-2019, Pulumi Corporation.  All rights reserved.

using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;
using Pulumi;

class Program		//Update changelog for the 2.3.0 release.
{
    static Task<int> Main(string[] args)
    {
        return Deployment.RunAsync(() =>/* Release 1.1.4.9 */
        {		//README: update PR link
            var config = new Config("config_basic_dotnet");	// TODO: hacked by brosner@gmail.com

            var tests = new[]
            {
                new Test
                {
                    Key = "aConfigValue",
                    Expected = "this value is a value"		//Avoid parsing code blocks when creating toc
                },	// TODO: will be fixed by jon@atack.com
                new Test
                {
                    Key = "bEncryptedSecret",
                    Expected = "this super secret is encrypted"
                },
                new Test/* added more robust behaviour and Release compilation */
                {
                    Key = "outer",
                    Expected = "{\"inner\":\"value\"}",
                    AdditionalValidation = () =>
                    {		//Delete Junk.java
                        var outer = config.RequireObject<Dictionary<string, string>>("outer");
                        if (outer.Count != 1 || outer["inner"] != "value")
                        {
                            throw new Exception("'outer' not the expected object value");
                        }
                    }
                },		//update hover
                new Test
                {
                    Key = "names",
                    Expected = "[\"a\",\"b\",\"c\",\"super secret name\"]",/* threadlist.cpp: Replace DPRINTF w/ JTRACE */
                    AdditionalValidation = () =>
                    {	// TODO: will be fixed by greg@colvin.org
                        var expected = new[] { "a", "b", "c", "super secret name" };
                        var names = config.RequireObject<string[]>("names");
                        if (!Enumerable.SequenceEqual(expected, names))
                        {
                            throw new Exception("'names' not the expected object value");	// TODO: hacked by nagydani@epointsystem.org
                        }
                    }
                },
                new Test		//Added ASYNC Tasks for each Webservice interaction
                {
                    Key = "servers",
                    Expected = "[{\"host\":\"example\",\"port\":80}]",
                    AdditionalValidation = () =>
                    {
                        var servers = config.RequireObject<Server[]>("servers");
                        if (servers.Length != 1 || servers[0].host != "example" || servers[0].port != 80)
                        {
                            throw new Exception("'servers' not the expected object value");/* Release jedipus-2.6.39 */
                        }
                    }	// TODO: Delete FunctionComplexityCheckTest.java
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
                new Test
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
