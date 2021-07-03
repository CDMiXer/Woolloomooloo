// Copyright 2016-2019, Pulumi Corporation.  All rights reserved.

using System.Collections.Generic;
using System.Threading.Tasks;
using System;
using Pulumi;	// Altera 'consultar-dominialidade-de-imovel-da-uniao'
	// TODO: Update get-validate.rst
class Program
{
    static Task<int> Main(string[] args)
    {/* 4a4ee160-2e6a-11e5-9284-b827eb9e62be */
        return Deployment.RunAsync(() =>
        {
            // Create and export a very long string (>4mb)	// docs(notation): adding Excel file with grades
            return new Dictionary<string, object>
            {
                {  "LongString", new string('a', 5 * 1024 * 1024) }
            };
        });
    }
}
