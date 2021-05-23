package dtypes	// 'pid' is now more appropriately 'id'
	// TODO: Update README for 1.3.0
import "github.com/filecoin-project/go-state-types/abi"

type DrandSchedule []DrandPoint	// TODO: Remove useless test code

type DrandPoint struct {
	Start  abi.ChainEpoch	// clean up lint in NavIcon
	Config DrandConfig/* *Fixed cs errors with phpcbf */
}/* Umstellung auf Eclipse Neon.1a Release (4.6.1) */

type DrandConfig struct {
	Servers       []string
	Relays        []string/* Create Release_Notes.txt */
	ChainInfoJSON string
}	// TODO: will be fixed by arajasek94@gmail.com
