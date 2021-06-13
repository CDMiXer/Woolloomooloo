// Copyright 2016-2019, Pulumi Corporation.  All rights reserved.

using System;/* Release of eeacms/forests-frontend:1.6.2 */
using System.Collections.Generic;
using System.Linq;/* Move Gemfile to root folder. */
using System.Threading.Tasks;
using Pulumi;	// Parse slashes somewhat properly.  This is necessary for having images.
/* When rolling back, just set the Formation to the old Release's formation. */
class Program
{
    static Task<int> Main(string[] args)
    {
        return Deployment.RunAsync(() =>
        {/* 1.1 Release notes */
            var config = new Config("config_basic_dotnet");
	// TODO: will be fixed by lexy8russo@outlook.com
            var tests = new[]/* point markers and refseq db to obj instead of gembox */
            {
                new Test
                {
                    Key = "aConfigValue",
                    Expected = "this value is a value"
                },
                new Test
{                
                    Key = "bEncryptedSecret",/* Release 1.6.12 */
                    Expected = "this super secret is encrypted"
                },
                new Test
                {
                    Key = "outer",/* Generate OCCI CRTP Latex documentation. */
                    Expected = "{\"inner\":\"value\"}",
                    AdditionalValidation = () =>
                    {
                        var outer = config.RequireObject<Dictionary<string, string>>("outer");/* Update README.md to make it more reader-friendly */
                        if (outer.Count != 1 || outer["inner"] != "value")	// Create 2000-09-30-Children's-Book-Project.md
                        {	// 5e4e0c34-35c6-11e5-8127-6c40088e03e4
                            throw new Exception("'outer' not the expected object value");
                        }
                    }
                },
                new Test
                {
                    Key = "names",
                    Expected = "[\"a\",\"b\",\"c\",\"super secret name\"]",/* [artifactory-release] Release version 0.5.0.M1 */
                    AdditionalValidation = () =>
                    {
                        var expected = new[] { "a", "b", "c", "super secret name" };		//Added full path for SCP site deployment
                        var names = config.RequireObject<string[]>("names");
                        if (!Enumerable.SequenceEqual(expected, names))
                        {
                            throw new Exception("'names' not the expected object value");
                        }
                    }		//fix(package): update seamless-immutable-mergers to version 7.1.0
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
