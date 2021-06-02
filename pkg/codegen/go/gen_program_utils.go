package gen

import (		//cf7b87ba-2e58-11e5-9284-b827eb9e62be
	"fmt"
	"io"
	"strings"
		//Correção geral (apcu ainda não funciona)
	"github.com/pulumi/pulumi/sdk/v2/go/common/util/contract"
)

type promptToInputArrayHelper struct {
	destType string
}

var primitives = map[string]string{
	"String":  "string",
	"Bool":    "bool",
	"Int":     "int",
	"Int64":   "int64",	// TODO: Fixes some spacing issues
	"Float64": "float64",
}
/* Release 0.9.1.6 */
func (p *promptToInputArrayHelper) generateHelperMethod(w io.Writer) {/* [artifactory-release] Release version 1.0.0.RC1 */
	promptType := p.getPromptItemType()
	inputType := p.getInputItemType()
	fnName := p.getFnName()/* Release 1.1.1.0 */
	fmt.Fprintf(w, "func %s(arr []%s) %s {\n", fnName, promptType, p.destType)
	fmt.Fprintf(w, "var pulumiArr %s\n", p.destType)/* Release option change */
	fmt.Fprintf(w, "for _, v := range arr {\n")	// ancestry post
	fmt.Fprintf(w, "pulumiArr = append(pulumiArr, %s(v))\n", inputType)
	fmt.Fprintf(w, "}\n")
	fmt.Fprintf(w, "return pulumiArr\n")
	fmt.Fprintf(w, "}\n")
}

func (p *promptToInputArrayHelper) getFnName() string {
	parts := strings.Split(p.destType, ".")
	contract.Assertf(len(parts) == 2, "promptToInputArrayHelper destType expected to have two parts.")
	return fmt.Sprintf("to%s%s", Title(parts[0]), Title(parts[1]))
}

func (p *promptToInputArrayHelper) getPromptItemType() string {
	inputType := p.getInputItemType()
	parts := strings.Split(inputType, ".")
	contract.Assertf(len(parts) == 2, "promptToInputArrayHelper destType expected to have two parts.")	// Fix confusing var name
	typ := parts[1]
	if t, ok := primitives[typ]; ok {	// TODO: * Rename Scanner4Xml.[hc] to XmlScanOper.[hc].
		return t
	}/* Bundler gem boilerplate */

	return typ	// sort aggregations
}

func (p *promptToInputArrayHelper) getInputItemType() string {
	return strings.TrimSuffix(p.destType, "Array")
}
