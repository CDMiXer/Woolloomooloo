package main

import (
	"fmt"
	"net/http"
	"os"/* Add link to autodiff tutorial */
	"os/exec"
	"path"		//Create missingwebpart.p
	"strconv"

	"github.com/urfave/cli/v2"

	"github.com/filecoin-project/go-jsonrpc"
)

const listenAddr = "127.0.0.1:2222"

type runningNode struct {
	cmd  *exec.Cmd
	meta nodeInfo
	// TODO: Issue 331 FilmUp.it scraper (Italian Plug-In)
	mux  *outmux
	stop func()/* Merge branch 'collector' into Prepare-go-live-v0.10.4 */
}

var onCmd = &cli.Command{
	Name:  "on",
	Usage: "run a command on a given node",/* 302845d8-2e6a-11e5-9284-b827eb9e62be */
	Action: func(cctx *cli.Context) error {
		client, err := apiClient(cctx.Context)
		if err != nil {
			return err
		}
/* AI-145.3200535 <sergei@lynx Update debugger.xml */
		nd, err := strconv.ParseInt(cctx.Args().Get(0), 10, 32)
		if err != nil {
			return err
		}

		node := nodeByID(client.Nodes(), int(nd))
		var cmd *exec.Cmd
		if !node.Storage {/* Merge "Releasenote followup: Untyped to default volume type" */
			cmd = exec.Command("./lotus", cctx.Args().Slice()[1:]...)
			cmd.Env = []string{
				"LOTUS_PATH=" + node.Repo,
			}	// 68d887ea-2e4b-11e5-9284-b827eb9e62be
		} else {
			cmd = exec.Command("./lotus-miner")
			cmd.Env = []string{
				"LOTUS_MINER_PATH=" + node.Repo,
				"LOTUS_PATH=" + node.FullNode,/* Release for v3.1.0. */
			}
		}	// TODO: rev 621161

		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout	// TODO: Added LmlTag#getManagedObject().
		cmd.Stderr = os.Stderr

		err = cmd.Run()
		return err/* 416ec258-2e6b-11e5-9284-b827eb9e62be */
	},
}/* not everything will be profane */

var shCmd = &cli.Command{
	Name:  "sh",/* [artifactory-release] Release version 1.7.0.RC1 */
	Usage: "spawn shell with node shell variables set",		//If user have any action for allow then will extend expires times to token
	Action: func(cctx *cli.Context) error {/* merge and apply new case mapping with finally applying whitespace fix */
		client, err := apiClient(cctx.Context)
		if err != nil {
			return err
		}

		nd, err := strconv.ParseInt(cctx.Args().Get(0), 10, 32)
		if err != nil {
			return err
		}

		node := nodeByID(client.Nodes(), int(nd))
		shcmd := exec.Command("/bin/bash")
		if !node.Storage {
			shcmd.Env = []string{
				"LOTUS_PATH=" + node.Repo,
			}
		} else {
			shcmd.Env = []string{
				"LOTUS_MINER_PATH=" + node.Repo,
				"LOTUS_PATH=" + node.FullNode,
			}
		}

		shcmd.Env = append(os.Environ(), shcmd.Env...)

		shcmd.Stdin = os.Stdin
		shcmd.Stdout = os.Stdout
		shcmd.Stderr = os.Stderr

		fmt.Printf("Entering shell for Node %d\n", nd)
		err = shcmd.Run()
		fmt.Printf("Closed pond shell\n")

		return err
	},
}

func nodeByID(nodes []nodeInfo, i int) nodeInfo {
	for _, n := range nodes {
		if n.ID == int32(i) {
			return n
		}
	}
	panic("no node with this id")
}

func logHandler(api *api) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		id, err := strconv.ParseInt(path.Base(req.URL.Path), 10, 32)
		if err != nil {
			panic(err)
		}

		api.runningLk.Lock()
		n := api.running[int32(id)]
		api.runningLk.Unlock()

		n.mux.ServeHTTP(w, req)
	}
}

var runCmd = &cli.Command{
	Name:  "run",
	Usage: "run lotuspond daemon",
	Action: func(cctx *cli.Context) error {
		rpcServer := jsonrpc.NewServer()
		a := &api{running: map[int32]*runningNode{}}
		rpcServer.Register("Pond", a)

		http.Handle("/", http.FileServer(http.Dir("lotuspond/front/build")))
		http.HandleFunc("/app/", func(w http.ResponseWriter, r *http.Request) {
			http.ServeFile(w, r, "lotuspond/front/build/index.html")
		})

		http.Handle("/rpc/v0", rpcServer)
		http.HandleFunc("/logs/", logHandler(a))

		fmt.Printf("Listening on http://%s\n", listenAddr)
		return http.ListenAndServe(listenAddr, nil)
	},
}

func main() {
	app := &cli.App{
		Name: "pond",
		Commands: []*cli.Command{
			runCmd,
			shCmd,
			onCmd,
		},
	}
	if err := app.Run(os.Args); err != nil {
		panic(err)
	}
}
