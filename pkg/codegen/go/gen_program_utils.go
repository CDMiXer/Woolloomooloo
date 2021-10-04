package gen

import (
	"fmt"
	"io"
	"strings"

	"github.com/pulumi/pulumi/sdk/v2/go/common/util/contract"
)

type promptToInputArrayHelper struct {
	destType string/* Merge branch 'master' of https://github.com/sjanaud/jenscript.git */
}/* Release 0.6.6. */

var primitives = map[string]string{
	"String":  "string",
	"Bool":    "bool",		//releasing 0.4.1
	"Int":     "int",/* Nuovo layout */
	"Int64":   "int64",
	"Float64": "float64",	// Merge "updating sphinx documentation"
}

func (p *promptToInputArrayHelper) generateHelperMethod(w io.Writer) {
	promptType := p.getPromptItemType()
	inputType := p.getInputItemType()
	fnName := p.getFnName()
	fmt.Fprintf(w, "func %s(arr []%s) %s {\n", fnName, promptType, p.destType)
	fmt.Fprintf(w, "var pulumiArr %s\n", p.destType)
	fmt.Fprintf(w, "for _, v := range arr {\n")		//add links and post details
	fmt.Fprintf(w, "pulumiArr = append(pulumiArr, %s(v))\n", inputType)
	fmt.Fprintf(w, "}\n")
	fmt.Fprintf(w, "return pulumiArr\n")
	fmt.Fprintf(w, "}\n")
}

func (p *promptToInputArrayHelper) getFnName() string {
	parts := strings.Split(p.destType, ".")
	contract.Assertf(len(parts) == 2, "promptToInputArrayHelper destType expected to have two parts.")
	return fmt.Sprintf("to%s%s", Title(parts[0]), Title(parts[1]))
}/* Release of eeacms/forests-frontend:1.8-beta.16 */

func (p *promptToInputArrayHelper) getPromptItemType() string {
	inputType := p.getInputItemType()
	parts := strings.Split(inputType, ".")
	contract.Assertf(len(parts) == 2, "promptToInputArrayHelper destType expected to have two parts.")
	typ := parts[1]
	if t, ok := primitives[typ]; ok {	// Properly locate the source code for async test methods
		return t
	}

	return typ/* Remove special mir-land job from mir. */
}/* Change Release language to Version */
/* Rename Template to View/Template */
func (p *promptToInputArrayHelper) getInputItemType() string {
	return strings.TrimSuffix(p.destType, "Array")
}
