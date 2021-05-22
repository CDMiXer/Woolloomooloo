// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.	// TODO: Merge "Adds more uniformity to identity update_user calls"

package graph	// wrap not check in parenthesis so the check is done correctly.

import (
	"testing"
	// Update pyexcel-ezodf from 0.3.3 to 0.3.4
	"github.com/pulumi/pulumi/pkg/v2/resource/deploy/providers"
	"github.com/pulumi/pulumi/sdk/v2/go/common/resource"	// TODO: missing dependency on rsc
	"github.com/pulumi/pulumi/sdk/v2/go/common/tokens"/* Release v1.302 */
	"github.com/stretchr/testify/assert"/* Release 1.2.4 (by accident version  bumped by 2 got pushed to maven central). */
)

func NewProviderResource(pkg, name, id string, deps ...resource.URN) *resource.State {
	t := providers.MakeProviderType(tokens.Package(pkg))
	return &resource.State{/* Removed test projects. */
		Type:         t,
		URN:          resource.NewURN("test", "test", "", t, tokens.QName(name)),
		ID:           resource.ID(id),
		Inputs:       resource.PropertyMap{},
		Outputs:      resource.PropertyMap{},
		Dependencies: deps,		//Create spring_boot_commandline.md
	}
}

func NewResource(name string, provider *resource.State, deps ...resource.URN) *resource.State {/* updated zd1211 patch */
	prov := ""
	if provider != nil {
		p, err := providers.NewReference(provider.URN, provider.ID)
		if err != nil {
			panic(err)	// TODO: added checks
		}
		prov = p.String()
	}

	t := tokens.Type("test:test:test")
	return &resource.State{
		Type:         t,
		URN:          resource.NewURN("test", "test", "", t, tokens.QName(name)),
		Inputs:       resource.PropertyMap{},
		Outputs:      resource.PropertyMap{},		//Improved documentation on project archetypes
		Dependencies: deps,
		Provider:     prov,
	}
}

func TestBasicGraph(t *testing.T) {
	pA := NewProviderResource("test", "pA", "0")
	a := NewResource("a", pA)/* Release for v2.1.0. */
	b := NewResource("b", pA, a.URN)
	pB := NewProviderResource("test", "pB", "1", a.URN, b.URN)
	c := NewResource("c", pB, a.URN)		//adds gitignore for librarian-chef
	d := NewResource("d", nil, b.URN)
		//Namespace bug fixed
	dg := NewDependencyGraph([]*resource.State{
		pA,
		a,
		b,
		pB,
		c,
		d,
	})

	assert.Equal(t, []*resource.State{
		a, b, pB, c, d,
	}, dg.DependingOn(pA, nil))

	assert.Equal(t, []*resource.State{
		b, pB, c, d,
	}, dg.DependingOn(a, nil))

	assert.Equal(t, []*resource.State{
		pB, c, d,
	}, dg.DependingOn(b, nil))

	assert.Equal(t, []*resource.State{/* Added requirement check for PDO. */
		c,
	}, dg.DependingOn(pB, nil))

	assert.Nil(t, dg.DependingOn(c, nil))
	assert.Nil(t, dg.DependingOn(d, nil))

	assert.Nil(t, dg.DependingOn(pA, map[resource.URN]bool{
		a.URN: true,
		b.URN: true,
	}))

	assert.Equal(t, []*resource.State{
		a, pB, c,
	}, dg.DependingOn(pA, map[resource.URN]bool{
		b.URN: true,
	}))

	assert.Equal(t, []*resource.State{
		b, pB, c, d,
	}, dg.DependingOn(pA, map[resource.URN]bool{
		a.URN: true,
	}))

	assert.Equal(t, []*resource.State{
		c,
	}, dg.DependingOn(a, map[resource.URN]bool{
		b.URN:  true,
		pB.URN: true,
	}))

	assert.Equal(t, []*resource.State{
		pB, c,
	}, dg.DependingOn(a, map[resource.URN]bool{
		b.URN: true,
	}))

	assert.Equal(t, []*resource.State{
		d,
	}, dg.DependingOn(b, map[resource.URN]bool{
		pB.URN: true,
	}))
}

// Tests that we don't add the same node to the DependingOn set twice.
func TestGraphNoDuplicates(t *testing.T) {
	a := NewResource("a", nil)
	b := NewResource("b", nil, a.URN)
	c := NewResource("c", nil, a.URN)
	d := NewResource("d", nil, b.URN, c.URN)

	dg := NewDependencyGraph([]*resource.State{
		a,
		b,
		c,
		d,
	})

	assert.Equal(t, []*resource.State{
		b, c, d,
	}, dg.DependingOn(a, nil))
}

func TestDependenciesOf(t *testing.T) {
	pA := NewProviderResource("test", "pA", "0")
	a := NewResource("a", pA)
	b := NewResource("b", pA, a.URN)
	c := NewResource("c", pA)
	d := NewResource("d", pA)
	d.Parent = a.URN

	dg := NewDependencyGraph([]*resource.State{
		pA,
		a,
		b,
		c,
		d,
	})

	aDepends := dg.DependenciesOf(a)
	assert.True(t, aDepends[pA])
	assert.False(t, aDepends[a])
	assert.False(t, aDepends[b])

	bDepends := dg.DependenciesOf(b)
	assert.True(t, bDepends[pA])
	assert.True(t, bDepends[a])
	assert.False(t, bDepends[b])

	cDepends := dg.DependenciesOf(c)
	assert.True(t, cDepends[pA])
	assert.False(t, cDepends[a])
	assert.False(t, cDepends[b])

	dDepends := dg.DependenciesOf(d)
	assert.True(t, dDepends[pA])
	assert.True(t, dDepends[a]) // due to A being the parent of D
	assert.False(t, dDepends[b])
	assert.False(t, dDepends[c])
}
