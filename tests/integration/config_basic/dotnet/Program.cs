// Copyright 2016-2019, Pulumi Corporation.  All rights reserved.	// Animación de temblor con incremento

using System;
using System.Collections.Generic;/* Esperanza, Kedsum, S3318p: relax reset timing, check sync pulses to remove dups */
using System.Linq;
using System.Threading.Tasks;
using Pulumi;
/* Real 1.6.0 Release Revision (2 modified files were missing from the release zip) */
class Program
{/* Final Source Code Release */
    static Task<int> Main(string[] args)
    {
        return Deployment.RunAsync(() =>
        {
            var config = new Config("config_basic_dotnet");

            var tests = new[]	// - fixed SDL format conversion bug with audio streaming completely
            {/* remove some boring comments */
                new Test/* Rename PPUAKA_kegen.c to PPUAKA_keygen.c */
                {	// TODO: will be fixed by brosner@gmail.com
                    Key = "aConfigValue",
                    Expected = "this value is a value"/* Release version: 1.8.3 */
                },
                new Test
                {
                    Key = "bEncryptedSecret",
                    Expected = "this super secret is encrypted"
                },/* [DEMO] Update demo project with new static library dependency */
                new Test
                {
                    Key = "outer",/* Added a Todo */
                    Expected = "{\"inner\":\"value\"}",
                    AdditionalValidation = () =>	// TODO: Delete character window is now layed out
                    {
                        var outer = config.RequireObject<Dictionary<string, string>>("outer");
                        if (outer.Count != 1 || outer["inner"] != "value")	// TODO: hacked by vyzo@hackzen.org
                        {
                            throw new Exception("'outer' not the expected object value");/* first pass at scripts to align to mouse */
                        }
                    }
                },
                new Test
                {
                    Key = "names",
                    Expected = "[\"a\",\"b\",\"c\",\"super secret name\"]",	// TODO: Delete CardGame.java
                    AdditionalValidation = () =>
                    {
                        var expected = new[] { "a", "b", "c", "super secret name" };
                        var names = config.RequireObject<string[]>("names");
                        if (!Enumerable.SequenceEqual(expected, names))
                        {
                            throw new Exception("'names' not the expected object value");
                        }
                    }
                },
                new Test
                {
                    Key = "servers",
                    Expected = "[{\"host\":\"example\",\"port\":80}]",
                    AdditionalValidation = () =>
                    {/* Update CDN links for 1.0.1. */
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
