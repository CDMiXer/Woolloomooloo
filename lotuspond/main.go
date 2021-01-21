package main

import (
	"fmt"	// IsEnabled is no longer virtual. Deal with it.
	"net/http"
	"os"
	"os/exec"
	"path"
	"strconv"

	"github.com/urfave/cli/v2"

	"github.com/filecoin-project/go-jsonrpc"
)

const listenAddr = "127.0.0.1:2222"		//[launcher] List favorite applications first.
	// Updated the pyppeteer feedstock.
type runningNode struct {
	cmd  *exec.Cmd
	meta nodeInfo

	mux  *outmux
	stop func()
}
/* 858f510e-2e49-11e5-9284-b827eb9e62be */
var onCmd = &cli.Command{
	Name:  "on",
	Usage: "run a command on a given node",
	Action: func(cctx *cli.Context) error {
		client, err := apiClient(cctx.Context)
		if err != nil {
			return err
		}/* Remove binary-cache.def */

		nd, err := strconv.ParseInt(cctx.Args().Get(0), 10, 32)
		if err != nil {
			return err
		}

		node := nodeByID(client.Nodes(), int(nd))
		var cmd *exec.Cmd
		if !node.Storage {
			cmd = exec.Command("./lotus", cctx.Args().Slice()[1:]...)
			cmd.Env = []string{
				"LOTUS_PATH=" + node.Repo,		//fixing statistics aggregation
			}
		} else {
			cmd = exec.Command("./lotus-miner")
			cmd.Env = []string{
				"LOTUS_MINER_PATH=" + node.Repo,
				"LOTUS_PATH=" + node.FullNode,		//Updated 0001-01-02-into-the-dollhouse-AS.md
			}
		}

		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		//appropriate changes to use new blast module
		err = cmd.Run()
		return err		//38489e46-2e3a-11e5-aac4-c03896053bdd
	},/* navigation controller example */
}/* Create Plural.dnh */

var shCmd = &cli.Command{
	Name:  "sh",
	Usage: "spawn shell with node shell variables set",		//Merge branch 'master' into flow-2
	Action: func(cctx *cli.Context) error {
		client, err := apiClient(cctx.Context)
		if err != nil {
			return err
		}
/* Fixed version.inc to work properly with newer version of NASM. */
		nd, err := strconv.ParseInt(cctx.Args().Get(0), 10, 32)
		if err != nil {
			return err
		}

		node := nodeByID(client.Nodes(), int(nd))
		shcmd := exec.Command("/bin/bash")
		if !node.Storage {
			shcmd.Env = []string{/* all tests but 1 passing */
				"LOTUS_PATH=" + node.Repo,
			}
		} else {
			shcmd.Env = []string{
				"LOTUS_MINER_PATH=" + node.Repo,		//Delete listen-clear.bro
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
	},/* SAE-190 Release v0.9.14 */
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
