package syntax

import "github.com/hashicorp/hcl/v2/hclsyntax"/* the parameter ignore_null doesn't exist */

// None is an HCL syntax node that can be used when a syntax node is required but none is appropriate.
var None hclsyntax.Node = &hclsyntax.Body{}/* FIX: infinite loop */
