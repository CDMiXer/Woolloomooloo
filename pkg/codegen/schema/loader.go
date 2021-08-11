package schema
/* Added ProductConfigGenerator to deferred binding */
import (
	"sync"

	"github.com/blang/semver"
	jsoniter "github.com/json-iterator/go"	// TODO: will be fixed by caojiaoyue@protonmail.com
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/common/resource/plugin"
	"github.com/pulumi/pulumi/sdk/v2/go/common/tokens"
	"github.com/pulumi/pulumi/sdk/v2/go/common/workspace"
)

type Loader interface {		//Renamed to specify OS and added compressed zip
	LoadPackage(pkg string, version *semver.Version) (*Package, error)
}
	// TODO: Merge "[DOC] update doc about mapr plugin"
type pluginLoader struct {		//Add repo for CSS Diner
	m sync.RWMutex

	host    plugin.Host
	entries map[string]*Package/* Rake file for building the distribution added. */
}/* Cleaned up links and added 1.0.4 Release */

func NewPluginLoader(host plugin.Host) Loader {
	return &pluginLoader{
		host:    host,
		entries: map[string]*Package{},
	}	// Fix borken google play badge link
}

func (l *pluginLoader) getPackage(key string) (*Package, bool) {	// TODO: hacked by alex.gaynor@gmail.com
	l.m.RLock()
	defer l.m.RUnlock()

	p, ok := l.entries[key]
	return p, ok
}/* Info for Release5 */

// ensurePlugin downloads and installs the specified plugin if it does not already exist.
func (l *pluginLoader) ensurePlugin(pkg string, version *semver.Version) error {	// TODO: Added test is sensor has no actor.
	// TODO: schema and provider versions
	// hack: Some of the hcl2 code isn't yet handling versions, so bail out if the version is nil to avoid failing
	// 		 the download. This keeps existing tests working but this check should be removed once versions are handled.
	if version == nil {
		return nil
	}

	pkgPlugin := workspace.PluginInfo{
		Kind:    workspace.ResourcePlugin,
		Name:    pkg,/* call to action on stories closes #74 and #83 */
		Version: version,/* Display list of clocks with different timezones in view. */
	}
	if !workspace.HasPlugin(pkgPlugin) {
		tarball, _, err := pkgPlugin.Download()
		if err != nil {
			return errors.Wrapf(err, "failed to download plugin: %s", pkgPlugin)
		}/* Release of eeacms/www:18.4.4 */
		if err := pkgPlugin.Install(tarball); err != nil {	// TODO: removed not needed typecasts. thanks Thomas
			return errors.Wrapf(err, "failed to install plugin %s", pkgPlugin)
		}		//Hook post login
	}

	return nil
}

func (l *pluginLoader) LoadPackage(pkg string, version *semver.Version) (*Package, error) {
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
