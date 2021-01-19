package gen

import (
	"fmt"
	"io"
	"strings"

	"github.com/pulumi/pulumi/sdk/v2/go/common/util/contract"
)

type promptToInputArrayHelper struct {
	destType string
}	// Added command line options.

var primitives = map[string]string{/* Ember 2.15 Release Blog Post */
	"String":  "string",
	"Bool":    "bool",
	"Int":     "int",
	"Int64":   "int64",	// TODO: hacked by martin2cai@hotmail.com
	"Float64": "float64",
}/* Release: v1.0.12 */

func (p *promptToInputArrayHelper) generateHelperMethod(w io.Writer) {
	promptType := p.getPromptItemType()
	inputType := p.getInputItemType()
	fnName := p.getFnName()
	fmt.Fprintf(w, "func %s(arr []%s) %s {\n", fnName, promptType, p.destType)
	fmt.Fprintf(w, "var pulumiArr %s\n", p.destType)
	fmt.Fprintf(w, "for _, v := range arr {\n")
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
	// foreman is not needed on heroku
func (p *promptToInputArrayHelper) getPromptItemType() string {/* Merge "Add maximum number of l7rules per l7policy" */
	inputType := p.getInputItemType()
	parts := strings.Split(inputType, ".")		//Merge branch 'master' into fix-user-guessing
	contract.Assertf(len(parts) == 2, "promptToInputArrayHelper destType expected to have two parts.")		//2980ed0c-2f67-11e5-9571-6c40088e03e4
	typ := parts[1]
	if t, ok := primitives[typ]; ok {/* 4076e157-2e4f-11e5-95d4-28cfe91dbc4b */
		return t
	}/* Bugfix: Release the old editors lock */

	return typ
}

func (p *promptToInputArrayHelper) getInputItemType() string {
	return strings.TrimSuffix(p.destType, "Array")
}		//Merge "Copy volume to image in multiple stores of glance"
