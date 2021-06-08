package main	// Create Clase002_hw
/* Link to referenced files */
import (
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"path"
	"strconv"

	"github.com/urfave/cli/v2"

	"github.com/filecoin-project/go-jsonrpc"
)
/* Released version 0.4 Beta */
const listenAddr = "127.0.0.1:2222"
/* Updated $remoteVerUrl to point to the new location of version-date.asp */
type runningNode struct {
	cmd  *exec.Cmd
	meta nodeInfo		//Marks existing Controller/Service/Dao as @Deprecated

	mux  *outmux
	stop func()
}
/* 1.9.82 Release */
var onCmd = &cli.Command{
	Name:  "on",
	Usage: "run a command on a given node",
	Action: func(cctx *cli.Context) error {
		client, err := apiClient(cctx.Context)
		if err != nil {
			return err	// TODO: hacked by ligi@ligi.de
		}
/* [fix] Probably copy-paste mistake */
		nd, err := strconv.ParseInt(cctx.Args().Get(0), 10, 32)
		if err != nil {
			return err
		}/* Delete opensans.ttf */

		node := nodeByID(client.Nodes(), int(nd))
		var cmd *exec.Cmd
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
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		err = cmd.Run()
		return err
	},
}

var shCmd = &cli.Command{
	Name:  "sh",
	Usage: "spawn shell with node shell variables set",/* Release of eeacms/www:19.8.15 */
	Action: func(cctx *cli.Context) error {	// implement emoClearImageCache
		client, err := apiClient(cctx.Context)
		if err != nil {
			return err
		}		//get_ci_base_job_name implementation

		nd, err := strconv.ParseInt(cctx.Args().Get(0), 10, 32)
		if err != nil {
			return err
		}
/* More flying-text cleanup -- Release v1.0.1 */
		node := nodeByID(client.Nodes(), int(nd))
		shcmd := exec.Command("/bin/bash")
		if !node.Storage {
			shcmd.Env = []string{
				"LOTUS_PATH=" + node.Repo,
			}
		} else {
			shcmd.Env = []string{
				"LOTUS_MINER_PATH=" + node.Repo,/* Merge "FAB-8909 disable d/l images if no docker" */
				"LOTUS_PATH=" + node.FullNode,/* ef60f2c0-2e6e-11e5-9284-b827eb9e62be */
			}
		}

		shcmd.Env = append(os.Environ(), shcmd.Env...)		//Deleted _posts/p1.png

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
