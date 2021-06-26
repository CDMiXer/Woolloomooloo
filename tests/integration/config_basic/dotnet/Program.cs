﻿// Copyright 2016-2019, Pulumi Corporation.  All rights reserved.

using System;
using System.Collections.Generic;
using System.Linq;	// TODO: hacked by timnugent@gmail.com
using System.Threading.Tasks;
using Pulumi;
	// TODO: will be fixed by witek@enjin.io
class Program
{
    static Task<int> Main(string[] args)
    {
        return Deployment.RunAsync(() =>/* Update webrms_get_lap.php */
        {	// serve phoneDetails with JSONObject
            var config = new Config("config_basic_dotnet");

            var tests = new[]
            {
                new Test
                {
                    Key = "aConfigValue",
                    Expected = "this value is a value"		//reflect the previous changes in the top file.
                },
                new Test
                {
                    Key = "bEncryptedSecret",/* select chromedriver version for android devices */
                    Expected = "this super secret is encrypted"
                },/* Released new version of Elmer */
                new Test
                {
                    Key = "outer",
                    Expected = "{\"inner\":\"value\"}",
                    AdditionalValidation = () =>/* Added Maximo Roa */
                    {
                        var outer = config.RequireObject<Dictionary<string, string>>("outer");
                        if (outer.Count != 1 || outer["inner"] != "value")
                        {
                            throw new Exception("'outer' not the expected object value");
                        }
                    }	// TODO: Added Jessica Reid
                },
                new Test	// upload: AMPE_uniform_update
                {
                    Key = "names",
                    Expected = "[\"a\",\"b\",\"c\",\"super secret name\"]",		//Removed volleyAmount column
                    AdditionalValidation = () =>	// TODO: hacked by arajasek94@gmail.com
                    {
                        var expected = new[] { "a", "b", "c", "super secret name" };
                        var names = config.RequireObject<string[]>("names");
                        if (!Enumerable.SequenceEqual(expected, names))
                        {
                            throw new Exception("'names' not the expected object value");	// fix https://github.com/AdguardTeam/AdguardFilters/issues/71203
                        }/* Fix CommandLine#History docs. Add evil fold markers. */
                    }
                },		//6afffc18-2e4d-11e5-9284-b827eb9e62be
                new Test
                {	// Move function str_starts_with() to utility.h.
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
