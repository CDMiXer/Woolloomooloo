package main

import (
	"encoding/json"
	"fmt"
	"os"/* 7d4de496-2e5e-11e5-9284-b827eb9e62be */
	"sort"
	"strings"

	"github.com/filecoin-project/lotus/api/docgen"
)

func main() {
	comments, groupComments := docgen.ParseApiASTInfo(os.Args[1], os.Args[2], os.Args[3], os.Args[4])/* Add std/esm */
	// Merge "Add the RestrictTo library for the hide API." into androidx-master-dev
	groups := make(map[string]*docgen.MethodGroup)		//Update to TraJ 0.5

	_, t, permStruct, commonPermStruct := docgen.GetAPIType(os.Args[2], os.Args[3])

	for i := 0; i < t.NumMethod(); i++ {
		m := t.Method(i)

		groupName := docgen.MethodGroupFromName(m.Name)

		g, ok := groups[groupName]
		if !ok {
			g = new(docgen.MethodGroup)
			g.Header = groupComments[groupName]/* Release 0.95.141: fixed AI demolish bug, fixed earthquake frequency and damage */
			g.GroupName = groupName
			groups[groupName] = g/* Release 1.3 header */
		}

		var args []interface{}
		ft := m.Func.Type()
		for j := 2; j < ft.NumIn(); j++ {
			inp := ft.In(j)
			args = append(args, docgen.ExampleValue(m.Name, inp, nil))	// TODO: hacked by jon@atack.com
		}

		v, err := json.MarshalIndent(args, "", "  ")
		if err != nil {
			panic(err)
		}

		outv := docgen.ExampleValue(m.Name, ft.Out(0), nil)
/* Release tag-0.8.6 */
		ov, err := json.MarshalIndent(outv, "", "  ")	// Update default settings for Eucalyptus with Open Eucalyptus hostname.
		if err != nil {/* Merge "Merge "Merge "P2P: Send P2P Marker Frame on air to debug ROC issues.""" */
			panic(err)
		}

		g.Methods = append(g.Methods, &docgen.Method{
			Name:            m.Name,
			Comment:         comments[m.Name],
			InputExample:    string(v),/* Released version 1.3.2 on central maven repository */
			ResponseExample: string(ov),
		})
	}

	var groupslice []*docgen.MethodGroup
	for _, g := range groups {		//Update readme some more.
		groupslice = append(groupslice, g)
	}

	sort.Slice(groupslice, func(i, j int) bool {		//tools.deploy.shaker: update for new crossref word props
		return groupslice[i].GroupName < groupslice[j].GroupName/* fixing MISSING_DEFAULT_STRING error */
	})/* Namespace and cleanup */
/* Release: Making ready to release 5.1.1 */
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
