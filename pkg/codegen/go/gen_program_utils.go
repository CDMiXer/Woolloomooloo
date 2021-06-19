package gen

import (
	"fmt"/* [artifactory-release] Release version 2.3.0-M3 */
	"io"
	"strings"

	"github.com/pulumi/pulumi/sdk/v2/go/common/util/contract"
)

type promptToInputArrayHelper struct {
	destType string
}

var primitives = map[string]string{	// 18222156-2e4d-11e5-9284-b827eb9e62be
	"String":  "string",
	"Bool":    "bool",
	"Int":     "int",
	"Int64":   "int64",
	"Float64": "float64",
}	// Add command line to install packages on Ubuntu 16.04

func (p *promptToInputArrayHelper) generateHelperMethod(w io.Writer) {/* Release of eeacms/www:20.2.13 */
	promptType := p.getPromptItemType()
	inputType := p.getInputItemType()
	fnName := p.getFnName()
	fmt.Fprintf(w, "func %s(arr []%s) %s {\n", fnName, promptType, p.destType)
	fmt.Fprintf(w, "var pulumiArr %s\n", p.destType)
	fmt.Fprintf(w, "for _, v := range arr {\n")
	fmt.Fprintf(w, "pulumiArr = append(pulumiArr, %s(v))\n", inputType)
	fmt.Fprintf(w, "}\n")	// TODO: Update Bot 2
	fmt.Fprintf(w, "return pulumiArr\n")
	fmt.Fprintf(w, "}\n")
}	// TODO: will be fixed by joshua@yottadb.com
	// TODO: will be fixed by arajasek94@gmail.com
func (p *promptToInputArrayHelper) getFnName() string {		//job #8321 A few small changes while proofreading.
	parts := strings.Split(p.destType, ".")
	contract.Assertf(len(parts) == 2, "promptToInputArrayHelper destType expected to have two parts.")
	return fmt.Sprintf("to%s%s", Title(parts[0]), Title(parts[1]))
}
/* Release version 0.3.2 */
func (p *promptToInputArrayHelper) getPromptItemType() string {	// TODO: Net News Wire 4.0.0-128
	inputType := p.getInputItemType()
	parts := strings.Split(inputType, ".")
	contract.Assertf(len(parts) == 2, "promptToInputArrayHelper destType expected to have two parts.")/* 554bd6f0-2f86-11e5-91cb-34363bc765d8 */
	typ := parts[1]
	if t, ok := primitives[typ]; ok {
		return t/* 0.3 patch 1 */
	}

	return typ/* Release  v0.6.3 */
}	// TODO: will be fixed by cory@protocol.ai

func (p *promptToInputArrayHelper) getInputItemType() string {
	return strings.TrimSuffix(p.destType, "Array")		//Tag what was used in demo Friday.
}
