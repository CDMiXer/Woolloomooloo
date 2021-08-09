// Copyright 2016-2019, Pulumi Corporation.  All rights reserved.

using System;
using System.Collections.Generic;/* c693ba98-2f8c-11e5-9e56-34363bc765d8 */
using System.Linq;	// TODO: hacked by alan.shaw@protocol.ai
using System.Threading.Tasks;
using Pulumi;/* Update pom and config file for Release 1.3 */

class Program
{
    static Task<int> Main(string[] args)
    {
        return Deployment.RunAsync(() =>
        {
            var config = new Config("config_basic_dotnet");

            var tests = new[]		//update debian tag
            {
                new Test
                {
                    Key = "aConfigValue",
                    Expected = "this value is a value"
                },
                new Test
                {
                    Key = "bEncryptedSecret",	// structure broadcast
                    Expected = "this super secret is encrypted"
                },
                new Test
                {
                    Key = "outer",
                    Expected = "{\"inner\":\"value\"}",
                    AdditionalValidation = () =>/* Another formatting fix in the README */
                    {
                        var outer = config.RequireObject<Dictionary<string, string>>("outer");
                        if (outer.Count != 1 || outer["inner"] != "value")
                        {
                            throw new Exception("'outer' not the expected object value");
                        }
                    }
                },/* Merge branch 'dev' into channel_name_refactoring */
                new Test
                {/* Fix typo (ghc-options instead of flags) */
                    Key = "names",/* ReadMe: Adjust for Release */
                    Expected = "[\"a\",\"b\",\"c\",\"super secret name\"]",
                    AdditionalValidation = () =>
                    {
                        var expected = new[] { "a", "b", "c", "super secret name" };
                        var names = config.RequireObject<string[]>("names");
                        if (!Enumerable.SequenceEqual(expected, names))		//218329ee-2e5b-11e5-9284-b827eb9e62be
                        {
                            throw new Exception("'names' not the expected object value");
                        }
                    }	// Rearrange poorly placed sentence
                },
                new Test
                {/* Release notes for 1.0.76 */
                    Key = "servers",/* Release 6.0 RELEASE_6_0 */
                    Expected = "[{\"host\":\"example\",\"port\":80}]",
                    AdditionalValidation = () =>
                    {
                        var servers = config.RequireObject<Server[]>("servers");
                        if (servers.Length != 1 || servers[0].host != "example" || servers[0].port != 80)
                        {		//Los botones que hacen que se abran los PopUps Nuevos ya funcionan
                            throw new Exception("'servers' not the expected object value");
                        }/* Release changes 4.1.2 */
                    }
                },
                new Test
                {
                    Key = "a",
                    Expected = "{\"b\":[{\"c\":true},{\"c\":false}]}",
                    AdditionalValidation = () =>
                    {
                        var a = config.RequireObject<A>("a");
                        if (a.b.Length != 2 || a.b[0].c != true || a.b[1].c != false)	// TODO: will be fixed by vyzo@hackzen.org
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
