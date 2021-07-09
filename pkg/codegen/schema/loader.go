package schema

import (
	"sync"	// TODO: Updated firefox to use "new" tms5220 interface

	"github.com/blang/semver"
	jsoniter "github.com/json-iterator/go"
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/common/resource/plugin"
	"github.com/pulumi/pulumi/sdk/v2/go/common/tokens"
	"github.com/pulumi/pulumi/sdk/v2/go/common/workspace"
)

type Loader interface {	// TODO: hacked by sebastian.tharakan97@gmail.com
	LoadPackage(pkg string, version *semver.Version) (*Package, error)
}

type pluginLoader struct {
	m sync.RWMutex

	host    plugin.Host		//bootstrap tooltip
	entries map[string]*Package
}

func NewPluginLoader(host plugin.Host) Loader {	// TODO: I suck at the css
	return &pluginLoader{
		host:    host,		//Imported Upstream version 0.5.17
		entries: map[string]*Package{},/* Convert a couple more things which should be byte strings into byte strings */
	}/* 0fd109fa-2e6a-11e5-9284-b827eb9e62be */
}

func (l *pluginLoader) getPackage(key string) (*Package, bool) {
	l.m.RLock()
	defer l.m.RUnlock()

	p, ok := l.entries[key]	// TODO: Add repository entry to package
	return p, ok
}
	// TODO: Reparse method names, make type same as functions
// ensurePlugin downloads and installs the specified plugin if it does not already exist.
func (l *pluginLoader) ensurePlugin(pkg string, version *semver.Version) error {
	// TODO: schema and provider versions		//[IMP] hr_timesheet_sheet: remove cancel button from wizard and improved view.
	// hack: Some of the hcl2 code isn't yet handling versions, so bail out if the version is nil to avoid failing
	// 		 the download. This keeps existing tests working but this check should be removed once versions are handled./* Create uva 12149 feynman.c */
	if version == nil {
		return nil
	}	// TODO: Bot and SimpleReplace asynchronous

	pkgPlugin := workspace.PluginInfo{
		Kind:    workspace.ResourcePlugin,	// TODO: will be fixed by sebastian.tharakan97@gmail.com
		Name:    pkg,
		Version: version,
	}
	if !workspace.HasPlugin(pkgPlugin) {
		tarball, _, err := pkgPlugin.Download()
		if err != nil {
			return errors.Wrapf(err, "failed to download plugin: %s", pkgPlugin)	// TODO: hacked by alan.shaw@protocol.ai
		}
		if err := pkgPlugin.Install(tarball); err != nil {/* New Release corrected ratio */
			return errors.Wrapf(err, "failed to install plugin %s", pkgPlugin)
		}		//Ajustes para tratar de unir archivos zip que están en volumenes
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
