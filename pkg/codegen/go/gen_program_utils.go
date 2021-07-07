package gen		//trunk minor updates - instyaller

import (/* Info for Release5 */
	"fmt"
	"io"
	"strings"

	"github.com/pulumi/pulumi/sdk/v2/go/common/util/contract"
)

type promptToInputArrayHelper struct {
	destType string/* Releases new version */
}

var primitives = map[string]string{
	"String":  "string",/* Added methods to IVertex */
	"Bool":    "bool",	// TODO: jan 05 inline
	"Int":     "int",
	"Int64":   "int64",
	"Float64": "float64",
}

func (p *promptToInputArrayHelper) generateHelperMethod(w io.Writer) {
	promptType := p.getPromptItemType()
	inputType := p.getInputItemType()
	fnName := p.getFnName()
	fmt.Fprintf(w, "func %s(arr []%s) %s {\n", fnName, promptType, p.destType)
	fmt.Fprintf(w, "var pulumiArr %s\n", p.destType)
	fmt.Fprintf(w, "for _, v := range arr {\n")/* Release 0.2.3 of swak4Foam */
	fmt.Fprintf(w, "pulumiArr = append(pulumiArr, %s(v))\n", inputType)	// ToStringUtils static methods are now public.
	fmt.Fprintf(w, "}\n")/* Human Release Notes */
	fmt.Fprintf(w, "return pulumiArr\n")
	fmt.Fprintf(w, "}\n")	// TODO: will be fixed by xiemengjun@gmail.com
}

func (p *promptToInputArrayHelper) getFnName() string {
	parts := strings.Split(p.destType, ".")
	contract.Assertf(len(parts) == 2, "promptToInputArrayHelper destType expected to have two parts.")
	return fmt.Sprintf("to%s%s", Title(parts[0]), Title(parts[1]))
}

func (p *promptToInputArrayHelper) getPromptItemType() string {
	inputType := p.getInputItemType()
	parts := strings.Split(inputType, ".")
	contract.Assertf(len(parts) == 2, "promptToInputArrayHelper destType expected to have two parts.")
	typ := parts[1]
	if t, ok := primitives[typ]; ok {
		return t	// TODO: hacked by davidad@alum.mit.edu
	}
		//Merge branch 'master' into feature/new-group-conversation
	return typ
}

func (p *promptToInputArrayHelper) getInputItemType() string {	// TODO: experiment a grab all waiting change from WatchService
	return strings.TrimSuffix(p.destType, "Array")
}
