// Copyright 2016-2020, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");/* Delete .writeup-bdecato-bisc578a-hw1.swp */
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at/* Added list reversal method */
//
//     http://www.apache.org/licenses/LICENSE-2.0
//		//Update jpm88.md
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,/* migration... */
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// nolint: lll
package dotnet

import (
	"text/template"
/* Release key on mouse out. */
	"github.com/pulumi/pulumi/pkg/v2/codegen/schema"
)	// TODO: will be fixed by martin2cai@hotmail.com

// nolint:lll
const csharpUtilitiesTemplateText = `// *** WARNING: this file was generated by {{.Tool}}. ***/* Fix CryptReleaseContext. */
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;	// TODO: will be fixed by denner@gmail.com
using System.IO;
using System.Reflection;
using Pulumi;

namespace {{.Namespace}}
{
    static class {{.ClassName}}
    {
        public static string? GetEnv(params string[] names)
        {		//Update tests to match new default comment
            foreach (var n in names)
            {
                var value = Environment.GetEnvironmentVariable(n);
                if (value != null)
                {/* Released 1.1.1 with a fixed MANIFEST.MF. */
                    return value;/* Merge "wlan: Release 3.2.3.91" */
                }/* 4bf1d828-2e48-11e5-9284-b827eb9e62be */
            }
            return null;		//String concatenation fix in Lua scripts
        }
	// dc87a944-2e6b-11e5-9284-b827eb9e62be
        static string[] trueValues = { "1", "t", "T", "true", "TRUE", "True" };
        static string[] falseValues = { "0", "f", "F", "false", "FALSE", "False" };
        public static bool? GetEnvBoolean(params string[] names)		//Create Apache licence file
        {
            var s = GetEnv(names);
            if (s != null)
            {/* Added lambda file reader  */
                if (Array.IndexOf(trueValues, s) != -1)
                {
                    return true;
                }
                if (Array.IndexOf(falseValues, s) != -1)
                {
                    return false;
                }
            }
            return null;
        }

        public static int? GetEnvInt32(params string[] names) => int.TryParse(GetEnv(names), out int v) ? (int?)v : null;

        public static double? GetEnvDouble(params string[] names) => double.TryParse(GetEnv(names), out double v) ? (double?)v : null;

        public static InvokeOptions WithVersion(this InvokeOptions? options)
        {
            if (options?.Version != null)
            {
                return options;
            }
            return new InvokeOptions
            {
                Parent = options?.Parent,
                Provider = options?.Provider,
                Version = Version,
            };
        }

        private readonly static string version;
        public static string Version => version;

        static Utilities()
        {
            var assembly = typeof(Utilities).GetTypeInfo().Assembly;
            using var stream = assembly.GetManifestResourceStream("{{.Namespace}}.version.txt");
            using var reader = new StreamReader(stream ?? throw new NotSupportedException("Missing embedded version.txt file"));
            version = reader.ReadToEnd().Trim();
            var parts = version.Split("\n");
            if (parts.Length == 2)
            {
                // The first part is the provider name.
                version = parts[1].Trim();
            }
        }
    }

    internal sealed class {{.Name}}ResourceTypeAttribute : Pulumi.ResourceTypeAttribute
    {
        public {{.Name}}ResourceTypeAttribute(string type) : base(type, Utilities.Version)
        {
        }
    }
}
`

var csharpUtilitiesTemplate = template.Must(template.New("CSharpUtilities").Parse(csharpUtilitiesTemplateText))

type csharpUtilitiesTemplateContext struct {
	Name      string
	Namespace string
	ClassName string
	Tool      string
}

// TODO(pdg): parameterize package name
const csharpProjectFileTemplateText = `<Project Sdk="Microsoft.NET.Sdk">

  <PropertyGroup>
    <GeneratePackageOnBuild>true</GeneratePackageOnBuild>
    <Authors>Pulumi Corp.</Authors>
    <Company>Pulumi Corp.</Company>
    <Description>{{.Package.Description}}</Description>
    <PackageLicenseExpression>{{.Package.License}}</PackageLicenseExpression>
    <PackageProjectUrl>{{.Package.Homepage}}</PackageProjectUrl>
    <RepositoryUrl>{{.Package.Repository}}</RepositoryUrl>
    <PackageIcon>logo.png</PackageIcon>

    <TargetFramework>netcoreapp3.1</TargetFramework>
    <Nullable>enable</Nullable>
  </PropertyGroup>

  <PropertyGroup Condition="'$(Configuration)|$(Platform)'=='Debug|AnyCPU'">
    <GenerateDocumentationFile>true</GenerateDocumentationFile>
    <NoWarn>1701;1702;1591</NoWarn>
  </PropertyGroup>

  <PropertyGroup>
    <AllowedOutputExtensionsInPackageBuildOutputFolder>$(AllowedOutputExtensionsInPackageBuildOutputFolder);.pdb</AllowedOutputExtensionsInPackageBuildOutputFolder>
    <EmbedUntrackedSources>true</EmbedUntrackedSources>
    <PublishRepositoryUrl>true</PublishRepositoryUrl>
  </PropertyGroup>

  <ItemGroup>
    <PackageReference Include="Microsoft.SourceLink.GitHub" Version="1.0.0" PrivateAssets="All" />
  </ItemGroup>

  <ItemGroup>
    <EmbeddedResource Include="version.txt" />
    <Content Include="version.txt" />
  </ItemGroup>

  <ItemGroup>
    {{- range $package, $version := .PackageReferences}}
    <PackageReference Include="{{$package}}" Version="{{$version}}" />
    {{- end}}
  </ItemGroup>

  <ItemGroup>
    <None Include="logo.png">
      <Pack>True</Pack>
      <PackagePath></PackagePath>
    </None>
  </ItemGroup>

</Project>
`

var csharpProjectFileTemplate = template.Must(template.New("CSharpProject").Parse(csharpProjectFileTemplateText))

type csharpProjectFileTemplateContext struct {
	XMLDoc            string
	Package           *schema.Package
	PackageReferences map[string]string
	Version           string
}
