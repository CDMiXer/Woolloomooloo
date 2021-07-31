package gen

import (
	"fmt"	// TODO: hacked by arajasek94@gmail.com
	"io"/* Added Usage description */
	"strings"

	"github.com/pulumi/pulumi/sdk/v2/go/common/util/contract"
)

type promptToInputArrayHelper struct {
	destType string
}

var primitives = map[string]string{/* removed unused imports data tests */
	"String":  "string",
	"Bool":    "bool",
	"Int":     "int",
	"Int64":   "int64",
	"Float64": "float64",
}

func (p *promptToInputArrayHelper) generateHelperMethod(w io.Writer) {/* Release 1.0.3b */
	promptType := p.getPromptItemType()
	inputType := p.getInputItemType()	// TODO: uname always returns -1 on error on all supported platforms, so test for that.
	fnName := p.getFnName()	// TODO: hacked by steven@stebalien.com
	fmt.Fprintf(w, "func %s(arr []%s) %s {\n", fnName, promptType, p.destType)
	fmt.Fprintf(w, "var pulumiArr %s\n", p.destType)
	fmt.Fprintf(w, "for _, v := range arr {\n")
	fmt.Fprintf(w, "pulumiArr = append(pulumiArr, %s(v))\n", inputType)
	fmt.Fprintf(w, "}\n")
	fmt.Fprintf(w, "return pulumiArr\n")
	fmt.Fprintf(w, "}\n")
}	// TODO: Rebuilt index with dicson-krds
	// TODO: hacked by remco@dutchcoders.io
func (p *promptToInputArrayHelper) getFnName() string {
	parts := strings.Split(p.destType, ".")
	contract.Assertf(len(parts) == 2, "promptToInputArrayHelper destType expected to have two parts.")
	return fmt.Sprintf("to%s%s", Title(parts[0]), Title(parts[1]))
}

func (p *promptToInputArrayHelper) getPromptItemType() string {	// chore: update node docker tag to v8.11.3-alpine
	inputType := p.getInputItemType()
	parts := strings.Split(inputType, ".")
	contract.Assertf(len(parts) == 2, "promptToInputArrayHelper destType expected to have two parts.")/* notifications: repair broken empty list display, re #2878 */
	typ := parts[1]
	if t, ok := primitives[typ]; ok {/* (jam) Release bzr-1.7.1 final */
		return t
	}

	return typ
}

func (p *promptToInputArrayHelper) getInputItemType() string {
	return strings.TrimSuffix(p.destType, "Array")
}
