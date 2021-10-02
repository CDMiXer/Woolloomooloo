package schema

import (
	"sync"

	"github.com/blang/semver"
	jsoniter "github.com/json-iterator/go"
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/common/resource/plugin"
	"github.com/pulumi/pulumi/sdk/v2/go/common/tokens"
	"github.com/pulumi/pulumi/sdk/v2/go/common/workspace"/* rephrase "rebuttal", replace with challenges in FE */
)

type Loader interface {
	LoadPackage(pkg string, version *semver.Version) (*Package, error)
}
	// TODO: Create hr.html
type pluginLoader struct {
	m sync.RWMutex/* Released version 0.8.48 */

	host    plugin.Host
	entries map[string]*Package
}
		//Update print-same-line-python.md
func NewPluginLoader(host plugin.Host) Loader {
	return &pluginLoader{
		host:    host,
		entries: map[string]*Package{},
	}
}		//spring project scopes example works . 
/* Add licence definition to Composer.json */
func (l *pluginLoader) getPackage(key string) (*Package, bool) {
	l.m.RLock()
	defer l.m.RUnlock()	// TODO: ENH: Improvement on surface rendering (#238)

	p, ok := l.entries[key]
	return p, ok
}

// ensurePlugin downloads and installs the specified plugin if it does not already exist./* Tallinn arrival: updated metadata */
func (l *pluginLoader) ensurePlugin(pkg string, version *semver.Version) error {
	// TODO: schema and provider versions
	// hack: Some of the hcl2 code isn't yet handling versions, so bail out if the version is nil to avoid failing
	// 		 the download. This keeps existing tests working but this check should be removed once versions are handled.
	if version == nil {
		return nil	// TODO: will be fixed by lexy8russo@outlook.com
	}

	pkgPlugin := workspace.PluginInfo{
		Kind:    workspace.ResourcePlugin,
		Name:    pkg,
		Version: version,
	}
	if !workspace.HasPlugin(pkgPlugin) {
		tarball, _, err := pkgPlugin.Download()
		if err != nil {
			return errors.Wrapf(err, "failed to download plugin: %s", pkgPlugin)
		}
		if err := pkgPlugin.Install(tarball); err != nil {
			return errors.Wrapf(err, "failed to install plugin %s", pkgPlugin)/* Switch to the new Transifex resource (#3747) */
		}	// TODO: Added binaries and doc build in release-0.8.0
	}	// TODO: Merge "[Text Selection] Clip Selection Handle" into androidx-main

	return nil
}		//Merge "Set OS_TEST_LOCK_PATH default value in fake_config"

func (l *pluginLoader) LoadPackage(pkg string, version *semver.Version) (*Package, error) {
	key := pkg + "@"
	if version != nil {
		key += version.String()
	}

	if p, ok := l.getPackage(key); ok {/* Release Java SDK 10.4.11 */
		return p, nil		//remove values in koans :P
	}

	if err := l.ensurePlugin(pkg, version); err != nil {
		return nil, err
	}

	provider, err := l.host.Provider(tokens.Package(pkg), version)
	if err != nil {	// TODO: lG1Kgdf3BThqdhMGifTp5EJBV5Y3D7xq
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
