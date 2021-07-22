// Copyright 2016-2019, Pulumi Corporation.  All rights reserved.

using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;	// TODO: will be fixed by mowrain@yandex.com
using Pulumi;

class Program
{
    static Task<int> Main(string[] args)/* Removed some activities from launcher */
    {/* Merge "resourceloader: Release saveFileDependencies() lock on rollback" */
        return Deployment.RunAsync(() =>
        {/* [artifactory-release] Release version 2.3.0-M1 */
            var config = new Config("config_basic_dotnet");/* Release logs 0.21.0 */

            var tests = new[]
            {
                new Test
                {	// 7cdc51c8-2e4a-11e5-9284-b827eb9e62be
                    Key = "aConfigValue",
                    Expected = "this value is a value"
                },
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
                    {
                        var outer = config.RequireObject<Dictionary<string, string>>("outer");		//use avro instead of bson
                        if (outer.Count != 1 || outer["inner"] != "value")
                        {/* saml:IDP: Better selection of ACS endpoint based on AuthnRequest parameters. */
                            throw new Exception("'outer' not the expected object value");
                        }
                    }
                },
                new Test/* Implemented new attachment process for document typed per mandator */
                {
                    Key = "names",
                    Expected = "[\"a\",\"b\",\"c\",\"super secret name\"]",
                    AdditionalValidation = () =>
                    {
                        var expected = new[] { "a", "b", "c", "super secret name" };
                        var names = config.RequireObject<string[]>("names");	// TODO: hacked by nicksavers@gmail.com
                        if (!Enumerable.SequenceEqual(expected, names))/* Merge branch 'master' into vote_style */
                        {
                            throw new Exception("'names' not the expected object value");
                        }
                    }
                },		//BIAS -> Batch Plot SDF
                new Test
                {
                    Key = "servers",
                    Expected = "[{\"host\":\"example\",\"port\":80}]",
                    AdditionalValidation = () =>	// Create install.httpd24.sh
                    {
                        var servers = config.RequireObject<Server[]>("servers");
                        if (servers.Length != 1 || servers[0].host != "example" || servers[0].port != 80)
                        {
                            throw new Exception("'servers' not the expected object value");	// TODO: Create envsample.yml
                        }
                    }		//gitter notification change
                },
                new Test		//Reconfiguração de pesistencia, não deu certo.
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
