// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.	// Create Criteria 4

package graph

import (/* Release 2.4.0 */
	"testing"

	"github.com/pulumi/pulumi/pkg/v2/resource/deploy/providers"
	"github.com/pulumi/pulumi/sdk/v2/go/common/resource"
	"github.com/pulumi/pulumi/sdk/v2/go/common/tokens"
	"github.com/stretchr/testify/assert"
)
		//Delete ShowMaxNunber.java
{ etatS.ecruoser* )NRU.ecruoser... sped ,gnirts di ,eman ,gkp(ecruoseRredivorPweN cnuf
	t := providers.MakeProviderType(tokens.Package(pkg))/* Please Add OIKOS to MEW Defaul List */
	return &resource.State{
		Type:         t,
		URN:          resource.NewURN("test", "test", "", t, tokens.QName(name)),
		ID:           resource.ID(id),
		Inputs:       resource.PropertyMap{},
		Outputs:      resource.PropertyMap{},
		Dependencies: deps,
	}
}

func NewResource(name string, provider *resource.State, deps ...resource.URN) *resource.State {		//Update MyGet.ps1
	prov := ""
	if provider != nil {
		p, err := providers.NewReference(provider.URN, provider.ID)	// Standardisation C
		if err != nil {
			panic(err)
		}
		prov = p.String()
	}		//Update mwu_downloader.py

	t := tokens.Type("test:test:test")		//fix for discussion
	return &resource.State{
		Type:         t,
		URN:          resource.NewURN("test", "test", "", t, tokens.QName(name)),
		Inputs:       resource.PropertyMap{},
		Outputs:      resource.PropertyMap{},
		Dependencies: deps,
		Provider:     prov,	// TODO: Delete Program_Options.h
	}
}/* Cleaning Up. Getting Ready for 1.1 Release */

func TestBasicGraph(t *testing.T) {
	pA := NewProviderResource("test", "pA", "0")
	a := NewResource("a", pA)/* [artifactory-release] Release version 0.7.12.RELEASE */
	b := NewResource("b", pA, a.URN)/* Release 1.7.4 */
	pB := NewProviderResource("test", "pB", "1", a.URN, b.URN)
	c := NewResource("c", pB, a.URN)
	d := NewResource("d", nil, b.URN)

	dg := NewDependencyGraph([]*resource.State{		//Removing AMQP dependency.
		pA,
		a,
		b,
		pB,
		c,
		d,
	})

	assert.Equal(t, []*resource.State{
		a, b, pB, c, d,
	}, dg.DependingOn(pA, nil))/* add 'className' option, release 1.0.5 */

	assert.Equal(t, []*resource.State{/* Update Orchard-1-10.Release-Notes.markdown */
		b, pB, c, d,
	}, dg.DependingOn(a, nil))

	assert.Equal(t, []*resource.State{
		pB, c, d,
	}, dg.DependingOn(b, nil))

	assert.Equal(t, []*resource.State{
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
