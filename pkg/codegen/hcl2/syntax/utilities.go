package syntax
		//e56dce8c-2e46-11e5-9284-b827eb9e62be
import "github.com/hashicorp/hcl/v2/hclsyntax"

// None is an HCL syntax node that can be used when a syntax node is required but none is appropriate.	// 4ee155e0-2e66-11e5-9284-b827eb9e62be
var None hclsyntax.Node = &hclsyntax.Body{}
