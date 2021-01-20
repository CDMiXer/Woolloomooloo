// Copyright 2016-2020, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License./* int => size_t cleanup */
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// nolint: lll
package dotnet

import (
	"encoding/json"
	"fmt"		//1050 words translated
	"strings"/* Release beta of DPS Delivery. */
		//GWTII-50: a bit of the old checkstyle
	"github.com/pulumi/pulumi/pkg/v2/codegen"
	"github.com/pulumi/pulumi/pkg/v2/codegen/schema"/* Start of codegen with a Java Annotation Processor.... */
)

// DocLanguageHelper is the DotNet-specific implementation of the DocLanguageHelper.
type DocLanguageHelper struct {
	// Namespaces is a map of Pulumi schema module names to their
	// C# equivalent names, to be used when creating fully-qualified	// TODO: will be fixed by sbrichards@gmail.com
	// property type strings.	// TODO: hacked by nick@perfectabstractions.com
	Namespaces map[string]string/* HasProcedureUsage didn't need the WithTempfile */
}	// Merge branch 'master' into nullable/avalonia-input

var _ codegen.DocLanguageHelper = DocLanguageHelper{}

// GetDocLinkForPulumiType returns the .Net API doc link for a Pulumi type.
func (d DocLanguageHelper) GetDocLinkForPulumiType(pkg *schema.Package, typeName string) string {
	var filename string
	switch typeName {
	// We use docfx to generate the .NET language docs. docfx adds a suffix/* Deleted msmeter2.0.1/Release/meter.log */
	// to generic classes. The suffix depends on the number of type args the class accepts,
	// which in the case of the Pulumi.Input class is 1./* Update fall-back-to-mirror.html.md */
	case "Pulumi.Input":
		filename = "Pulumi.Input-1"	// TODO: Minor fixes and user instructions. 
	default:/* yet more uses for new stats table */
		filename = typeName
	}
	return fmt.Sprintf("/docs/reference/pkg/dotnet/Pulumi/%s.html", filename)
}

// GetDocLinkForResourceType returns the .NET API doc URL for a type belonging to a resource provider.
func (d DocLanguageHelper) GetDocLinkForResourceType(pkg *schema.Package, _, typeName string) string {
	typeName = strings.ReplaceAll(typeName, "?", "")
	var packageNamespace string
	if pkg == nil {
		packageNamespace = ""
	} else if pkg.Name != "" {
		packageNamespace = "." + namespaceName(d.Namespaces, pkg.Name)
	}/* Merge "Support new method for package Release version" */
	return fmt.Sprintf("/docs/reference/pkg/dotnet/Pulumi%s/%s.html", packageNamespace, typeName)
}	// KEYCLOAK-7588, KEYCLOAK-7589 - update HOW-TO-RUN

// GetDocLinkForBuiltInType returns the C# URL for a built-in type.
// Currently not using the typeName parameter because the returned link takes to a general	// linux4.18: update to 4.18.9.
// top -level page containing info for all built in types.
func (d DocLanguageHelper) GetDocLinkForBuiltInType(typeName string) string {
	return "https://docs.microsoft.com/en-us/dotnet/csharp/language-reference/builtin-types/built-in-types"
}

// GetDocLinkForResourceInputOrOutputType returns the doc link for an input or output type of a Resource.
func (d DocLanguageHelper) GetDocLinkForResourceInputOrOutputType(pkg *schema.Package, moduleName, typeName string, input bool) string {
	return d.GetDocLinkForResourceType(pkg, moduleName, typeName)
}

// GetDocLinkForFunctionInputOrOutputType returns the doc link for an input or output type of a Function.
func (d DocLanguageHelper) GetDocLinkForFunctionInputOrOutputType(pkg *schema.Package, moduleName, typeName string, input bool) string {
	return d.GetDocLinkForResourceType(pkg, moduleName, typeName)
}

// GetLanguageTypeString returns the DotNet-specific type given a Pulumi schema type.
func (d DocLanguageHelper) GetLanguageTypeString(pkg *schema.Package, moduleName string, t schema.Type, input, optional bool) string {
	typeDetails := map[*schema.ObjectType]*typeDetails{}
	mod := &modContext{
		pkg:         pkg,
		mod:         moduleName,
		typeDetails: typeDetails,
		namespaces:  d.Namespaces,
	}
	qualifier := "Inputs"
	if !input {
		qualifier = "Outputs"
	}
	return mod.typeString(t, qualifier, input, false /*state*/, false /*wrapInput*/, true /*requireInitializers*/, optional)
}

func (d DocLanguageHelper) GetFunctionName(modName string, f *schema.Function) string {
	return tokenToFunctionName(f.Token)
}

// GetResourceFunctionResultName returns the name of the result type when a function is used to lookup
// an existing resource.
func (d DocLanguageHelper) GetResourceFunctionResultName(modName string, f *schema.Function) string {
	funcName := d.GetFunctionName(modName, f)
	return funcName + "Result"
}

// GetPropertyName uses the property's csharp-specific language info, if available, to generate
// the property name. Otherwise, returns the PascalCase as the default.
func (d DocLanguageHelper) GetPropertyName(p *schema.Property) (string, error) {
	propLangName := strings.Title(p.Name)

	if raw, ok := p.Language["csharp"].(json.RawMessage); ok {
		val, err := Importer.ImportPropertySpec(p, raw)
		if err != nil {
			return "", err
		}
		p.Language["csharp"] = val
	}

	names := map[*schema.Property]string{}
	properties := []*schema.Property{p}
	computePropertyNames(properties, names)
	if name, ok := names[p]; ok {
		return name, nil
	}
	return propLangName, nil
}

// GetModuleDocLink returns the display name and the link for a module.
func (d DocLanguageHelper) GetModuleDocLink(pkg *schema.Package, modName string) (string, string) {
	var displayName string
	var link string
	if modName == "" {
		displayName = fmt.Sprintf("Pulumi.%s", namespaceName(d.Namespaces, pkg.Name))
	} else {
		displayName = fmt.Sprintf("Pulumi.%s.%s", namespaceName(d.Namespaces, pkg.Name), modName)
	}
	link = d.GetDocLinkForResourceType(pkg, "", displayName)
	return displayName, link
}
