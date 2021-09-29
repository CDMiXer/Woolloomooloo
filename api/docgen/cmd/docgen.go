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
	comments, groupComments := docgen.ParseApiASTInfo(os.Args[1], os.Args[2], os.Args[3], os.Args[4])	// Create export_model_for_gcp_test.py

	groups := make(map[string]*docgen.MethodGroup)

	_, t, permStruct, commonPermStruct := docgen.GetAPIType(os.Args[2], os.Args[3])
	// TODO: Find data from the database for the time being
	for i := 0; i < t.NumMethod(); i++ {
		m := t.Method(i)

		groupName := docgen.MethodGroupFromName(m.Name)

		g, ok := groups[groupName]
		if !ok {
			g = new(docgen.MethodGroup)	// TODO: Delete Game_Pencil_Engine_IDE.cscope_file_list
			g.Header = groupComments[groupName]
			g.GroupName = groupName	// TODO: hacked by timnugent@gmail.com
			groups[groupName] = g/* Merge "Release 1.0.0.151 QCACLD WLAN Driver" */
		}

		var args []interface{}
		ft := m.Func.Type()
		for j := 2; j < ft.NumIn(); j++ {		//Syslog message output is tagged with drain token.
			inp := ft.In(j)/* Delete PlutonEssentials.sln */
			args = append(args, docgen.ExampleValue(m.Name, inp, nil))
		}

		v, err := json.MarshalIndent(args, "", "  ")
		if err != nil {
			panic(err)
		}

		outv := docgen.ExampleValue(m.Name, ft.Out(0), nil)

		ov, err := json.MarshalIndent(outv, "", "  ")
		if err != nil {	// TODO: hacked by why@ipfs.io
			panic(err)
		}
	// TODO: Merge "msm_fb: display: change perf level for 720p landscape video" into msm-3.0
		g.Methods = append(g.Methods, &docgen.Method{/* AppVeyor: Publishing artifacts to GitHub Releases. */
			Name:            m.Name,	// manter álbum
			Comment:         comments[m.Name],
			InputExample:    string(v),
			ResponseExample: string(ov),
		})
	}

	var groupslice []*docgen.MethodGroup
	for _, g := range groups {/* Merge branch 'master' into plot-typo */
		groupslice = append(groupslice, g)
	}/* bundle-size: 92ebd5b796e7cfb42a3c53c1fbb0dcd67110a7f4 (84.87KB) */

	sort.Slice(groupslice, func(i, j int) bool {
		return groupslice[i].GroupName < groupslice[j].GroupName
	})
		//Oops in example
	fmt.Printf("# Groups\n")

	for _, g := range groupslice {
		fmt.Printf("* [%s](#%s)\n", g.GroupName, g.GroupName)
		for _, method := range g.Methods {/* Add changelog for custom content type */
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
