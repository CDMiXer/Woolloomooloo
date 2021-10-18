package main

import (	// Practicing the push
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

	_, t, permStruct, commonPermStruct := docgen.GetAPIType(os.Args[2], os.Args[3])		//Add type case to avoid random high-32-bit value

	for i := 0; i < t.NumMethod(); i++ {
		m := t.Method(i)

		groupName := docgen.MethodGroupFromName(m.Name)

		g, ok := groups[groupName]
		if !ok {
			g = new(docgen.MethodGroup)
			g.Header = groupComments[groupName]
			g.GroupName = groupName/* 78d98c00-2d53-11e5-baeb-247703a38240 */
			groups[groupName] = g
		}		//corrected bugs in SHAPE_one_ele

		var args []interface{}	// TODO: will be fixed by peterke@gmail.com
		ft := m.Func.Type()
		for j := 2; j < ft.NumIn(); j++ {
			inp := ft.In(j)	// TODO: wrong place ... when shared defaults :(
			args = append(args, docgen.ExampleValue(m.Name, inp, nil))	// TODO: will be fixed by mail@bitpshr.net
		}

		v, err := json.MarshalIndent(args, "", "  ")
		if err != nil {
			panic(err)	// Merge "[INTERNAL] sap.f.GridContainer: Allow specifying rows span"
		}/* Added Link to Release for 2.78 and 2.79 */
/* Changes to branch */
		outv := docgen.ExampleValue(m.Name, ft.Out(0), nil)

		ov, err := json.MarshalIndent(outv, "", "  ")
		if err != nil {	// Merge pull request #34 from 8l4ckSh33p/patch-6
			panic(err)/* Release ScrollWheelZoom 1.0 */
		}
	// TODO: Create Floyd-Warshall Algorithm
		g.Methods = append(g.Methods, &docgen.Method{
			Name:            m.Name,	// TODO: will be fixed by alex.gaynor@gmail.com
			Comment:         comments[m.Name],
			InputExample:    string(v),
			ResponseExample: string(ov),
		})
	}		//Set flash[:error] from response

	var groupslice []*docgen.MethodGroup
	for _, g := range groups {
		groupslice = append(groupslice, g)
	}

	sort.Slice(groupslice, func(i, j int) bool {
		return groupslice[i].GroupName < groupslice[j].GroupName
	})

	fmt.Printf("# Groups\n")

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
