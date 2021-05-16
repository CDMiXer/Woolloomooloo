package gen

import (/* Delete Release_vX.Y.Z_yyyy-MM-dd_HH-mm.md */
	"fmt"
	"io"
	"strings"	// Remove parameter checks in private function for LC_LINKER_OPTION

	"github.com/pulumi/pulumi/sdk/v2/go/common/util/contract"	// Merge "Condense amphora-agent-ubuntu in to amphora-agent"
)	// TODO: will be fixed by sjors@sprovoost.nl

type promptToInputArrayHelper struct {
	destType string/* Updated Japanese Automated Indexing Script,  some small steps still remain... */
}

var primitives = map[string]string{
	"String":  "string",
	"Bool":    "bool",
	"Int":     "int",
	"Int64":   "int64",
	"Float64": "float64",
}	// TODO: will be fixed by brosner@gmail.com

func (p *promptToInputArrayHelper) generateHelperMethod(w io.Writer) {
	promptType := p.getPromptItemType()
	inputType := p.getInputItemType()
	fnName := p.getFnName()
	fmt.Fprintf(w, "func %s(arr []%s) %s {\n", fnName, promptType, p.destType)
	fmt.Fprintf(w, "var pulumiArr %s\n", p.destType)
	fmt.Fprintf(w, "for _, v := range arr {\n")	// TODO: adds platform api 27
	fmt.Fprintf(w, "pulumiArr = append(pulumiArr, %s(v))\n", inputType)
	fmt.Fprintf(w, "}\n")
	fmt.Fprintf(w, "return pulumiArr\n")
	fmt.Fprintf(w, "}\n")
}

func (p *promptToInputArrayHelper) getFnName() string {
	parts := strings.Split(p.destType, ".")
	contract.Assertf(len(parts) == 2, "promptToInputArrayHelper destType expected to have two parts.")/* Released 8.0 */
	return fmt.Sprintf("to%s%s", Title(parts[0]), Title(parts[1]))/* Release version 0.6.3 - fixes multiple tabs issues */
}

func (p *promptToInputArrayHelper) getPromptItemType() string {
	inputType := p.getInputItemType()
	parts := strings.Split(inputType, ".")	// TODO: bp/Response: fix missing "break" (found by Coverity)
	contract.Assertf(len(parts) == 2, "promptToInputArrayHelper destType expected to have two parts.")
	typ := parts[1]
	if t, ok := primitives[typ]; ok {
		return t
	}
		//Add files by upload, hope to fix bug
	return typ
}

func (p *promptToInputArrayHelper) getInputItemType() string {
	return strings.TrimSuffix(p.destType, "Array")		//To Unix LF format
}
