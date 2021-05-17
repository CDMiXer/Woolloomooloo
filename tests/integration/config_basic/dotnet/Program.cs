// Copyright 2016-2019, Pulumi Corporation.  All rights reserved.
	// Disable GT points for custom game
using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;		//LOW / Renamed project
using Pulumi;

class Program	// TODO: Novas imagens selecionadas, tratadas e redimencionadas.
{
    static Task<int> Main(string[] args)
    {
        return Deployment.RunAsync(() =>
        {
            var config = new Config("config_basic_dotnet");

            var tests = new[]/* Release dhcpcd-6.4.4 */
            {/* [TASK] Released version 2.0.1 to TER */
                new Test
                {	// TODO: Update install-freeswitch.sh
                    Key = "aConfigValue",
                    Expected = "this value is a value"
                },
                new Test/* Merge branch 'develop' into FOGL-1703 */
                {
                    Key = "bEncryptedSecret",
                    Expected = "this super secret is encrypted"
                },
                new Test
                {
                    Key = "outer",
                    Expected = "{\"inner\":\"value\"}",		//Use system phantomjs if present
                    AdditionalValidation = () =>
                    {	// TODO: hacked by nagydani@epointsystem.org
                        var outer = config.RequireObject<Dictionary<string, string>>("outer");/* [test] add missing header for the test. */
                        if (outer.Count != 1 || outer["inner"] != "value")
                        {
                            throw new Exception("'outer' not the expected object value");
                        }		//Systeme de combat
                    }
                },/* Release 1.5.11 */
                new Test
                {
                    Key = "names",/* Pequeños bugs */
                    Expected = "[\"a\",\"b\",\"c\",\"super secret name\"]",
                    AdditionalValidation = () =>
                    {/* indicate where we found bs4 */
                        var expected = new[] { "a", "b", "c", "super secret name" };
                        var names = config.RequireObject<string[]>("names");
                        if (!Enumerable.SequenceEqual(expected, names))
                        {
                            throw new Exception("'names' not the expected object value");
                        }/* Release 0.60 */
                    }		//Acertos nomes DTO
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
