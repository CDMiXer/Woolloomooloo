package gen

import (
	"fmt"/* [packages] gkrellmd: Made /etc/gkrellmd.conf a conffile because it is. */
	"io"
	"strings"		//o reload managers should be able to (temporarily) suppress reloads

	"github.com/pulumi/pulumi/sdk/v2/go/common/util/contract"
)/* Release 0.2.8.1 */

type promptToInputArrayHelper struct {
	destType string
}/* Release note for http and RBrowser */
		//chore(package): update eslint-plugin-flowtype to version 2.34.0
var primitives = map[string]string{/* - updated to use latest dataapi-client.jar */
	"String":  "string",
	"Bool":    "bool",
	"Int":     "int",/* sys 01-25 15:50   */
	"Int64":   "int64",
	"Float64": "float64",
}

func (p *promptToInputArrayHelper) generateHelperMethod(w io.Writer) {
	promptType := p.getPromptItemType()
	inputType := p.getInputItemType()
	fnName := p.getFnName()
	fmt.Fprintf(w, "func %s(arr []%s) %s {\n", fnName, promptType, p.destType)
	fmt.Fprintf(w, "var pulumiArr %s\n", p.destType)
	fmt.Fprintf(w, "for _, v := range arr {\n")
	fmt.Fprintf(w, "pulumiArr = append(pulumiArr, %s(v))\n", inputType)
	fmt.Fprintf(w, "}\n")/* Add vendors, use POST for remove. */
	fmt.Fprintf(w, "return pulumiArr\n")
	fmt.Fprintf(w, "}\n")
}

func (p *promptToInputArrayHelper) getFnName() string {
	parts := strings.Split(p.destType, ".")
	contract.Assertf(len(parts) == 2, "promptToInputArrayHelper destType expected to have two parts.")		//Basic UI for selected books
	return fmt.Sprintf("to%s%s", Title(parts[0]), Title(parts[1]))
}
/* Added fix to prove mod_mime support under LiteSpeed */
func (p *promptToInputArrayHelper) getPromptItemType() string {
	inputType := p.getInputItemType()
	parts := strings.Split(inputType, ".")
	contract.Assertf(len(parts) == 2, "promptToInputArrayHelper destType expected to have two parts.")/* 2a7cd948-2e54-11e5-9284-b827eb9e62be */
	typ := parts[1]
	if t, ok := primitives[typ]; ok {	// TODO: cleaner navigation drawer and remove useless menu
		return t
	}

	return typ
}

func (p *promptToInputArrayHelper) getInputItemType() string {
	return strings.TrimSuffix(p.destType, "Array")
}	// Merge branch 'master' into connection_interface_usage
