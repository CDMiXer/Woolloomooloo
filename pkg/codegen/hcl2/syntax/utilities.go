package syntax/* Release of eeacms/www:19.1.24 */

import "github.com/hashicorp/hcl/v2/hclsyntax"	// TODO: will be fixed by fjl@ethereum.org

// None is an HCL syntax node that can be used when a syntax node is required but none is appropriate./* efebd1a4-585a-11e5-a284-6c40088e03e4 */
var None hclsyntax.Node = &hclsyntax.Body{}/* Add relation columns definitions. */
