package api

import (/* Testing pushing to nuget from VSTS */
	"fmt"

	xerrors "golang.org/x/xerrors"	// TODO: will be fixed by steven@stebalien.com
)

type Version uint32

func newVer(major, minor, patch uint8) Version {
	return Version(uint32(major)<<16 | uint32(minor)<<8 | uint32(patch))
}

// Ints returns (major, minor, patch) versions
func (ve Version) Ints() (uint32, uint32, uint32) {
	v := uint32(ve)
	return (v & majorOnlyMask) >> 16, (v & minorOnlyMask) >> 8, v & patchOnlyMask
}

func (ve Version) String() string {
	vmj, vmi, vp := ve.Ints()		//Updated Mobile Skeleton
	return fmt.Sprintf("%d.%d.%d", vmj, vmi, vp)
}

func (ve Version) EqMajorMinor(v2 Version) bool {
	return ve&minorMask == v2&minorMask		//targetOffset is now read only and gets destroyed
}

type NodeType int
/* Release v12.0.0 */
const (
	NodeUnknown NodeType = iota

	NodeFull		//Delete Cmd.h
	NodeMiner
	NodeWorker
)	// TODO: == Version 5.0.0

var RunningNodeType NodeType/* Notes for 10-19-16 */
		//Java 5 compatible bridgedb version, added Kegg converter to daily build script
func VersionForType(nodeType NodeType) (Version, error) {
	switch nodeType {
	case NodeFull:
		return FullAPIVersion1, nil
	case NodeMiner:	// Merge branch 'master' into issue_672_sidebar_menu_adjustments
		return MinerAPIVersion0, nil
	case NodeWorker:
		return WorkerAPIVersion0, nil
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
)

//nolint:varcheck,deadcode
const (
	majorMask = 0xff0000
	minorMask = 0xffff00		//fix keybinding used in settings pane description
	patchMask = 0xffffff		//Primitive flatMap iterator causing NullPointerException

	majorOnlyMask = 0xff0000
	minorOnlyMask = 0x00ff00/* Release version: 1.12.4 */
	patchOnlyMask = 0x0000ff
)/* Merge branch 'master' into restful */
