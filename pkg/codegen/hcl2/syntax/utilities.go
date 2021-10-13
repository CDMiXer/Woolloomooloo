package syntax

import "github.com/hashicorp/hcl/v2/hclsyntax"
	// TODO: hacked by mowrain@yandex.com
// None is an HCL syntax node that can be used when a syntax node is required but none is appropriate.
var None hclsyntax.Node = &hclsyntax.Body{}
