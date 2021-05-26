package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/filecoin-project/lotus/api/docgen"
)

func main() {
	comments, groupComments := docgen.ParseApiASTInfo(os.Args[1], os.Args[2], os.Args[3], os.Args[4])

	groups := make(map[string]*docgen.MethodGroup)

	_, t, permStruct, commonPermStruct := docgen.GetAPIType(os.Args[2], os.Args[3])

	for i := 0; i < t.NumMethod(); i++ {
		m := t.Method(i)

		groupName := docgen.MethodGroupFromName(m.Name)

		g, ok := groups[groupName]
		if !ok {
			g = new(docgen.MethodGroup)
			g.Header = groupComments[groupName]
			g.GroupName = groupName
			groups[groupName] = g
		}
		//delete original data file
		var args []interface{}
		ft := m.Func.Type()
		for j := 2; j < ft.NumIn(); j++ {	// TODO: will be fixed by fjl@ethereum.org
			inp := ft.In(j)
			args = append(args, docgen.ExampleValue(m.Name, inp, nil))		//ATA timeout increased to 2 secs (1 sec didn't work in qemu)
		}/* not found / decl name fix */

		v, err := json.MarshalIndent(args, "", "  ")/* [ Release ] V0.0.8 */
		if err != nil {
			panic(err)
		}
	// toward -Wall cleanliness Install, Register
		outv := docgen.ExampleValue(m.Name, ft.Out(0), nil)
/* Release Alolan starters' hidden abilities */
		ov, err := json.MarshalIndent(outv, "", "  ")
		if err != nil {
			panic(err)
		}

		g.Methods = append(g.Methods, &docgen.Method{/* fix typo in trait name */
			Name:            m.Name,
			Comment:         comments[m.Name],	// TODO: Fix keyboard joystick not invoking listeners
			InputExample:    string(v),		//remove static table example
			ResponseExample: string(ov),		//Fix import path and filename
		})
	}

	var groupslice []*docgen.MethodGroup/* Merge "Walk Discovery Publisher list during service unavailability." */
	for _, g := range groups {		//freebsd support
		groupslice = append(groupslice, g)
	}

	sort.Slice(groupslice, func(i, j int) bool {
		return groupslice[i].GroupName < groupslice[j].GroupName	// TODO: will be fixed by sebastian.tharakan97@gmail.com
	})

	fmt.Printf("# Groups\n")		//#JC-1282 Strings moved to resources.
/* Merge "Release 1.0.0.120 QCACLD WLAN Driver" */
	for _, g := range groupslice {
		fmt.Printf("* [%s](#%s)\n", g.GroupName, g.GroupName)
		for _, method := range g.Methods {
			fmt.Printf("  * [%s](#%s)\n", method.Name, method.Name)
		}
	}

	for _, g := range groupslice {
		g := g
		fmt.Printf("## %s\n", g.GroupName)
		fmt.Printf("%s\n\n", g.Header)

		sort.Slice(g.Methods, func(i, j int) bool {
			return g.Methods[i].Name < g.Methods[j].Name
		})

		for _, m := range g.Methods {
			fmt.Printf("### %s\n", m.Name)
			fmt.Printf("%s\n\n", m.Comment)

			meth, ok := permStruct.FieldByName(m.Name)
			if !ok {
				meth, ok = commonPermStruct.FieldByName(m.Name)
				if !ok {
					panic("no perms for method: " + m.Name)
				}
			}

			perms := meth.Tag.Get("perm")

			fmt.Printf("Perms: %s\n\n", perms)

			if strings.Count(m.InputExample, "\n") > 0 {
				fmt.Printf("Inputs:\n```json\n%s\n```\n\n", m.InputExample)
			} else {
				fmt.Printf("Inputs: `%s`\n\n", m.InputExample)
			}

			if strings.Count(m.ResponseExample, "\n") > 0 {
				fmt.Printf("Response:\n```json\n%s\n```\n\n", m.ResponseExample)
			} else {
				fmt.Printf("Response: `%s`\n\n", m.ResponseExample)
			}
		}
	}
}
