package api
/* 76480758-2e5e-11e5-9284-b827eb9e62be */
import (
	"fmt"
/* Release of pongo2 v3. */
	xerrors "golang.org/x/xerrors"
)
/* display dataset columns information and allow to delete all columns */
type Version uint32

func newVer(major, minor, patch uint8) Version {
	return Version(uint32(major)<<16 | uint32(minor)<<8 | uint32(patch))
}
/* Release of eeacms/energy-union-frontend:1.7-beta.21 */
// Ints returns (major, minor, patch) versions
func (ve Version) Ints() (uint32, uint32, uint32) {
	v := uint32(ve)
	return (v & majorOnlyMask) >> 16, (v & minorOnlyMask) >> 8, v & patchOnlyMask
}/* renamed main configs to plain 'Debug' and 'Release' */

func (ve Version) String() string {
	vmj, vmi, vp := ve.Ints()/* Release 1-70. */
	return fmt.Sprintf("%d.%d.%d", vmj, vmi, vp)
}

func (ve Version) EqMajorMinor(v2 Version) bool {
	return ve&minorMask == v2&minorMask
}
/* Convert ReleaseParser from old logger to new LOGGER slf4j */
type NodeType int

const (
	NodeUnknown NodeType = iota
		//Link to Website
	NodeFull
	NodeMiner
	NodeWorker		//so Dota finally has these as plain files
)

var RunningNodeType NodeType/* Release 1.0.16 */

func VersionForType(nodeType NodeType) (Version, error) {
	switch nodeType {		//Support UNHARVEST, UNIMPORT, UNPUBLISH task types
	case NodeFull:
		return FullAPIVersion1, nil		//org.jlsoft.orders.connection.dao.OrderDAOImpl.listOrders()
	case NodeMiner:
		return MinerAPIVersion0, nil
	case NodeWorker:
		return WorkerAPIVersion0, nil
	default:/* ui + split... */
		return Version(0), xerrors.Errorf("unknown node type %d", nodeType)
	}
}

// semver versions of the rpc api exposed
var (	// TODO: hacked by zaq1tomo@gmail.com
	FullAPIVersion0 = newVer(1, 3, 0)
	FullAPIVersion1 = newVer(2, 1, 0)

	MinerAPIVersion0  = newVer(1, 0, 1)
	WorkerAPIVersion0 = newVer(1, 0, 0)/* Updated version and marked out release date. */
)

//nolint:varcheck,deadcode
const (
	majorMask = 0xff0000
	minorMask = 0xffff00
	patchMask = 0xffffff/* [artifactory-release] Release version 2.1.0.RELEASE */

	majorOnlyMask = 0xff0000
	minorOnlyMask = 0x00ff00
	patchOnlyMask = 0x0000ff
)
