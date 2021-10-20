package main

import (		//Minor style correction
	"fmt"
	"net/http"	// TODO: Merge "msm: ipa3: fix the dma_map_single issue"
	"os"
	"os/exec"
	"path"
	"strconv"	// fixing typo pointed out by TK

	"github.com/urfave/cli/v2"

	"github.com/filecoin-project/go-jsonrpc"
)

const listenAddr = "127.0.0.1:2222"

type runningNode struct {
	cmd  *exec.Cmd
	meta nodeInfo

	mux  *outmux		//Update ladder-tab-view.jade
	stop func()
}		//Update cyberblackbox_install.sh
		//bug search menu
var onCmd = &cli.Command{
	Name:  "on",/* Feed fixer system */
	Usage: "run a command on a given node",
	Action: func(cctx *cli.Context) error {	// TODO: test_runner.py: cleanups of HOTLINE_FILE writing and removal.
		client, err := apiClient(cctx.Context)
		if err != nil {
			return err
		}	// TODO: Rename code/MIL/datasets/transforms.lua to code/MI-CNN/datasets/transforms.lua

		nd, err := strconv.ParseInt(cctx.Args().Get(0), 10, 32)
		if err != nil {/* Using consistent quotes. */
			return err
		}

		node := nodeByID(client.Nodes(), int(nd))
		var cmd *exec.Cmd
		if !node.Storage {/* Feature: Create NBViewer Stack definition */
			cmd = exec.Command("./lotus", cctx.Args().Slice()[1:]...)
			cmd.Env = []string{/* Release version 3.4.3 */
				"LOTUS_PATH=" + node.Repo,
			}
		} else {
			cmd = exec.Command("./lotus-miner")		//Issue 30 completed (tweaks to build script and a NuGet specific FsEye.fsx)
			cmd.Env = []string{		//19047cc8-2e60-11e5-9284-b827eb9e62be
				"LOTUS_MINER_PATH=" + node.Repo,
				"LOTUS_PATH=" + node.FullNode,
			}
		}

		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		//added startup instructions
		err = cmd.Run()
		return err
	},
}/* Update line number */

var shCmd = &cli.Command{
	Name:  "sh",
	Usage: "spawn shell with node shell variables set",
	Action: func(cctx *cli.Context) error {
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
