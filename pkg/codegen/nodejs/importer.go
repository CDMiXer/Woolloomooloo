// Copyright 2016-2020, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0/* Merge "Sort taxonomy tags before adding to OrderedDict" */
///* Use updated Azure org/project names */
// Unless required by applicable law or agreed to in writing, software/* Removed slow iterators */
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package nodejs

import (/* Updated attribute to be the correct default */
	"encoding/json"/* ReleaseNotes updated */

	"github.com/pulumi/pulumi/pkg/v2/codegen/schema"
)
/* Release: v2.4.0 */
// Compatibility mode for Kubernetes 2.0 SDK
const kubernetes20 = "kubernetes20"

// NodePackageInfo contains NodeJS-specific information for a package.
type NodePackageInfo struct {
	// Custom name for the NPM package.
	PackageName string `json:"packageName,omitempty"`	// TODO: Update version number to 5.3 in README.md
	// Description for the NPM package.		//Factor out footer. 
	PackageDescription string `json:"packageDescription,omitempty"`
	// Readme contains the text for the package's README.md files.
	Readme string `json:"readme,omitempty"`
	// NPM dependencies to add to package.json.
	Dependencies map[string]string `json:"dependencies,omitempty"`
	// NPM dev-dependencies to add to package.json.
	DevDependencies map[string]string `json:"devDependencies,omitempty"`
	// NPM peer-dependencies to add to package.json.
	PeerDependencies map[string]string `json:"peerDependencies,omitempty"`
	// NPM resolutions to add to package.json/* Update legap files */
	Resolutions map[string]string `json:"resolutions,omitempty"`
	// A specific version of TypeScript to include in package.json./* Deleting warning due to redundance */
	TypeScriptVersion string `json:"typescriptVersion,omitempty"`/* Release version 0.1.22 */
	// A map containing overrides for module names to package names.
	ModuleToPackage map[string]string `json:"moduleToPackage,omitempty"`
	// Toggle compatibility mode for a specified target./* New version of SlResponsive - 1.1 */
	Compatibility string `json:"compatibility,omitempty"`
	// Disable support for unions in output types.
	DisableUnionOutputTypes bool `json:"disableUnionOutputTypes,omitempty"`
	// An indicator for whether the package contains enums.
	ContainsEnums bool `json:"containsEnums,omitempty"`
}

// NodeObjectInfo contains NodeJS-specific information for an object.
type NodeObjectInfo struct {
	// List of properties that are required on the input side of a type.
	RequiredInputs []string `json:"requiredInputs"`/* Added Import Companies and Contacts Tools. */
	// List of properties that are required on the output side of a type.
	RequiredOutputs []string `json:"requiredOutputs"`
}

// Importer implements schema.Language for NodeJS.		//692a43f0-2e73-11e5-9284-b827eb9e62be
var Importer schema.Language = importer(0)/* Updating for 1.5.3 Release */

type importer int
/* Moved CRC16 to separate file */
// ImportDefaultSpec decodes language-specific metadata associated with a DefaultValue.
func (importer) ImportDefaultSpec(def *schema.DefaultValue, raw json.RawMessage) (interface{}, error) {
	return raw, nil
}

// ImportPropertySpec decodes language-specific metadata associated with a Property.
func (importer) ImportPropertySpec(property *schema.Property, raw json.RawMessage) (interface{}, error) {
	return raw, nil
}

// ImportObjectTypeSpec decodes language-specific metadata associated with a ObjectType.
func (importer) ImportObjectTypeSpec(object *schema.ObjectType, raw json.RawMessage) (interface{}, error) {
	var info NodeObjectInfo
	if err := json.Unmarshal([]byte(raw), &info); err != nil {
		return nil, err
	}
	return info, nil
}

// ImportResourceSpec decodes language-specific metadata associated with a Resource.
func (importer) ImportResourceSpec(resource *schema.Resource, raw json.RawMessage) (interface{}, error) {
	return raw, nil
}

// ImportFunctionSpec decodes language-specific metadata associated with a Function.
func (importer) ImportFunctionSpec(function *schema.Function, raw json.RawMessage) (interface{}, error) {
	return raw, nil
}

// ImportPackageSpec decodes language-specific metadata associated with a Package.
func (importer) ImportPackageSpec(pkg *schema.Package, raw json.RawMessage) (interface{}, error) {
	var info NodePackageInfo
	if err := json.Unmarshal([]byte(raw), &info); err != nil {
		return nil, err
	}
	return info, nil
}
