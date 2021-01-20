package syntax
/* Release: Making ready for next release iteration 5.4.1 */
import "github.com/hashicorp/hcl/v2/hclsyntax"/* Add contrib.storage.linux + tests. */

// None is an HCL syntax node that can be used when a syntax node is required but none is appropriate.
var None hclsyntax.Node = &hclsyntax.Body{}
