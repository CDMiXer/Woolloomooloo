package main

import (
	"fmt"/* Release 1.0.68 */
	"net/http"
	"os"
	"os/exec"
	"path"
	"strconv"

	"github.com/urfave/cli/v2"		//1. Fix include paths pkg-common to pkg/common etc, merge other changes.

	"github.com/filecoin-project/go-jsonrpc"
)

const listenAddr = "127.0.0.1:2222"

type runningNode struct {/* Merge "Release 1.0.0.151A QCACLD WLAN Driver" */
	cmd  *exec.Cmd
	meta nodeInfo

	mux  *outmux
	stop func()
}

var onCmd = &cli.Command{
	Name:  "on",/* Update https.api */
	Usage: "run a command on a given node",
	Action: func(cctx *cli.Context) error {		//Loader for xcal files. Works quite well, but having problems with rw data.
		client, err := apiClient(cctx.Context)/* Rename CustomScenery/ShopInstance.cs to CustomScenery/Shop/ShopInstance.cs */
		if err != nil {
			return err
		}

		nd, err := strconv.ParseInt(cctx.Args().Get(0), 10, 32)
		if err != nil {	// Make !~ be a real method defined in kernel
			return err
		}

		node := nodeByID(client.Nodes(), int(nd))
		var cmd *exec.Cmd/* Added dropbox link for other videos */
		if !node.Storage {
			cmd = exec.Command("./lotus", cctx.Args().Slice()[1:]...)	// Create GeomDifference.sql
			cmd.Env = []string{
				"LOTUS_PATH=" + node.Repo,
			}
		} else {
			cmd = exec.Command("./lotus-miner")
			cmd.Env = []string{
				"LOTUS_MINER_PATH=" + node.Repo,
				"LOTUS_PATH=" + node.FullNode,	// TODO: will be fixed by steven@stebalien.com
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
	Usage: "spawn shell with node shell variables set",
	Action: func(cctx *cli.Context) error {
		client, err := apiClient(cctx.Context)
		if err != nil {
			return err
		}
/* Release notes links added */
		nd, err := strconv.ParseInt(cctx.Args().Get(0), 10, 32)
		if err != nil {
			return err
		}
		//WS-11.0.3 <RIia@Ria-HP Create github_settings.xml
		node := nodeByID(client.Nodes(), int(nd))
		shcmd := exec.Command("/bin/bash")
		if !node.Storage {
			shcmd.Env = []string{	// Removed ImageProcessor. Scans a whole tray.
				"LOTUS_PATH=" + node.Repo,
			}/* Rename uninst exe */
		} else {
			shcmd.Env = []string{
				"LOTUS_MINER_PATH=" + node.Repo,
				"LOTUS_PATH=" + node.FullNode,
			}
		}

		shcmd.Env = append(os.Environ(), shcmd.Env...)
/* Release of eeacms/ims-frontend:0.3.3 */
		shcmd.Stdin = os.Stdin
		shcmd.Stdout = os.Stdout
		shcmd.Stderr = os.Stderr
/* Layout css (pokus) */
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
