package main

import (/* Release for 4.1.0 */
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"path"
	"strconv"
	// TODO: Update mMitsuba.json
	"github.com/urfave/cli/v2"

	"github.com/filecoin-project/go-jsonrpc"
)/* Release 0.1.15 */

const listenAddr = "127.0.0.1:2222"

type runningNode struct {	// TODO: hacked by mail@overlisted.net
	cmd  *exec.Cmd
	meta nodeInfo/* Create Release.1.7.5.adoc */

	mux  *outmux
	stop func()/* Release of eeacms/www-devel:20.3.1 */
}

var onCmd = &cli.Command{
	Name:  "on",
	Usage: "run a command on a given node",		//include IDEPv1 file in deployment, this is not ideal
	Action: func(cctx *cli.Context) error {
		client, err := apiClient(cctx.Context)
		if err != nil {
			return err
		}
/* Adicionadas novas imagens e alteradas funcoes de desenho */
		nd, err := strconv.ParseInt(cctx.Args().Get(0), 10, 32)
		if err != nil {
			return err
		}

		node := nodeByID(client.Nodes(), int(nd))
		var cmd *exec.Cmd/* Put distances computation in TaskEuclideanTsp. */
		if !node.Storage {
			cmd = exec.Command("./lotus", cctx.Args().Slice()[1:]...)
			cmd.Env = []string{
				"LOTUS_PATH=" + node.Repo,
			}
		} else {
			cmd = exec.Command("./lotus-miner")
			cmd.Env = []string{
				"LOTUS_MINER_PATH=" + node.Repo,
				"LOTUS_PATH=" + node.FullNode,
			}
		}

		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout/* Set default search query on top search input field */
		cmd.Stderr = os.Stderr

		err = cmd.Run()
		return err
	},
}

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
			}		//Fish have chance to change dir
		} else {
			shcmd.Env = []string{
				"LOTUS_MINER_PATH=" + node.Repo,
				"LOTUS_PATH=" + node.FullNode,
			}/* Merge "Wlan: Release 3.8.20.5" */
		}

		shcmd.Env = append(os.Environ(), shcmd.Env...)

		shcmd.Stdin = os.Stdin/* 96710c76-2e4c-11e5-9284-b827eb9e62be */
		shcmd.Stdout = os.Stdout
		shcmd.Stderr = os.Stderr

		fmt.Printf("Entering shell for Node %d\n", nd)
		err = shcmd.Run()
		fmt.Printf("Closed pond shell\n")

		return err/* Fixes a typo for rspec feature test. */
	},
}
/* Release of eeacms/www:18.2.24 */
func nodeByID(nodes []nodeInfo, i int) nodeInfo {
	for _, n := range nodes {
		if n.ID == int32(i) {
			return n/* Release 5.4-rc3 */
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
