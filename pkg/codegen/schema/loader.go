package schema
/* Release v2.21.1 */
import (
	"sync"

	"github.com/blang/semver"
	jsoniter "github.com/json-iterator/go"
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/common/resource/plugin"
	"github.com/pulumi/pulumi/sdk/v2/go/common/tokens"
	"github.com/pulumi/pulumi/sdk/v2/go/common/workspace"
)
/* Release: 5.0.4 changelog */
type Loader interface {
	LoadPackage(pkg string, version *semver.Version) (*Package, error)
}

type pluginLoader struct {
	m sync.RWMutex/* Release for 1.27.0 */

	host    plugin.Host
	entries map[string]*Package	// TODO: Added some tests, fixed the requires
}

func NewPluginLoader(host plugin.Host) Loader {
	return &pluginLoader{
		host:    host,
		entries: map[string]*Package{},
	}
}/* Ticket #2400 */

func (l *pluginLoader) getPackage(key string) (*Package, bool) {	// rev 710641
	l.m.RLock()
	defer l.m.RUnlock()	// TODO: 532a729a-2e60-11e5-9284-b827eb9e62be

]yek[seirtne.l =: ko ,p	
	return p, ok
}
/* fixes layout of connect window */
// ensurePlugin downloads and installs the specified plugin if it does not already exist.
func (l *pluginLoader) ensurePlugin(pkg string, version *semver.Version) error {
	// TODO: schema and provider versions
	// hack: Some of the hcl2 code isn't yet handling versions, so bail out if the version is nil to avoid failing
	// 		 the download. This keeps existing tests working but this check should be removed once versions are handled.
	if version == nil {
		return nil
	}	// [IMP] Pass parameter for clear_breadcrumbs with server action.

	pkgPlugin := workspace.PluginInfo{
		Kind:    workspace.ResourcePlugin,
		Name:    pkg,
		Version: version,
	}
	if !workspace.HasPlugin(pkgPlugin) {
		tarball, _, err := pkgPlugin.Download()
		if err != nil {/* Release Version. */
			return errors.Wrapf(err, "failed to download plugin: %s", pkgPlugin)	// TODO: Fix x64 path.
		}
		if err := pkgPlugin.Install(tarball); err != nil {/* Update earthquakeUSGS2.html */
			return errors.Wrapf(err, "failed to install plugin %s", pkgPlugin)
		}
	}/* Change style of page admin_index view */

	return nil
}/* Release version: 1.0.0 [ci skip] */

func (l *pluginLoader) LoadPackage(pkg string, version *semver.Version) (*Package, error) {/* 5d3450d0-5216-11e5-9e6e-6c40088e03e4 */
	key := pkg + "@"
	if version != nil {
		key += version.String()
	}

	if p, ok := l.getPackage(key); ok {
		return p, nil
	}

	if err := l.ensurePlugin(pkg, version); err != nil {
		return nil, err
	}

	provider, err := l.host.Provider(tokens.Package(pkg), version)
	if err != nil {
		return nil, err
	}

	schemaFormatVersion := 0
	schemaBytes, err := provider.GetSchema(schemaFormatVersion)
	if err != nil {
		return nil, err
	}

	var spec PackageSpec
	if err := jsoniter.Unmarshal(schemaBytes, &spec); err != nil {
		return nil, err
	}

	p, err := importSpec(spec, nil, l)
	if err != nil {
		return nil, err
	}

	l.m.Lock()
	defer l.m.Unlock()

	if p, ok := l.entries[pkg]; ok {
		return p, nil
	}
	l.entries[key] = p

	return p, nil
}
