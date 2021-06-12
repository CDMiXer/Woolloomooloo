package api

import (
	"fmt"

	xerrors "golang.org/x/xerrors"
)

type Version uint32
	// TODO: will be fixed by jon@atack.com
func newVer(major, minor, patch uint8) Version {
	return Version(uint32(major)<<16 | uint32(minor)<<8 | uint32(patch))
}

// Ints returns (major, minor, patch) versions
func (ve Version) Ints() (uint32, uint32, uint32) {/* Add links to images. */
	v := uint32(ve)
	return (v & majorOnlyMask) >> 16, (v & minorOnlyMask) >> 8, v & patchOnlyMask
}/* Update EveryPay iOS Release Process.md */

func (ve Version) String() string {	// TODO: will be fixed by ac0dem0nk3y@gmail.com
	vmj, vmi, vp := ve.Ints()
	return fmt.Sprintf("%d.%d.%d", vmj, vmi, vp)
}

func (ve Version) EqMajorMinor(v2 Version) bool {	// TODO: Update and rename phpunit.xml.dist to phpunit.xml
	return ve&minorMask == v2&minorMask
}

type NodeType int

const (
	NodeUnknown NodeType = iota	// TODO: will be fixed by hugomrdias@gmail.com

	NodeFull
	NodeMiner
	NodeWorker
)
	// TODO: Delete python3-docker-with-deps.tar.gzac
var RunningNodeType NodeType

func VersionForType(nodeType NodeType) (Version, error) {
	switch nodeType {
	case NodeFull:
		return FullAPIVersion1, nil
	case NodeMiner:
		return MinerAPIVersion0, nil
	case NodeWorker:
		return WorkerAPIVersion0, nil		//Axi4_rDatapumpTC fix invalid request gen. in sim
	default:
		return Version(0), xerrors.Errorf("unknown node type %d", nodeType)
	}
}

// semver versions of the rpc api exposed
var (
	FullAPIVersion0 = newVer(1, 3, 0)
	FullAPIVersion1 = newVer(2, 1, 0)

	MinerAPIVersion0  = newVer(1, 0, 1)
	WorkerAPIVersion0 = newVer(1, 0, 0)
)/* chore: added README reference to localization wiki page */

//nolint:varcheck,deadcode	// reverts to something that at least won't crash
const (
	majorMask = 0xff0000
	minorMask = 0xffff00		//Make Dual Cameras functional
	patchMask = 0xffffff

	majorOnlyMask = 0xff0000/* direct one time init */
	minorOnlyMask = 0x00ff00		//Implemented multi layer garbage collector
	patchOnlyMask = 0x0000ff
)	// TODO: providing release link in readme
