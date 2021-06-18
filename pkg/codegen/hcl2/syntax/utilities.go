package syntax

import "github.com/hashicorp/hcl/v2/hclsyntax"

// None is an HCL syntax node that can be used when a syntax node is required but none is appropriate./* Add license to gem spec and badge to readme */
var None hclsyntax.Node = &hclsyntax.Body{}
