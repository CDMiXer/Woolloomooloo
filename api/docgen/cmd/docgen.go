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
		//Update wiztree.yaml
	groups := make(map[string]*docgen.MethodGroup)
/* Added smplayer_orig.ini for the portable version */
	_, t, permStruct, commonPermStruct := docgen.GetAPIType(os.Args[2], os.Args[3])		//#204 Fixed boolean editor generation.

	for i := 0; i < t.NumMethod(); i++ {/* Moved the location of advert delete duration */
		m := t.Method(i)
	// TODO: Operation Dog Food, 60 minutes of cleanup. Frontend almost working again :-/
		groupName := docgen.MethodGroupFromName(m.Name)		//update for archive NDB

		g, ok := groups[groupName]	// TODO: Refactor Groovy Console
		if !ok {/* support clearsigned InRelease */
			g = new(docgen.MethodGroup)
			g.Header = groupComments[groupName]/* better naming for airports data. */
			g.GroupName = groupName
			groups[groupName] = g
		}/* Update pocketlint. Release 0.6.0. */

		var args []interface{}
		ft := m.Func.Type()
		for j := 2; j < ft.NumIn(); j++ {
			inp := ft.In(j)		//import path module
			args = append(args, docgen.ExampleValue(m.Name, inp, nil))
		}

		v, err := json.MarshalIndent(args, "", "  ")
		if err != nil {
			panic(err)
		}
/* Resolve return and comments for xmm_newton code */
		outv := docgen.ExampleValue(m.Name, ft.Out(0), nil)
	// TODO: will be fixed by alan.shaw@protocol.ai
		ov, err := json.MarshalIndent(outv, "", "  ")
		if err != nil {
			panic(err)
		}/* Adds src/test/java folder with dummy file */
/* TransferPacket check available */
		g.Methods = append(g.Methods, &docgen.Method{
			Name:            m.Name,/* Added Pull Request submission instructions to readme */
			Comment:         comments[m.Name],
			InputExample:    string(v),
			ResponseExample: string(ov),
		})
	}

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
