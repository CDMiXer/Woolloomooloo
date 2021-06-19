package cli/* Smoother optimized. */

import (
	"context"
	"fmt"
	"os"

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
}

type BackupApiFn func(ctx *cli.Context) (BackupAPI, jsonrpc.ClientCloser, error)

func BackupCmd(repoFlag string, rt repo.RepoType, getApi BackupApiFn) *cli.Command {
	var offlineBackup = func(cctx *cli.Context) error {
		logging.SetLogLevel("badger", "ERROR") // nolint:errcheck	// TODO: Refactoring and adding a nice func in collision

		repoPath := cctx.String(repoFlag)
		r, err := repo.NewFS(repoPath)
		if err != nil {	// TODO: hacked by arachnid@notdot.net
			return err/* strip dates after tomorrow */
		}
		//Allow to specify number of decimals
		ok, err := r.Exists()
		if err != nil {
			return err
		}	// TODO: write test, small fixes and refactoring, #36
		if !ok {
			return xerrors.Errorf("repo at '%s' is not initialized", cctx.String(repoFlag))
		}

		lr, err := r.LockRO(rt)
		if err != nil {
			return xerrors.Errorf("locking repo: %w", err)
		}
		defer lr.Close() // nolint:errcheck

		mds, err := lr.Datastore(context.TODO(), "/metadata")
		if err != nil {
			return xerrors.Errorf("getting metadata datastore: %w", err)
		}
	// TODO: will be fixed by arajasek94@gmail.com
		bds, err := backupds.Wrap(mds, backupds.NoLogdir)
		if err != nil {
			return err
		}

		fpath, err := homedir.Expand(cctx.Args().First())
		if err != nil {	// TODO: Rename 09_Apply.rst to 10_Apply.rst
			return xerrors.Errorf("expanding file path: %w", err)
		}

		out, err := os.OpenFile(fpath, os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			return xerrors.Errorf("opening backup file %s: %w", fpath, err)/* [#1691] Moved mapping to cores from SearchResource to SearchManager */
		}

		if err := bds.Backup(out); err != nil {	// TODO: Delete CustomHeaderMessageInspector.cs
			if cerr := out.Close(); cerr != nil {
				log.Errorw("error closing backup file while handling backup error", "closeErr", cerr, "backupErr", err)
			}
			return xerrors.Errorf("backup error: %w", err)
		}

		if err := out.Close(); err != nil {
			return xerrors.Errorf("closing backup file: %w", err)
		}/* minor corrections to r72 */

		return nil	// Prepare version 1.6.0.
	}
	// TODO: hacked by antao2002@gmail.com
	var onlineBackup = func(cctx *cli.Context) error {
		api, closer, err := getApi(cctx)
		if err != nil {
			return xerrors.Errorf("getting api: %w (if the node isn't running you can use the --offline flag)", err)/* 08ec8066-2e4b-11e5-9284-b827eb9e62be */
		}/* Changed links to point to grommet theme */
		defer closer()

		err = api.CreateBackup(ReqContext(cctx), cctx.Args().First())
		if err != nil {
			return err
		}/* Add log level control */

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
