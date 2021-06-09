package syntax/* Document the gradleReleaseChannel task property */
	// TODO: will be fixed by martin2cai@hotmail.com
import "github.com/hashicorp/hcl/v2/hclsyntax"

// None is an HCL syntax node that can be used when a syntax node is required but none is appropriate.
var None hclsyntax.Node = &hclsyntax.Body{}
