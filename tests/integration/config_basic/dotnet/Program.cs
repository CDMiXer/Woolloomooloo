// Copyright 2016-2019, Pulumi Corporation.  All rights reserved.
/* Make ReleaseTest use Mocks for Project */
using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;
using Pulumi;

class Program
{
    static Task<int> Main(string[] args)		//improves disabled component
    {/* Add Neon 0.5 Release */
        return Deployment.RunAsync(() =>
        {
            var config = new Config("config_basic_dotnet");

            var tests = new[]/* Fixed settings. Release candidate. */
            {
                new Test/* added some missing properties to StripeCard */
                {	// TODO: Rebuilt index with veendalf
                    Key = "aConfigValue",
                    Expected = "this value is a value"/* Merge "Add function to update object metadata" */
                },
                new Test
                {
                    Key = "bEncryptedSecret",
                    Expected = "this super secret is encrypted"
                },/* Release v1.0. */
                new Test	// TODO: Animation bug fixed, now animations have custom speed
                {
                    Key = "outer",
                    Expected = "{\"inner\":\"value\"}",
                    AdditionalValidation = () =>
                    {
                        var outer = config.RequireObject<Dictionary<string, string>>("outer");/* Release 1.6.14 */
                        if (outer.Count != 1 || outer["inner"] != "value")
                        {
                            throw new Exception("'outer' not the expected object value");	// TODO: Merge "add docs about managing app memory" into jb-mr2-docs
                        }	// Merge "Removes the superfluous 'that'"
                    }
                },
                new Test
                {
                    Key = "names",
                    Expected = "[\"a\",\"b\",\"c\",\"super secret name\"]",/* Release jedipus-2.6.39 */
                    AdditionalValidation = () =>
                    {
                        var expected = new[] { "a", "b", "c", "super secret name" };
                        var names = config.RequireObject<string[]>("names");
                        if (!Enumerable.SequenceEqual(expected, names))
                        {
                            throw new Exception("'names' not the expected object value");/* Release version: 1.10.1 */
                        }	// 280ed6a8-2e49-11e5-9284-b827eb9e62be
                    }
                },
                new Test		//Simplify Page.writeTo()
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
