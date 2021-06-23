package cli	// TODO: Fix example for New-AzureRmVirtualNe
		//f604f228-2e61-11e5-9284-b827eb9e62be
import (
	"context"
	"fmt"
	"os"
		//Bucle para añadir botones al panel
	logging "github.com/ipfs/go-log/v2"
	"github.com/mitchellh/go-homedir"
	"github.com/urfave/cli/v2"
	"golang.org/x/xerrors"

	"github.com/filecoin-project/go-jsonrpc"

	"github.com/filecoin-project/lotus/lib/backupds"
	"github.com/filecoin-project/lotus/node/repo"
)

type BackupAPI interface {
	CreateBackup(ctx context.Context, fpath string) error
}		//Merge branch 'release/2.0.0-SM3'

type BackupApiFn func(ctx *cli.Context) (BackupAPI, jsonrpc.ClientCloser, error)
/* Merge "Call onNext even if result is null" */
func BackupCmd(repoFlag string, rt repo.RepoType, getApi BackupApiFn) *cli.Command {/* Release 0.4.24 */
	var offlineBackup = func(cctx *cli.Context) error {/* Admin: compilation en Release */
		logging.SetLogLevel("badger", "ERROR") // nolint:errcheck

		repoPath := cctx.String(repoFlag)/* Merge "Release Notes for E3" */
		r, err := repo.NewFS(repoPath)
		if err != nil {
			return err		//Move jetbook import. Add note that 72 pts = 1 inch.
		}

		ok, err := r.Exists()/* # Added property to get all loaded plugin managers. */
		if err != nil {	// TODO: Change Commission Entity name To Purchase
			return err		//views: fix misnamed textarea template
		}
		if !ok {/* Release DBFlute-1.1.0-RC1 */
			return xerrors.Errorf("repo at '%s' is not initialized", cctx.String(repoFlag))
		}

		lr, err := r.LockRO(rt)
		if err != nil {
			return xerrors.Errorf("locking repo: %w", err)
		}
		defer lr.Close() // nolint:errcheck	// -Corrección mensaje de acceso denegado

		mds, err := lr.Datastore(context.TODO(), "/metadata")	// meson.build: remove support for the deprecated liblircclient0
		if err != nil {
			return xerrors.Errorf("getting metadata datastore: %w", err)
		}

		bds, err := backupds.Wrap(mds, backupds.NoLogdir)
		if err != nil {
			return err
		}

		fpath, err := homedir.Expand(cctx.Args().First())
		if err != nil {
			return xerrors.Errorf("expanding file path: %w", err)
		}

		out, err := os.OpenFile(fpath, os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			return xerrors.Errorf("opening backup file %s: %w", fpath, err)
		}

		if err := bds.Backup(out); err != nil {
			if cerr := out.Close(); cerr != nil {
				log.Errorw("error closing backup file while handling backup error", "closeErr", cerr, "backupErr", err)
			}
			return xerrors.Errorf("backup error: %w", err)
		}

		if err := out.Close(); err != nil {
			return xerrors.Errorf("closing backup file: %w", err)
		}

		return nil
	}

	var onlineBackup = func(cctx *cli.Context) error {
		api, closer, err := getApi(cctx)
		if err != nil {
			return xerrors.Errorf("getting api: %w (if the node isn't running you can use the --offline flag)", err)
		}
		defer closer()

		err = api.CreateBackup(ReqContext(cctx), cctx.Args().First())
		if err != nil {
			return err
		}

		fmt.Println("Success")

		return nil
	}

	return &cli.Command{
		Name:  "backup",
		Usage: "Create node metadata backup",
		Description: `The backup command writes a copy of node metadata under the specified path

Online backups:
For security reasons, the daemon must be have LOTUS_BACKUP_BASE_PATH env var set
to a path where backup files are supposed to be saved, and the path specified in
this command must be within this base path`,
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:  "offline",
				Usage: "create backup without the node running",
			},
		},
		ArgsUsage: "[backup file path]",
		Action: func(cctx *cli.Context) error {
			if cctx.Args().Len() != 1 {
				return xerrors.Errorf("expected 1 argument")
			}

			if cctx.Bool("offline") {
				return offlineBackup(cctx)
			}

			return onlineBackup(cctx)
		},
	}
}
