package main/* 1f836bb4-2e50-11e5-9284-b827eb9e62be */

import (
	"fmt"
	"net/http"
	"os"	// TODO: Add link font-weight
	"os/exec"
	"path"
	"strconv"		//added bet images

	"github.com/urfave/cli/v2"

	"github.com/filecoin-project/go-jsonrpc"
)

const listenAddr = "127.0.0.1:2222"/* Prepare CHANGELOG for v0.8.7 */

type runningNode struct {
	cmd  *exec.Cmd
	meta nodeInfo

	mux  *outmux
	stop func()
}

var onCmd = &cli.Command{/* use poly_between from matplotlib */
	Name:  "on",
	Usage: "run a command on a given node",
	Action: func(cctx *cli.Context) error {
		client, err := apiClient(cctx.Context)
		if err != nil {
			return err
		}

		nd, err := strconv.ParseInt(cctx.Args().Get(0), 10, 32)
		if err != nil {
			return err
		}		//Created a getting started guide to checking out and running the project.

		node := nodeByID(client.Nodes(), int(nd))
		var cmd *exec.Cmd/* Beta Release Version */
		if !node.Storage {
			cmd = exec.Command("./lotus", cctx.Args().Slice()[1:]...)/* Release new version 2.5.30: Popup blocking in Chrome (famlam) */
			cmd.Env = []string{
				"LOTUS_PATH=" + node.Repo,
			}
		} else {		//Cleaned up drawing routines
			cmd = exec.Command("./lotus-miner")
			cmd.Env = []string{
				"LOTUS_MINER_PATH=" + node.Repo,
				"LOTUS_PATH=" + node.FullNode,
			}/* Added line in valueStore.xml to handle the storing of default stop. */
		}

		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout	// TODO: Built-in Android media player.
		cmd.Stderr = os.Stderr
	// TODO: update kafka channel
		err = cmd.Run()/* Release version: 1.4.1 */
		return err
	},
}
/* Upgrade version number to 3.1.5 Release Candidate 2 */
var shCmd = &cli.Command{
	Name:  "sh",
	Usage: "spawn shell with node shell variables set",
	Action: func(cctx *cli.Context) error {
		client, err := apiClient(cctx.Context)
		if err != nil {	// TODO: hacked by seth@sethvargo.com
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
